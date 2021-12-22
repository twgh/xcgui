// 字体.
package font

import (
	"github.com/twgh/xcgui/objectbase"
	"github.com/twgh/xcgui/res"
	"github.com/twgh/xcgui/xc"
)

// 字体.
type Font struct {
	objectbase.ObjectBase
}

// 字体_创建, 创建炫彩字体, 当字体句柄与元素关联后, 会自动释放.
//
// size: 字体大小.
func NewFont(size int) *Font {
	p := &Font{}
	p.SetHandle(xc.XFont_Create(size))
	return p
}

// 字体_创建扩展, 创建炫彩字体, 返回字体句柄.
//
// pName: 字体名称.
//
// size: 字体大小, 单位(pt,磅).
//
// style: 字体样式, FontStyle_.
func NewFontEX(pName string, size int, style int) *Font {
	p := &Font{}
	p.SetHandle(xc.XFont_CreateEx(pName, size, style))
	return p
}

// 字体_创建从LOGFONT, 创建炫彩字体.
//
// pFontInfo: 字体信息.
func NewFontLOGFONTW(pFontInfo *xc.LOGFONTW) *Font {
	p := &Font{}
	p.SetHandle(xc.XFont_CreateLOGFONTW(pFontInfo))
	return p
}

// 字体_创建从HFONT, 创建炫彩字体从现有HFONT字体.
//
// hFont: 字体句柄.
func NewFontFromHFONT(hFont int) *Font {
	p := &Font{}
	p.SetHandle(xc.XFont_CreateFromHFONT(hFont))
	return p
}

// 字体_创建从Font, 创建炫彩字体从GDI+字体(Font).
//
// pFont: GDI+字体指针(Font*).
func NewFontFromFont(pFont int) *Font {
	p := &Font{}
	p.SetHandle(xc.XFont_CreateFromFont(pFont))
	return p
}

// 字体_创建从文件, 创建字体从文件.
//
// pFontFile: 字体文件名.
//
// size: 字体大小, 单位(pt,磅).
//
// style: 字体样式, FontStyle_.
func NewFontFromFile(pFontFile string, size int, style int) *Font {
	p := &Font{}
	p.SetHandle(xc.XFont_CreateFromFile(pFontFile, size, style))
	return p
}

// 字体_创建从内存, 创建炫彩字体从内存, 返回字体句柄.
//
// data: 字体文件数据.
//
// fontSize: 字体大小, 单位(pt,磅).
//
// style: 字体样式, FontStyle_.
func NewFontFromMem(data []byte, fontSize, style int) *Font {
	p := &Font{}
	p.SetHandle(xc.XFont_CreateFromMem(data, fontSize, style))
	return p
}

// 字体_创建从资源, 创建字体从资源, 返回炫彩字体句柄.
//
// id: xx.
//
// pType: xx.
//
// fontSize: 字体大小, 单位(pt,磅).
//
// style: 字体样式, FontStyle_.
//
// hModule: xx.
func NewFontFromRes(id int, pType string, fontSize int, style int, hModule int) *Font {
	p := &Font{}
	p.SetHandle(xc.XFont_CreateFromRes(id, pType, fontSize, style, hModule))
	return p
}

// 从句柄创建对象.
func NewFontByHandle(handle int) *Font {
	p := &Font{}
	p.SetHandle(handle)
	return p
}

// 根据资源文件中的name创建对象, 失败返回nil.
func NewFontByName(name string) *Font {
	handle := res.GetFont(name)
	if handle > 0 {
		p := &Font{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 字体_启用自动销毁, 是否自动销毁.
//
// bEnable: 是否启用.
func (f *Font) EnableAutoDestroy(bEnable bool) int {
	return xc.XFont_EnableAutoDestroy(f.Handle, bEnable)
}

// 字体_取Font, 获取字体, 返回GDI+ Font指针.
func (f *Font) GetFont() int {
	return xc.XFont_GetFont(f.Handle)
}

// 字体_取信息, 获取字体信息.
//
// pInfo: 接收返回的字体信息.
func (f *Font) GetFontInfo(pInfo *xc.Font_Info_) int {
	return xc.XFont_GetFontInfo(f.Handle, pInfo)
}

// 字体_取LOGFONTW, 获取字体LOGFONTW.
//
// hdc: hdc句柄.
//
// pOut: 接收返回信息.
func (f *Font) GetLOGFONTW(hdc int, pOut *xc.LOGFONTW) bool {
	return xc.XFont_GetLOGFONTW(f.Handle, hdc, pOut)
}

// 字体_销毁, 强制销毁炫彩字体, 谨慎使用, 建议使用 XFont_Release() 释放.
func (f *Font) Destroy() int {
	return xc.XFont_Destroy(f.Handle)
}

// 字体_增加引用计数.
func (f *Font) AddRef() int {
	return xc.XFont_AddRef(f.Handle)
}

// 字体_取引用计数.
func (f *Font) GetRefCount() int {
	return xc.XFont_GetRefCount(f.Handle)
}

// 字体_释放引用计数, 释放引用计数, 当引用计数为0时自动销毁.
func (f *Font) Release() int {
	return xc.XFont_Release(f.Handle)
}
