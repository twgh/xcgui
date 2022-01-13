package xc

import "unsafe"

// SVG_加载从文件, 返回SVG句柄.
//
// pFileName: 文件名.
func XSvg_LoadFile(pFileName string) int {
	r, _, _ := xSvg_LoadFile.Call(StrPtr(pFileName))
	return int(r)
}

// SVG_加载从字符串, 返回SVG句柄.
//
// pString: 字符串指针.
func XSvg_LoadString(pString string) int {
	r, _, _ := xSvg_LoadString.Call(XC_wtoa(pString))
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
	r, _, _ := xSvg_LoadZip.Call(StrPtr(pZipFileName), StrPtr(pFileName), StrPtr(pPassword))
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
	r, _, _ := xSvg_LoadRes.Call(uintptr(id), StrPtr(pType), uintptr(hModule))
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
func XSvg_SetPosition(hSvg int, x, y int) int {
	r, _, _ := xSvg_SetPosition.Call(uintptr(hSvg), uintptr(x), uintptr(y))
	return int(r)
}

// SVG_取偏移.
//
// hSvg: SVG句柄.
//
// pX: x轴偏移.
//
// pY: y轴偏移.
func XSvg_GetPosition(hSvg int, pX, pY *int) int {
	r, _, _ := xSvg_GetPosition.Call(uintptr(hSvg), uintptr(unsafe.Pointer(pX)), uintptr(unsafe.Pointer(pY)))
	return int(r)
}

// SVG_置偏移F.
//
// hSvg: SVG句柄.
//
// x: x轴偏移.
//
// y: y轴偏移.
func XSvg_SetPositionF(hSvg int, x, y float32) int {
	r, _, _ := xSvg_SetPositionF.Call(uintptr(hSvg), Float32Ptr(x), Float32Ptr(y))
	return int(r)
}

// SVG_取偏移F.
//
// hSvg: SVG句柄.
//
// pX: x轴偏移.
//
// pY: y轴偏移.
func XSvg_GetPositionF(hSvg int, pX, pY *float32) int {
	r, _, _ := xSvg_GetPositionF.Call(uintptr(hSvg), uintptr(unsafe.Pointer(pX)), uintptr(unsafe.Pointer(pY)))
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
	r, _, _ := xSvg_EnableAutoDestroy.Call(uintptr(hSvg), BoolPtr(bEnable))
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

// SVG_置透明度.
//
// hSvg: SVG句柄.
//
// alpha: 透明度.
func XSvg_SetAlpha(hSvg int, alpha byte) int {
	r, _, _ := xSvg_SetAlpha.Call(uintptr(hSvg), uintptr(alpha))
	return int(r)
}

// SVG_取透明度, 返回透明度.
//
// hSvg: SVG句柄.
func XSvg_GetAlpha(hSvg int) int {
	r, _, _ := xSvg_GetAlpha.Call(uintptr(hSvg))
	return int(r)
}

// SVG_置用户填充颜色, 用户颜色将覆盖默认样式.
//
// hSvg: SVG句柄.
//
// color: 颜色, AGBR颜色.
//
// bEnable: 是否有效.
func XSvg_SetUserFillColor(hSvg int, color int, bEnable bool) int {
	r, _, _ := xSvg_SetUserFillColor.Call(uintptr(hSvg), uintptr(color), BoolPtr(bEnable))
	return int(r)
}

// SVG_置用户笔触颜色, 用户颜色将覆盖默认样式.
//
// hSvg: SVG句柄.
//
// color: 颜色, AGBR颜色.
//
// strokeWidth: 笔触宽度.
//
// bEnable: 是否有效.
func XSvg_SetUserStrokeColor(hSvg int, color int, strokeWidth float32, bEnable bool) int {
	r, _, _ := xSvg_SetUserStrokeColor.Call(uintptr(hSvg), uintptr(color), Float32Ptr(strokeWidth), BoolPtr(bEnable))
	return int(r)
}

// SVG_取用户填充颜色.
//
// hSvg: SVG句柄.
//
// pColor: 返回颜色值, AGBR颜色.
func XSvg_GetUserFillColor(hSvg int, pColor *int) bool {
	r, _, _ := xSvg_GetUserFillColor.Call(uintptr(hSvg), uintptr(unsafe.Pointer(pColor)))
	return int(r) != 0
}

// SVG_取用户笔触颜色.
//
// hSvg: SVG句柄.
//
// pColor: 返回颜色值, AGBR颜色.
//
// pStrokeWidth: .
func XSvg_GetUserStrokeColor(hSvg int, pColor *int, pStrokeWidth *float32) bool {
	r, _, _ := xSvg_GetUserStrokeColor.Call(uintptr(hSvg), uintptr(unsafe.Pointer(pColor)), uintptr(unsafe.Pointer(pStrokeWidth)))
	return int(r) != 0
}

// SVG_置旋转角度, 默认以自身中心点旋转.
//
// hSvg: SVG句柄.
//
// angle: 转角度.
func XSvg_SetRotateAngle(hSvg int, angle float32) int {
	r, _, _ := xSvg_SetRotateAngle.Call(uintptr(hSvg), Float32Ptr(angle))
	return int(r)
}

// SVG_取旋转角度, 返回旋转角度.
//
// hSvg: SVG句柄.
func XSvg_GetRotateAngle(hSvg int) float32 {
	_, r, _ := xSvg_GetRotateAngle.Call(uintptr(hSvg))
	return UintPtrToFloat32(r)
}

// SVG_置旋转.
//
// hSvg: SVG句柄.
//
// angle: 角度.
//
// x: 旋转中心点X.
//
// y: 旋转中心点Y.
//
// bOffset: TRUE: 旋转中心点相对于自身中心偏移, FALSE:使用绝对坐标.
func XSvg_SetRotate(hSvg int, angle float32, x float32, y float32, bOffset bool) int {
	r, _, _ := xSvg_SetRotate.Call(uintptr(hSvg), Float32Ptr(angle), Float32Ptr(x), Float32Ptr(y), BoolPtr(bOffset))
	return int(r)
}

// SVG_取旋转.
//
// hSvg: SVG句柄.
//
// pAngle: 返回角度.
//
// pX: 返回 旋转中心点X.
//
// pY: 返回 旋转中心点Y.
//
// pbOffset: 返回TRUE: 旋转中心点相对于自身中心偏移, FALSE:使用绝对坐标.
func XSvg_GetRotate(hSvg int, pAngle *float32, pX *float32, pY *float32, pbOffset *bool) int {
	r, _, _ := xSvg_GetRotate.Call(uintptr(hSvg), uintptr(unsafe.Pointer(pAngle)), uintptr(unsafe.Pointer(pX)), uintptr(unsafe.Pointer(pY)), uintptr(unsafe.Pointer(pbOffset)))
	return int(r)
}

// SVG_显示, 显示或隐藏.
//
// hSvg: SVG句柄.
//
// bShow: 是否显示.
func XSvg_Show(hSvg int, bShow bool) int {
	r, _, _ := xSvg_Show.Call(uintptr(hSvg), BoolPtr(bShow))
	return int(r)
}

// SVG_加载从字符串W.
//
// pString: 字符串指针.
func XSvg_LoadStringW(pString string) int {
	r, _, _ := xSvg_LoadStringW.Call(StrPtr(pString))
	return int(r)
}

// SVG_加载从字符串UTF8.
//
// pString: 字符串指针.
func XSvg_LoadStringUtf8(pString string) int {
	r, _, _ := xSvg_LoadStringUtf8.Call(XC_wtoutf8(pString))
	return int(r)
}
