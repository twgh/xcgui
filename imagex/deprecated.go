package imagex

import "github.com/twgh/xcgui/xc"

// Deprecated
//
// !这是旧版函数, 请使用 NewBySvgString
func NewBySvgStringUtf8(pString string) *Image {
	p := &Image{}
	p.SetHandle(xc.XImage_LoadSvgString(pString))
	return p
}
