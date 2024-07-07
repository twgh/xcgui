package xc

import (
	"github.com/twgh/xcgui/common"
	"unsafe"

	"github.com/twgh/xcgui/xcc"
)

// XFont_Create 字体_创建, 创建炫彩字体. 当字体句柄与元素关联后, 会自动释放. 返回炫彩字体句柄.
//
// size: 字体大小,单位(pt,磅).
func XFont_Create(size int32) int {
	r, _, _ := xFont_Create.Call(uintptr(size))
	return int(r)
}

// XFont_CreateEx 字体_创建扩展. 创建炫彩字体. 返回炫彩字体句柄.
//
// pName: 字体名称.
//
// size: 字体大小, 单位(pt, 磅).
//
// style: 字体样式, xcc.FontStyle_ .
func XFont_CreateEx(pName string, size int32, style xcc.FontStyle_) int {
	r, _, _ := xFont_CreateEx.Call(common.StrPtr(pName), uintptr(size), uintptr(style))
	return int(r)
}

// XFont_CreateLOGFONTW 字体_创建从LOGFONT. 创建炫彩字体. 返回炫彩字体句柄.
//
// pFontInfo: 字体信息.
func XFont_CreateFromLOGFONTW(pFontInfo *LOGFONTW) int {
	r, _, _ := xFont_CreateFromLOGFONTW.Call(uintptr(unsafe.Pointer(pFontInfo)))
	return int(r)
}

// XFont_CreateFromHFONT 字体_创建从HFONT. 创建炫彩字体从现有HFONT字体. 返回炫彩字体句柄.
//
// hFont: 字体句柄.
func XFont_CreateFromHFONT(hFont uintptr) int {
	r, _, _ := xFont_CreateFromHFONT.Call(hFont)
	return int(r)
}

// XFont_CreateFromFont 字体_创建从Font. 创建炫彩字体从GDI+字体. 返回字体句柄.
//
// pFont: GDI+字体指针.
func XFont_CreateFromFont(pFont uintptr) int {
	r, _, _ := xFont_CreateFromFont.Call(pFont)
	return int(r)
}

// XFont_CreateFromFile 字体_创建从文件. 创建字体从文件. 返回炫彩字体句柄.
//
// pFontFile: 字体文件名.
//
// size: 字体大小, 单位(pt,磅).
//
// style: 字体样式, xcc.FontStyle_ .
func XFont_CreateFromFile(pFontFile string, size int32, style xcc.FontStyle_) int {
	r, _, _ := xFont_CreateFromFile.Call(common.StrPtr(pFontFile), uintptr(size), uintptr(style))
	return int(r)
}

// XFont_CreateFromZip 字体_创建从ZIP. 返回炫彩字体句柄.
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
func XFont_CreateFromZip(pZipFileName, pFileName, pPassword string, fontSize int32, style xcc.FontStyle_) int {
	r, _, _ := xFont_CreateFromZip.Call(common.StrPtr(pZipFileName), common.StrPtr(pFileName), common.StrPtr(pPassword), uintptr(fontSize), uintptr(style))
	return int(r)
}

// XFont_CreateFromZipMem 字体_创建从内存ZIP. 返回炫彩字体句柄.
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
func XFont_CreateFromZipMem(data []byte, pFileName, pPassword string, fontSize int32, style xcc.FontStyle_) int {
	r, _, _ := xFont_CreateFromZipMem.Call(common.ByteSliceDataPtr(&data), common.StrPtr(pFileName), common.StrPtr(pPassword), uintptr(fontSize), uintptr(style))
	return int(r)
}

// XFont_CreateFromMem 字体_创建从内存. 创建炫彩字体从内存. 返回炫彩字体句柄.
//
// data: 字体文件数据.
//
// fontSize: 字体大小, 单位(pt,磅).
//
// style: 字体样式, xcc.FontStyle_ .
func XFont_CreateFromMem(data []byte, fontSize int32, style xcc.FontStyle_) int {
	r, _, _ := xFont_CreateFromMem.Call(common.ByteSliceDataPtr(&data), uintptr(len(data)), uintptr(fontSize), uintptr(style))
	return int(r)
}

// XFont_CreateFromRes 字体_创建从资源. 创建字体从资源. 返回炫彩字体句柄.
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
func XFont_CreateFromRes(id int32, pType string, fontSize int32, style xcc.FontStyle_, hModule uintptr) int {
	r, _, _ := xFont_CreateFromRes.Call(uintptr(id), common.StrPtr(pType), uintptr(fontSize), uintptr(style), hModule)
	return int(r)
}

// XFont_EnableAutoDestroy 字体_启用自动销毁. 是否自动销毁.
//
// hFontX: 字体句柄.
//
// bEnable: 是否启用.
func XFont_EnableAutoDestroy(hFontX int, bEnable bool) {
	xFont_EnableAutoDestroy.Call(uintptr(hFontX), common.BoolPtr(bEnable))
}

// XFont_GetFont 字体_取Font. 获取字体. 返回 GDI+ Font指针.
//
// hFontX: 字体句柄.
func XFont_GetFont(hFontX int) uintptr {
	r, _, _ := xFont_GetFont.Call(uintptr(hFontX))
	return r
}

// XFont_GetFontInfo 字体_取信息. 获取字体信息.
//
// hFontX: 字体句柄.
//
// pInfo: 接收返回的字体信息.
func XFont_GetFontInfo(hFontX int, pInfo *Font_Info_) {
	xFont_GetFontInfo.Call(uintptr(hFontX), uintptr(unsafe.Pointer(pInfo)))
}

// XFont_GetLOGFONTW 字体_取LOGFONTW. 获取字体LOGFONTW.
//
// hFontX: 字体句柄.
//
// hdc: hdc句柄.
//
// pOut: 接收返回信息.
func XFont_GetLOGFONTW(hFontX int, hdc uintptr, pOut *LOGFONTW) bool {
	r, _, _ := xFont_GetLOGFONTW.Call(uintptr(hFontX), hdc, uintptr(unsafe.Pointer(pOut)))
	return r != 0
}

// XFont_Destroy 字体_销毁. 强制销毁炫彩字体, 谨慎使用, 建议使用 XFont_Release() 释放.
//
// hFontX: 字体句柄.
func XFont_Destroy(hFontX int) {
	xFont_Destroy.Call(uintptr(hFontX))
}

// XFont_AddRef 字体_增加引用计数.
//
// hFontX: 字体句柄.
func XFont_AddRef(hFontX int) {
	xFont_AddRef.Call(uintptr(hFontX))
}

// XFont_GetRefCount 字体_取引用计数.
//
// hFontX: 字体句柄.
func XFont_GetRefCount(hFontX int) int32 {
	r, _, _ := xFont_GetRefCount.Call(uintptr(hFontX))
	return int32(r)
}

// XFont_Release 字体_释放引用计数. 释放引用计数, 当引用计数为0时自动销毁.
//
// hFontX: 字体句柄.
func XFont_Release(hFontX int) {
	xFont_Release.Call(uintptr(hFontX))
}

// XFont_SetUnderlineEdit 字体_置下划线. 仅供edit字体使用, 因为edit不支持下划线字体, 所以需要单独设置.
//
// hFontX: 字体句柄.
//
// bUnderline: 是否启用下划线.
//
// bStrikeout: 是否启用删除线.
func XFont_SetUnderlineEdit(hFontX int, bUnderline, bStrikeout bool) {
	xFont_SetUnderlineEdit.Call(uintptr(hFontX), common.BoolPtr(bUnderline), common.BoolPtr(bStrikeout))
}

// XFont_GetUnderlineEdit 字体_取下划线. 仅供edit字体使用, 因为edit不支持下划线字体, 所以需要单独设置. bUnderline: 返回是否启用下划线. bStrikeout: 返回是否启用删除线.
//
// hFontX: 字体句柄.
func XFont_GetUnderlineEdit(hFontX int) (bUnderline, bStrikeout bool) {
	xFont_GetUnderlineEdit.Call(uintptr(hFontX), uintptr(unsafe.Pointer(&bUnderline)), uintptr(unsafe.Pointer(&bStrikeout)))
	return
}
