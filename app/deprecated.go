package app

import (
	"github.com/twgh/xcgui/imagex"
	"github.com/twgh/xcgui/svg"
)

// Deprecated
//
// !这是旧版函数, 请使用 NewSvgByString
func NewSvgByStringUtf8(pString string) *svg.Svg {
	return NewSvgByString(pString)
}

// Deprecated
//
// !这是旧版函数, 请使用 NewImageBySvgString
func NewImageBySvgStringUtf8(pString string) *imagex.Image {
	return NewImageBySvgString(pString)
}

// Deprecated
//
// !这是旧版函数, 请使用 NewSvgByString
func NewSvgByStringW(pString string) *svg.Svg {
	return NewSvgByString(pString)
}

// Deprecated
//
// !这是旧版函数, 请使用 NewImageBySvgString
func NewImageBySvgStringW(pString string) *imagex.Image {
	return NewImageBySvgString(pString)
}
