package adwaita

//go:generate go run ./cmd/gir-generate -o ./pkg/

import (
	"github.com/diamondburned/gotk4/gir/cmd/gir-generate/gendata"
	"github.com/diamondburned/gotk4/gir/cmd/gir-generate/genmain"
	"github.com/diamondburned/gotk4/gir/girgen/types"
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
		Filters: []types.FilterMatcher{
			types.AbsoluteFilter("C.AtspiDeviceLegacy"),
			types.AbsoluteFilter("C.AtspiDeviceLegacyClass"),
			types.AbsoluteFilter("C.atspi_device_legacy_get_type"),
			types.AbsoluteFilter("C.atspi_device_legacy_new"),
			types.AbsoluteFilter("C.AtspiDeviceX11"),
			types.AbsoluteFilter("C.AtspiDeviceX11Class"),
			types.AbsoluteFilter("C.atspi_device_x11_get_type"),
			types.AbsoluteFilter("C.atspi_device_x11_new"),
			types.AbsoluteFilter("C.AtspiDeviceClass"),
			types.AbsoluteFilter("C.AtspiDevice"),
			types.RegexFilter(`Atk.*Matches*`),
			types.RegexFilter(`Atk.*RelationSet*`),
		},
	},
)
