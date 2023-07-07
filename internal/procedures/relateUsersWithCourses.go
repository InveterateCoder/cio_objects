package procedures

import (
	"cio_objects/internal/cio"
	"cio_objects/internal/courses"
	"cio_objects/internal/helpers"
	"cio_objects/internal/paths"
)

func RelateUsersWithCourses(region paths.Region) {
	pths := paths.NewPaths(region)
	if !helpers.FileExists(pths.CustomerioPath) {
		dump_customerio(pths)
	}
	customerio := helpers.ReadJson[[]helpers.Customer](pths.CustomerioPath)
	cioCourses := courses.NewCioCourses(region)
	users := cioCourses.GetUsersWithActiveCourses()
	cioCourses.Close()
	cio := cio.NewCIO(region)
	cio.RelateUsersWithCourseObjects(users, customerio)
}
