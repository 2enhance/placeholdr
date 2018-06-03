package svg

import (
	"fmt"

	"github.com/2enhance/placeholdr/config"
	svgo "github.com/ajstarks/svgo"
)

func logo(canvas *svgo.SVG, options *config.Options, unit int) *svgo.SVG {
	bg := options.BackgroundColor.Invert()
	style := fmt.Sprintf(
		"fill:%s;stroke:%s;stroke-width:6px",
		bg.Format(),
		options.TextColor.Format(),
	)

	squareOffset := unit / 3

	for i := 0; i < 3; i++ {
		offset := squareOffset * (i + 1/3)
		canvas.Rect(options.Width/2-offset, options.Height/3-offset, unit, unit, style)
	}

	return canvas
}
