// 图形绘制
package draw

import "github.com/twgh/xcgui/xc"

// 图形绘制
type Draw struct {
	HDRAW int
}

// 绘制_创建, 创建图形绘制模块实例, 返回句柄
// hdc: 设备上下文句柄HDC.
func NewDraw(hdc int) *Draw {
	p := &Draw{
		HDRAW: xc.XDraw_Create(hdc),
	}
	return p
}

// 绘制_销毁, 销毁图形绘制模块实例句柄
func (d *Draw) Destroy() int {
	return xc.XDraw_Destroy(d.HDRAW)
}

// 绘制_置偏移, 设置坐标偏移量, X向左偏移为负数, 向右偏移为正数
// x: X轴偏移量.
// y: Y轴偏移量.
func (d *Draw) SetOffset(x int, y int) int {
	return xc.XDraw_SetOffset(d.HDRAW, x, y)
}

// 绘制_取偏移, 获取坐标偏移量, X向左偏移为负数, 向右偏移为正数
// pX: 接收X轴偏移量.
// pY: 接收Y轴偏移量.
func (d *Draw) GetOffset(pX *int, pY *int) int {
	return xc.XDraw_GetOffset(d.HDRAW, pX, pY)
}

// 绘制_还原状态, 还原状态, 释放用户绑定的GDI对象, 例如画刷, 画笔
func (d *Draw) RestoreGDIOBJ() int {
	return xc.XDraw_RestoreGDIOBJ(d.HDRAW)
}

// 绘制_取HDC, 获取绑定的设备上下文HDC, 返回HDC句柄
func (d *Draw) GetHDC() int {
	return xc.XDraw_GetHDC(d.HDRAW)
}

// 绘制_置画刷颜色, 设置画刷颜色
// color: RGB颜色值
// alpha: 透明度, 0-255, 0完全透明, 255不透明
func (d *Draw) SetBrushColor(color int, alpha uint8) int {
	return xc.XDraw_SetBrushColor(d.HDRAW, color, alpha)
}

// 绘制_置文本垂直, 设置文本垂直显示
// bVertical: 是否垂直显示文本.
func (d *Draw) SetTextVertical(bVertical bool) int {
	return xc.XDraw_SetTextVertical(d.HDRAW, bVertical)
}

// 绘制_置文本对齐, 设置文本对齐
// nFlags: 对齐标识, TextFormatFlag_, TextAlignFlag_, TextTrimming_
func (d *Draw) SetTextAlign(nFlags int) int {
	return xc.XDraw_SetTextAlign(d.HDRAW, nFlags)
}

// 绘制_置字体
// hFontx: 炫彩字体.
func (d *Draw) SetFontX(hFontx int) int {
	return xc.XDraw_SetFontX(d.HDRAW, hFontx)
}

// 绘制_置线宽
// nWidth: 宽度.
func (d *Draw) SetLineWidth(nWidth int) int {
	return xc.XDraw_SetLineWidth(d.HDRAW, nWidth)
}

// 绘制_置背景模式
// bTransparent: 参见MSDN.
func (d *Draw) SetBkMode(bTransparent bool) int {
	return xc.XDraw_SetBkMode(d.HDRAW, bTransparent)
}

// 绘制_置裁剪区域, 设置裁剪区域
// pRect: 区域坐标.
func (d *Draw) SetClipRect(pRect *xc.RECT) int {
	return xc.XDraw_SetClipRect(d.HDRAW, pRect)
}

// 绘制_清除裁剪区域
func (d *Draw) ClearClip() int {
	return xc.XDraw_ClearClip(d.HDRAW)
}

// 绘制_启用平滑模式
// bEnable: 是否启用.
func (d *Draw) EnableSmoothingMode(bEnable bool) int {
	return xc.XDraw_EnableSmoothingMode(d.HDRAW, bEnable)
}

// 绘制_启用窗口透明判断, 当启用之后, 调用GDI+函数时, 如果参数alpha=255, 将自动修改为254, 应对GDI+的bug, 否则透明通道异常
// bTransparent: 是否启用
func (d *Draw) EnableWndTransparent(bTransparent bool) int {
	return xc.XDraw_EnableWndTransparent(d.HDRAW, bTransparent)
}

// 绘制_创建实心画刷, GDI创建具有指定的纯色逻辑刷
// crColor: 画刷颜色.
func (d *Draw) CreateSolidBrush(crColor int) int {
	return xc.XDraw_CreateSolidBrush(d.HDRAW, crColor)
}

// 绘制_创建画笔, GDI创建一个逻辑笔, 指定的样式, 宽度和颜色, 随后的笔可以选择到设备上下文, 用于绘制线条和曲线
// fnPenStyle: 画笔样式.
// nWidth: 画笔宽度.
// crColor: 颜色.
func (d *Draw) CreatePen(fnPenStyle int, nWidth int, crColor int) int {
	return xc.XDraw_CreatePen(d.HDRAW, fnPenStyle, nWidth, crColor)
}

// 绘制_创建矩形区域, GDI创建矩形区域, 成功返回区域句柄, 失败返回NULL
// nLeftRect: 左上角X坐标.
// nTopRect: 左上角Y坐标.
// nRightRect: 右下角X坐标.
// nBottomRect: 右下角Y坐标.
func (d *Draw) CreateRectRgn(nLeftRect int, nTopRect int, nRightRect int, nBottomRect int) int {
	return xc.XDraw_CreateRectRgn(d.HDRAW, nLeftRect, nTopRect, nRightRect, nBottomRect)
}

// 绘制_创建圆角矩形区域, GDI创建一个圆角的矩形区域, 成功返回区域句柄, 失败返回NULL
// nLeftRect: X-坐标的左上角.
// nTopRect: Y-坐标左上角坐标
// nRightRect: X-坐标右下角
// nBottomRect: Y-坐标右下角
// nWidthEllipse: 椭圆的宽度.
// nHeightEllipse: 椭圆的高度.
func (d *Draw) CreateRoundRectRgn(nLeftRect int, nTopRect int, nRightRect int, nBottomRect int, nWidthEllipse int, nHeightEllipse int) int {
	return xc.XDraw_CreateRoundRectRgn(d.HDRAW, nLeftRect, nTopRect, nRightRect, nBottomRect, nWidthEllipse, nHeightEllipse)
}

// 绘制_创建多边形区域, GDI创建一个多边形区域, 成功返回区域句柄, 失败返回NULL
// pPt: POINT数组.
// cPoints: 数组大小.
// fnPolyFillMode: 多边形填充模式
func (d *Draw) CreatePolygonRgn(pPt *xc.POINT, cPoints int, fnPolyFillMode int) int {
	return xc.XDraw_CreatePolygonRgn(d.HDRAW, pPt, cPoints, fnPolyFillMode)
}

// 绘制_选择裁剪区域, 选择一个区域作为当前裁剪区域, 注意: 该函数只对GDI有效
// hRgn: 区域句柄.
func (d *Draw) SelectClipRgn(hRgn int) int {
	return xc.XDraw_SelectClipRgn(d.HDRAW, hRgn)
}

// 绘制_填充矩形, 通过使用指定的刷子填充一个矩形, 此功能包括左侧和顶部的边界, 但不包括矩形的右边和底部边界
// pRect: 矩形区域.
func (d *Draw) FillRect(pRect *xc.RECT) int {
	return xc.XDraw_FillRect(d.HDRAW, pRect)
}

// 绘制_填充矩形指定颜色
// pRect: 矩形区域.
// color: RGB颜色.
// alpha: 透明度.
func (d *Draw) FillRectColor(pRect *xc.RECT, color int, alpha uint8) int {
	return xc.XDraw_FillRectColor(d.HDRAW, pRect, color, alpha)
}

// 绘制_填充区域, 通过使用指定的画刷填充一个区域
// hrgn: 区域句柄.
// hbr: 画刷句柄.
func (d *Draw) FillRgn(hrgn int, hbr int) bool {
	return xc.XDraw_FillRgn(d.HDRAW, hrgn, hbr)
}

// 绘制_填充圆形
// pRect: 矩形区域
func (d *Draw) FillEllipse(pRect *xc.RECT) int {
	return xc.XDraw_FillEllipse(d.HDRAW, pRect)
}

// 绘制_圆形, 绘制圆边框
// pRect: 矩形区域.
func (d *Draw) DrawEllipse(pRect *xc.RECT) int {
	return xc.XDraw_DrawEllipse(d.HDRAW, pRect)
}

// 绘制_填充圆角矩形
// pRect: 矩形坐标.
// nWidth: 圆角宽度.
// nHeight: 圆角高度.
func (d *Draw) FillRoundRect(pRect *xc.RECT, nWidth int, nHeight int) int {
	return xc.XDraw_FillRoundRect(d.HDRAW, pRect, nWidth, nHeight)
}

// 绘制_圆角矩形, 绘制圆角矩形边框
// pRect: 矩形坐标.
// nWidth: 圆角宽度.
// nHeight: 圆角高度.
func (d *Draw) DrawRoundRect(pRect *xc.RECT, nWidth int, nHeight int) int {
	return xc.XDraw_DrawRoundRect(d.HDRAW, pRect, nWidth, nHeight)
}

// 绘制_填充圆角矩形扩展
// pRect: 坐标.
// nLeftTop: 圆角大小.
// nRightTop: 圆角大小.
// nRightBottom: 圆角大小.
// nLeftBottom: 圆角大小.
func (d *Draw) FillRoundRectEx(pRect *xc.RECT, nLeftTop int, nRightTop int, nRightBottom int, nLeftBottom int) int {
	return xc.XDraw_FillRoundRectEx(d.HDRAW, pRect, nLeftTop, nRightTop, nRightBottom, nLeftBottom)
}

// 绘制_圆角矩形扩展, 绘制圆角矩形边框
// pRect: 坐标.
// nLeftTop: 圆角大小.
// nRightTop: 圆角大小.
// nRightBottom: 圆角大小.
// nLeftBottom: 圆角大小.
func (d *Draw) DrawRoundRectEx(pRect *xc.RECT, nLeftTop int, nRightTop int, nRightBottom int, nLeftBottom int) int {
	return xc.XDraw_DrawRoundRectEx(d.HDRAW, pRect, nLeftTop, nRightTop, nRightBottom, nLeftBottom)
}

// 绘制_矩形, 绘制矩形, 使用当前的画刷和画笔. 如果函数成功, 返回非零值, 如果函数失败, 返回值是零
// nLeftRect: 左上角X坐标.
// nTopRect: 左上角Y坐标.
// nRightRect: 右下角X坐标.
// nBottomRect: 右下角Y坐标.
func (d *Draw) Rectangle(nLeftRect int, nTopRect int, nRightRect int, nBottomRect int) bool {
	return xc.XDraw_Rectangle(d.HDRAW, nLeftRect, nTopRect, nRightRect, nBottomRect)
}

// 绘制_组框, 绘制组框, 矩形边框
// pRect: 矩形坐标.
// pName: 标题名称.
// textColor: 文本颜色.
// textAlpha: 文本透明度.
// pOffset: 文本偏移
func (d *Draw) DrawGroupBox_Rect(pRect *xc.RECT, pName string, textColor int, textAlpha uint8, pOffset *xc.POINT) int {
	return xc.XDraw_DrawGroupBox_Rect(d.HDRAW, pRect, pName, textColor, textAlpha, pOffset)
}

// 绘制_组框圆角, 绘制组框, 圆角边框
// pRect: 矩形坐标.
// pName: 标题名称.
// textColor: 文本颜色.
// textAlpha: 文本透明度.
// pOffset: 文本偏移
// nWidth: 圆角宽度.
// nHeight: 圆角高度.
func (d *Draw) DrawGroupBox_RoundRect(pRect *xc.RECT, pName string, textColor int, textAlpha uint8, pOffset *xc.POINT, nWidth int, nHeight int) int {
	return xc.XDraw_DrawGroupBox_RoundRect(d.HDRAW, pRect, pName, textColor, textAlpha, pOffset, nWidth, nHeight)
}

// 绘制_渐变填充2, 渐变填充, 从一种颜色过渡到另一种颜色
// color1: 开始颜色.
// alpha1: 透明度.
// color2: 结束颜色.
// alpha2: 透明度.
// pRect: 矩形坐标.
// mode: 模式, GRADIENT_FILL_
func (d *Draw) GradientFill2(color1 int, alpha1 uint8, color2 int, alpha2 uint8, pRect *xc.RECT, mode int) int {
	return xc.XDraw_GradientFill2(d.HDRAW, color1, alpha1, color2, alpha2, pRect, mode)
}

// 绘制_渐变填充4, 渐变填充,从一种颜色过渡到另一种颜色
// color1: 开始颜色.
// color2: 结束颜色
// color3: 开始颜色
// color4: 结束颜色.
// pRect: 矩形坐标.
// mode: 模式, GRADIENT_FILL_
func (d *Draw) GradientFill4(color1 int, color2 int, color3 int, color4 int, pRect *xc.RECT, mode int) bool {
	return xc.XDraw_GradientFill4(d.HDRAW, color1, color2, color3, color4, pRect, mode)
}

// 绘制_边框区域, 绘制边框, 使用指定的画刷绘制指定的区域的边框. 如果函数成功, 返回非零值, 如果函数失败, 返回值是零
// hrgn: 区域句柄.
// hbr: 画刷句柄.
// nWidth: 边框宽度
// nHeight: 边框高度
func (d *Draw) FrameRgn(hrgn int, hbr int, nWidth int, nHeight int) bool {
	return xc.XDraw_FrameRgn(d.HDRAW, hrgn, hbr, nWidth, nHeight)
}

// 绘制_边框矩形, 绘制矩形边框, 使用指定的画刷. 如果函数成功,返回非零值,如果函数失败,返回值是零
// pRect: 矩形坐标
func (d *Draw) FrameRect(pRect *xc.RECT) int {
	return xc.XDraw_FrameRect(d.HDRAW, pRect)
}

// 绘制_线条
// x1: 坐标.
// y1: 坐标.
// x2: 坐标.
// y2: 坐标.
func (d *Draw) DrawLine(x1 int, y1 int, x2 int, y2 int) int {
	return xc.XDraw_DrawLine(d.HDRAW, x1, y1, x2, y2)
}

// 绘制_曲线
// points: 坐标点数组
// count: 数组大小
// tension: 大于或等于0.0F的值，指定曲线的张力。
func (d *Draw) DrawCurve(points int, count int, tension float32) int {
	return xc.XDraw_DrawCurve(d.HDRAW, points, count, tension)
}

// 绘制_焦点矩形
// pRect: 矩形坐标.
func (d *Draw) FocusRect(pRect *xc.RECT) int {
	return xc.XDraw_FocusRect(d.HDRAW, pRect)
}

// 绘制_移动到起点, 更新当前位置到指定点，并返回以前的位置. 如果函数成功, 返回非零值.
// X: 坐标.
// Y: 坐标.
// pPoint: 接收以前的当前位置到一个POINT结构的指针, 如果这个参数是NULL指针, 没有返回原来的位置
func (d *Draw) MoveToEx(X int, Y int, pPoint *xc.POINT) bool {
	return xc.XDraw_MoveToEx(d.HDRAW, X, Y, pPoint)
}

// 绘制_线终点, 函数绘制一条线从当前位置到, 但不包括指定点. 如果函数成功, 返回非零值.
// nXEnd: X坐标
// nYEnd: Y坐标
func (d *Draw) LineTo(nXEnd int, nYEnd int) bool {
	return xc.XDraw_LineTo(d.HDRAW, nXEnd, nYEnd)
}

// 绘制_折线, Polyline() 参见MSDN.
// pArrayPt: 参见MSDN.
// arrayPtSize: 参见MSDN.
func (d *Draw) Polyline(pArrayPt int, arrayPtSize int) bool {
	return xc.XDraw_Polyline(d.HDRAW, pArrayPt, arrayPtSize)
}

// 绘制_点线, 绘制水平或垂直虚线.
// x1: 起点x坐标.
// y1: 起点y坐标.
// x2: 结束点x坐标.
// y2: 结束点y坐标.
func (d *Draw) Dottedline(x1 int, y1 int, x2 int, y2 int) int {
	return xc.XDraw_Dottedline(d.HDRAW, x1, y1, x2, y2)
}

// 绘制_置像素颜色, 函数设置在指定的坐标到指定的颜色的像素. 如果函数成功返回RGB值, 如果失败返回-1
// X: 坐标
// Y: 坐标
// crColor: RGB颜色值
func (d *Draw) SetPixel(X int, Y int, crColor int) int {
	return xc.XDraw_SetPixel(d.HDRAW, X, Y, crColor)
}

// 绘制_复选框, 绘制复选框.
// x: 坐标.
// y: 坐标.
// color: 边框颜色.
// bCheck: 是否选中状态.
func (d *Draw) Check(x int, y int, color int, bCheck bool) int {
	return xc.XDraw_Check(d.HDRAW, x, y, color, bCheck)
}

// 绘制_图标, 绘制图标, DrawIconEx()参见MSDN.
// xLeft: .
// yTop: .
// hIcon: .
// cxWidth: .
// cyWidth: .
// istepIfAniCur: .
// hbrFlickerFreeDraw: .
// diFlags: .
func (d *Draw) DrawIconEx(xLeft int, yTop int, hIcon int, cxWidth int, cyWidth int, istepIfAniCur int, hbrFlickerFreeDraw int, diFlags int) bool {
	return xc.XDraw_DrawIconEx(d.HDRAW, xLeft, yTop, hIcon, cxWidth, cyWidth, istepIfAniCur, hbrFlickerFreeDraw, diFlags)
}

// 绘制_复制, BitBlt() 参见MSDN.
// nXDest: XX.
// nYDest: XX.
// nWidth: XX.
// nHeight: XX.
// hdcSrc: XX.
// nXSrc: XX.
// nYSrc: XX.
// dwRop: XX.
func (d *Draw) BitBlt(nXDest int, nYDest int, nWidth int, nHeight int, hdcSrc int, nXSrc int, nYSrc int, dwRop int) bool {
	return xc.XDraw_BitBlt(d.HDRAW, nXDest, nYDest, nWidth, nHeight, hdcSrc, nXSrc, nYSrc, dwRop)
}

// 绘制_复制2, BitBlt() 参见MSDN.
// nXDest: XX.
// nYDest: XX.
// nWidth: XX.
// nHeight: XX.
// hDrawSrc: XX.
// nXSrc: XX.
// nYSrc: XX.
// dwRop: XX.
func (d *Draw) BitBlt2(nXDest int, nYDest int, nWidth int, nHeight int, hDrawSrc int, nXSrc int, nYSrc int, dwRop int) bool {
	return xc.XDraw_BitBlt2(d.HDRAW, nXDest, nYDest, nWidth, nHeight, hDrawSrc, nXSrc, nYSrc, dwRop)
}

// 绘制_带透明复制, AlphaBlend() 参见MSDN.
// nXOriginDest: XX.
// nYOriginDest: XX.
// nWidthDest: XX.
// nHeightDest: XX.
// hdcSrc: XX.
// nXOriginSrc: XX.
// nYOriginSrc: XX.
// nWidthSrc: XX.
// nHeightSrc: XX.
// alpha: XX.
func (d *Draw) AlphaBlend(nXOriginDest int, nYOriginDest int, nWidthDest int, nHeightDest int, hdcSrc int, nXOriginSrc int, nYOriginSrc int, nWidthSrc int, nHeightSrc int, alpha int) bool {
	return xc.XDraw_AlphaBlend(d.HDRAW, nXOriginDest, nYOriginDest, nWidthDest, nHeightDest, hdcSrc, nXOriginSrc, nYOriginSrc, nWidthSrc, nHeightSrc, alpha)
}

// 绘制_三角箭头, 绘制三角型箭头.
// align: 箭头方向, 左1, 上2, 右3, 下4.
// x: 中心点X坐标.
// y: 中心点Y坐标.
// width: 三角形宽度.
// height: 三角形高度.
func (d *Draw) TriangularArrow(align int, x int, y int, width int, height int) int {
	return xc.XDraw_TriangularArrow(d.HDRAW, align, x, y, width, height)
}

// 绘制_多边形, 绘制多边形.
// points: 顶点坐标数组.
// nCount: 顶点数量.
func (d *Draw) DrawPolygon(points int, nCount int) int {
	return xc.XDraw_DrawPolygon(d.HDRAW, points, nCount)
}

// 绘制_多边形F, 绘制多边形.
// points: 顶点坐标数组.
// nCount: 顶点数量.
func (d *Draw) DrawPolygonF(points int, nCount int) int {
	return xc.XDraw_DrawPolygonF(d.HDRAW, points, nCount)
}

// 绘制_填充多边形, 填充多边形.
// points: 顶点坐标数组.
// nCount: 顶点数量.
func (d *Draw) FillPolygon(points int, nCount int) int {
	return xc.XDraw_FillPolygon(d.HDRAW, points, nCount)
}

// 绘制_填充多边形F, 填充多边形.
// points: 顶点坐标数组.
// nCount: 顶点数量.
func (d *Draw) FillPolygonF(points int, nCount int) int {
	return xc.XDraw_FillPolygonF(d.HDRAW, points, nCount)
}

// 绘制_图片
// hImageFrame: 图片句柄.
// x: x坐标.
// y: y坐标.
func (d *Draw) Image(hImageFrame int, x int, y int) int {
	return xc.XDraw_Image(d.HDRAW, hImageFrame, x, y)
}

// 绘制_图片2
// hImageFrame: 图片句柄.
// x: x坐标.
// y: y坐标.
// width: 宽度.
// height: 高度.
func (d *Draw) Image2(hImageFrame int, x int, y int, width int, height int) int {
	return xc.XDraw_Image2(d.HDRAW, hImageFrame, x, y, width, height)
}

// 绘制_图片缩放
// hImageFrame: 图片句柄.
// x: x坐标.
// y: y坐标.
// width: 宽度.
// height: 高度.
func (d *Draw) ImageStretch(hImageFrame int, x int, y int, width int, height int) int {
	return xc.XDraw_ImageStretch(d.HDRAW, hImageFrame, x, y, width, height)
}

// 绘制_图片自适应
// hImageFrame: 图片句柄.
// pRect: 坐标.
// bOnlyBorder: 是否只绘制边缘区域.
func (d *Draw) ImageAdaptive(hImageFrame int, pRect *xc.RECT, bOnlyBorder bool) int {
	return xc.XDraw_ImageAdaptive(d.HDRAW, hImageFrame, pRect, bOnlyBorder)
}

// 绘制_图片平铺
// hImageFrame: 图片句柄.
// pRect: 坐标.
// flag: 标识, 0:从左上角开始平铺, 1:从左下角开始平铺.
func (d *Draw) ImageExTile(hImageFrame int, pRect *xc.RECT, flag int) int {
	return xc.XDraw_ImageExTile(d.HDRAW, hImageFrame, pRect, flag)
}

// 绘制_图片增强
// hImageFrame: 图片句柄.
// pRect: 坐标.
// bClip: 是否裁剪区域.
func (d *Draw) ImageSuper(hImageFrame int, pRect *xc.RECT, bClip bool) int {
	return xc.XDraw_ImageSuper(d.HDRAW, hImageFrame, pRect, bClip)
}

// 绘制_图片增强2
// hImageFrame: 图片句柄.
// prcDest: 目标坐标.
// prcSrc: 源坐标.
func (d *Draw) ImageSuper2(hImageFrame int, prcDest *xc.RECT, prcSrc *xc.RECT) int {
	return xc.XDraw_ImageSuper2(d.HDRAW, hImageFrame, prcDest, prcSrc)
}

// 绘制_图片增强遮盖, 绘制带遮盖的图片.
// hImageFrame: 图片句柄.
// hImageFrameMask: 图片句柄, 遮盖.
// pRect: 坐标.
// pRectMask: 坐标, 遮盖.
// bClip: 是否裁剪区域.
func (d *Draw) ImageSuperMask(hImageFrame int, hImageFrameMask int, pRect *xc.RECT, pRectMask *xc.RECT, bClip bool) int {
	return xc.XDraw_ImageSuperMask(d.HDRAW, hImageFrame, hImageFrameMask, pRect, pRectMask, bClip)
}

// 绘制_图片遮盖, 绘制带遮盖的图片.
// hImageFrame: 图片句柄.
// hImageFrameMask: 图片句柄
// x: hImageFrame X坐标.
// y: hImageFrame Y坐标.
// x2: hImageFrameMask X坐标.
// y2: hImageFrameMask Y坐标.
func (d *Draw) ImageMask(hImageFrame int, hImageFrameMask int, x int, y int, x2 int, y2 int) int {
	return xc.XDraw_ImageMask(d.HDRAW, hImageFrame, hImageFrameMask, x, y, x2, y2)
}

// 绘制_文本指定矩形, DrawText() 参见MSDN.
// lpString: 字符串.
// nCount: 字符串长度.
// lpRect: 坐标.
func (d *Draw) DrawText(lpString string, nCount int, lpRect *xc.RECT) int {
	return xc.XDraw_DrawText(d.HDRAW, lpString, nCount, lpRect)
}

// 绘制_文本下划线
// lpString: 字符串.
// nCount: 字符串长度.
// lpRect: 坐标.
// colorLine: 下划线颜色.
// alphaLine: 下划线透明度.
func (d *Draw) DrawTextUnderline(lpString string, nCount int, lpRect *xc.RECT, colorLine int, alphaLine uint8) int {
	return xc.XDraw_DrawTextUnderline(d.HDRAW, lpString, nCount, lpRect, colorLine, alphaLine)
}

// 绘制_文本, TextOut() 参见MSDN.
// nXStart: XX.
// nYStart: XX.
// lpString: XX.
// cbString: XX.
func (d *Draw) TextOut(nXStart int, nYStart int, lpString string, cbString string) int {
	return xc.XDraw_TextOut(d.HDRAW, nXStart, nYStart, lpString, cbString)
}

// 绘制_文本扩展, TextOut() 参见MSDN.
// nXStart: XX.
// nYStart: XX.
// lpString: XX.
func (d *Draw) TextOutEx(nXStart int, nYStart int, lpString string) int {
	return xc.XDraw_TextOutEx(d.HDRAW, nXStart, nYStart, lpString)
}

// 绘制_文本A, TextOut() 参见MSDN.
// nXStart: XX.
// nYStart: XX.
// lpString: XX.
func (d *Draw) TextOutA(nXStart int, nYStart int, lpString string) int {
	return xc.XDraw_TextOutA(d.HDRAW, nXStart, nYStart, lpString)
}

// 绘制_圆弧.
// x: 坐标.
// y: 坐标.
// nWidth: 宽度.
// nHeight: 高度.
// startAngle: 起始角度.
// sweepAngle: 绘制角度, 从起始角度开始计算.
func (d *Draw) DrawArc(x, y int, nWidth int, nHeight int, startAngle float32, sweepAngle float32) int {
	return xc.XDraw_DrawArc(d.HDRAW, x, y, nWidth, nHeight, startAngle, sweepAngle)
}
