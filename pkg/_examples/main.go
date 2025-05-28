package main

import (
	"fmt"

	atspi "github.com/c-loftus/go-atspi/pkg/atspi/v2"
)

func main() {
	t := atspi.GTypeTextGranularity
	fmt.Print(t)
	return
}
