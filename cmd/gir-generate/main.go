package main

import (
	atspi "github.com/c-loftus/go-atspi"
	"github.com/diamondburned/gotk4/gir/cmd/gir-generate/genmain"
)

func main() {
	genmain.Run(atspi.Data)
}
