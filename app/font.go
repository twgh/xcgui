package app

import (
	"github.com/twgh/xcgui/font"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

// NewFont 字体_创建, 创建炫彩字体. 当字体句柄与元素关联后, 会自动释放, 失败返回 nil.
//
// size: 字体大小,单位(pt, 磅).
func NewFont(size int32) *font.Font {
	return font.New(size)
}

// NewFontEX 字体_创建扩展. 创建炫彩字体, 失败返回 nil.
//
// name: 字体名称.
//
// size: 字体大小, 单位(pt, 磅).
//
// style: 字体样式, xcc.FontStyle_ .
func NewFontEX(name string, size int32, style xcc.FontStyle_) *font.Font {
	return font.NewEX(name, size, style)
}

// NewFontByLOGFONTW 字体_创建从LOGFONT. 创建炫彩字体, 失败返回 nil.
//
// pFontInfo: 字体信息.
func NewFontByLOGFONTW(pFontInfo *xc.LOGFONTW) *font.Font {
	return font.NewByLOGFONTW(pFontInfo)
}

// NewFontByHFONT 字体_创建从HFONT. 创建炫彩字体从现有HFONT字体, 失败返回 nil.
//
// hFont: 字体句柄.
func NewFontByHFONT(hFont uintptr) *font.Font {
	return font.NewByHFONT(hFont)
}

// NewFontByFont 字体_创建从Font. 创建炫彩字体从GDI+字体, 失败返回 nil.
//
// pFont: GDI+ 字体指针.
func NewFontByFont(pFont uintptr) *font.Font {
	return font.NewByFont(pFont)
}

// NewFontByFile 字体_创建从文件. 创建字体从文件, 失败返回 nil.
//
// fontFile: 字体文件名.
//
// size: 字体大小, 单位(pt, 磅).
//
// style: 字体样式, xcc.FontStyle_ .
func NewFontByFile(fontFile string, size int32, style xcc.FontStyle_) *font.Font {
	return font.NewByFile(fontFile, size, style)
}

// NewFontByZip 字体_创建从ZIP, 失败返回 nil.
//
// zipFileName: zip文件名.
//
// fileName: 字体文件名.
//
// password: zip密码.
//
// fontSize: 字体大小, 单位(pt, 磅).
//
// style: 字体样式: xcc.FontStyle_ .
func NewFontByZip(zipFileName, fileName, password string, fontSize int32, style xcc.FontStyle_) *font.Font {
	return font.NewByZip(zipFileName, fileName, password, fontSize, style)
}

// NewFontByZipMem 字体_创建从内存ZIP, 失败返回 nil.
//
// data: zip数据.
//
// fileName: 字体文件名.
//
// password: zip密码.
//
// fontSize: 字体大小, 单位(pt, 磅).
//
// style: 字体样式: xcc.FontStyle_ .
func NewFontByZipMem(data []byte, fileName, password string, fontSize int32, style xcc.FontStyle_) *font.Font {
	return font.NewByZipMem(data, fileName, password, fontSize, style)
}

// NewFontByMem 字体_创建从内存. 创建炫彩字体从内存, 失败返回 nil.
//
// data: 字体文件数据.
//
// fontSize: 字体大小, 单位(pt, 磅).
//
// style: 字体样式, xcc.FontStyle_ .
func NewFontByMem(data []byte, fontSize int32, style xcc.FontStyle_) *font.Font {
	return font.NewByMem(data, fontSize, style)
}

// NewFontByRes 字体_创建从资源. 创建字体从资源, 失败返回 nil.
//
// id: xx.
//
// Type: xx.
//
// fontSize: 字体大小, 单位(pt, 磅).
//
// style: 字体样式, xcc.FontStyle_ .
//
// hModule: xx.
func NewFontByRes(id int32, Type string, fontSize int32, style xcc.FontStyle_, hModule uintptr) *font.Font {
	return font.NewByRes(id, Type, fontSize, style, hModule)
}

// NewFontByHandle 从句柄创建对象, 失败返回 nil.
//
// handle: 炫彩字体句柄.
func NewFontByHandle(handle int) *font.Font {
	return font.NewByHandle(handle)
}

// NewFontByName 根据资源文件中的 name 创建对象, 失败返回 nil.
//
// name: name名称.
func NewFontByName(name string) *font.Font {
	return font.NewByName(name)
}
