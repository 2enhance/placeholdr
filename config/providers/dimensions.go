package providers

import (
	"strconv"
	"strings"
)

var defaultWidth = 450
var defaultHeight = 450

func Dimensions(dim string) (int, int) {
	if dim == "" {
		return defaultWidth, defaultHeight
	}

	dimSlice := strings.Split(strings.ToLower(dim), "x")

	if len(dimSlice) < 2 {
		return defaultWidth, defaultHeight
	}

	width, _ := strconv.ParseInt(dimSlice[0], 10, 64)
	height, _ := strconv.ParseInt(dimSlice[1], 10, 64)

	if width == 0 || height == 0 {
		return defaultWidth, defaultHeight
	}

	return int(width), int(height)
}
