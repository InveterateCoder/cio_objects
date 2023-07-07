package main

import (
	"cio_objects/internal/cio"
	"cio_objects/internal/paths"
	"fmt"
)

func main() {
	cio := cio.NewCIO(paths.BR)
	fmt.Printf("%+v\n", cio)
}
