package xc

import "unsafe"

// SVG_加载从文件, 返回SVG句柄.
//
// pFileName: 文件名.
func XSvg_LoadFile(pFileName string) int {
	r, _, _ := xSvg_LoadFile.Call(strPtr(pFileName))
	return int(r)
}

// SVG_加载从字符串, 返回SVG句柄.
//
// pString: 字符串指针.
//
// nLength: 字符串长度.
func XSvg_LoadString(pString string, nLength int) int {
	r, _, _ := xSvg_LoadString.Call(XC_wtoa(pString), uintptr(nLength))
	return int(r)
}

// SVG_加载从ZIP, 返回SVG句柄.
//
// pZipFileName: zip文件名.
//
// pFileName: svg文件名.
//
// pPassword: zip密码.
func XSvg_LoadZip(pZipFileName, pFileName, pPassword string) int {
	r, _, _ := xSvg_LoadZip.Call(strPtr(pZipFileName), strPtr(pFileName), strPtr(pPassword))
	return int(r)
}

// SVG_加载从资源, 返回SVG句柄.
//
// id: 资源ID.
//
// pType: 资源类型.在rc资源文件中.
//
// hModule: 从指定模块加载.
func XSvg_LoadRes(id int, pType string, hModule int) int {
	r, _, _ := xSvg_LoadRes.Call(uintptr(id), strPtr(pType), uintptr(hModule))
	return int(r)
}

// SVG_置大小.
//
// hSvg: SVG句柄.
//
// nWidth: 宽度.
//
// nHeight: 高度.
func XSvg_SetSize(hSvg int, nWidth, nHeight int) int {
	r, _, _ := xSvg_SetSize.Call(uintptr(hSvg), uintptr(nWidth), uintptr(nHeight))
	return int(r)
}

// SVG_取大小.
//
// hSvg: SVG句柄.
//
// pWidth: 接收返回宽度.
//
// pHeight: 接收返回高度.
func XSvg_GetSize(hSvg int, pWidth, pHeight *int) int {
	r, _, _ := xSvg_GetSize.Call(uintptr(hSvg), uintptr(unsafe.Pointer(pWidth)), uintptr(unsafe.Pointer(pHeight)))
	return int(r)
}

// SVG_取宽度.
//
// hSvg: SVG句柄.
func XSvg_GetWidth(hSvg int) int {
	r, _, _ := xSvg_GetWidth.Call(uintptr(hSvg))
	return int(r)
}

// SVG_取高度.
//
// hSvg: SVG句柄.
func XSvg_GetHeight(hSvg int) int {
	r, _, _ := xSvg_GetHeight.Call(uintptr(hSvg))
	return int(r)
}

// SVG_置偏移.
//
// hSvg: SVG句柄.
//
// x: x轴偏移.
//
// y: y轴偏移.
func XSvg_SetOffset(hSvg int, x int, y int) int {
	r, _, _ := xSvg_SetOffset.Call(uintptr(hSvg), uintptr(x), uintptr(y))
	return int(r)
}

// SVG_取偏移.
//
// hSvg: SVG句柄.
//
// pX: x轴偏移.
//
// pY: y轴偏移.
func XSvg_GetOffset(hSvg int, pX, pY *int) int {
	r, _, _ := xSvg_GetOffset.Call(uintptr(hSvg), uintptr(unsafe.Pointer(pX)), uintptr(unsafe.Pointer(pY)))
	return int(r)
}

// SVG_取视图框.
//
// hSvg: SVG句柄.
//
// pViewBox: 接收返回视图框.
func XSvg_GetViewBox(hSvg int, pViewBox *RECT) int {
	r, _, _ := xSvg_GetViewBox.Call(uintptr(hSvg), uintptr(unsafe.Pointer(pViewBox)))
	return int(r)
}

// SVG_启用自动销毁.
//
// hSvg: SVG句柄.
//
// bEnable: 是否自动销毁.
func XSvg_EnableAutoDestroy(hSvg int, bEnable bool) int {
	r, _, _ := xSvg_EnableAutoDestroy.Call(uintptr(hSvg), boolPtr(bEnable))
	return int(r)
}

// SVG_增加引用计数.
//
// hSvg: SVG句柄.
func XSvg_AddRef(hSvg int) int {
	r, _, _ := xSvg_AddRef.Call(uintptr(hSvg))
	return int(r)
}

// SVG_释放引用计数.
//
// hSvg: SVG句柄.
func XSvg_Release(hSvg int) int {
	r, _, _ := xSvg_Release.Call(uintptr(hSvg))
	return int(r)
}

// SVG_取引用计数.
//
// hSvg: SVG句柄.
func XSvg_GetRefCount(hSvg int) int {
	r, _, _ := xSvg_GetRefCount.Call(uintptr(hSvg))
	return int(r)
}

// SVG_销毁.
//
// hSvg: SVG句柄.
func XSvg_Destroy(hSvg int) int {
	r, _, _ := xSvg_Destroy.Call(uintptr(hSvg))
	return int(r)
}
