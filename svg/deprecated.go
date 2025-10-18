package svg

import (
	"github.com/twgh/xcgui/xc"
)

// Deprecated
//
// !这是旧版函数, 请使用 NewByString
func NewByStringUtf8(pString string) *Svg {
	p := &Svg{}
	p.SetHandle(xc.XSvg_LoadString(pString))
	return p
}
