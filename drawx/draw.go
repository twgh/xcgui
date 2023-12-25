package drawx

import (
	"github.com/twgh/xcgui/objectbase"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

// Draw 图形绘制.
type Draw struct {
	objectbase.ObjectBase
}

// New 绘制_创建, 创建图形绘制模块实例, 返回句柄.
//
// hWindow: 窗口句柄.
func New(hWindow int) *Draw {
	p := &Draw{}
	p.SetHandle(xc.XDraw_Create(hWindow))
	return p
}

// NewGDI 绘制_创建GDI, 创建图形绘制模块实例, 返回图形绘制模块实例句柄.
//
// hWindow: 窗口句柄.
//
// hdc: hdc句柄.
func NewGDI(hWindow int, hdc uintptr) *Draw {
	p := &Draw{}
	p.SetHandle(xc.XDraw_CreateGDI(hWindow, hdc))
	return p
}

// NewByHandle 从图形绘制模块实例句柄创建对象.
func NewByHandle(handle int) *Draw {
	p := &Draw{}
	p.SetHandle(handle)
	return p
}

// 绘制_销毁, 销毁图形绘制模块实例句柄.
func (d *Draw) Destroy() int {
	return xc.XDraw_Destroy(d.Handle)
}

// 绘制_虚线, 绘制水平或垂直虚线.
//
// x1: 起点x坐标.
//
// y1: 起点y坐标.
//
// x2: 结束点x坐标.
//
// y2: 结束点y坐标.
func (d *Draw) Dottedline(x1 int, y1 int, x2 int, y2 int) int {
	return xc.XDraw_Dottedline(d.Handle, x1, y1, x2, y2)
}

// 绘制_虚线F, 绘制水平或垂直虚线.
//
// x1: 起点x坐标.
//
// y1: 起点y坐标.
//
// x2: 结束点x坐标.
//
// y2: 结束点y坐标.
func (d *Draw) DottedlineF(x1, y1, x2, y2 float32) int {
	return xc.XDraw_DottedlineF(d.Handle, x1, y1, x2, y2)
}

// 绘制_圆弧.
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
func (d *Draw) DrawArc(x, y int, nWidth int, nHeight int, startAngle float32, sweepAngle float32) int {
	return xc.XDraw_DrawArc(d.Handle, x, y, nWidth, nHeight, startAngle, sweepAngle)
}

// 绘制_圆弧F.
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
func (d *Draw) DrawArcF(x, y, nWidth, nHeight, startAngle, sweepAngle float32) int {
	return xc.XDraw_DrawArcF(d.Handle, x, y, nWidth, nHeight, startAngle, sweepAngle)
}

// 绘制_曲线, D2D暂时留空.
//
// points: 坐标点数组.
//
// count: 数组大小.
//
// tension: 大于或等于0.0F的值，指定曲线的张力, D2D 忽略此参数。.
func (d *Draw) DrawCurve(points []xc.POINT, count int, tension float32) int {
	return xc.XDraw_DrawCurve(d.Handle, points, count, tension)
}

// 绘制_曲线F, D2D暂时留空.
//
// points: 坐标点数组.
//
// count: 数组大小.
//
// tension: 大于或等于0.0F的值，指定曲线的张力, D2D 忽略此参数。.
func (d *Draw) DrawCurveF(points []xc.POINTF, count int, tension float32) int {
	return xc.XDraw_DrawCurveF(d.Handle, points, count, tension)
}

// 绘制_线条.
//
// x1: 坐标.
//
// y1: 坐标.
//
// x2: 坐标.
//
// y2: 坐标.
func (d *Draw) DrawLine(x1 int, y1 int, x2 int, y2 int) int {
	return xc.XDraw_DrawLine(d.Handle, x1, y1, x2, y2)
}

// 绘制_线条F.
//
// x1: 坐标.
//
// y1: 坐标.
//
// x2: 坐标.
//
// y2: 坐标.
func (d *Draw) DrawLineF(x1, y1, x2, y2 float32) int {
	return xc.XDraw_DrawLineF(d.Handle, x1, y1, x2, y2)
}

// 绘制_多边形, 绘制多边形.
//
// points: 顶点坐标数组.
//
// nCount: 顶点数量.
func (d *Draw) DrawPolygon(points []xc.POINT, nCount int) int {
	return xc.XDraw_DrawPolygon(d.Handle, points, nCount)
}

// 绘制_多边形F, 绘制多边形.
//
// points: 顶点坐标数组.
//
// nCount: 顶点数量.
func (d *Draw) DrawPolygonF(points []xc.POINTF, nCount int) int {
	return xc.XDraw_DrawPolygonF(d.Handle, points, nCount)
}

// 绘制_矩形, 绘制矩形边框.
//
// pRect: 矩形坐标 .
func (d *Draw) DrawRect(pRect *xc.RECT) int {
	return xc.XDraw_DrawRect(d.Handle, pRect)
}

// 绘制_矩形F, 绘制矩形边框.
//
// pRect: 矩形坐标 .
func (d *Draw) DrawRectF(pRect *xc.RECTF) int {
	return xc.XDraw_DrawRectF(d.Handle, pRect)
}

// 绘制_置偏移, 设置坐标偏移量, X向左偏移为负数, 向右偏移为正数.
//
// x: X轴偏移量.
//
// y: Y轴偏移量.
func (d *Draw) SetOffset(x, y int32) int {
	return xc.XDraw_SetOffset(d.Handle, x, y)
}

// 绘制_取偏移, 获取坐标偏移量, X向左偏移为负数, 向右偏移为正数.
//
// pX: 接收X轴偏移量.
//
// pY: 接收Y轴偏移量.
func (d *Draw) GetOffset(pX, pY *int32) int {
	return xc.XDraw_GetOffset(d.Handle, pX, pY)
}

// 绘制_还原状态, 还原状态, 释放用户绑定的GDI对象, 例如画刷, 画笔.
func (d *Draw) GDI_RestoreGDIOBJ() int {
	return xc.XDraw_GDI_RestoreGDIOBJ(d.Handle)
}

// 绘制_取HDC, 获取绑定的设备上下文HDC, 返回HDC句柄.
func (d *Draw) GetHDC() uintptr {
	return xc.XDraw_GetHDC(d.Handle)
}

// 绘制_置画刷颜色, 设置画刷颜色.
//
// color: ABGR 颜色值.
func (d *Draw) SetBrushColor(color int) int {
	return xc.XDraw_SetBrushColor(d.Handle, color)
}

// 绘制_置文本垂直, 设置文本垂直显示.
//
// bVertical: 是否垂直显示文本.
func (d *Draw) SetTextVertical(bVertical bool) int {
	return xc.XDraw_SetTextVertical(d.Handle, bVertical)
}

// 绘制_置文本对齐, 设置文本对齐.
//
// nFlags: 对齐标识, TextFormatFlag_, TextAlignFlag_, TextTrimming_.
func (d *Draw) SetTextAlign(nFlags xcc.TextFormatFlag_) int {
	return xc.XDraw_SetTextAlign(d.Handle, nFlags)
}

// 绘制_置字体.
//
// hFontx: 炫彩字体.
func (d *Draw) SetFont(hFontx int) int {
	return xc.XDraw_SetFont(d.Handle, hFontx)
}

// 绘制_置线宽.
//
// nWidth: 宽度.
func (d *Draw) SetLineWidth(nWidth int) int {
	return xc.XDraw_SetLineWidth(d.Handle, nWidth)
}

// 绘制_置线宽F.
//
// nWidth: 宽度.
func (d *Draw) SetLineWidthF(nWidth float32) int {
	return xc.XDraw_SetLineWidthF(d.Handle, nWidth)
}

// 绘制_置背景模式, SetBkMode() 参见MSDN.
//
// bTransparent: 参见MSDN.
func (d *Draw) GDI_SetBkMode(bTransparent bool) int {
	return xc.XDraw_GDI_SetBkMode(d.Handle, bTransparent)
}

// 绘制_置裁剪区域, 设置裁剪区域.
//
// pRect: 区域坐标.
func (d *Draw) SetClipRect(pRect *xc.RECT) int {
	return xc.XDraw_SetClipRect(d.Handle, pRect)
}

// 绘制_置D2D文本渲染模式.
//
// mode	渲染模式 XC_DWRITE_RENDERING_MODE_.
func (d *Draw) SetD2dTextRenderingMode(mode xcc.XC_DWRITE_RENDERING_MODE_) int {
	return xc.XDraw_SetD2dTextRenderingMode(d.Handle, mode)
}

// 绘制_清除裁剪区域.
func (d *Draw) ClearClip() int {
	return xc.XDraw_ClearClip(d.Handle)
}

// 绘制_启用平滑模式.
//
// bEnable: 是否启用.
func (d *Draw) EnableSmoothingMode(bEnable bool) int {
	return xc.XDraw_EnableSmoothingMode(d.Handle, bEnable)
}

// 绘制_启用窗口透明判断, 当启用之后, 调用GDI+函数时, 如果参数alpha=255, 将自动修改为254, 应对GDI+的bug, 否则透明通道异常.
//
// bTransparent: 是否启用.
func (d *Draw) EnableWndTransparent(bTransparent bool) int {
	return xc.XDraw_EnableWndTransparent(d.Handle, bTransparent)
}

// 绘制_创建实心画刷, GDI创建具有指定的纯色逻辑刷.
//
// crColor: 画刷颜色.
func (d *Draw) GDI_CreateSolidBrush(crColor int) int {
	return xc.XDraw_GDI_CreateSolidBrush(d.Handle, crColor)
}

// 绘制_创建画笔, GDI创建一个逻辑笔, 指定的样式, 宽度和颜色, 随后的笔可以选择到设备上下文, 用于绘制线条和曲线.
//
// fnPenStyle: 画笔样式, PS_SOLID:实线, PS_DASH:段线, PS_DOT:点线, PS_DASHDOT:段线_点线, PS_DASHDOTDOT:段线_点_点, PS_NULL:空, PS_INSIDEFRAME:实线_笔宽是向里扩展.
//
// nWidth: 画笔宽度.
//
// crColor: ABGR 颜色.
func (d *Draw) GDI_CreatePen(fnPenStyle int, nWidth int, crColor int) int {
	return xc.XDraw_GDI_CreatePen(d.Handle, fnPenStyle, nWidth, crColor)
}

// 绘制_创建矩形区域, GDI创建矩形区域, 成功返回区域句柄, 失败返回NULL.
//
// nLeftRect: 左上角X坐标.
//
// nTopRect: 左上角Y坐标.
//
// nRightRect: 右下角X坐标.
//
// nBottomRect: 右下角Y坐标.
func (d *Draw) GDI_CreateRectRgn(nLeftRect int, nTopRect int, nRightRect int, nBottomRect int) int {
	return xc.XDraw_GDI_CreateRectRgn(d.Handle, nLeftRect, nTopRect, nRightRect, nBottomRect)
}

// 绘制_创建圆角矩形区域, GDI创建一个圆角的矩形区域, 成功返回区域句柄, 失败返回NULL.
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
func (d *Draw) GDI_CreateRoundRectRgn(nLeftRect int, nTopRect int, nRightRect int, nBottomRect int, nWidthEllipse int, nHeightEllipse int) int {
	return xc.XDraw_GDI_CreateRoundRectRgn(d.Handle, nLeftRect, nTopRect, nRightRect, nBottomRect, nWidthEllipse, nHeightEllipse)
}

// 绘制_创建多边形区域, GDI创建一个多边形区域, 成功返回区域句柄, 失败返回NULL.
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
func (d *Draw) GDI_CreatePolygonRgn(pPt []xc.POINT, cPoints, fnPolyFillMode int) int {
	return xc.XDraw_GDI_CreatePolygonRgn(d.Handle, pPt, cPoints, fnPolyFillMode)
}

// 绘制_GDI_椭圆.
//
// pRect: 矩形区域.
func (d *Draw) GDI_Ellipse(pRect *xc.RECT) bool {
	return xc.XDraw_GDI_Ellipse(d.Handle, pRect)
}

// 绘制_选择裁剪区域, 选择一个区域作为当前裁剪区域, 注意: 该函数只对GDI有效.
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
func (d *Draw) GDI_SelectClipRgn(hRgn int) int {
	return xc.XDraw_GDI_SelectClipRgn(d.Handle, hRgn)
}

// 绘制_填充矩形, 通过使用指定的刷子填充一个矩形, 此功能包括左侧和顶部的边界, 但不包括矩形的右边和底部边界.
//
// pRect: 矩形区域.
func (d *Draw) FillRect(pRect *xc.RECT) int {
	return xc.XDraw_FillRect(d.Handle, pRect)
}

// 绘制_填充矩形F, 通过使用指定的刷子填充一个矩形, 此功能包括左侧和顶部的边界, 但不包括矩形的右边和底部边界.
//
// pRect: 矩形区域.
func (d *Draw) FillRectF(pRect *xc.RECTF) int {
	return xc.XDraw_FillRectF(d.Handle, pRect)
}

// 绘制_填充矩形指定颜色.
//
// pRect: 矩形区域.
//
// color: ABGR 颜色.
func (d *Draw) FillRectColor(pRect *xc.RECT, color int) int {
	return xc.XDraw_FillRectColor(d.Handle, pRect, color)
}

// 绘制_填充矩形指定颜色F.
//
// pRect: 矩形区域.
//
// color: ABGR 颜色.
func (d *Draw) FillRectColorF(pRect *xc.RECTF, color int) int {
	return xc.XDraw_FillRectColorF(d.Handle, pRect, color)
}

// 绘制_填充区域, 通过使用指定的画刷填充一个区域.
//
// hrgn: 区域句柄.
//
// hbr: 画刷句柄.
func (d *Draw) GDI_FillRgn(hrgn int, hbr int) bool {
	return xc.XDraw_GDI_FillRgn(d.Handle, hrgn, hbr)
}

// 绘制_填充圆形.
//
// pRect: 矩形区域.
func (d *Draw) FillEllipse(pRect *xc.RECT) int {
	return xc.XDraw_FillEllipse(d.Handle, pRect)
}

// 绘制_填充圆形F.
//
// pRect: 矩形区域.
func (d *Draw) FillEllipseF(pRect *xc.RECTF) int {
	return xc.XDraw_FillEllipseF(d.Handle, pRect)
}

// 绘制_圆形, 绘制圆边框.
//
// pRect: 矩形区域.
func (d *Draw) DrawEllipse(pRect *xc.RECT) int {
	return xc.XDraw_DrawEllipse(d.Handle, pRect)
}

// 绘制_填充圆角矩形.
//
// pRect: 矩形坐标.
//
// nWidth: 圆角宽度.
//
// nHeight: 圆角高度.
func (d *Draw) FillRoundRect(pRect *xc.RECT, nWidth, nHeight int) int {
	return xc.XDraw_FillRoundRect(d.Handle, pRect, nWidth, nHeight)
}

// 绘制_填充圆角矩形F.
//
// pRect: 矩形坐标.
//
// nWidth: 圆角宽度.
//
// nHeight: 圆角高度.
func (d *Draw) FillRoundRectF(pRect *xc.RECTF, nWidth, nHeight float32) int {
	return xc.XDraw_FillRoundRectF(d.Handle, pRect, nWidth, nHeight)
}

// 绘制_圆角矩形, 绘制圆角矩形边框.
//
// pRect: 矩形坐标.
//
// nWidth: 圆角宽度.
//
// nHeight: 圆角高度.
func (d *Draw) DrawRoundRect(pRect *xc.RECT, nWidth int, nHeight int) int {
	return xc.XDraw_DrawRoundRect(d.Handle, pRect, nWidth, nHeight)
}

// 绘制_圆角矩形F, 绘制圆角矩形边框.
//
// pRect: 矩形坐标.
//
// nWidth: 圆角宽度.
//
// nHeight: 圆角高度.
func (d *Draw) DrawRoundRectF(pRect *xc.RECT, nWidth, nHeight float32) int {
	return xc.XDraw_DrawRoundRectF(d.Handle, pRect, nWidth, nHeight)
}

// 绘制_填充圆角矩形扩展.
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
func (d *Draw) FillRoundRectEx(pRect *xc.RECT, nLeftTop, nRightTop, nRightBottom, nLeftBottom int) int {
	return xc.XDraw_FillRoundRectEx(d.Handle, pRect, nLeftTop, nRightTop, nRightBottom, nLeftBottom)
}

// 绘制_填充圆角矩形扩展F.
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
func (d *Draw) FillRoundRectExF(pRect *xc.RECTF, nLeftTop, nRightTop, nRightBottom, nLeftBottom float32) int {
	return xc.XDraw_FillRoundRectExF(d.Handle, pRect, nLeftTop, nRightTop, nRightBottom, nLeftBottom)
}

// 绘制_圆角矩形扩展, 绘制圆角矩形边框.
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
func (d *Draw) DrawRoundRectEx(pRect *xc.RECT, nLeftTop int, nRightTop int, nRightBottom int, nLeftBottom int) int {
	return xc.XDraw_DrawRoundRectEx(d.Handle, pRect, nLeftTop, nRightTop, nRightBottom, nLeftBottom)
}

// 绘制_圆角矩形扩展F, 绘制圆角矩形边框.
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
func (d *Draw) DrawRoundRectExF(pRect *xc.RECT, nLeftTop, nRightTop, nRightBottom, nLeftBottom float32) int {
	return xc.XDraw_DrawRoundRectExF(d.Handle, pRect, nLeftTop, nRightTop, nRightBottom, nLeftBottom)
}

// 绘制_矩形, 绘制矩形, 使用当前的画刷和画笔. 如果函数成功, 返回非零值, 如果函数失败, 返回值是零.
//
// nLeftRect: 左上角X坐标.
//
// nTopRect: 左上角Y坐标.
//
// nRightRect: 右下角X坐标.
//
// nBottomRect: 右下角Y坐标.
func (d *Draw) GDI_Rectangle(nLeftRect int, nTopRect int, nRightRect int, nBottomRect int) bool {
	return xc.XDraw_GDI_Rectangle(d.Handle, nLeftRect, nTopRect, nRightRect, nBottomRect)
}

// 绘制_渐变填充2, 渐变填充, 从一种颜色过渡到另一种颜色.
//
// pRect: 矩形坐标.
//
// color1: 开始颜色, ABGR 颜色.
//
// color2: 结束颜色, ABGR 颜色.
//
// mode: 模式, GRADIENT_FILL_.
func (d *Draw) GradientFill2(pRect *xc.RECT, color1 int, color2 int, mode xcc.GRADIENT_FILL_) int {
	return xc.XDraw_GradientFill2(d.Handle, pRect, color1, color2, mode)
}

// 绘制_渐变填充2F, 渐变填充, 从一种颜色过渡到另一种颜色.
//
// pRect: 矩形坐标.
//
// color1: 开始颜色, ABGR 颜色.
//
// color2: 结束颜色, ABGR 颜色.
//
// mode: 模式, GRADIENT_FILL_.
func (d *Draw) GradientFill2F(pRect *xc.RECTF, color1 int, color2 int, mode xcc.GRADIENT_FILL_) int {
	return xc.XDraw_GradientFill2F(d.Handle, pRect, color1, color2, mode)
}

// 绘制_渐变填充4, 渐变填充,从一种颜色过渡到另一种颜色.
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
func (d *Draw) GradientFill4(pRect *xc.RECT, color1 int, color2 int, color3 int, color4 int, mode xcc.GRADIENT_FILL_) bool {
	return xc.XDraw_GradientFill4(d.Handle, pRect, color1, color2, color3, color4, mode)
}

// 绘制_渐变填充4F, 渐变填充,从一种颜色过渡到另一种颜色.
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
func (d *Draw) GradientFill4F(pRect *xc.RECTF, color1 int, color2 int, color3 int, color4 int, mode xcc.GRADIENT_FILL_) bool {
	return xc.XDraw_GradientFill4F(d.Handle, pRect, color1, color2, color3, color4, mode)
}

// 绘制_边框区域, 绘制边框, 使用指定的画刷绘制指定的区域的边框. 如果函数成功, 返回非零值, 如果函数失败, 返回值是零.
//
// hrgn: 区域句柄.
//
// hbr: 画刷句柄.
//
// nWidth: 边框宽度, 垂直边.
//
// nHeight: 边框高度, 水平边.
func (d *Draw) GDI_FrameRgn(hrgn int, hbr int, nWidth int, nHeight int) bool {
	return xc.XDraw_GDI_FrameRgn(d.Handle, hrgn, hbr, nWidth, nHeight)
}

// 绘制_焦点矩形.
//
// pRect: 矩形坐标.
func (d *Draw) FocusRect(pRect *xc.RECT) int {
	return xc.XDraw_FocusRect(d.Handle, pRect)
}

// 绘制_焦点矩形F.
//
// pRect: 矩形坐标.
func (d *Draw) FocusRectF(pRect *xc.RECTF) int {
	return xc.XDraw_FocusRectF(d.Handle, pRect)
}

// 绘制_移动到起点, 更新当前位置到指定点，并返回以前的位置. 如果函数成功, 返回非零值.
//
// X: 坐标.
//
// Y: 坐标.
//
// pPoint: 接收以前的当前位置到一个POINT结构的指针, 如果这个参数是NULL指针, 没有返回原来的位置.
func (d *Draw) GDI_MoveToEx(X int, Y int, pPoint *xc.POINT) bool {
	return xc.XDraw_GDI_MoveToEx(d.Handle, X, Y, pPoint)
}

// 绘制_线终点, 函数绘制一条线从当前位置到, 但不包括指定点. 如果函数成功, 返回非零值.
//
// nXEnd: X坐标, 线结束点.
//
// nYEnd: Y坐标, 线结束点.
func (d *Draw) GDI_LineTo(nXEnd int, nYEnd int) bool {
	return xc.XDraw_GDI_LineTo(d.Handle, nXEnd, nYEnd)
}

// 绘制_折线, Polyline() 参见MSDN.
//
// pArrayPt: 参见MSDN.
//
// arrayPtSize: 参见MSDN.
func (d *Draw) GDI_Polyline(pArrayPt []xc.POINT, arrayPtSize int) bool {
	return xc.XDraw_GDI_Polyline(d.Handle, pArrayPt, arrayPtSize)
}

// 绘制_置像素颜色, 函数设置在指定的坐标到指定的颜色的像素. 如果函数成功返回RGB值, 如果失败返回-1.
//
// X: 坐标.
//
// Y: 坐标.
//
// crColor: RGB颜色值.
func (d *Draw) GDI_SetPixel(X int, Y int, crColor int) int {
	return xc.XDraw_GDI_SetPixel(d.Handle, X, Y, crColor)
}

// 绘制_取D2D渲染目标, 返回 *ID2D1RenderTarget.
func (d *Draw) XDraw_GetD2dRenderTarget() int {
	return xc.XDraw_GetD2dRenderTarget(d.Handle)
}

// 绘制_图标, 绘制图标, DrawIconEx()参见MSDN.
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
func (d *Draw) GDI_DrawIconEx(xLeft int, yTop int, hIcon uintptr, cxWidth int, cyWidth int, istepIfAniCur int, hbrFlickerFreeDraw int, diFlags int) bool {
	return xc.XDraw_GDI_DrawIconEx(d.Handle, xLeft, yTop, hIcon, cxWidth, cyWidth, istepIfAniCur, hbrFlickerFreeDraw, diFlags)
}

// 绘制_复制, BitBlt() 参见MSDN.
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
func (d *Draw) GDI_BitBlt(nXDest, nYDest, nWidth, nHeight int32, hdcSrc uintptr, nXSrc, nYSrc int32, dwRop uint32) bool {
	return xc.XDraw_GDI_BitBlt(d.Handle, nXDest, nYDest, nWidth, nHeight, hdcSrc, nXSrc, nYSrc, dwRop)
}

// 绘制_复制2, BitBlt() 参见MSDN.
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
func (d *Draw) GDI_BitBlt2(nXDest, nYDest, nWidth, nHeight int32, hDrawSrc uintptr, nXSrc, nYSrc int32, dwRop uint32) bool {
	return xc.XDraw_GDI_BitBlt2(d.Handle, nXDest, nYDest, nWidth, nHeight, hDrawSrc, nXSrc, nYSrc, dwRop)
}

// 绘制_带透明复制, AlphaBlend() 参见MSDN.
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
func (d *Draw) GDI_AlphaBlend(nXOriginDest, nYOriginDest, nWidthDest, nHeightDest int32, hdcSrc uintptr, nXOriginSrc, nYOriginSrc, nWidthSrc, nHeightSrc, alpha int32) bool {
	return xc.XDraw_GDI_AlphaBlend(d.Handle, nXOriginDest, nYOriginDest, nWidthDest, nHeightDest, hdcSrc, nXOriginSrc, nYOriginSrc, nWidthSrc, nHeightSrc, alpha)
}

// 绘制_填充多边形, 填充多边形.
//
// points: 顶点坐标数组.
//
// nCount: 顶点数量.
func (d *Draw) FillPolygon(points []xc.POINT, nCount int) int {
	return xc.XDraw_FillPolygon(d.Handle, points, nCount)
}

// 绘制_填充多边形F, 填充多边形.
//
// points: 顶点坐标数组.
//
// nCount: 顶点数量.
func (d *Draw) FillPolygonF(points []xc.POINTF, nCount int) int {
	return xc.XDraw_FillPolygonF(d.Handle, points, nCount)
}

// 绘制_图片.
//
// hImageFrame: 图片句柄.
//
// x: x坐标.
//
// y: y坐标.
func (d *Draw) Image(hImageFrame int, x, y int32) {
	xc.XDraw_Image(d.Handle, hImageFrame, x, y)
}

// 绘制_图片F.
//
// hImageFrame: 图片句柄.
//
// x: x坐标.
//
// y: y坐标.
func (d *Draw) ImageF(hImageFrame int, x, y float32) int {
	return xc.XDraw_ImageF(d.Handle, hImageFrame, x, y)
}

// 绘制_图片自适应.
//
// hImageFrame: 图片句柄.
//
// pRect: 坐标.
//
// bOnlyBorder: 是否只绘制边缘区域.
func (d *Draw) ImageAdaptive(hImageFrame int, pRect *xc.RECT, bOnlyBorder bool) int {
	return xc.XDraw_ImageAdaptive(d.Handle, hImageFrame, pRect, bOnlyBorder)
}

// 绘制_图片自适应F.
//
// hImageFrame: 图片句柄.
//
// pRect: 坐标.
//
// bOnlyBorder: 是否只绘制边缘区域.
func (d *Draw) ImageAdaptiveF(hImageFrame int, pRect *xc.RECTF, bOnlyBorder bool) int {
	return xc.XDraw_ImageAdaptiveF(d.Handle, hImageFrame, pRect, bOnlyBorder)
}

// 绘制_图片扩展, 绘制图片.
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
func (d *Draw) XDraw_ImageEx(hImageFrame int, x, y, width, height int) int {
	return xc.XDraw_ImageEx(d.Handle, hImageFrame, x, y, width, height)
}

// 绘制_图片扩展F, 绘制图片.
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
func (d *Draw) XDraw_ImageExF(hImageFrame int, x, y, width, height float32) int {
	return xc.XDraw_ImageExF(d.Handle, hImageFrame, x, y, width, height)
}

// 绘制_图片增强.
//
// hImageFrame: 图片句柄.
//
// pRect: 坐标.
//
// bClip: 是否裁剪区域.
func (d *Draw) ImageSuper(hImageFrame int, pRect *xc.RECT, bClip bool) int {
	return xc.XDraw_ImageSuper(d.Handle, hImageFrame, pRect, bClip)
}

// 绘制_图片增强F.
//
// hImageFrame: 图片句柄.
//
// pRect: 坐标.
//
// bClip: 是否裁剪区域.
func (d *Draw) ImageSuperF(hImageFrame int, pRect *xc.RECTF, bClip bool) int {
	return xc.XDraw_ImageSuperF(d.Handle, hImageFrame, pRect, bClip)
}

// 绘制_图片增强扩展.
//
// hImageFrame: 图片句柄.
//
// prcDest: 目标坐标.
//
// prcSrc: 源坐标.
func (d *Draw) ImageSuperEx(hImageFrame int, prcDest *xc.RECT, prcSrc *xc.RECT) int {
	return xc.XDraw_ImageSuperEx(d.Handle, hImageFrame, prcDest, prcSrc)
}

// 绘制_图片增强扩展F.
//
// hImageFrame: 图片句柄.
//
// prcDest: 目标坐标.
//
// prcSrc: 源坐标.
func (d *Draw) ImageSuperExF(hImageFrame int, prcDest *xc.RECTF, prcSrc *xc.RECT) int {
	return xc.XDraw_ImageSuperExF(d.Handle, hImageFrame, prcDest, prcSrc)
}

// 绘制_图片增强遮盖, 绘制带遮盖的图片. D2D留空.
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
func (d *Draw) ImageSuperMask(hImageFrame int, hImageFrameMask int, pRect *xc.RECT, pRectMask *xc.RECT, bClip bool) int {
	return xc.XDraw_ImageSuperMask(d.Handle, hImageFrame, hImageFrameMask, pRect, pRectMask, bClip)
}

// 绘制_图片平铺, 绘制图片.
//
// hImageFrame: 图片句柄.
//
// pRect: 坐标.
//
// flag: 标识, 0:从左上角开始平铺, 1:从左下角开始平铺.
func (d *Draw) ImageTile(hImageFrame int, hImageFrameMask int, pRect *xc.RECT, flag int) int {
	return xc.XDraw_ImageTile(d.Handle, hImageFrame, hImageFrameMask, pRect, flag)
}

// 绘制_图片平铺F, 绘制图片.
//
// hImageFrame: 图片句柄.
//
// pRect: 坐标.
//
// flag: 标识, 0:从左上角开始平铺, 1:从左下角开始平铺.
func (d *Draw) ImageTileF(hImageFrame int, hImageFrameMask int, pRect *xc.RECTF, flag int) int {
	return xc.XDraw_ImageTileF(d.Handle, hImageFrame, hImageFrameMask, pRect, flag)
}

// 绘制_图片遮盖, 绘制带遮盖的图片, D2D留空.
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
func (d *Draw) ImageMask(hImageFrame int, hImageFrameMask int, x int, y int, x2 int, y2 int) int {
	return xc.XDraw_ImageMask(d.Handle, hImageFrame, hImageFrameMask, x, y, x2, y2)
}

// 绘制_文本指定矩形, DrawText() 参见MSDN.
//
// lpString: 字符串.
//
// lpRect: 坐标.
func (d *Draw) DrawText(lpString string, lpRect *xc.RECT) int {
	return xc.XDraw_DrawText(d.Handle, lpString, lpRect)
}

// 绘制_文本指定矩形F, DrawText() 参见MSDN.
//
// lpString: 字符串.
//
// lpRect: 坐标.
func (d *Draw) DrawTextF(lpString string, lpRect *xc.RECTF) int {
	return xc.XDraw_DrawTextF(d.Handle, lpString, lpRect)
}

// 绘制_文本下划线.
//
// lpString: 字符串.
//
// lpRect: 坐标.
//
// colorLine: 下划线颜色, ABGR 颜色.
func (d *Draw) DrawTextUnderline(lpString string, lpRect *xc.RECT, colorLine int) int {
	return xc.XDraw_DrawTextUnderline(d.Handle, lpString, lpRect, colorLine)
}

// 绘制_文本下划线F.
//
// lpString: 字符串.
//
// lpRect: 坐标.
//
// colorLine: 下划线颜色, ABGR 颜色.
func (d *Draw) DrawTextUnderlineF(lpString string, lpRect *xc.RECTF, colorLine int) int {
	return xc.XDraw_DrawTextUnderlineF(d.Handle, lpString, lpRect, colorLine)
}

// 绘制_文本, TextOut() 参见MSDN.
//
// nXStart: XX.
//
// nYStart: XX.
//
// lpString: XX.
//
// cbString: XX.
func (d *Draw) TextOut(nXStart int, nYStart int, lpString string, cbString string) int {
	return xc.XDraw_TextOut(d.Handle, nXStart, nYStart, lpString, cbString)
}

// 绘制_文本F, TextOut() 参见MSDN.
//
// nXStart: XX.
//
// nYStart: XX.
//
// lpString: XX.
//
// cbString: XX.
func (d *Draw) TextOutF(nXStart, nYStart float32, lpString string, cbString string) int {
	return xc.XDraw_TextOutF(d.Handle, nXStart, nYStart, lpString, cbString)
}

// 绘制_文本扩展, TextOut() 参见MSDN.
//
// nXStart: XX.
//
// nYStart: XX.
//
// lpString: XX.
func (d *Draw) TextOutEx(nXStart int, nYStart int, lpString string) int {
	return xc.XDraw_TextOutEx(d.Handle, nXStart, nYStart, lpString)
}

// 绘制_文本扩展F, TextOut() 参见MSDN.
//
// nXStart: XX.
//
// nYStart: XX.
//
// lpString: XX.
func (d *Draw) TextOutExF(nXStart, nYStart float32, lpString string) int {
	return xc.XDraw_TextOutExF(d.Handle, nXStart, nYStart, lpString)
}

// 绘制_文本A, TextOut() 参见MSDN.
//
// nXStart: XX.
//
// nYStart: XX.
//
// lpString: XX.
func (d *Draw) TextOutA(nXStart int, nYStart int, lpString string) int {
	return xc.XDraw_TextOutA(d.Handle, nXStart, nYStart, lpString)
}

// 绘制_文本AF, TextOut() 参见MSDN.
//
// nXStart: XX.
//
// nYStart: XX.
//
// lpString: XX.
func (d *Draw) TextOutAF(nXStart, nYStart float32, lpString string) int {
	return xc.XDraw_TextOutAF(d.Handle, nXStart, nYStart, lpString)
}

// 绘制_设置文本渲染提示.
//
// nType: XX.
func (d *Draw) SetTextRenderingHint(nType int) int {
	return xc.XDraw_SetTextRenderingHint(d.Handle, nType)
}

// 绘制_SVG源.
//
// hSvg: SVG句柄.
func (d *Draw) DrawSvgSrc(hSvg int) int {
	return xc.XDraw_DrawSvgSrc(d.Handle, hSvg)
}

// 绘制_SVG.
//
// hSvg: SVG句柄.
//
// x: x坐标.
//
// y: y坐标.
func (d *Draw) DrawSvg(hSvg int, x int, y int) int {
	return xc.XDraw_DrawSvg(d.Handle, hSvg, x, y)
}

// 绘制_SVG扩展.
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
func (d *Draw) DrawSvgEx(hSvg int, x int, y int, nWidth int, nHeight int) int {
	return xc.XDraw_DrawSvgEx(d.Handle, hSvg, x, y, nWidth, nHeight)
}

// 绘制_SVG大小.
//
// hSvg: SVG句柄.
//
// nWidth: 宽度.
//
// nHeight: 高度.
func (d *Draw) DrawSvgSize(hSvg int, nWidth int, nHeight int) int {
	return xc.XDraw_DrawSvgSize(d.Handle, hSvg, nWidth, nHeight)
}

// 绘制_D2D_清理, 使用指定颜色清理画布.
//
// color: ABGR 颜色值.
func (d *Draw) D2D_Clear(color int) int {
	return xc.XDraw_D2D_Clear(d.Handle, color)
}

// 绘制_图片遮盖矩形, 使用矩形作为遮罩.
//
// hImageFrame: 图片句柄.
//
// pRect: 矩形坐标.
//
// pRcMask: 遮罩坐标.
//
// pRcRoundAngle: 遮罩圆角.
func (d *Draw) ImageMaskRect(hImageFrame int, pRect *xc.RECT, pRcMask *xc.RECT, pRcRoundAngle *xc.RECT) int {
	return xc.XDraw_ImageMaskRect(d.Handle, hImageFrame, pRect, pRcMask, pRcRoundAngle)
}

// 绘制_图片遮盖圆型, 使用圆形作为遮罩.
//
// hImageFrame: 图片句柄.
//
// pRect: 矩形坐标.
//
// pRcMask: 遮罩坐标.
func (d *Draw) ImageMaskEllipse(hImageFrame int, pRect *xc.RECT, pRcMask *xc.RECT) int {
	return xc.XDraw_ImageMaskEllipse(d.Handle, hImageFrame, pRect, pRcMask)
}

// 绘制_取字体, 返回字体句柄.
func (d *Draw) GetFont() int {
	return xc.XDraw_GetFont(d.Handle)
}
