package procedures

import (
	"cio_objects/internal/cio"
	"cio_objects/internal/courses"
	"cio_objects/internal/paths"
	"fmt"
)

func DumpCoursesToCio(region paths.Region) {
	cioCourses := courses.NewCioCourses(region)
	defer cioCourses.Close()
	list := cioCourses.GetAllCourses()
	resp := cio.NewCIO(region).CreateOrUpdateClasses(list)
	fmt.Println(string(resp))
}
