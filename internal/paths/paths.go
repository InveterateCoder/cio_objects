package paths

import "fmt"

type Region string

var (
	BR Region = "BR"
	MX Region = "MX"
	CO Region = "CO"
	PE Region = "PE"
)

type Paths struct {
	Region         Region
	CustomerioPath string
}

func NewPaths(region Region) *Paths {
	if region != BR && region != MX && region != CO && region != PE {
		panic("wrong region")
	}
	buildPath := func(name string) string {
		return fmt.Sprintf("assets/%s/%s.json", region, name)
	}
	pths := Paths{
		Region:         region,
		CustomerioPath: buildPath("customerio"),
	}
	return &pths
}
