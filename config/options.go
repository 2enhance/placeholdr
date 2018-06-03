package config

import (
	"github.com/2enhance/placeholdr/config/providers"
)

type Options struct {
	Width           int
	Height          int
	BackgroundColor *providers.Color
	Gender          string
	SkinColor       *providers.Color
	HairColor       *providers.Color
	ShirtColor      *providers.Color
	TextColor       *providers.Color
	BrandName       string
}

func NewPlaceholderOptions(dimensions string, id string, isGrey bool) *Options {
	width, height := providers.Dimensions(dimensions)

	backgroundColor := providers.BackgroundColor(providers.Seed(id))

	if isGrey {
		backgroundColor = providers.GreyColor()
	}

	return &Options{
		Width:           width,
		Height:          height,
		BackgroundColor: backgroundColor,
		TextColor:       providers.TextColor(backgroundColor),
	}
}

func NewLogoOptions(dimensions string, id string) *Options {
	width, height := providers.Dimensions(dimensions)
	brandName := providers.BrandName(id)
	backgroundColor := providers.BackgroundColor(providers.Seed(brandName))

	return &Options{
		Width:           width,
		Height:          height,
		BackgroundColor: backgroundColor,
		TextColor:       providers.TextColor(backgroundColor),
		BrandName:       brandName,
	}
}

func NewAvatarOptions(dimensions string, id string) *Options {
	seed := providers.Seed(id)
	width, height := providers.Dimensions(dimensions)
	backgroundColor := providers.BackgroundColor(seed)

	return &Options{
		Width:           width,
		Height:          height,
		BackgroundColor: backgroundColor,
		Gender:          providers.Gender(seed),
		ShirtColor:      providers.ShirtColor(backgroundColor),
		SkinColor:       providers.SkinColor(seed),
		HairColor:       providers.HairColor(seed),
	}
}
