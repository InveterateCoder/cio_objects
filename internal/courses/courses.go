package courses

import (
	"cio_objects/internal/bqtools"
	"cio_objects/internal/paths"
	"fmt"
	"strings"

	"google.golang.org/api/iterator"
)

func NewCioCourses(region paths.Region) *CioCourses {
	if region != paths.BR && region != paths.MX && region != paths.CO && region != paths.PE {
		panic("wrong region")
	}
	cioCourse := CioCourses{
		region:          region,
		regionLowerText: strings.ToLower(string(region)),
		bq:              bqtools.NewBQ(),
	}
	if region == paths.BR {
		cioCourse.courseTable = "`ebac-287911.prod_lms_br.course_view`"
		cioCourse.nomenclatureTable = "`ebac-287911.prod_cms_br.nomenclature_view`"
		cioCourse.enrollmentTable = "`ebac-287911.prod_lms_br.enrollment_view`"
		cioCourse.userTable = "`ebac-287911.prod_lms_br.user_view`"
	} else {
		cioCourse.courseTable = "`ebac-287911.prod_lms_mx.course_view`"
		cioCourse.nomenclatureTable = "`ebac-287911.prod_cms_mx.nomenclature_view`"
		cioCourse.enrollmentTable = "`ebac-287911.prod_lms_mx.enrollment_view`"
		cioCourse.userTable = "`ebac-287911.prod_lms_mx.user_view`"
	}
	return &cioCourse
}

func (cioCourses *CioCourses) Close() {
	cioCourses.bq.Close()
}

func (cioCourses *CioCourses) GetAllCourses() (courses []Course) {
	it, err := cioCourses.bq.Query(fmt.Sprintf(`select c.id, c.slug,
c.title as name, n.short_name, c.type, c.status, c.active, c.published,
c.published_at, c.created, c.modified, c.external_id, c.cms_id, c.duration,
c.start_at, c.not_for_sale, c.graduation_criteria  from %s c
left join %s n on n.id = cast(c.cms_id as int64);`, cioCourses.courseTable, cioCourses.nomenclatureTable))
	if err != nil {
		panic(err)
	}
	for {
		var c bqCourse
		if err := it.Next(&c); err != nil {
			if err == iterator.Done {
				break
			} else {
				panic(err)
			}
		}
		courses = append(courses, c.parseToCourse())
	}
	return
}

func (cioCourses *CioCourses) GetUsersWithActiveCourses() (users []UserWithActiveCourses) {
	var domainQuery string
	switch cioCourses.region {
	case paths.MX:
		domainQuery = "\nand (u.domain is null or u.domain = 'MX')"
	case paths.CO:
		domainQuery = "\nand u.domain = 'CO'"
	case paths.PE:
		domainQuery = "\nand u.domain = 'PE'"
	}
	it, err := cioCourses.bq.Query(fmt.Sprintf(`select u.id as user_id, u.email,
string_agg(e.course_id, ',') as courses from %s e
join %s u on u.id = e.user_id
where e.active is true%s
and (e.date_from is null or e.date_from < current_datetime())
and (e.date_to is null or e.date_to > current_datetime())
group by u.id, u.email`, cioCourses.enrollmentTable, cioCourses.userTable, domainQuery))
	if err != nil {
		panic(err)
	}
	for {
		var user bqUserWithActiveCourses
		if err := it.Next(&user); err != nil {
			if err == iterator.Done {
				break
			} else {
				panic(err)
			}
		}
		users = append(users, user.parseToUserWithActiveCourses())
	}
	return
}
