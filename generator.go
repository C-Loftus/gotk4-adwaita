package adwaita

//go:generate go run ./cmd/gir-generate -o ./pkg/

import (
	"github.com/diamondburned/gotk4/gir/cmd/gir-generate/gendata"
	"github.com/diamondburned/gotk4/gir/cmd/gir-generate/genmain"
)

const atspiModule = "github.com/c-loftus/go-atspi/pkg"

var Data = genmain.Overlay(
	gendata.Main,
	genmain.Data{
		Module: atspiModule,
		Packages: []genmain.Package{
			{
				Name:       "atk",
				Namespaces: []string{"Atspi-2"},
			},
		},
		PkgGenerated: []string{"atspi"},
		PkgExceptions: []string{
			"go.mod",
			"go.sum",
			"LICENSE",
			"_examples",
		},
	},
)
