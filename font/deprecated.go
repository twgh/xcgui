package font

import (
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

// Deprecated
//
// !已废弃, 请使用 font.New().
func NewFont(size int) *Font {
	return New(size)
}

// Deprecated
//
// !已废弃, 请使用 font.NewEX().
func NewFontEX(pName string, size int, style xcc.FontStyle_) *Font {
	return NewEX(pName, size, style)
}

// Deprecated
//
// !已废弃, 请使用 font.NewLOGFONTW().
func NewFontLOGFONTW(pFontInfo *xc.LOGFONTW) *Font {
	return NewLOGFONTW(pFontInfo)
}

// Deprecated
//
// !已废弃, 请使用 font.NewByHFONT().
func NewFontFromHFONT(hFont int) *Font {
	return NewByHFONT(hFont)
}

// Deprecated
//
// !已废弃, 请使用 font.NewByFont().
func NewFontFromFont(pFont int) *Font {
	return NewByFont(pFont)
}

// Deprecated
//
// !已废弃, 请使用 font.NewByFile().
func NewFontFromFile(pFontFile string, size int, style xcc.FontStyle_) *Font {
	return NewByFile(pFontFile, size, style)
}

// Deprecated
//
// !已废弃, 请使用 font.NewByZip().
func NewFontFromZip(pZipFileName, pFileName, pPassword string, fontSize int, style xcc.FontStyle_) *Font {
	return NewByZip(pZipFileName, pFileName, pPassword, fontSize, style)
}

// Deprecated
//
// !已废弃, 请使用 font.NewByZipMem().
func NewFontFromZipMem(data []byte, pFileName, pPassword string, fontSize int, style xcc.FontStyle_) *Font {
	return NewByZipMem(data, pFileName, pPassword, fontSize, style)
}

// Deprecated
//
// !已废弃, 请使用 font.NewByMem().
func NewFontFromMem(data []byte, fontSize int, style xcc.FontStyle_) *Font {
	return NewByMem(data, fontSize, style)
}

// Deprecated
//
// !已废弃, 请使用 font.NewByRes().
func NewFontFromRes(id int, pType string, fontSize int, style xcc.FontStyle_, hModule int) *Font {
	return NewByRes(id, pType, fontSize, style, hModule)
}

// Deprecated
//
// !已废弃, 请使用 font.NewByHandle().
func NewFontByHandle(handle int) *Font {
	return NewByHandle(handle)
}

// Deprecated
//
// !已废弃, 请使用 font.NewByName().
func NewFontByName(name string) *Font {
	return NewByName(name)
}
