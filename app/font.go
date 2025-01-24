package app

import (
	"github.com/twgh/xcgui/font"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

// NewFont 字体_创建, 创建炫彩字体. 当字体句柄与元素关联后, 会自动释放.
//
// size: 字体大小,单位(pt, 磅).
func NewFont(size int32) *font.Font {
	return font.New(size)
}

// NewFontEX 字体_创建扩展. 创建炫彩字体.
//
// pName: 字体名称.
//
// size: 字体大小, 单位(pt, 磅).
//
// style: 字体样式, xcc.FontStyle_ .
func NewFontEX(pName string, size int32, style xcc.FontStyle_) *font.Font {
	return font.NewEX(pName, size, style)
}

// NewFontByLOGFONTW 字体_创建从LOGFONT. 创建炫彩字体.
//
// pFontInfo: 字体信息.
func NewFontByLOGFONTW(pFontInfo *xc.LOGFONTW) *font.Font {
	return font.NewByLOGFONTW(pFontInfo)
}

// NewFontByHFONT 字体_创建从HFONT. 创建炫彩字体从现有HFONT字体.
//
// hFont: 字体句柄.
func NewFontByHFONT(hFont uintptr) *font.Font {
	return font.NewByHFONT(hFont)
}

// NewFontByFont 字体_创建从Font. 创建炫彩字体从GDI+字体.
//
// pFont: GDI+ 字体指针.
func NewFontByFont(pFont uintptr) *font.Font {
	return font.NewByFont(pFont)
}

// NewFontByFile 字体_创建从文件. 创建字体从文件.
//
// pFontFile: 字体文件名.
//
// size: 字体大小, 单位(pt, 磅).
//
// style: 字体样式, xcc.FontStyle_ .
func NewFontByFile(pFontFile string, size int32, style xcc.FontStyle_) *font.Font {
	return font.NewByFile(pFontFile, size, style)
}

// NewFontByZip 字体_创建从ZIP.
//
// pZipFileName: zip文件名.
//
// pFileName: 字体文件名.
//
// pPassword: zip密码.
//
// fontSize: 字体大小, 单位(pt, 磅).
//
// style: 字体样式: xcc.FontStyle_ .
func NewFontByZip(pZipFileName, pFileName, pPassword string, fontSize int32, style xcc.FontStyle_) *font.Font {
	return font.NewByZip(pZipFileName, pFileName, pPassword, fontSize, style)
}

// NewFontByZipMem 字体_创建从内存ZIP.
//
// data: zip数据.
//
// pFileName: 字体文件名.
//
// pPassword: zip密码.
//
// fontSize: 字体大小, 单位(pt, 磅).
//
// style: 字体样式: xcc.FontStyle_ .
func NewFontByZipMem(data []byte, pFileName, pPassword string, fontSize int32, style xcc.FontStyle_) *font.Font {
	return font.NewByZipMem(data, pFileName, pPassword, fontSize, style)
}

// NewFontByMem 字体_创建从内存. 创建炫彩字体从内存.
//
// data: 字体文件数据.
//
// fontSize: 字体大小, 单位(pt, 磅).
//
// style: 字体样式, xcc.FontStyle_ .
func NewFontByMem(data []byte, fontSize int32, style xcc.FontStyle_) *font.Font {
	return font.NewByMem(data, fontSize, style)
}

// NewFontByRes 字体_创建从资源. 创建字体从资源.
//
// id: xx.
//
// pType: xx.
//
// fontSize: 字体大小, 单位(pt, 磅).
//
// style: 字体样式, xcc.FontStyle_ .
//
// hModule: xx.
func NewFontByRes(id int32, pType string, fontSize int32, style xcc.FontStyle_, hModule uintptr) *font.Font {
	return font.NewByRes(id, pType, fontSize, style, hModule)
}

// NewFontByHandle 从句柄创建对象.
//
// handle: 炫彩字体句柄.
func NewFontByHandle(handle int) *font.Font {
	return font.NewByHandle(handle)
}

// NewFontByName 根据资源文件中的name创建对象, 失败返回nil.
//
// name: name名称.
func NewFontByName(name string) *font.Font {
	return font.NewByName(name)
}
