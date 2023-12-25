package xc

import (
	"github.com/twgh/xcgui/common"
	"unsafe"

	"github.com/twgh/xcgui/xcc"
)

// XFont_Create 字体_创建, 创建炫彩字体. 当字体句柄与元素关联后, 会自动释放.
//
//	@param size 字体大小,单位(pt,磅).
//	@return int 返回字体句柄.
func XFont_Create(size int32) int {
	r, _, _ := xFont_Create.Call(uintptr(size))
	return int(r)
}

// XFont_CreateEx 字体_创建扩展. 创建炫彩字体.
//
//	@param pName 字体名称.
//	@param size 字体大小, 单位(pt,磅).
//	@param style 字体样式, xcc.FontStyle_.
//	@return int 返回字体句柄.
func XFont_CreateEx(pName string, size int32, style xcc.FontStyle_) int {
	r, _, _ := xFont_CreateEx.Call(common.StrPtr(pName), uintptr(size), uintptr(style))
	return int(r)
}

// XFont_CreateLOGFONTW 字体_创建从LOGFONT. 创建炫彩字体.
//
//	@param pFontInfo 字体信息.
//	@return int 返回字体句柄.
func XFont_CreateLOGFONTW(pFontInfo *LOGFONTW) int {
	r, _, _ := xFont_CreateLOGFONTW.Call(uintptr(unsafe.Pointer(pFontInfo)))
	return int(r)
}

// XFont_CreateFromHFONT 字体_创建从HFONT. 创建炫彩字体从现有HFONT字体.
//
//	@param hFont 字体句柄.
//	@return int 返回字体句柄.
func XFont_CreateFromHFONT(hFont uintptr) int {
	r, _, _ := xFont_CreateFromHFONT.Call(hFont)
	return int(r)
}

// XFont_CreateFromFont 字体_创建从Font. 创建炫彩字体从GDI+字体.
//
//	@param pFont GDI+字体指针.
//	@return int 返回字体句柄.
func XFont_CreateFromFont(pFont uintptr) int {
	r, _, _ := xFont_CreateFromFont.Call(pFont)
	return int(r)
}

// XFont_CreateFromFile 字体_创建从文件. 创建字体从文件.
//
//	@param pFontFile 字体文件名.
//	@param size 字体大小, 单位(pt,磅).
//	@param style 字体样式, xcc.FontStyle_.
//	@return int 返回炫彩字体句柄.
func XFont_CreateFromFile(pFontFile string, size int32, style xcc.FontStyle_) int {
	r, _, _ := xFont_CreateFromFile.Call(common.StrPtr(pFontFile), uintptr(size), uintptr(style))
	return int(r)
}

// XFont_CreateFromZip 字体_创建从ZIP.
//
//	@param pZipFileName zip文件名.
//	@param pFileName 字体文件名.
//	@param pPassword zip密码.
//	@param fontSize 字体大小, 单位(pt,磅).
//	@param style 字体样式: xcc.FontStyle_.
//	@return int 返回炫彩字体句柄.
func XFont_CreateFromZip(pZipFileName, pFileName, pPassword string, fontSize int32, style xcc.FontStyle_) int {
	r, _, _ := xFont_CreateFromZip.Call(common.StrPtr(pZipFileName), common.StrPtr(pFileName), common.StrPtr(pPassword), uintptr(fontSize), uintptr(style))
	return int(r)
}

// XFont_CreateFromZipMem 字体_创建从内存ZIP.
//
//	@param data zip数据.
//	@param pFileName 字体文件名.
//	@param pPassword zip密码.
//	@param fontSize 字体大小, 单位(pt,磅).
//	@param style 字体样式: xcc.FontStyle_.
//	@return int 返回炫彩字体句柄.
func XFont_CreateFromZipMem(data []byte, pFileName, pPassword string, fontSize int32, style xcc.FontStyle_) int {
	r, _, _ := xFont_CreateFromZipMem.Call(common.ByteSliceDataPtr(&data), common.StrPtr(pFileName), common.StrPtr(pPassword), uintptr(fontSize), uintptr(style))
	return int(r)
}

// XFont_CreateFromMem 字体_创建从内存. 创建炫彩字体从内存.
//
//	@param data 字体文件数据.
//	@param fontSize 字体大小, 单位(pt,磅).
//	@param style 字体样式, xcc.FontStyle_.
//	@return int 返回字体句柄.
func XFont_CreateFromMem(data []byte, fontSize int32, style xcc.FontStyle_) int {
	r, _, _ := xFont_CreateFromMem.Call(common.ByteSliceDataPtr(&data), uintptr(len(data)), uintptr(fontSize), uintptr(style))
	return int(r)
}

// XFont_CreateFromRes 字体_创建从资源. 创建字体从资源.
//
//	@param id xx.
//	@param pType xx.
//	@param fontSize 字体大小, 单位(pt,磅).
//	@param style 字体样式, xcc.FontStyle_.
//	@param hModule xx.
//	@return int 返回炫彩字体句柄.
func XFont_CreateFromRes(id int32, pType string, fontSize int32, style xcc.FontStyle_, hModule uintptr) int {
	r, _, _ := xFont_CreateFromRes.Call(uintptr(id), common.StrPtr(pType), uintptr(fontSize), uintptr(style), hModule)
	return int(r)
}

// XFont_EnableAutoDestroy 字体_启用自动销毁. 是否自动销毁.
//
//	@param hFontX 字体句柄.
//	@param bEnable 是否启用.
//	@return int
func XFont_EnableAutoDestroy(hFontX int, bEnable bool) int {
	r, _, _ := xFont_EnableAutoDestroy.Call(uintptr(hFontX), common.BoolPtr(bEnable))
	return int(r)
}

// XFont_GetFont 字体_取Font. 获取字体.
//
//	@param hFontX 字体句柄.
//	@return int 返回GDI+ Font指针.
func XFont_GetFont(hFontX int) int {
	r, _, _ := xFont_GetFont.Call(uintptr(hFontX))
	return int(r)
}

// XFont_GetFontInfo 字体_取信息. 获取字体信息.
//
//	@param hFontX 字体句柄.
//	@param pInfo 接收返回的字体信息.
//	@return int
func XFont_GetFontInfo(hFontX int, pInfo *Font_Info_) int {
	r, _, _ := xFont_GetFontInfo.Call(uintptr(hFontX), uintptr(unsafe.Pointer(pInfo)))
	return int(r)
}

// XFont_GetLOGFONTW 字体_取LOGFONTW. 获取字体LOGFONTW.
//
//	@param hFontX 字体句柄.
//	@param hdc hdc句柄.
//	@param pOut 接收返回信息.
//	@return bool
func XFont_GetLOGFONTW(hFontX int, hdc uintptr, pOut *LOGFONTW) bool {
	r, _, _ := xFont_GetLOGFONTW.Call(uintptr(hFontX), hdc, uintptr(unsafe.Pointer(pOut)))
	return r != 0
}

// XFont_Destroy 字体_销毁. 强制销毁炫彩字体, 谨慎使用, 建议使用 XFont_Release() 释放.
//
//	@param hFontX 字体句柄.
//	@return int
func XFont_Destroy(hFontX int) {
	xFont_Destroy.Call(uintptr(hFontX))
}

// XFont_AddRef 字体_增加引用计数.
//
//	@param hFontX 字体句柄.
//	@return int
func XFont_AddRef(hFontX int) {
	xFont_AddRef.Call(uintptr(hFontX))
}

// XFont_GetRefCount 字体_取引用计数.
//
//	@param hFontX 字体句柄.
//	@return int
func XFont_GetRefCount(hFontX int) int32 {
	r, _, _ := xFont_GetRefCount.Call(uintptr(hFontX))
	return int32(r)
}

// XFont_Release 字体_释放引用计数. 释放引用计数, 当引用计数为0时自动销毁.
//
//	@param hFontX 字体句柄.
//	@return int
func XFont_Release(hFontX int) {
	xFont_Release.Call(uintptr(hFontX))
}
