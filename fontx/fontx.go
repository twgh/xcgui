// 字体
package fontx

import (
	"github.com/twgh/xcgui/xc"
)

// 字体
type FontX struct {
	HFontX int //字体句柄
}

// 字体_创建, 创建炫彩字体, 当字体句柄与元素关联后, 会自动释放
// size: 字体大小
func NewFontX(size int) *FontX {
	p := &FontX{
		HFontX: xc.XFont_Create(size),
	}
	return p
}

// 字体_创建2, 创建炫彩字体
// pName: 字体名称.
// size: 字体大小
// style: 字体样式, FontStyle_
func NewFontX2(pName string, size int, style int) *FontX {
	p := &FontX{
		HFontX: xc.XFont_Create2(pName, size, style),
	}
	return p
}

// 字体_创建3, 创建炫彩字体
// pInfo: 字体信息.
func NewFontX3(pInfo *xc.Font_Info_) *FontX {
	p := &FontX{
		HFontX: xc.XFont_Create3(pInfo),
	}
	return p
}

// 字体_创建扩展, 创建炫彩字体
// pFontInfo: 字体信息.
func NewFontXEx(pFontInfo *xc.LOGFONTW) *FontX {
	p := &FontX{
		HFontX: xc.XFont_CreateEx(pFontInfo),
	}
	return p
}

// 字体_创建从HFONT, 创建炫彩字体从现有HFONT字体
// hFont: 字体句柄.
func NewFontXFromHFONT(hFont int) *FontX {
	p := &FontX{
		HFontX: xc.XFont_CreateFromHFONT(hFont),
	}
	return p
}

// 字体_创建从Font, 创建炫彩字体从GDI+字体(Font)
// pFont: GDI+字体指针(Font*).
func NewFontXFromFont(pFont int) *FontX {
	p := &FontX{
		HFontX: xc.XFont_CreateFromFont(pFont),
	}
	return p
}

// 字体_创建从文件, 创建字体从文件
// pFontFile: 字体文件名.
// size: 字体大小.
// style: 样式, FontStyle_
func NewFontXFromFile(pFontFile string, size int, style int) *FontX {
	p := &FontX{
		HFontX: xc.XFont_CreateFromFile(pFontFile, size, style),
	}
	return p
}

// 给本类的HFontX赋值
func (f *FontX) SetHFontX(hFontX int) {
	f.HFontX = hFontX
}

// 字体_启用自动销毁, 是否自动销毁
// bEnable: 是否启用.
func (f *FontX) EnableAutoDestroy(bEnable bool) int {
	return xc.XFont_EnableAutoDestroy(f.HFontX, bEnable)
}

// 字体_取Font, 获取字体, 返回GDI+ Font指针.
func (f *FontX) GetFont() int {
	return xc.XFont_GetFont(f.HFontX)
}

// 字体_取信息, 获取字体信息.
// pInfo: 接收返回的字体信息.
func (f *FontX) GetFontInfo(pInfo *xc.Font_Info_) int {
	return xc.XFont_GetFontInfo(f.HFontX, pInfo)
}

// 字体_取LOGFONTW, 获取字体LOGFONTW
// hdc: hdc句柄
// pOut: 接收返回信息
func (f *FontX) GetLOGFONTW(hdc int, pOut *xc.LOGFONTW) bool {
	return xc.XFont_GetLOGFONTW(f.HFontX, hdc, pOut)
}

// 字体_销毁, 强制销毁炫彩字体, 谨慎使用, 建议使用 XFont_Release() 释放
func (f *FontX) Destroy() int {
	return xc.XFont_Destroy(f.HFontX)
}

// 字体_增加引用计数
func (f *FontX) AddRef() int {
	return xc.XFont_AddRef(f.HFontX)
}

// 字体_取引用计数
func (f *FontX) GetRefCount() int {
	return xc.XFont_GetRefCount(f.HFontX)
}

// 字体_释放引用计数, 释放引用计数, 当引用计数为0时自动销毁.
func (f *FontX) Release() int {
	return xc.XFont_Release(f.HFontX)
}
