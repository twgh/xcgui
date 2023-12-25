package xc

import (
	"github.com/twgh/xcgui/common"
	"unsafe"
)

// SVG_加载从文件, 返回SVG句柄.
//
// pFileName: 文件名.
func XSvg_LoadFile(pFileName string) int {
	r, _, _ := xSvg_LoadFile.Call(common.StrPtr(pFileName))
	return int(r)
}

// SVG_加载从字符串, 返回SVG句柄.
//
// pString: 字符串.
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
	r, _, _ := xSvg_LoadZip.Call(common.StrPtr(pZipFileName), common.StrPtr(pFileName), common.StrPtr(pPassword))
	return int(r)
}

// SVG_加载从资源ZIP, 返回SVG句柄.
//
// id: 资源ID.
//
// pFileName: svg文件名.
//
// pPassword: zip密码.
//
// hModule: 模块句柄, 可填0.
func XSvg_LoadZipRes(id int32, pFileName, pPassword string, hModule uintptr) int {
	r, _, _ := xSvg_LoadZipRes.Call(uintptr(id), common.StrPtr(pFileName), common.StrPtr(pPassword), hModule)
	return int(r)
}

// SVG_加载从资源, 返回SVG句柄.
//
// id: 资源ID.
//
// pType: 资源类型.在rc资源文件中.
//
// hModule: 从指定模块加载.
func XSvg_LoadRes(id int32, pType string, hModule uintptr) int {
	r, _, _ := xSvg_LoadRes.Call(uintptr(id), common.StrPtr(pType), hModule)
	return int(r)
}

// SVG_置大小.
//
// hSvg: SVG句柄.
//
// nWidth: 宽度.
//
// nHeight: 高度.
func XSvg_SetSize(hSvg int, nWidth, nHeight int32) {
	xSvg_SetSize.Call(uintptr(hSvg), uintptr(nWidth), uintptr(nHeight))
}

// SVG_取大小.
//
// hSvg: SVG句柄.
//
// pWidth: 接收返回宽度.
//
// pHeight: 接收返回高度.
func XSvg_GetSize(hSvg int, pWidth, pHeight *int32) {
	xSvg_GetSize.Call(uintptr(hSvg), uintptr(unsafe.Pointer(pWidth)), uintptr(unsafe.Pointer(pHeight)))
}

// SVG_取宽度.
//
// hSvg: SVG句柄.
func XSvg_GetWidth(hSvg int) int32 {
	r, _, _ := xSvg_GetWidth.Call(uintptr(hSvg))
	return int32(r)
}

// SVG_取高度.
//
// hSvg: SVG句柄.
func XSvg_GetHeight(hSvg int) int32 {
	r, _, _ := xSvg_GetHeight.Call(uintptr(hSvg))
	return int32(r)
}

// SVG_置偏移.
//
// hSvg: SVG句柄.
//
// x: x轴偏移.
//
// y: y轴偏移.
func XSvg_SetPosition(hSvg int, x, y int32) {
	xSvg_SetPosition.Call(uintptr(hSvg), uintptr(x), uintptr(y))
}

// SVG_取偏移.
//
// hSvg: SVG句柄.
//
// pX: x轴偏移.
//
// pY: y轴偏移.
func XSvg_GetPosition(hSvg int, pX, pY *int32) {
	xSvg_GetPosition.Call(uintptr(hSvg), uintptr(unsafe.Pointer(pX)), uintptr(unsafe.Pointer(pY)))
}

// SVG_置偏移F.
//
// hSvg: SVG句柄.
//
// x: x轴偏移.
//
// y: y轴偏移.
func XSvg_SetPositionF(hSvg int, x, y float32) {
	xSvg_SetPositionF.Call(uintptr(hSvg), common.Float32Ptr(x), common.Float32Ptr(y))
}

// SVG_取偏移F.
//
// hSvg: SVG句柄.
//
// pX: x轴偏移.
//
// pY: y轴偏移.
func XSvg_GetPositionF(hSvg int, pX, pY *float32) {
	xSvg_GetPositionF.Call(uintptr(hSvg), uintptr(unsafe.Pointer(pX)), uintptr(unsafe.Pointer(pY)))
}

// SVG_取视图框.
//
// hSvg: SVG句柄.
//
// pViewBox: 接收返回视图框.
func XSvg_GetViewBox(hSvg int, pViewBox *RECT) {
	xSvg_GetViewBox.Call(uintptr(hSvg), uintptr(unsafe.Pointer(pViewBox)))
}

// SVG_启用自动销毁.
//
// hSvg: SVG句柄.
//
// bEnable: 是否自动销毁.
func XSvg_EnableAutoDestroy(hSvg int, bEnable bool) {
	xSvg_EnableAutoDestroy.Call(uintptr(hSvg), common.BoolPtr(bEnable))
}

// SVG_增加引用计数.
//
// hSvg: SVG句柄.
func XSvg_AddRef(hSvg int) {
	xSvg_AddRef.Call(uintptr(hSvg))
}

// SVG_释放引用计数.
//
// hSvg: SVG句柄.
func XSvg_Release(hSvg int) {
	xSvg_Release.Call(uintptr(hSvg))
}

// SVG_取引用计数.
//
// hSvg: SVG句柄.
func XSvg_GetRefCount(hSvg int) int32 {
	r, _, _ := xSvg_GetRefCount.Call(uintptr(hSvg))
	return int32(r)
}

// SVG_销毁.
//
// hSvg: SVG句柄.
func XSvg_Destroy(hSvg int) {
	xSvg_Destroy.Call(uintptr(hSvg))
}

// SVG_置透明度.
//
// hSvg: SVG句柄.
//
// alpha: 透明度.
func XSvg_SetAlpha(hSvg int, alpha byte) {
	xSvg_SetAlpha.Call(uintptr(hSvg), uintptr(alpha))
}

// SVG_取透明度, 返回透明度.
//
// hSvg: SVG句柄.
func XSvg_GetAlpha(hSvg int) byte {
	r, _, _ := xSvg_GetAlpha.Call(uintptr(hSvg))
	return byte(r)
}

// SVG_置用户填充颜色, 用户颜色将覆盖默认样式.
//
// hSvg: SVG句柄.
//
// color: 颜色, AGBR颜色.
//
// bEnable: 是否有效.
func XSvg_SetUserFillColor(hSvg int, color int, bEnable bool) {
	xSvg_SetUserFillColor.Call(uintptr(hSvg), uintptr(color), common.BoolPtr(bEnable))
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
func XSvg_SetUserStrokeColor(hSvg int, color int, strokeWidth float32, bEnable bool) {
	xSvg_SetUserStrokeColor.Call(uintptr(hSvg), uintptr(color), common.Float32Ptr(strokeWidth), common.BoolPtr(bEnable))
}

// SVG_取用户填充颜色.
//
// hSvg: SVG句柄.
//
// pColor: 返回颜色值, AGBR颜色.
func XSvg_GetUserFillColor(hSvg int, pColor *int) bool {
	r, _, _ := xSvg_GetUserFillColor.Call(uintptr(hSvg), uintptr(unsafe.Pointer(pColor)))
	return r != 0
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
	return r != 0
}

// SVG_置旋转角度, 默认以自身中心点旋转.
//
// hSvg: SVG句柄.
//
// angle: 转角度.
func XSvg_SetRotateAngle(hSvg int, angle float32) {
	xSvg_SetRotateAngle.Call(uintptr(hSvg), common.Float32Ptr(angle))
}

// SVG_取旋转角度, 返回旋转角度.
//
// hSvg: SVG句柄.
func XSvg_GetRotateAngle(hSvg int) float32 {
	_, r, _ := xSvg_GetRotateAngle.Call(uintptr(hSvg))
	return common.UintPtrToFloat32(r)
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
func XSvg_SetRotate(hSvg int, angle float32, x float32, y float32, bOffset bool) {
	xSvg_SetRotate.Call(uintptr(hSvg), common.Float32Ptr(angle), common.Float32Ptr(x), common.Float32Ptr(y), common.BoolPtr(bOffset))
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
func XSvg_GetRotate(hSvg int, pAngle *float32, pX *float32, pY *float32, pbOffset *bool) {
	xSvg_GetRotate.Call(uintptr(hSvg), uintptr(unsafe.Pointer(pAngle)), uintptr(unsafe.Pointer(pX)), uintptr(unsafe.Pointer(pY)), uintptr(unsafe.Pointer(pbOffset)))
}

// SVG_显示, 显示或隐藏.
//
// hSvg: SVG句柄.
//
// bShow: 是否显示.
func XSvg_Show(hSvg int, bShow bool) {
	xSvg_Show.Call(uintptr(hSvg), common.BoolPtr(bShow))
}

// SVG_加载从字符串W.
//
// pString: 字符串.
func XSvg_LoadStringW(pString string) int {
	r, _, _ := xSvg_LoadStringW.Call(common.StrPtr(pString))
	return int(r)
}

// SVG_加载从字符串UTF8.
//
// pString: 字符串.
func XSvg_LoadStringUtf8(pString string) int {
	r, _, _ := xSvg_LoadStringUtf8.Call(XC_wtoutf8(pString))
	return int(r)
}

// SVG_加载从内存ZIP.
//
// data: zip数据.
//
// pFileName: svg文件名.
//
// pPassword: zip密码.
func XSvg_LoadZipMem(data []byte, pFileName, pPassword string) int {
	r, _, _ := xSvg_LoadZipMem.Call(common.ByteSliceDataPtr(&data), uintptr(len(data)), common.StrPtr(pFileName), common.StrPtr(pPassword))
	return int(r)
}
