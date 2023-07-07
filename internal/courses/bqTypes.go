package courses

import (
	"strings"

	"cloud.google.com/go/bigquery"
)

type bqCourse struct {
	ID                 string                `bigquery:"id"`
	Slug               bigquery.NullString   `bigquery:"slug"`
	Name               bigquery.NullString   `bigquery:"name"`
	ShortName          bigquery.NullString   `bigquery:"short_name"`
	Type               bigquery.NullString   `bigquery:"type"`
	Status             bigquery.NullString   `bigquery:"status"`
	Active             bigquery.NullBool     `bigquery:"active"`
	Published          bigquery.NullBool     `bigquery:"published"`
	PublishedAt        bigquery.NullDateTime `bigquery:"published_at"`
	Created            bigquery.NullDateTime `bigquery:"created"`
	Modified           bigquery.NullDateTime `bigquery:"modified"`
	ExternalID         bigquery.NullString   `bigquery:"external_id"`
	CmsID              bigquery.NullString   `bigquery:"cms_id"`
	Duration           bigquery.NullInt64    `bigquery:"duration"`
	StartAt            bigquery.NullDateTime `bigquery:"start_at"`
	NotForSale         bigquery.NullBool     `bigquery:"not_for_sale"`
	GraduationCriteria bigquery.NullString   `bigquery:"graduation_criteria"`
}

func (c *bqCourse) parseToCourse() Course {
	course := Course{}
	course.ID = c.ID
	if c.Name.Valid {
		course.Name = &c.Name.StringVal
	}
	if c.ShortName.Valid {
		course.ShortName = &c.ShortName.StringVal
	}
	if c.Type.Valid {
		course.Type = &c.Type.StringVal
	}
	if c.Status.Valid {
		course.Status = &c.Status.StringVal
	}
	if c.Active.Valid {
		course.Active = &c.Active.Bool
	}
	if c.CmsID.Valid {
		course.CmsID = &c.CmsID.StringVal
	}
	if c.Duration.Valid {
		course.Duration = &c.Duration.Int64
	}
	if c.GraduationCriteria.Valid {
		course.GraduationCriteria = &c.GraduationCriteria.StringVal
	}
	return course
}

type bqUserWithActiveCourses struct {
	UserID     string `bigquery:"user_id"`
	Email      string `bigquery:"email"`
	CoursesIds string `bigquery:"courses"`
}

func (u *bqUserWithActiveCourses) parseToUserWithActiveCourses() UserWithActiveCourses {
	return UserWithActiveCourses{
		UserID:     u.UserID,
		Email:      u.Email,
		CoursesIds: strings.Split(u.CoursesIds, ","),
	}
}
