package xc

import (
	"unsafe"

	"github.com/twgh/xcgui/common"

	"github.com/twgh/xcgui/xcc"
)

// 绘制_创建, 创建图形绘制模块实例, 返回句柄.
//
// hWindow: 窗口句柄.
func XDraw_Create(hWindow int) int {
	r, _, _ := xDraw_Create.Call(uintptr(hWindow))
	return int(r)
}

// 绘制_创建GDI, 创建图形绘制模块实例, 返回图形绘制模块实例句柄.
//
// hWindow: 窗口句柄.
//
// hdc: hdc句柄.
func XDraw_CreateGDI(hWindow int, hdc uintptr) int {
	r, _, _ := xDraw_CreateGDI.Call(uintptr(hWindow), hdc)
	return int(r)
}

// 绘制_销毁, 销毁图形绘制模块实例句柄.
//
// hDraw: 图形绘制句柄.
func XDraw_Destroy(hDraw int) int {
	r, _, _ := xDraw_Destroy.Call(uintptr(hDraw))
	return int(r)
}

// 绘制_虚线, 绘制水平或垂直虚线.
//
// hDraw: 图形绘制句柄.
//
// x1: 起点x坐标.
//
// y1: 起点y坐标.
//
// x2: 结束点x坐标.
//
// y2: 结束点y坐标.
func XDraw_Dottedline(hDraw int, x1 int, y1 int, x2 int, y2 int) int {
	r, _, _ := xDraw_Dottedline.Call(uintptr(hDraw), uintptr(x1), uintptr(y1), uintptr(x2), uintptr(y2))
	return int(r)
}

// 绘制_虚线F, 绘制水平或垂直虚线.
//
// hDraw: 图形绘制句柄.
//
// x1: 起点x坐标.
//
// y1: 起点y坐标.
//
// x2: 结束点x坐标.
//
// y2: 结束点y坐标.
func XDraw_DottedlineF(hDraw int, x1, y1, x2, y2 float32) int {
	r, _, _ := xDraw_DottedlineF.Call(uintptr(hDraw), common.Float32Ptr(x1), common.Float32Ptr(y1), common.Float32Ptr(x2), common.Float32Ptr(y2))
	return int(r)
}

// 绘制_圆弧.
//
// hDraw: 图形绘制句柄.
//
// x: 坐标.
//
// y: 坐标.
//
// nWidth: 宽度.
//
// nHeight: 高度.
//
// startAngle: 起始角度.
//
// sweepAngle: 绘制角度, 从起始角度开始计算.
func XDraw_DrawArc(hDraw int, x, y int, nWidth int, nHeight int, startAngle float32, sweepAngle float32) int {
	r, _, _ := xDraw_DrawArc.Call(uintptr(hDraw), uintptr(x), uintptr(y), uintptr(nWidth), uintptr(nHeight), common.Float32Ptr(startAngle), common.Float32Ptr(sweepAngle))
	return int(r)
}

// 绘制_圆弧F.
//
// hDraw: 图形绘制句柄.
//
// x: 坐标.
//
// y: 坐标.
//
// nWidth: 宽度.
//
// nHeight: 高度.
//
// startAngle: 起始角度.
//
// sweepAngle: 绘制角度, 从起始角度开始计算.
func XDraw_DrawArcF(hDraw int, x, y, nWidth, nHeight, startAngle, sweepAngle float32) int {
	r, _, _ := xDraw_DrawArcF.Call(uintptr(hDraw), common.Float32Ptr(x), common.Float32Ptr(y), common.Float32Ptr(nWidth), common.Float32Ptr(nHeight), common.Float32Ptr(startAngle), common.Float32Ptr(sweepAngle))
	return int(r)
}

// 绘制_曲线, D2D暂时留空.
//
// hDraw: 图形绘制句柄.
//
// points: 坐标点数组.
//
// count: 数组大小.
//
// tension: 大于或等于0.0F的值，指定曲线的张力, D2D 忽略此参数.
func XDraw_DrawCurve(hDraw int, points []POINT, count int, tension float32) int {
	r, _, _ := xDraw_DrawCurve.Call(uintptr(hDraw), uintptr(unsafe.Pointer(&points[0])), uintptr(count), common.Float32Ptr(tension))
	return int(r)
}

// 绘制_曲线F, D2D暂时留空.
//
// hDraw: 图形绘制句柄.
//
// points: 坐标点数组.
//
// count: 数组大小.
//
// tension: 大于或等于0.0F的值，指定曲线的张力, D2D 忽略此参数.
func XDraw_DrawCurveF(hDraw int, points []POINTF, count int, tension float32) int {
	r, _, _ := xDraw_DrawCurveF.Call(uintptr(hDraw), uintptr(unsafe.Pointer(&points[0])), uintptr(count), common.Float32Ptr(tension))
	return int(r)
}

// 绘制_线条.
//
// hDraw: 图形绘制句柄.
//
// x1: 坐标.
//
// y1: 坐标.
//
// x2: 坐标.
//
// y2: 坐标.
func XDraw_DrawLine(hDraw int, x1 int, y1 int, x2 int, y2 int) int {
	r, _, _ := xDraw_DrawLine.Call(uintptr(hDraw), uintptr(x1), uintptr(y1), uintptr(x2), uintptr(y2))
	return int(r)
}

// 绘制_线条F.
//
// hDraw: 图形绘制句柄.
//
// x1: 坐标.
//
// y1: 坐标.
//
// x2: 坐标.
//
// y2: 坐标.
func XDraw_DrawLineF(hDraw int, x1, y1, x2, y2 float32) int {
	r, _, _ := xDraw_DrawLineF.Call(uintptr(hDraw), common.Float32Ptr(x1), common.Float32Ptr(y1), common.Float32Ptr(x2), common.Float32Ptr(y2))
	return int(r)
}

// 绘制_多边形, 绘制多边形.
//
// hDraw: 图形绘制句柄.
//
// points: 顶点坐标数组.
//
// nCount: 顶点数量.
func XDraw_DrawPolygon(hDraw int, points []POINT, nCount int) int {
	r, _, _ := xDraw_DrawPolygon.Call(uintptr(hDraw), uintptr(unsafe.Pointer(&points[0])), uintptr(nCount))
	return int(r)
}

// 绘制_多边形F, 绘制多边形.
//
// hDraw: 图形绘制句柄.
//
// points: 顶点坐标数组.
//
// nCount: 顶点数量.
func XDraw_DrawPolygonF(hDraw int, points []POINTF, nCount int) int {
	r, _, _ := xDraw_DrawPolygonF.Call(uintptr(hDraw), uintptr(unsafe.Pointer(&points[0])), uintptr(nCount))
	return int(r)
}

// 绘制_矩形, 绘制矩形边框.
//
// hDraw: 图形绘制句柄.
//
// pRect: 矩形坐标 .
func XDraw_DrawRect(hDraw int, pRect *RECT) int {
	r, _, _ := xDraw_DrawRect.Call(uintptr(hDraw), uintptr(unsafe.Pointer(pRect)))
	return int(r)
}

// 绘制_矩形F, 绘制矩形边框.
//
// hDraw: 图形绘制句柄.
//
// pRect: 矩形坐标 .
func XDraw_DrawRectF(hDraw int, pRect *RECTF) int {
	r, _, _ := xDraw_DrawRectF.Call(uintptr(hDraw), uintptr(unsafe.Pointer(pRect)))
	return int(r)
}

// 绘制_置偏移, 设置坐标偏移量, X向左偏移为负数, 向右偏移为正数.
//
// hDraw: 图形绘制句柄.
//
// x: X轴偏移量.
//
// y: Y轴偏移量.
func XDraw_SetOffset(hDraw int, x, y int32) int {
	r, _, _ := xDraw_SetOffset.Call(uintptr(hDraw), uintptr(x), uintptr(y))
	return int(r)
}

// 绘制_取偏移, 获取坐标偏移量, X向左偏移为负数, 向右偏移为正数.
//
// hDraw: 图形绘制句柄.
//
// pX: 接收X轴偏移量.
//
// pY: 接收Y轴偏移量.
func XDraw_GetOffset(hDraw int, pX, pY *int32) int {
	r, _, _ := xDraw_GetOffset.Call(uintptr(hDraw), uintptr(unsafe.Pointer(pX)), uintptr(unsafe.Pointer(pY)))
	return int(r)
}

// 绘制_还原状态, 还原状态, 释放用户绑定的GDI对象, 例如画刷, 画笔.
//
// hDraw: 图形绘制句柄.
func XDraw_GDI_RestoreGDIOBJ(hDraw int) int {
	r, _, _ := xDraw_GDI_RestoreGDIOBJ.Call(uintptr(hDraw))
	return int(r)
}

// 绘制_取HDC, 获取绑定的设备上下文HDC, 返回HDC句柄.
//
// hDraw: 图形绘制句柄.
func XDraw_GetHDC(hDraw int) uintptr {
	r, _, _ := xDraw_GetHDC.Call(uintptr(hDraw))
	return r
}

// 绘制_置画刷颜色, 设置画刷颜色.
//
// hDraw: 图形绘制句柄.
//
// color: ABGR 颜色值.
func XDraw_SetBrushColor(hDraw int, color int) int {
	r, _, _ := xDraw_SetBrushColor.Call(uintptr(hDraw), uintptr(color))
	return int(r)
}

// 绘制_置文本垂直, 设置文本垂直显示.
//
// hDraw: 图形绘制句柄.
//
// bVertical: 是否垂直显示文本.
func XDraw_SetTextVertical(hDraw int, bVertical bool) int {
	r, _, _ := xDraw_SetTextVertical.Call(uintptr(hDraw), common.BoolPtr(bVertical))
	return int(r)
}

// 绘制_置文本对齐, 设置文本对齐.
//
// hDraw: 图形绘制句柄.
//
// nFlags: 对齐标识, TextFormatFlag_, TextAlignFlag_, TextTrimming_.
func XDraw_SetTextAlign(hDraw int, nFlags xcc.TextFormatFlag_) int {
	r, _, _ := xDraw_SetTextAlign.Call(uintptr(hDraw), uintptr(nFlags))
	return int(r)
}

// 绘制_置字体.
//
// hDraw: 图形绘制句柄.
//
// hFontx: 炫彩字体.
func XDraw_SetFont(hDraw int, hFontx int) int {
	r, _, _ := xDraw_SetFont.Call(uintptr(hDraw), uintptr(hFontx))
	return int(r)
}

// 绘制_置线宽.
//
// hDraw: 图形绘制句柄.
//
// nWidth: 宽度.
func XDraw_SetLineWidth(hDraw int, nWidth int) int {
	r, _, _ := xDraw_SetLineWidth.Call(uintptr(hDraw), uintptr(nWidth))
	return int(r)
}

// 绘制_置线宽F.
//
// hDraw: 图形绘制句柄.
//
// nWidth: 宽度.
func XDraw_SetLineWidthF(hDraw int, nWidth float32) int {
	r, _, _ := xDraw_SetLineWidthF.Call(uintptr(hDraw), common.Float32Ptr(nWidth))
	return int(r)
}

// 绘制_置背景模式, SetBkMode() 参见MSDN.
//
// hDraw: 图形绘制句柄.
//
// bTransparent: 参见MSDN.
func XDraw_GDI_SetBkMode(hDraw int, bTransparent bool) int {
	r, _, _ := xDraw_GDI_SetBkMode.Call(uintptr(hDraw), common.BoolPtr(bTransparent))
	return int(r)
}

// 绘制_置裁剪区域, 设置裁剪区域.
//
// hDraw: 图形绘制句柄.
//
// pRect: 区域坐标.
func XDraw_SetClipRect(hDraw int, pRect *RECT) int {
	r, _, _ := xDraw_SetClipRect.Call(uintptr(hDraw), uintptr(unsafe.Pointer(pRect)))
	return int(r)
}

// 绘制_置D2D文本渲染模式.
//
// hDraw: 图形绘制句柄.
//
// mode	渲染模式 XC_DWRITE_RENDERING_MODE_.
func XDraw_SetD2dTextRenderingMode(hDraw int, mode xcc.XC_DWRITE_RENDERING_MODE_) int {
	r, _, _ := xDraw_SetD2dTextRenderingMode.Call(uintptr(hDraw), uintptr(mode))
	return int(r)
}

// 绘制_清除裁剪区域.
//
// hDraw: 图形绘制句柄.
func XDraw_ClearClip(hDraw int) int {
	r, _, _ := xDraw_ClearClip.Call(uintptr(hDraw))
	return int(r)
}

// 绘制_启用平滑模式.
//
// hDraw: 图形绘制句柄.
//
// bEnable: 是否启用.
func XDraw_EnableSmoothingMode(hDraw int, bEnable bool) int {
	r, _, _ := xDraw_EnableSmoothingMode.Call(uintptr(hDraw), common.BoolPtr(bEnable))
	return int(r)
}

// 绘制_启用窗口透明判断, 当启用之后, 调用GDI+函数时, 如果参数alpha=255, 将自动修改为254, 应对GDI+的bug, 否则透明通道异常.
//
// hDraw: 图形绘制句柄.
//
// bTransparent: 是否启用.
func XDraw_EnableWndTransparent(hDraw int, bTransparent bool) int {
	r, _, _ := xDraw_EnableWndTransparent.Call(uintptr(hDraw), common.BoolPtr(bTransparent))
	return int(r)
}

// 绘制_创建实心画刷, GDI创建具有指定的纯色逻辑刷.
//
// hDraw: 图形绘制句柄.
//
// crColor: 画刷颜色, RGB颜色.
func XDraw_GDI_CreateSolidBrush(hDraw int, crColor int) int {
	r, _, _ := xDraw_GDI_CreateSolidBrush.Call(uintptr(hDraw), uintptr(crColor))
	return int(r)
}

// 绘制_创建画笔, GDI创建一个逻辑笔, 指定的样式, 宽度和颜色, 随后的笔可以选择到设备上下文, 用于绘制线条和曲线.
//
// hDraw: 图形绘制句柄.
//
// fnPenStyle: 画笔样式, PS_SOLID:实线, PS_DASH:段线, PS_DOT:点线, PS_DASHDOT:段线_点线, PS_DASHDOTDOT:段线_点_点, PS_NULL:空, PS_INSIDEFRAME:实线_笔宽是向里扩展.
//
// nWidth: 画笔宽度.
//
// crColor: RGB颜色.
func XDraw_GDI_CreatePen(hDraw int, fnPenStyle int, nWidth int, crColor int) int {
	r, _, _ := xDraw_GDI_CreatePen.Call(uintptr(hDraw), uintptr(fnPenStyle), uintptr(nWidth), uintptr(crColor))
	return int(r)
}

// 绘制_创建矩形区域, GDI创建矩形区域, 成功返回区域句柄, 失败返回NULL.
//
// hDraw: 图形绘制句柄.
//
// nLeftRect: 左上角X坐标.
//
// nTopRect: 左上角Y坐标.
//
// nRightRect: 右下角X坐标.
//
// nBottomRect: 右下角Y坐标.
func XDraw_GDI_CreateRectRgn(hDraw int, nLeftRect int, nTopRect int, nRightRect int, nBottomRect int) int {
	r, _, _ := xDraw_GDI_CreateRectRgn.Call(uintptr(hDraw), uintptr(nLeftRect), uintptr(nTopRect), uintptr(nRightRect), uintptr(nBottomRect))
	return int(r)
}

// 绘制_创建圆角矩形区域, GDI创建一个圆角的矩形区域, 成功返回区域句柄, 失败返回NULL.
//
// hDraw: 图形绘制句柄.
//
// nLeftRect: X-坐标的左上角.
//
// nTopRect: Y-坐标左上角坐标.
//
// nRightRect: X-坐标右下角.
//
// nBottomRect: Y-坐标右下角.
//
// nWidthEllipse: 椭圆的宽度.
//
// nHeightEllipse: 椭圆的高度.
func XDraw_GDI_CreateRoundRectRgn(hDraw int, nLeftRect int, nTopRect int, nRightRect int, nBottomRect int, nWidthEllipse int, nHeightEllipse int) int {
	r, _, _ := xDraw_GDI_CreateRoundRectRgn.Call(uintptr(hDraw), uintptr(nLeftRect), uintptr(nTopRect), uintptr(nRightRect), uintptr(nBottomRect), uintptr(nWidthEllipse), uintptr(nHeightEllipse))
	return int(r)
}

// 绘制_创建多边形区域, GDI创建一个多边形区域, 成功返回区域句柄, 失败返回NULL.
//
// hDraw: 图形绘制句柄.
//
// pPt: POINT数组.
//
// cPoints: 数组大小.
//
// fnPolyFillMode: 多边形填充模式, 指定用于确定在该地区的像素填充模式,这个参数可以是下列值之一.
//
// ALTERNATE Selects alternate mode (fills area between odd-numbered and even-numbered polygon sides on each scan line).
//
// WINDING Selects winding mode (fills any region with a nonzero winding value).
func XDraw_GDI_CreatePolygonRgn(hDraw int, pPt []POINT, cPoints int, fnPolyFillMode int) int {
	r, _, _ := xDraw_GDI_CreatePolygonRgn.Call(uintptr(hDraw), uintptr(unsafe.Pointer(&pPt[0])), uintptr(cPoints), uintptr(fnPolyFillMode))
	return int(r)
}

// 绘制_选择裁剪区域, 选择一个区域作为当前裁剪区域, 注意: 该函数只对GDI有效.
//
// hDraw: 图形绘制句柄.
//
// hRgn: 区域句柄.
//
// 返回: 返回值指定地区的复杂性，可以是下列值之一.
//
// NULLREGION Region is empty.
//
// SIMPLEREGION Region is a single rectangle.
//
// COMPLEXREGION Region is more than one rectangle.
//
// ERROR An error occurred. (The previous clipping region is unaffected).
func XDraw_GDI_SelectClipRgn(hDraw int, hRgn int) int {
	r, _, _ := xDraw_GDI_SelectClipRgn.Call(uintptr(hDraw), uintptr(hRgn))
	return int(r)
}

// 绘制_填充矩形, 通过使用指定的刷子填充一个矩形, 此功能包括左侧和顶部的边界, 但不包括矩形的右边和底部边界.
//
// hDraw: 图形绘制句柄.
//
// pRect: 矩形区域.
func XDraw_FillRect(hDraw int, pRect *RECT) int {
	r, _, _ := xDraw_FillRect.Call(uintptr(hDraw), uintptr(unsafe.Pointer(pRect)))
	return int(r)
}

// 绘制_填充矩形F, 通过使用指定的刷子填充一个矩形, 此功能包括左侧和顶部的边界, 但不包括矩形的右边和底部边界.
//
// hDraw: 图形绘制句柄.
//
// pRect: 矩形区域.
func XDraw_FillRectF(hDraw int, pRect *RECTF) int {
	r, _, _ := xDraw_FillRectF.Call(uintptr(hDraw), uintptr(unsafe.Pointer(pRect)))
	return int(r)
}

// 绘制_填充矩形指定颜色.
//
// hDraw: 图形绘制句柄.
//
// pRect: 矩形区域.
//
// color: ABGR 颜色.
func XDraw_FillRectColor(hDraw int, pRect *RECT, color int) int {
	r, _, _ := xDraw_FillRectColor.Call(uintptr(hDraw), uintptr(unsafe.Pointer(pRect)), uintptr(color))
	return int(r)
}

// 绘制_填充矩形指定颜色F.
//
// hDraw: 图形绘制句柄.
//
// pRect: 矩形区域.
//
// color: ABGR 颜色.
func XDraw_FillRectColorF(hDraw int, pRect *RECTF, color int) int {
	r, _, _ := xDraw_FillRectColorF.Call(uintptr(hDraw), uintptr(unsafe.Pointer(pRect)), uintptr(color))
	return int(r)
}

// 绘制_填充区域, 通过使用指定的画刷填充一个区域.
//
// hDraw: 图形绘制句柄.
//
// hrgn: 区域句柄.
//
// hbr: 画刷句柄.
func XDraw_GDI_FillRgn(hDraw int, hrgn int, hbr int) bool {
	r, _, _ := xDraw_GDI_FillRgn.Call(uintptr(hDraw), uintptr(hrgn), uintptr(hbr))
	return r != 0
}

// 绘制_填充圆形.
//
// hDraw: 图形绘制句柄.
//
// pRect: 矩形区域.
func XDraw_FillEllipse(hDraw int, pRect *RECT) int {
	r, _, _ := xDraw_FillEllipse.Call(uintptr(hDraw), uintptr(unsafe.Pointer(pRect)))
	return int(r)
}

// 绘制_填充圆形F.
//
// hDraw: 图形绘制句柄.
//
// pRect: 矩形区域.
func XDraw_FillEllipseF(hDraw int, pRect *RECTF) int {
	r, _, _ := xDraw_FillEllipseF.Call(uintptr(hDraw), uintptr(unsafe.Pointer(pRect)))
	return int(r)
}

// 绘制_圆形, 绘制圆边框.
//
// hDraw: 图形绘制句柄.
//
// pRect: 矩形区域.
func XDraw_DrawEllipse(hDraw int, pRect *RECT) int {
	r, _, _ := xDraw_DrawEllipse.Call(uintptr(hDraw), uintptr(unsafe.Pointer(pRect)))
	return int(r)
}

// 绘制_填充圆角矩形.
//
// hDraw: 图形绘制句柄.
//
// pRect: 矩形坐标.
//
// nWidth: 圆角宽度.
//
// nHeight: 圆角高度.
func XDraw_FillRoundRect(hDraw int, pRect *RECT, nWidth, nHeight int) int {
	r, _, _ := xDraw_FillRoundRect.Call(uintptr(hDraw), uintptr(unsafe.Pointer(pRect)), uintptr(nWidth), uintptr(nHeight))
	return int(r)
}

// 绘制_填充圆角矩形F.
//
// hDraw: 图形绘制句柄.
//
// pRect: 矩形坐标.
//
// nWidth: 圆角宽度.
//
// nHeight: 圆角高度.
func XDraw_FillRoundRectF(hDraw int, pRect *RECTF, nWidth, nHeight float32) int {
	r, _, _ := xDraw_FillRoundRectF.Call(uintptr(hDraw), uintptr(unsafe.Pointer(pRect)), common.Float32Ptr(nWidth), common.Float32Ptr(nHeight))
	return int(r)
}

// 绘制_圆角矩形, 绘制圆角矩形边框.
//
// hDraw: 图形绘制句柄.
//
// pRect: 矩形坐标.
//
// nWidth: 圆角宽度.
//
// nHeight: 圆角高度.
func XDraw_DrawRoundRect(hDraw int, pRect *RECT, nWidth int, nHeight int) int {
	r, _, _ := xDraw_DrawRoundRect.Call(uintptr(hDraw), uintptr(unsafe.Pointer(pRect)), uintptr(nWidth), uintptr(nHeight))
	return int(r)
}

// 绘制_圆角矩形F, 绘制圆角矩形边框.
//
// hDraw: 图形绘制句柄.
//
// pRect: 矩形坐标.
//
// nWidth: 圆角宽度.
//
// nHeight: 圆角高度.
func XDraw_DrawRoundRectF(hDraw int, pRect *RECT, nWidth, nHeight float32) int {
	r, _, _ := xDraw_DrawRoundRectF.Call(uintptr(hDraw), uintptr(unsafe.Pointer(pRect)), common.Float32Ptr(nWidth), common.Float32Ptr(nHeight))
	return int(r)
}

// 绘制_填充圆角矩形扩展.
//
// hDraw: 图形绘制句柄.
//
// pRect: 坐标.
//
// nLeftTop: 圆角大小.
//
// nRightTop: 圆角大小.
//
// nRightBottom: 圆角大小.
//
// nLeftBottom: 圆角大小.
func XDraw_FillRoundRectEx(hDraw int, pRect *RECT, nLeftTop, nRightTop, nRightBottom, nLeftBottom int) int {
	r, _, _ := xDraw_FillRoundRectEx.Call(uintptr(hDraw), uintptr(unsafe.Pointer(pRect)), uintptr(nLeftTop), uintptr(nRightTop), uintptr(nRightBottom), uintptr(nLeftBottom))
	return int(r)
}

// 绘制_填充圆角矩形扩展F.
//
// hDraw: 图形绘制句柄.
//
// pRect: 坐标.
//
// nLeftTop: 圆角大小.
//
// nRightTop: 圆角大小.
//
// nRightBottom: 圆角大小.
//
// nLeftBottom: 圆角大小.
func XDraw_FillRoundRectExF(hDraw int, pRect *RECTF, nLeftTop, nRightTop, nRightBottom, nLeftBottom float32) int {
	r, _, _ := xDraw_FillRoundRectExF.Call(uintptr(hDraw), uintptr(unsafe.Pointer(pRect)), common.Float32Ptr(nLeftTop), common.Float32Ptr(nRightTop), common.Float32Ptr(nRightBottom), common.Float32Ptr(nLeftBottom))
	return int(r)
}

// 绘制_圆角矩形扩展, 绘制圆角矩形边框.
//
// hDraw: 图形绘制句柄.
//
// pRect: 坐标.
//
// nLeftTop: 圆角大小.
//
// nRightTop: 圆角大小.
//
// nRightBottom: 圆角大小.
//
// nLeftBottom: 圆角大小.
func XDraw_DrawRoundRectEx(hDraw int, pRect *RECT, nLeftTop int, nRightTop int, nRightBottom int, nLeftBottom int) int {
	r, _, _ := xDraw_DrawRoundRectEx.Call(uintptr(hDraw), uintptr(unsafe.Pointer(pRect)), uintptr(nLeftTop), uintptr(nRightTop), uintptr(nRightBottom), uintptr(nLeftBottom))
	return int(r)
}

// 绘制_圆角矩形扩展F, 绘制圆角矩形边框.
//
// hDraw: 图形绘制句柄.
//
// pRect: 坐标.
//
// nLeftTop: 圆角大小.
//
// nRightTop: 圆角大小.
//
// nRightBottom: 圆角大小.
//
// nLeftBottom: 圆角大小.
func XDraw_DrawRoundRectExF(hDraw int, pRect *RECT, nLeftTop, nRightTop, nRightBottom, nLeftBottom float32) int {
	r, _, _ := xDraw_DrawRoundRectExF.Call(uintptr(hDraw), uintptr(unsafe.Pointer(pRect)), common.Float32Ptr(nLeftTop), common.Float32Ptr(nRightTop), common.Float32Ptr(nRightBottom), common.Float32Ptr(nLeftBottom))
	return int(r)
}

// 绘制_矩形, 绘制矩形, 使用当前的画刷和画笔. 如果函数成功, 返回非零值, 如果函数失败, 返回值是零.
//
// hDraw: 图形绘制句柄.
//
// nLeftRect: 左上角X坐标.
//
// nTopRect: 左上角Y坐标.
//
// nRightRect: 右下角X坐标.
//
// nBottomRect: 右下角Y坐标.
func XDraw_GDI_Rectangle(hDraw int, nLeftRect int, nTopRect int, nRightRect int, nBottomRect int) bool {
	r, _, _ := xDraw_GDI_Rectangle.Call(uintptr(hDraw), uintptr(nLeftRect), uintptr(nTopRect), uintptr(nRightRect), uintptr(nBottomRect))
	return r != 0
}

// 绘制_渐变填充2, 渐变填充, 从一种颜色过渡到另一种颜色.
//
// hDraw: 图形绘制句柄.
//
// pRect: 矩形坐标.
//
// color1: 开始颜色, ABGR 颜色.
//
// color2: 结束颜色, ABGR 颜色.
//
// mode: 模式, GRADIENT_FILL_.
func XDraw_GradientFill2(hDraw int, pRect *RECT, color1 int, color2 int, mode xcc.GRADIENT_FILL_) int {
	r, _, _ := xDraw_GradientFill2.Call(uintptr(hDraw), uintptr(unsafe.Pointer(pRect)), uintptr(color1), uintptr(color2), uintptr(mode))
	return int(r)
}

// 绘制_渐变填充2F, 渐变填充, 从一种颜色过渡到另一种颜色.
//
// hDraw: 图形绘制句柄.
//
// pRect: 矩形坐标.
//
// color1: 开始颜色, ABGR 颜色.
//
// color2: 结束颜色, ABGR 颜色.
//
// mode: 模式, GRADIENT_FILL_.
func XDraw_GradientFill2F(hDraw int, pRect *RECTF, color1 int, color2 int, mode xcc.GRADIENT_FILL_) int {
	r, _, _ := xDraw_GradientFill2F.Call(uintptr(hDraw), uintptr(unsafe.Pointer(pRect)), uintptr(color1), uintptr(color2), uintptr(mode))
	return int(r)
}

// 绘制_渐变填充4, 渐变填充,从一种颜色过渡到另一种颜色.
//
// hDraw: 图形绘制句柄.
//
// pRect: 矩形坐标.
//
// color1: 开始颜色, ABGR 颜色.
//
// color2: 结束颜色, ABGR 颜色.
//
// color3: 开始颜色, ABGR 颜色.
//
// color4: 结束颜色, ABGR 颜色.
//
// mode: 模式, GRADIENT_FILL_.
func XDraw_GradientFill4(hDraw int, pRect *RECT, color1 int, color2 int, color3 int, color4 int, mode xcc.GRADIENT_FILL_) bool {
	r, _, _ := xDraw_GradientFill4.Call(uintptr(hDraw), uintptr(unsafe.Pointer(pRect)), uintptr(color1), uintptr(color2), uintptr(color3), uintptr(color4), uintptr(mode))
	return r != 0
}

// 绘制_渐变填充4F, 渐变填充,从一种颜色过渡到另一种颜色.
//
// hDraw: 图形绘制句柄.
//
// pRect: 矩形坐标.
//
// color1: 开始颜色, ABGR 颜色.
//
// color2: 结束颜色, ABGR 颜色.
//
// color3: 开始颜色, ABGR 颜色.
//
// color4: 结束颜色, ABGR 颜色.
//
// mode: 模式, GRADIENT_FILL_.
func XDraw_GradientFill4F(hDraw int, pRect *RECTF, color1 int, color2 int, color3 int, color4 int, mode xcc.GRADIENT_FILL_) bool {
	r, _, _ := xDraw_GradientFill4F.Call(uintptr(hDraw), uintptr(unsafe.Pointer(pRect)), uintptr(color1), uintptr(color2), uintptr(color3), uintptr(color4), uintptr(mode))
	return r != 0
}

// 绘制_边框区域, 绘制边框, 使用指定的画刷绘制指定的区域的边框. 如果函数成功, 返回非零值, 如果函数失败, 返回值是零.
//
// hDraw: 图形绘制句柄.
//
// hrgn: 区域句柄.
//
// hbr: 画刷句柄.
//
// nWidth: 边框宽度, 垂直边.
//
// nHeight: 边框高度, 水平边.
func XDraw_GDI_FrameRgn(hDraw int, hrgn int, hbr int, nWidth int, nHeight int) bool {
	r, _, _ := xDraw_GDI_FrameRgn.Call(uintptr(hDraw), uintptr(hrgn), uintptr(hbr), uintptr(nWidth), uintptr(nHeight))
	return r != 0
}

// 绘制_焦点矩形.
//
// hDraw: 图形绘制句柄.
//
// pRect: 矩形坐标.
func XDraw_FocusRect(hDraw int, pRect *RECT) int {
	r, _, _ := xDraw_FocusRect.Call(uintptr(hDraw), uintptr(unsafe.Pointer(pRect)))
	return int(r)
}

// 绘制_焦点矩形F.
//
// hDraw: 图形绘制句柄.
//
// pRect: 矩形坐标.
func XDraw_FocusRectF(hDraw int, pRect *RECTF) int {
	r, _, _ := xDraw_FocusRectF.Call(uintptr(hDraw), uintptr(unsafe.Pointer(pRect)))
	return int(r)
}

// 绘制_移动到起点, 更新当前位置到指定点，并返回以前的位置. 如果函数成功, 返回非零值.
//
// hDraw: 图形绘制句柄.
//
// X: 坐标.
//
// Y: 坐标.
//
// pPoint: 接收以前的当前位置到一个POINT结构的指针, 如果这个参数是NULL指针, 没有返回原来的位置.
func XDraw_GDI_MoveToEx(hDraw int, X int, Y int, pPoint *POINT) bool {
	r, _, _ := xDraw_GDI_MoveToEx.Call(uintptr(hDraw), uintptr(X), uintptr(Y), uintptr(unsafe.Pointer(pPoint)))
	return r != 0
}

// 绘制_线终点, 函数绘制一条线从当前位置到, 但不包括指定点. 如果函数成功, 返回非零值.
//
// hDraw: 图形绘制句柄.
//
// nXEnd: X坐标, 线结束点.
//
// nYEnd: Y坐标, 线结束点.
func XDraw_GDI_LineTo(hDraw int, nXEnd int, nYEnd int) bool {
	r, _, _ := xDraw_GDI_LineTo.Call(uintptr(hDraw), uintptr(nXEnd), uintptr(nYEnd))
	return r != 0
}

// 绘制_折线, Polyline() 参见MSDN.
//
// hDraw: 图形绘制句柄.
//
// pArrayPt: 参见MSDN.
//
// arrayPtSize: 参见MSDN.
func XDraw_GDI_Polyline(hDraw int, pArrayPt []POINT, arrayPtSize int) bool {
	r, _, _ := xDraw_GDI_Polyline.Call(uintptr(hDraw), uintptr(unsafe.Pointer(&pArrayPt[0])), uintptr(arrayPtSize))
	return r != 0
}

// 绘制_置像素颜色, 函数设置在指定的坐标到指定的颜色的像素. 如果函数成功返回RGB值, 如果失败返回-1.
//
// hDraw: 图形绘制句柄.
//
// X: 坐标.
//
// Y: 坐标.
//
// crColor: RGB颜色值.
func XDraw_GDI_SetPixel(hDraw int, X int, Y int, crColor int) int {
	r, _, _ := xDraw_GDI_SetPixel.Call(uintptr(hDraw), uintptr(X), uintptr(Y), uintptr(crColor))
	return int(r)
}

// 绘制_取D2D渲染目标, 返回 *ID2D1RenderTarget.
//
// hDraw: 图形绘制句柄.
func XDraw_GetD2dRenderTarget(hDraw int) int {
	r, _, _ := xDraw_GetD2dRenderTarget.Call(uintptr(hDraw))
	return int(r)
}

// 绘制_图标, 绘制图标, DrawIconEx()参见MSDN.
//
// hDraw: .
//
// xLeft: .
//
// yTop: .
//
// hIcon: .
//
// cxWidth: .
//
// cyWidth: .
//
// istepIfAniCur: .
//
// hbrFlickerFreeDraw: .
//
// diFlags: .
func XDraw_GDI_DrawIconEx(hDraw int, xLeft int, yTop int, hIcon uintptr, cxWidth int, cyWidth int, istepIfAniCur int, hbrFlickerFreeDraw int, diFlags int) bool {
	r, _, _ := xDraw_GDI_DrawIconEx.Call(uintptr(hDraw), uintptr(xLeft), uintptr(yTop), hIcon, uintptr(cxWidth), uintptr(cyWidth), uintptr(istepIfAniCur), uintptr(hbrFlickerFreeDraw), uintptr(diFlags))
	return r != 0
}

// 绘制_复制, BitBlt() 参见MSDN.
//
// hDrawDest: XX.
//
// nXDest: XX.
//
// nYDest: XX.
//
// nWidth: XX.
//
// nHeight: XX.
//
// hdcSrc: XX.
//
// nXSrc: XX.
//
// nYSrc: XX.
//
// dwRop: XX.
func XDraw_GDI_BitBlt(hDrawDest int, nXDest, nYDest, nWidth, nHeight int32, hdcSrc uintptr, nXSrc, nYSrc int32, dwRop uint32) bool {
	r, _, _ := xDraw_GDI_BitBlt.Call(uintptr(hDrawDest), uintptr(nXDest), uintptr(nYDest), uintptr(nWidth), uintptr(nHeight), hdcSrc, uintptr(nXSrc), uintptr(nYSrc), uintptr(dwRop))
	return r != 0
}

// 绘制_复制2, BitBlt() 参见MSDN.
//
// hDrawDest: XX.
//
// nXDest: XX.
//
// nYDest: XX.
//
// nWidth: XX.
//
// nHeight: XX.
//
// hDrawSrc: XX.
//
// nXSrc: XX.
//
// nYSrc: XX.
//
// dwRop: XX.
func XDraw_GDI_BitBlt2(hDrawDest int, nXDest, nYDest, nWidth, nHeight int32, hDrawSrc uintptr, nXSrc, nYSrc int32, dwRop uint32) bool {
	r, _, _ := xDraw_GDI_BitBlt2.Call(uintptr(hDrawDest), uintptr(nXDest), uintptr(nYDest), uintptr(nWidth), uintptr(nHeight), hDrawSrc, uintptr(nXSrc), uintptr(nYSrc), uintptr(dwRop))
	return r != 0
}

// 绘制_带透明复制, AlphaBlend() 参见MSDN.
//
// hDraw: XX.
//
// nXOriginDest: XX.
//
// nYOriginDest: XX.
//
// nWidthDest: XX.
//
// nHeightDest: XX.
//
// hdcSrc: XX.
//
// nXOriginSrc: XX.
//
// nYOriginSrc: XX.
//
// nWidthSrc: XX.
//
// nHeightSrc: XX.
//
// alpha: XX.
func XDraw_GDI_AlphaBlend(hDraw int, nXOriginDest, nYOriginDest, nWidthDest, nHeightDest int32, hdcSrc uintptr, nXOriginSrc, nYOriginSrc, nWidthSrc, nHeightSrc, alpha int32) bool {
	r, _, _ := xDraw_GDI_AlphaBlend.Call(uintptr(hDraw), uintptr(nXOriginDest), uintptr(nYOriginDest), uintptr(nWidthDest), uintptr(nHeightDest), hdcSrc, uintptr(nXOriginSrc), uintptr(nYOriginSrc), uintptr(nWidthSrc), uintptr(nHeightSrc), uintptr(alpha))
	return r != 0
}

// 绘制_GDI_椭圆.
//
// hDraw: 图形绘制句柄.
//
// pRect: 矩形区域.
func XDraw_GDI_Ellipse(hDraw int, pRect *RECT) bool {
	r, _, _ := xDraw_GDI_Ellipse.Call(uintptr(hDraw), uintptr(unsafe.Pointer(pRect)))
	return r != 0
}

// 绘制_填充多边形, 填充多边形.
//
// hDraw: 图形绘制句柄.
//
// points: 顶点坐标数组.
//
// nCount: 顶点数量.
func XDraw_FillPolygon(hDraw int, points []POINT, nCount int) int {
	r, _, _ := xDraw_FillPolygon.Call(uintptr(hDraw), uintptr(unsafe.Pointer(&points[0])), uintptr(nCount))
	return int(r)
}

// 绘制_填充多边形F, 填充多边形.
//
// hDraw: 图形绘制句柄.
//
// points: 顶点坐标数组.
//
// nCount: 顶点数量.
func XDraw_FillPolygonF(hDraw int, points []POINTF, nCount int) int {
	r, _, _ := xDraw_FillPolygonF.Call(uintptr(hDraw), uintptr(unsafe.Pointer(&points[0])), uintptr(nCount))
	return int(r)
}

// 绘制_图片.
//
// hDraw: 图形绘制句柄.
//
// hImageFrame: 图片句柄.
//
// x: x坐标.
//
// y: y坐标.
func XDraw_Image(hDraw int, hImageFrame int, x, y int32) {
	xDraw_Image.Call(uintptr(hDraw), uintptr(hImageFrame), uintptr(x), uintptr(y))
}

// 绘制_图片F.
//
// hDraw: 图形绘制句柄.
//
// hImageFrame: 图片句柄.
//
// x: x坐标.
//
// y: y坐标.
func XDraw_ImageF(hDraw int, hImageFrame int, x, y float32) int {
	r, _, _ := xDraw_ImageF.Call(uintptr(hDraw), uintptr(hImageFrame), common.Float32Ptr(x), common.Float32Ptr(y))
	return int(r)
}

// 绘制_图片自适应.
//
// hDraw: 图形绘制句柄.
//
// hImageFrame: 图片句柄.
//
// pRect: 坐标.
//
// bOnlyBorder: 是否只绘制边缘区域.
func XDraw_ImageAdaptive(hDraw int, hImageFrame int, pRect *RECT, bOnlyBorder bool) int {
	r, _, _ := xDraw_ImageAdaptive.Call(uintptr(hDraw), uintptr(hImageFrame), uintptr(unsafe.Pointer(pRect)), common.BoolPtr(bOnlyBorder))
	return int(r)
}

// 绘制_图片自适应F.
//
// hDraw: 图形绘制句柄.
//
// hImageFrame: 图片句柄.
//
// pRect: 坐标.
//
// bOnlyBorder: 是否只绘制边缘区域.
func XDraw_ImageAdaptiveF(hDraw int, hImageFrame int, pRect *RECTF, bOnlyBorder bool) int {
	r, _, _ := xDraw_ImageAdaptiveF.Call(uintptr(hDraw), uintptr(hImageFrame), uintptr(unsafe.Pointer(pRect)), common.BoolPtr(bOnlyBorder))
	return int(r)
}

// 绘制_图片扩展, 绘制图片.
//
// hDraw: 图形绘制句柄.
//
// hImageFrame: 图片句柄.
//
// x: x坐标.
//
// y: y坐标.
//
// width: 宽度.
//
// height: 高度.
func XDraw_ImageEx(hDraw int, hImageFrame int, x, y, width, height int) int {
	r, _, _ := xDraw_ImageEx.Call(uintptr(hDraw), uintptr(hImageFrame), uintptr(x), uintptr(y), uintptr(width), uintptr(height))
	return int(r)
}

// 绘制_图片扩展F, 绘制图片.
//
// hDraw: 图形绘制句柄.
//
// hImageFrame: 图片句柄.
//
// x: x坐标.
//
// y: y坐标.
//
// width: 宽度.
//
// height: 高度.
func XDraw_ImageExF(hDraw int, hImageFrame int, x, y, width, height float32) int {
	r, _, _ := xDraw_ImageExF.Call(uintptr(hDraw), uintptr(hImageFrame), common.Float32Ptr(x), common.Float32Ptr(y), common.Float32Ptr(width), common.Float32Ptr(height))
	return int(r)
}

// 绘制_图片增强.
//
// hDraw: 图形绘制句柄.
//
// hImageFrame: 图片句柄.
//
// pRect: 坐标.
//
// bClip: 是否裁剪区域.
func XDraw_ImageSuper(hDraw int, hImageFrame int, pRect *RECT, bClip bool) int {
	r, _, _ := xDraw_ImageSuper.Call(uintptr(hDraw), uintptr(hImageFrame), uintptr(unsafe.Pointer(pRect)), common.BoolPtr(bClip))
	return int(r)
}

// 绘制_图片增强F.
//
// hDraw: 图形绘制句柄.
//
// hImageFrame: 图片句柄.
//
// pRect: 坐标.
//
// bClip: 是否裁剪区域.
func XDraw_ImageSuperF(hDraw int, hImageFrame int, pRect *RECTF, bClip bool) int {
	r, _, _ := xDraw_ImageSuperF.Call(uintptr(hDraw), uintptr(hImageFrame), uintptr(unsafe.Pointer(pRect)), common.BoolPtr(bClip))
	return int(r)
}

// 绘制_图片增强扩展.
//
// hDraw: 图形绘制句柄.
//
// hImageFrame: 图片句柄.
//
// prcDest: 目标坐标.
//
// prcSrc: 源坐标.
func XDraw_ImageSuperEx(hDraw int, hImageFrame int, prcDest *RECT, prcSrc *RECT) int {
	r, _, _ := xDraw_ImageSuperEx.Call(uintptr(hDraw), uintptr(hImageFrame), uintptr(unsafe.Pointer(prcDest)), uintptr(unsafe.Pointer(prcSrc)))
	return int(r)
}

// 绘制_图片增强扩展F.
//
// hDraw: 图形绘制句柄.
//
// hImageFrame: 图片句柄.
//
// prcDest: 目标坐标.
//
// prcSrc: 源坐标.
func XDraw_ImageSuperExF(hDraw int, hImageFrame int, prcDest *RECTF, prcSrc *RECT) int {
	r, _, _ := xDraw_ImageSuperExF.Call(uintptr(hDraw), uintptr(hImageFrame), uintptr(unsafe.Pointer(prcDest)), uintptr(unsafe.Pointer(prcSrc)))
	return int(r)
}

// 绘制_图片增强遮盖, 绘制带遮盖的图片. D2D留空.
//
// hDraw: 图形绘制句柄.
//
// hImageFrame: 图片句柄.
//
// hImageFrameMask: 图片句柄, 遮盖.
//
// pRect: 坐标.
//
// pRectMask: 坐标, 遮盖.
//
// bClip: 是否裁剪区域.
func XDraw_ImageSuperMask(hDraw int, hImageFrame int, hImageFrameMask int, pRect *RECT, pRectMask *RECT, bClip bool) int {
	r, _, _ := xDraw_ImageSuperMask.Call(uintptr(hDraw), uintptr(hImageFrame), uintptr(hImageFrameMask), uintptr(unsafe.Pointer(pRect)), uintptr(unsafe.Pointer(pRectMask)), common.BoolPtr(bClip))
	return int(r)
}

// 绘制_图片平铺, 绘制图片.
//
// hDraw: 图形绘制句柄.
//
// hImageFrame: 图片句柄.
//
// pRect: 坐标.
//
// flag: 标识, 0:从左上角开始平铺, 1:从左下角开始平铺.
func XDraw_ImageTile(hDraw int, hImageFrame int, hImageFrameMask int, pRect *RECT, flag int) int {
	r, _, _ := xDraw_ImageTile.Call(uintptr(hDraw), uintptr(hImageFrame), uintptr(hImageFrameMask), uintptr(unsafe.Pointer(pRect)), uintptr(flag))
	return int(r)
}

// 绘制_图片平铺F, 绘制图片.
//
// hDraw: 图形绘制句柄.
//
// hImageFrame: 图片句柄.
//
// pRect: 坐标.
//
// flag: 标识, 0:从左上角开始平铺, 1:从左下角开始平铺.
func XDraw_ImageTileF(hDraw int, hImageFrame int, hImageFrameMask int, pRect *RECTF, flag int) int {
	r, _, _ := xDraw_ImageTileF.Call(uintptr(hDraw), uintptr(hImageFrame), uintptr(hImageFrameMask), uintptr(unsafe.Pointer(pRect)), uintptr(flag))
	return int(r)
}

// 绘制_图片遮盖, 绘制带遮盖的图片, D2D留空.
//
// hDraw: 图形绘制句柄.
//
// hImageFrame: 图片句柄.
//
// hImageFrameMask: 图片句柄, 遮盖.
//
// x: hImageFrame X坐标.
//
// y: hImageFrame Y坐标.
//
// x2: hImageFrameMask X坐标.
//
// y2: hImageFrameMask Y坐标.
func XDraw_ImageMask(hDraw int, hImageFrame int, hImageFrameMask int, x int, y int, x2 int, y2 int) int {
	r, _, _ := xDraw_ImageMask.Call(uintptr(hDraw), uintptr(hImageFrame), uintptr(hImageFrameMask), uintptr(x), uintptr(y), uintptr(x2), uintptr(y2))
	return int(r)
}

// 绘制_图片遮盖矩形, 使用矩形作为遮罩.
//
// hDraw: 图形绘制句柄.
//
// hImageFrame: 图片句柄.
//
// pRect: 矩形坐标.
//
// pRcMask: 遮罩坐标.
//
// pRcRoundAngle: 遮罩圆角.
func XDraw_ImageMaskRect(hDraw int, hImageFrame int, pRect *RECT, pRcMask *RECT, pRcRoundAngle *RECT) int {
	r, _, _ := xDraw_ImageMaskRect.Call(uintptr(hDraw), uintptr(hImageFrame), uintptr(unsafe.Pointer(pRect)), uintptr(unsafe.Pointer(pRcMask)), uintptr(unsafe.Pointer(pRcRoundAngle)))
	return int(r)
}

// 绘制_图片遮盖圆型, 使用圆形作为遮罩.
//
// hDraw: 图形绘制句柄.
//
// hImageFrame: 图片句柄.
//
// pRect: 矩形坐标.
//
// pRcMask: 遮罩坐标.
func XDraw_ImageMaskEllipse(hDraw int, hImageFrame int, pRect *RECT, pRcMask *RECT) int {
	r, _, _ := xDraw_ImageMaskEllipse.Call(uintptr(hDraw), uintptr(hImageFrame), uintptr(unsafe.Pointer(pRect)), uintptr(unsafe.Pointer(pRcMask)))
	return int(r)
}

// 绘制_文本指定矩形, DrawText() 参见MSDN.
//
// hDraw: 图形绘制句柄.
//
// lpString: 字符串.
//
// lpRect: 坐标.
func XDraw_DrawText(hDraw int, lpString string, lpRect *RECT) int {
	r, _, _ := xDraw_DrawText.Call(uintptr(hDraw), common.StrPtr(lpString), uintptr(len([]rune(lpString))), uintptr(unsafe.Pointer(lpRect)))
	return int(r)
}

// 绘制_文本指定矩形F, DrawText() 参见MSDN.
//
// hDraw: 图形绘制句柄.
//
// lpString: 字符串.
//
// lpRect: 坐标.
func XDraw_DrawTextF(hDraw int, lpString string, lpRect *RECTF) int {
	r, _, _ := xDraw_DrawTextF.Call(uintptr(hDraw), common.StrPtr(lpString), uintptr(len([]rune(lpString))), uintptr(unsafe.Pointer(lpRect)))
	return int(r)
}

// 绘制_文本下划线.
//
// hDraw: 图形绘制句柄.
//
// lpString: 字符串.
//
// lpRect: 坐标.
//
// colorLine: 下划线颜色, ABGR 颜色.
func XDraw_DrawTextUnderline(hDraw int, lpString string, lpRect *RECT, colorLine int) int {
	r, _, _ := xDraw_DrawTextUnderline.Call(uintptr(hDraw), common.StrPtr(lpString), uintptr(len([]rune(lpString))), uintptr(unsafe.Pointer(lpRect)), uintptr(colorLine))
	return int(r)
}

// 绘制_文本下划线F.
//
// hDraw: 图形绘制句柄.
//
// lpString: 字符串.
//
// lpRect: 坐标.
//
// colorLine: 下划线颜色, ABGR 颜色.
func XDraw_DrawTextUnderlineF(hDraw int, lpString string, lpRect *RECTF, colorLine int) int {
	r, _, _ := xDraw_DrawTextUnderlineF.Call(uintptr(hDraw), common.StrPtr(lpString), uintptr(len([]rune(lpString))), uintptr(unsafe.Pointer(lpRect)), uintptr(colorLine))
	return int(r)
}

// 绘制_文本, TextOut() 参见MSDN.
//
// hDraw: 图形绘制句柄.
//
// nXStart: XX.
//
// nYStart: XX.
//
// lpString: XX.
//
// cbString: XX.
func XDraw_TextOut(hDraw int, nXStart int, nYStart int, lpString string, cbString string) int {
	r, _, _ := xDraw_TextOut.Call(uintptr(hDraw), uintptr(nXStart), uintptr(nYStart), common.StrPtr(lpString), common.StrPtr(cbString))
	return int(r)
}

// 绘制_文本F, TextOut() 参见MSDN.
//
// hDraw: 图形绘制句柄.
//
// nXStart: XX.
//
// nYStart: XX.
//
// lpString: XX.
//
// cbString: XX.
func XDraw_TextOutF(hDraw int, nXStart, nYStart float32, lpString string, cbString string) int {
	r, _, _ := xDraw_TextOutF.Call(uintptr(hDraw), common.Float32Ptr(nXStart), common.Float32Ptr(nYStart), common.StrPtr(lpString), common.StrPtr(cbString))
	return int(r)
}

// 绘制_文本扩展, TextOut() 参见MSDN.
//
// hDraw: 图形绘制句柄.
//
// nXStart: XX.
//
// nYStart: XX.
//
// lpString: XX.
func XDraw_TextOutEx(hDraw int, nXStart int, nYStart int, lpString string) int {
	r, _, _ := xDraw_TextOutEx.Call(uintptr(hDraw), uintptr(nXStart), uintptr(nYStart), common.StrPtr(lpString))
	return int(r)
}

// 绘制_文本扩展F, TextOut() 参见MSDN.
//
// hDraw: 图形绘制句柄.
//
// nXStart: XX.
//
// nYStart: XX.
//
// lpString: XX.
func XDraw_TextOutExF(hDraw int, nXStart, nYStart float32, lpString string) int {
	r, _, _ := xDraw_TextOutExF.Call(uintptr(hDraw), common.Float32Ptr(nXStart), common.Float32Ptr(nYStart), common.StrPtr(lpString))
	return int(r)
}

// 绘制_文本A, TextOut() 参见MSDN.
//
// hDraw: 图形绘制句柄.
//
// nXStart: XX.
//
// nYStart: XX.
//
// lpString: XX.
func XDraw_TextOutA(hDraw int, nXStart int, nYStart int, lpString string) int {
	r, _, _ := xDraw_TextOutA.Call(uintptr(hDraw), uintptr(nXStart), uintptr(nYStart), common.StrPtr(lpString))
	return int(r)
}

// 绘制_文本AF, TextOut() 参见MSDN.
//
// hDraw: 图形绘制句柄.
//
// nXStart: XX.
//
// nYStart: XX.
//
// lpString: XX.
func XDraw_TextOutAF(hDraw int, nXStart, nYStart float32, lpString string) int {
	r, _, _ := xDraw_TextOutAF.Call(uintptr(hDraw), common.Float32Ptr(nXStart), common.Float32Ptr(nYStart), common.StrPtr(lpString))
	return int(r)
}

// 绘制_设置文本渲染提示.
//
// hDraw: 图形绘制句柄.
//
// nType: XX.
func XDraw_SetTextRenderingHint(hDraw int, nType int) int {
	r, _, _ := xDraw_SetTextRenderingHint.Call(uintptr(hDraw), uintptr(nType))
	return int(r)
}

// 绘制_SVG源.
//
// hDraw: 图形绘制句柄.
//
// hSvg: SVG句柄.
func XDraw_DrawSvgSrc(hDraw int, hSvg int) int {
	r, _, _ := xDraw_DrawSvgSrc.Call(uintptr(hDraw), uintptr(hSvg))
	return int(r)
}

// 绘制_SVG.
//
// hDraw: 图形绘制句柄.
//
// hSvg: SVG句柄.
//
// x: x坐标.
//
// y: y坐标.
func XDraw_DrawSvg(hDraw int, hSvg int, x int, y int) int {
	r, _, _ := xDraw_DrawSvg.Call(uintptr(hDraw), uintptr(hSvg), uintptr(x), uintptr(y))
	return int(r)
}

// 绘制_SVG扩展.
//
// hDraw: 图形绘制句柄.
//
// hSvg: SVG句柄.
//
// x: x坐标.
//
// y: y坐标.
//
// nWidth: 宽度.
//
// nHeight: 高度.
func XDraw_DrawSvgEx(hDraw int, hSvg int, x int, y int, nWidth int, nHeight int) int {
	r, _, _ := xDraw_DrawSvgEx.Call(uintptr(hDraw), uintptr(hSvg), uintptr(x), uintptr(y), uintptr(nWidth), uintptr(nHeight))
	return int(r)
}

// 绘制_SVG大小.
//
// hDraw: 图形绘制句柄.
//
// hSvg: SVG句柄.
//
// nWidth: 宽度.
//
// nHeight: 高度.
func XDraw_DrawSvgSize(hDraw int, hSvg int, nWidth int, nHeight int) int {
	r, _, _ := xDraw_DrawSvgSize.Call(uintptr(hDraw), uintptr(hSvg), uintptr(nWidth), uintptr(nHeight))
	return int(r)
}

// 绘制_D2D_清理, 使用指定颜色清理画布.
//
// hDraw: 图形绘制句柄.
//
// color: ABGR 颜色值.
func XDraw_D2D_Clear(hDraw int, color int) int {
	r, _, _ := xDraw_D2D_Clear.Call(uintptr(hDraw), uintptr(color))
	return int(r)
}

// 绘制_取字体, 返回字体句柄.
//
// hDraw: 图形绘制句柄.
func XDraw_GetFont(hDraw int) int {
	r, _, _ := xDraw_GetFont.Call(uintptr(hDraw))
	return int(r)
}
