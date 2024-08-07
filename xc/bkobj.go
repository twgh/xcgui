package xc

import (
	"github.com/twgh/xcgui/common"
	"unsafe"

	"github.com/twgh/xcgui/xcc"
)

// 背景对象_置外间距, 外间距与对齐标识(BkObject_Align_Flag_)互相依赖.
//
// hObj: 背景对象句柄.
//
// left: 左边间距.
//
// top: 顶部间距.
//
// right: 右边间距.
//
// bottom: 底部间距.
func XBkObj_SetMargin(hObj int, left, top, right, bottom int32) {
	xBkObj_SetMargin.Call(uintptr(hObj), uintptr(left), uintptr(top), uintptr(right), uintptr(bottom))
}

// 背景对象_置对齐.
//
// hObj: 背景对象句柄.
//
// nFlags: 对齐方式: xcc.BkObject_Align_Flag_.
func XBkObj_SetAlign(hObj int, nFlags xcc.BkObject_Align_Flag_) {
	xBkObj_SetAlign.Call(uintptr(hObj), uintptr(nFlags))
}

// 背景对象_置图片.
//
// hObj: 背景对象句柄.
//
// hImage: 图片句柄.
func XBkObj_SetImage(hObj int, hImage int) {
	xBkObj_SetImage.Call(uintptr(hObj), uintptr(hImage))
}

// 背景对象_置旋转.
//
// hObj: 背景对象句柄.
//
// angle: 旋转角度.
func XBkObj_SetRotate(hObj int, angle float32) {
	xBkObj_SetRotate.Call(uintptr(hObj), common.Float32Ptr(angle))
}

// 背景对象_置填充颜色.
//
// hObj: 背景对象句柄.
//
// color: xc.RGBA 颜色值.
func XBkObj_SetFillColor(hObj int, color int) {
	xBkObj_SetFillColor.Call(uintptr(hObj), uintptr(color))
}

// 背景对象_置边框宽度.
//
// hObj: 背景对象句柄.
//
// width: 宽度.
func XBkObj_SetBorderWidth(hObj int, width int32) {
	xBkObj_SetBorderWidth.Call(uintptr(hObj), uintptr(width))
}

// 背景对象_置边框颜色.
//
// hObj: 背景对象句柄.
//
// color: xc.RGBA 颜色值.
func XBkObj_SetBorderColor(hObj int, color int) {
	xBkObj_SetBorderColor.Call(uintptr(hObj), uintptr(color))
}

// 背景对象_置矩形圆角.
//
// hObj: 背景对象句柄.
//
// leftTop: 左上角.
//
// leftBottom: 左下角.
//
// rightTop: 右上角.
//
// rightBottom: 右下角.
func XBkObj_SetRectRoundAngle(hObj int, leftTop, leftBottom, rightTop, rightBottom int32) {
	xBkObj_SetRectRoundAngle.Call(uintptr(hObj), uintptr(leftTop), uintptr(leftBottom), uintptr(rightTop), uintptr(rightBottom))
}

// 背景对象_启用填充, 启用绘制填充.
//
// hObj: 背景对象句柄.
//
// bEnable: 是否启用.
func XBkObj_EnableFill(hObj int, bEnable bool) {
	xBkObj_EnableFill.Call(uintptr(hObj), common.BoolPtr(bEnable))
}

// 背景对象_启用边框, 启用绘制边框.
//
// hObj: 背景对象句柄.
//
// bEnable: 是否启用.
func XBkObj_EnableBorder(hObj int, bEnable bool) {
	xBkObj_EnableBorder.Call(uintptr(hObj), common.BoolPtr(bEnable))
}

// 背景对象_置文本.
//
// hObj: 背景对象句柄.
//
// pText: 文本字符串.
func XBkObj_SetText(hObj int, pText string) {
	xBkObj_SetText.Call(uintptr(hObj), common.StrPtr(pText))
}

// 背景对象_置字体.
//
// hObj: 背景对象句柄.
//
// hFont: 字体句柄.
func XBkObj_SetFont(hObj int, hFont int) {
	xBkObj_SetFont.Call(uintptr(hObj), uintptr(hFont))
}

// 背景对象_置文本对齐.
//
// hObj: 背景对象句柄.
//
// nAlign: 文本对齐方式: xcc.TextFormatFlag_, xcc.TextAlignFlag_, xcc.TextTrimming_.
func XBkObj_SetTextAlign(hObj int, nAlign xcc.TextFormatFlag_) {
	xBkObj_SetTextAlign.Call(uintptr(hObj), uintptr(nAlign))
}

// 背景对象_取外间距.
//
// hObj: 背景对象句柄.
//
// pMargin: 接收返回外间距.
func XBkObj_GetMargin(hObj int, pMargin *RECT) {
	xBkObj_GetMargin.Call(uintptr(hObj), uintptr(unsafe.Pointer(pMargin)))
}

// 背景对象_取对齐, 返回对齐标识: xcc.BkObject_Align_Flag_.
//
// hObj: 背景对象句柄.
func XBkObj_GetAlign(hObj int) xcc.BkObject_Align_Flag_ {
	r, _, _ := xBkObj_GetAlign.Call(uintptr(hObj))
	return xcc.BkObject_Align_Flag_(r)
}

// 背景对象_取图片, 返回图片句柄.
//
// hObj: 背景对象句柄.
func XBkObj_GetImage(hObj int) int {
	r, _, _ := xBkObj_GetImage.Call(uintptr(hObj))
	return int(r)
}

// 背景对象_取旋转角度, 返回旋转角度.
//
// hObj: 背景对象句柄.
func XBkObj_GetRotate(hObj int) int32 {
	r, _, _ := xBkObj_GetRotate.Call(uintptr(hObj))
	return int32(r)
}

// 背景对象_取填充色, 返回 RGBA 填充色.
//
// hObj: 背景对象句柄.
func XBkObj_GetFillColor(hObj int) int {
	r, _, _ := xBkObj_GetFillColor.Call(uintptr(hObj))
	return int(r)
}

// 背景对象_取边框色, 返回 RGBA 边框色.
//
// hObj: 背景对象句柄.
func XBkObj_GetBorderColor(hObj int) int {
	r, _, _ := xBkObj_GetBorderColor.Call(uintptr(hObj))
	return int(r)
}

// 背景对象_取边框宽度, 返回边框宽度.
//
// hObj: 背景对象句柄.
func XBkObj_GetBorderWidth(hObj int) int32 {
	r, _, _ := xBkObj_GetBorderWidth.Call(uintptr(hObj))
	return int32(r)
}

// 背景对象_取矩形圆角.
//
// hObj: 背景对象句柄.
//
// pRect: 接收返回圆角大小.
func XBkObj_GetRectRoundAngle(hObj int, pRect *RECT) {
	xBkObj_GetRectRoundAngle.Call(uintptr(hObj), uintptr(unsafe.Pointer(pRect)))
}

// 背景对象_是否填充.
//
// hObj: 背景对象句柄.
func XBkObj_IsFill(hObj int) bool {
	r, _, _ := xBkObj_IsFill.Call(uintptr(hObj))
	return r != 0
}

// 背景对象_是否边框.
//
// hObj: 背景对象句柄.
func XBkObj_IsBorder(hObj int) bool {
	r, _, _ := xBkObj_IsBorder.Call(uintptr(hObj))
	return r != 0
}

// 背景对象_取文本.
//
// hObj: 背景对象句柄.
func XBkObj_GetText(hObj int) string {
	r, _, _ := xBkObj_GetText.Call(uintptr(hObj))
	return common.UintPtrToString(r)
}

// 背景对象_取字体, 返回字体句柄.
//
// hObj: 背景对象句柄.
func XBkObj_GetFont(hObj int) int {
	r, _, _ := xBkObj_GetFont.Call(uintptr(hObj))
	return int(r)
}

// 背景对象_取文本对齐, 返回文本对齐方式: xcc.TextFormatFlag_.
//
// hObj: 背景对象句柄.
func XBkObj_GetTextAlign(hObj int) xcc.TextFormatFlag_ {
	r, _, _ := xBkObj_GetTextAlign.Call(uintptr(hObj))
	return xcc.TextFormatFlag_(r)
}
