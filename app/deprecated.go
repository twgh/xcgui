package app

import (
	"github.com/twgh/xcgui/imagex"
	"github.com/twgh/xcgui/svg"
)

// Deprecated
//
// !这是旧版函数, 请使用 NewSvgByString
func NewSvgByStringUtf8(pString string) *svg.Svg {
	return svg.NewByString(pString)
}

// Deprecated
//
// !这是旧版函数, 请使用 NewImageBySvgString
func NewImageBySvgStringUtf8(pString string) *imagex.Image {
	return imagex.NewBySvgString(pString)
}
