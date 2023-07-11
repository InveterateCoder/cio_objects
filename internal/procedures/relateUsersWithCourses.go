package procedures

import (
	"cio_objects/internal/cio"
	"cio_objects/internal/courses"
	"cio_objects/internal/helpers"
	"cio_objects/internal/paths"
	"fmt"
	"math"
	"strings"
)

func RelateUsersWithCourses(region paths.Region, dumpNew bool) {
	pths := paths.NewPaths(region)
	if dumpNew || !helpers.FileExists(pths.CustomerioPath) {
		dump_customerio(pths)
	}
	customerio := helpers.ReadJson[[]helpers.Customer](pths.CustomerioPath)
	cioCourses := courses.NewCioCourses(region)
	users := cioCourses.GetUsersWithActiveCourses()
	cioCourses.Close()
	cio := cio.NewCIO(region)
	var totalProcessed float64 = 0
	state := cio.RelateUsersWithCourseObjects(users, customerio)
	totalItems := float64(state.TotalItems)
	for status := range state.Progress {
		totalProcessed += float64(status.ChunkLen)
		fmt.Printf("\rProcessed: %.0f - %.0f%%, ret: %s", totalProcessed, math.Floor(totalProcessed/totalItems*100), strings.Trim(string(status.CioRet), "\n"))
	}
	fmt.Println()
}
