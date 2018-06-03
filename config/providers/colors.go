package providers

import (
	"fmt"
	"math/rand"
	"time"
)

type Color struct {
	Red   int
	Green int
	Blue  int
}

func (c *Color) Desaturate(sat float64) {
	gray := float64(c.Red)*0.3086 + float64(c.Green)*0.6094 + float64(c.Blue)*0.0820

	c.Red = int(float64(c.Red)*sat + gray*(1-sat))
	c.Green = int(float64(c.Green)*sat + gray*(1-sat))
	c.Blue = int(float64(c.Blue)*sat + gray*(1-sat))
}

func (c *Color) Invert() *Color {
	return &Color{255 - c.Red, 255 - c.Green, 255 - c.Blue}
}

func (c *Color) Format() string {
	return fmt.Sprintf("rgb(%d,%d,%d)", c.Red, c.Green, c.Blue)
}

func BlackColor() *Color {
	return &Color{0, 0, 0}
}

func GreyColor() *Color {
	return &Color{197, 197, 197}
}

func WhiteColor() *Color {
	return &Color{255, 255, 255}
}

func RandomColor() *Color {
	rand.Seed(time.Now().UnixNano())

	return &Color{
		rand.Intn(255),
		rand.Intn(255),
		rand.Intn(255),
	}
}

var skinColors = []*Color{
	&Color{141, 85, 36},
	&Color{198, 134, 66},
	&Color{224, 172, 105},
	&Color{241, 194, 125},
	&Color{255, 219, 172},
}

var hairColors = []*Color{
	&Color{9, 8, 6},
	&Color{44, 34, 43},
	&Color{113, 99, 90},
	&Color{183, 166, 158},
	&Color{214, 196, 194},
	&Color{202, 191, 177},
	&Color{220, 208, 186},
	&Color{255, 245, 225},
	&Color{230, 206, 168},
	&Color{229, 200, 168},
	&Color{222, 188, 153},
	&Color{184, 151, 120},
	&Color{165, 107, 70},
	&Color{181, 82, 57},
	&Color{145, 85, 61},
	&Color{83, 61, 50},
	&Color{59, 48, 36},
	&Color{85, 72, 56},
	&Color{78, 67, 63},
	&Color{80, 68, 68},
	&Color{106, 78, 66},
	&Color{167, 133, 106},
	&Color{151, 121, 97},
}

func BackgroundColor(seed int) *Color {
	color := &Color{seed % 75, seed % 150, seed % 255}

	color.Desaturate(0.5)

	return color
}

func TextColor(bg *Color) *Color {
	if (float64(bg.Red)*0.299 + float64(bg.Green)*0.587 + float64(bg.Blue)*0.114) > 186 {
		return BlackColor()
	}

	return WhiteColor()
}

func SkinColor(seed int) *Color {
	return skinColors[seed%len(skinColors)]
}

func HairColor(seed int) *Color {
	return hairColors[seed%len(hairColors)]
}

func ShirtColor(bg *Color) *Color {
	return bg.Invert()
}
