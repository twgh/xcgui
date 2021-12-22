package xc

import "unsafe"

// 字体_创建, 创建炫彩字体, 当字体句柄与元素关联后, 会自动释放, 返回字体句柄.
//
// size: 字体大小,单位(pt,磅).
func XFont_Create(size int) int {
	r, _, _ := xFont_Create.Call(uintptr(size))
	return int(r)
}

// 字体_创建扩展, 创建炫彩字体, 返回字体句柄.
//
// pName: 字体名称.
//
// size: 字体大小, 单位(pt,磅).
//
// style: 字体样式, FontStyle_.
func XFont_CreateEx(pName string, size int, style int) int {
	r, _, _ := xFont_CreateEx.Call(strPtr(pName), uintptr(size), uintptr(style))
	return int(r)
}

// 字体_创建从LOGFONT, 创建炫彩字体, 返回字体句柄.
//
// pFontInfo: 字体信息.
func XFont_CreateLOGFONTW(pFontInfo *LOGFONTW) int {
	r, _, _ := xFont_CreateLOGFONTW.Call(uintptr(unsafe.Pointer(pFontInfo)))
	return int(r)
}

// 字体_创建从HFONT, 创建炫彩字体从现有HFONT字体, 返回字体句柄.
//
// hFont: 字体句柄.
func XFont_CreateFromHFONT(hFont int) int {
	r, _, _ := xFont_CreateFromHFONT.Call(uintptr(hFont))
	return int(r)
}

// 字体_创建从Font, 创建炫彩字体从GDI+字体(Font), 返回字体句柄.
//
// pFont: GDI+字体指针(Font*).
func XFont_CreateFromFont(pFont int) int {
	r, _, _ := xFont_CreateFromFont.Call(uintptr(pFont))
	return int(r)
}

// 字体_创建从文件, 创建字体从文件, 返回炫彩字体句柄.
//
// pFontFile: 字体文件名.
//
// size: 字体大小, 单位(pt,磅).
//
// style: 字体样式, FontStyle_.
func XFont_CreateFromFile(pFontFile string, size int, style int) int {
	r, _, _ := xFont_CreateFromFile.Call(strPtr(pFontFile), uintptr(size), uintptr(style))
	return int(r)
}

// 字体_创建从内存, 创建炫彩字体从内存, 返回字体句柄.
//
// data: 字体文件数据.
//
// fontSize: 字体大小, 单位(pt,磅).
//
// style: 字体样式, FontStyle_.
func XFont_CreateFromMem(data []byte, fontSize, style int) int {
	r, _, _ := xFont_CreateFromMem.Call(bytePtr2(&data), uintptr(len(data)), uintptr(fontSize), uintptr(style))
	return int(r)
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
func XFont_CreateFromRes(id int, pType string, fontSize, style, hModule int) int {
	r, _, _ := xFont_CreateFromRes.Call(uintptr(id), strPtr(pType), uintptr(fontSize), uintptr(style), uintptr(hModule))
	return int(r)
}

// 字体_启用自动销毁, 是否自动销毁.
//
// hFontX: 字体句柄.
//
// bEnable: 是否启用.
func XFont_EnableAutoDestroy(hFontX int, bEnable bool) int {
	r, _, _ := xFont_EnableAutoDestroy.Call(uintptr(hFontX), boolPtr(bEnable))
	return int(r)
}

// 字体_取Font, 获取字体, 返回GDI+ Font指针.
//
// hFontX: 字体句柄.
func XFont_GetFont(hFontX int) int {
	r, _, _ := xFont_GetFont.Call(uintptr(hFontX))
	return int(r)
}

// 字体_取信息, 获取字体信息.
//
// hFontX: 字体句柄.
//
// pInfo: 接收返回的字体信息.
func XFont_GetFontInfo(hFontX int, pInfo *Font_Info_) int {
	r, _, _ := xFont_GetFontInfo.Call(uintptr(hFontX), uintptr(unsafe.Pointer(pInfo)))
	return int(r)
}

// 字体_取LOGFONTW, 获取字体LOGFONTW.
//
// hFontX: 字体句柄.
//
// hdc: hdc句柄.
//
// pOut: 接收返回信息.
func XFont_GetLOGFONTW(hFontX int, hdc int, pOut *LOGFONTW) bool {
	r, _, _ := xFont_GetLOGFONTW.Call(uintptr(hFontX), uintptr(hdc), uintptr(unsafe.Pointer(pOut)))
	return int(r) != 0
}

// 字体_销毁, 强制销毁炫彩字体, 谨慎使用, 建议使用 XFont_Release() 释放.
//
// hFontX: 字体句柄.
func XFont_Destroy(hFontX int) int {
	r, _, _ := xFont_Destroy.Call(uintptr(hFontX))
	return int(r)
}

// 字体_增加引用计数.
//
// hFontX: 字体句柄.
func XFont_AddRef(hFontX int) int {
	r, _, _ := xFont_AddRef.Call(uintptr(hFontX))
	return int(r)
}

// 字体_取引用计数.
//
// hFontX: 字体句柄.
func XFont_GetRefCount(hFontX int) int {
	r, _, _ := xFont_GetRefCount.Call(uintptr(hFontX))
	return int(r)
}

// 字体_释放引用计数, 释放引用计数, 当引用计数为0时自动销毁.
//
// hFontX: 字体句柄.
func XFont_Release(hFontX int) int {
	r, _, _ := xFont_Release.Call(uintptr(hFontX))
	return int(r)
}
