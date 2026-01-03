package app

import (
	"github.com/twgh/xcgui/imagex"
	"github.com/twgh/xcgui/svg"
)

// Deprecated
//
// !这是旧版函数, 请使用 NewSvgByString
func NewSvgByStringUtf8(str string) *svg.Svg {
	return NewSvgByString(str)
}

// Deprecated
//
// !这是旧版函数, 请使用 NewImageBySvgString
func NewImageBySvgStringUtf8(str string) *imagex.Image {
	return NewImageBySvgString(str)
}

// Deprecated
//
// !这是旧版函数, 请使用 NewSvgByString
func NewSvgByStringW(str string) *svg.Svg {
	return NewSvgByString(str)
}

// Deprecated
//
// !这是旧版函数, 请使用 NewImageBySvgString
func NewImageBySvgStringW(str string) *imagex.Image {
	return NewImageBySvgString(str)
}
