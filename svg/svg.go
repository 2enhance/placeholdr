package svg

import (
	"fmt"
	"net/http"

	"github.com/2enhance/placeholdr/config"
	svgo "github.com/ajstarks/svgo"
)

func NewPlaceholder(w http.ResponseWriter, options *config.Options) {
	canvas := svgo.New(w)
	canvas.Start(options.Width, options.Height)

	background(canvas, options)
	text(canvas, options, options.Width/2, options.Height/2, fmt.Sprintf("%d x %d", options.Width, options.Height))

	canvas.End()
}

func NewLogo(w http.ResponseWriter, options *config.Options) {
	canvas := svgo.New(w)
	canvas.Start(options.Width, options.Height)

	unit := extractUnit(options)

	background(canvas, options)
	logo(canvas, options, unit)
	text(canvas, options, options.Width/2, options.Height/2+unit, options.BrandName)

	canvas.End()
}

func NewAvatar(w http.ResponseWriter, options *config.Options) {
	canvas := svgo.New(w)
	canvas.Startview(options.Width, options.Height, 0, 0, 120, 120)

	background(canvas, options)

	canvas.Gtransform("translate(26.000000, 5.000000)")

	if options.Gender == "female" {
		femaleHairBottom(canvas, options.HairColor)
		body(canvas, options.SkinColor)
		shirt(canvas, options.ShirtColor)
		femaleHairTop(canvas, options.HairColor)
	} else {
		body(canvas, options.SkinColor)
		shirt(canvas, options.ShirtColor)
		maleHair(canvas, options.HairColor)
	}

	canvas.Gend()
	canvas.End()
}

func extractUnit(options *config.Options) int {
	return (options.Width + options.Height) / 10
}

func background(canvas *svgo.SVG, options *config.Options) {
	styletyle := fmt.Sprintf("fill:%s", options.BackgroundColor.Format())

	canvas.Rect(0, 0, options.Width, options.Height, styletyle)
}

func text(canvas *svgo.SVG, options *config.Options, x int, y int, text string) {
	fontSize := extractUnit(options) / 2
	style := fmt.Sprintf("fill:%s;font-size:%dpt;text-anchor:middle;font-family:sans-serif", options.TextColor.Format(), fontSize)

	canvas.Text(x, y, text, style)
}
