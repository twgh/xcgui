package xcgui

import "unsafe"

// 绘制_创建, 创建图形绘制模块实例, 返回句柄
// hdc: 设备上下文句柄HDC.
func XDraw_Create(hdc int) int {
	r, _, _ := xDraw_Create.Call(uintptr(hdc))
	return int(r)
}

// 绘制_销毁, 销毁图形绘制模块实例句柄
// hDraw: 图形绘制句柄.
func XDraw_Destroy(hDraw int) int {
	r, _, _ := xDraw_Destroy.Call(uintptr(hDraw))
	return int(r)
}

// 绘制_置偏移, 设置坐标偏移量, X向左偏移为负数, 向右偏移为正数
// hDraw: 图形绘制句柄.
// x: X轴偏移量.
// y: Y轴偏移量.
func XDraw_SetOffset(hDraw int, x int, y int) int {
	r, _, _ := xDraw_SetOffset.Call(uintptr(hDraw), uintptr(x), uintptr(y))
	return int(r)
}

// 绘制_取偏移, 获取坐标偏移量, X向左偏移为负数, 向右偏移为正数
// hDraw: 图形绘制句柄.
// pX: 接收X轴偏移量.
// pY: 接收Y轴偏移量.
func XDraw_GetOffset(hDraw int, pX *int, pY *int) int {
	r, _, _ := xDraw_GetOffset.Call(uintptr(hDraw), uintptr(unsafe.Pointer(pX)), uintptr(unsafe.Pointer(pY)))
	return int(r)
}

// 绘制_还原状态, 还原状态, 释放用户绑定的GDI对象, 例如画刷, 画笔
// hDraw: 图形绘制句柄.
func XDraw_RestoreGDIOBJ(hDraw int) int {
	r, _, _ := xDraw_RestoreGDIOBJ.Call(uintptr(hDraw))
	return int(r)
}

// 绘制_取HDC, 获取绑定的设备上下文HDC, 返回HDC句柄
// hDraw: 图形绘制句柄.
func XDraw_GetHDC(hDraw int) int {
	r, _, _ := xDraw_GetHDC.Call(uintptr(hDraw))
	return int(r)
}

// 绘制_置画刷颜色, 设置画刷颜色
// hDraw: 图形绘制句柄.
// color: RGB颜色值
// alpha: 透明度, 0-255, 0完全透明, 255不透明
func XDraw_SetBrushColor(hDraw int, color int, alpha uint8) int {
	r, _, _ := xDraw_SetBrushColor.Call(uintptr(hDraw), uintptr(color), uintptr(alpha))
	return int(r)
}

// 绘制_置文本垂直, 设置文本垂直显示
// hDraw: 图形绘制句柄.
// bVertical: 是否垂直显示文本.
func XDraw_SetTextVertical(hDraw int, bVertical bool) int {
	r, _, _ := xDraw_SetTextVertical.Call(uintptr(hDraw), boolPtr(bVertical))
	return int(r)
}

// 绘制_置文本对齐, 设置文本对齐
// hDraw: 图形绘制句柄.
// nFlags: 对齐标识, TextFormatFlag_, TextAlignFlag_, TextTrimming_
func XDraw_SetTextAlign(hDraw int, nFlags int) int {
	r, _, _ := xDraw_SetTextAlign.Call(uintptr(hDraw), uintptr(nFlags))
	return int(r)
}

// 绘制_置字体
// hDraw: 图形绘制句柄.
// hFontx: 炫彩字体.
func XDraw_SetFontX(hDraw int, hFontx int) int {
	r, _, _ := xDraw_SetFontX.Call(uintptr(hDraw), uintptr(hFontx))
	return int(r)
}

// 绘制_置线宽
// hDraw: 图形绘制句柄.
// nWidth: 宽度.
func XDraw_SetLineWidth(hDraw int, nWidth int) int {
	r, _, _ := xDraw_SetLineWidth.Call(uintptr(hDraw), uintptr(nWidth))
	return int(r)
}

// 绘制_置背景模式
// hDraw: 图形绘制句柄.
// bTransparent: 参见MSDN.
func XDraw_SetBkMode(hDraw int, bTransparent bool) int {
	r, _, _ := xDraw_SetBkMode.Call(uintptr(hDraw), boolPtr(bTransparent))
	return int(r)
}

// 绘制_置裁剪区域, 设置裁剪区域
// hDraw: 图形绘制句柄.
// pRect: 区域坐标.
func XDraw_SetClipRect(hDraw int, pRect *RECT) int {
	r, _, _ := xDraw_SetClipRect.Call(uintptr(hDraw), uintptr(unsafe.Pointer(pRect)))
	return int(r)
}

// 绘制_清除裁剪区域
// hDraw: 图形绘制句柄.
func XDraw_ClearClip(hDraw int) int {
	r, _, _ := xDraw_ClearClip.Call(uintptr(hDraw))
	return int(r)
}

// 绘制_启用平滑模式
// hDraw: 图形绘制句柄.
// bEnable: 是否启用.
func XDraw_EnableSmoothingMode(hDraw int, bEnable bool) int {
	r, _, _ := xDraw_EnableSmoothingMode.Call(uintptr(hDraw), boolPtr(bEnable))
	return int(r)
}

// 绘制_启用窗口透明判断, 当启用之后, 调用GDI+函数时, 如果参数alpha=255, 将自动修改为254, 应对GDI+的bug, 否则透明通道异常
// hDraw: 图形绘制句柄
// bTransparent: 是否启用
func XDraw_EnableWndTransparent(hDraw int, bTransparent bool) int {
	r, _, _ := xDraw_EnableWndTransparent.Call(uintptr(hDraw), boolPtr(bTransparent))
	return int(r)
}

// 绘制_创建实心画刷, GDI创建具有指定的纯色逻辑刷
// hDraw: 图形绘制句柄.
// crColor: 画刷颜色.
func XDraw_CreateSolidBrush(hDraw int, crColor int) int {
	r, _, _ := xDraw_CreateSolidBrush.Call(uintptr(hDraw), uintptr(crColor))
	return int(r)
}

// 绘制_创建画笔, GDI创建一个逻辑笔, 指定的样式, 宽度和颜色, 随后的笔可以选择到设备上下文, 用于绘制线条和曲线
// hDraw: 图形绘制句柄.
// fnPenStyle: 画笔样式.
// nWidth: 画笔宽度.
// crColor: 颜色.
func XDraw_CreatePen(hDraw int, fnPenStyle int, nWidth int, crColor int) int {
	r, _, _ := xDraw_CreatePen.Call(uintptr(hDraw), uintptr(fnPenStyle), uintptr(nWidth), uintptr(crColor))
	return int(r)
}

// 绘制_创建矩形区域, GDI创建矩形区域, 成功返回区域句柄, 失败返回NULL
// hDraw: 图形绘制句柄.
// nLeftRect: 左上角X坐标.
// nTopRect: 左上角Y坐标.
// nRightRect: 右下角X坐标.
// nBottomRect: 右下角Y坐标.
func XDraw_CreateRectRgn(hDraw int, nLeftRect int, nTopRect int, nRightRect int, nBottomRect int) int {
	r, _, _ := xDraw_CreateRectRgn.Call(uintptr(hDraw), uintptr(nLeftRect), uintptr(nTopRect), uintptr(nRightRect), uintptr(nBottomRect))
	return int(r)
}

// 绘制_创建圆角矩形区域, GDI创建一个圆角的矩形区域, 成功返回区域句柄, 失败返回NULL
// hDraw: 图形绘制句柄.
// nLeftRect: X-坐标的左上角.
// nTopRect: Y-坐标左上角坐标
// nRightRect: X-坐标右下角
// nBottomRect: Y-坐标右下角
// nWidthEllipse: 椭圆的宽度.
// nHeightEllipse: 椭圆的高度.
func XDraw_CreateRoundRectRgn(hDraw int, nLeftRect int, nTopRect int, nRightRect int, nBottomRect int, nWidthEllipse int, nHeightEllipse int) int {
	r, _, _ := xDraw_CreateRoundRectRgn.Call(uintptr(hDraw), uintptr(nLeftRect), uintptr(nTopRect), uintptr(nRightRect), uintptr(nBottomRect), uintptr(nWidthEllipse), uintptr(nHeightEllipse))
	return int(r)
}

// 绘制_创建多边形区域, GDI创建一个多边形区域, 成功返回区域句柄, 失败返回NULL
// hDraw: 图形绘制句柄.
// pPt: POINT数组.
// cPoints: 数组大小.
// fnPolyFillMode: 多边形填充模式
func XDraw_CreatePolygonRgn(hDraw int, pPt *POINT, cPoints int, fnPolyFillMode int) int {
	r, _, _ := xDraw_CreatePolygonRgn.Call(uintptr(hDraw), uintptr(unsafe.Pointer(pPt)), uintptr(cPoints), uintptr(fnPolyFillMode))
	return int(r)
}

// 绘制_选择裁剪区域, 选择一个区域作为当前裁剪区域, 注意: 该函数只对GDI有效
// hDraw: 图形绘制句柄.
// hRgn: 区域句柄.
func XDraw_SelectClipRgn(hDraw int, hRgn int) int {
	r, _, _ := xDraw_SelectClipRgn.Call(uintptr(hDraw), uintptr(hRgn))
	return int(r)
}

// 绘制_填充矩形, 通过使用指定的刷子填充一个矩形, 此功能包括左侧和顶部的边界, 但不包括矩形的右边和底部边界
// hDraw: 图形绘制句柄.
// pRect: 矩形区域.
func XDraw_FillRect(hDraw int, pRect *RECT) int {
	r, _, _ := xDraw_FillRect.Call(uintptr(hDraw), uintptr(unsafe.Pointer(pRect)))
	return int(r)
}

// 绘制_填充矩形指定颜色
// hDraw: 图形绘制句柄.
// pRect: 矩形区域.
// color: RGB颜色.
// alpha: 透明度.
func XDraw_FillRectColor(hDraw int, pRect *RECT, color int, alpha uint8) int {
	r, _, _ := xDraw_FillRectColor.Call(uintptr(hDraw), uintptr(unsafe.Pointer(pRect)), uintptr(color), uintptr(alpha))
	return int(r)
}

// 绘制_填充区域, 通过使用指定的画刷填充一个区域
// hDraw: 图形绘制句柄.
// hrgn: 区域句柄.
// hbr: 画刷句柄.
func XDraw_FillRgn(hDraw int, hrgn int, hbr int) bool {
	r, _, _ := xDraw_FillRgn.Call(uintptr(hDraw), uintptr(hrgn), uintptr(hbr))
	return int(r) != 0
}

// 绘制_填充圆形
// hDraw: 图形绘制句柄.
// pRect: 矩形区域
func XDraw_FillEllipse(hDraw int, pRect *RECT) int {
	r, _, _ := xDraw_FillEllipse.Call(uintptr(hDraw), uintptr(unsafe.Pointer(pRect)))
	return int(r)
}

// 绘制_圆形, 绘制圆边框
// hDraw: 图形绘制句柄.
// pRect: 矩形区域.
func XDraw_DrawEllipse(hDraw int, pRect *RECT) int {
	r, _, _ := xDraw_DrawEllipse.Call(uintptr(hDraw), uintptr(unsafe.Pointer(pRect)))
	return int(r)
}

// 绘制_填充圆角矩形
// hDraw: 图形绘制句柄.
// pRect: 矩形坐标.
// nWidth: 圆角宽度.
// nHeight: 圆角高度.
func XDraw_FillRoundRect(hDraw int, pRect *RECT, nWidth int, nHeight int) int {
	r, _, _ := xDraw_FillRoundRect.Call(uintptr(hDraw), uintptr(unsafe.Pointer(pRect)), uintptr(nWidth), uintptr(nHeight))
	return int(r)
}

// 绘制_圆角矩形, 绘制圆角矩形边框
// hDraw: 图形绘制句柄.
// pRect: 矩形坐标.
// nWidth: 圆角宽度.
// nHeight: 圆角高度.
func XDraw_DrawRoundRect(hDraw int, pRect *RECT, nWidth int, nHeight int) int {
	r, _, _ := xDraw_DrawRoundRect.Call(uintptr(hDraw), uintptr(unsafe.Pointer(pRect)), uintptr(nWidth), uintptr(nHeight))
	return int(r)
}

// 绘制_填充圆角矩形扩展
// hDraw: 图形绘制句柄.
// pRect: 坐标.
// nLeftTop: 圆角大小.
// nRightTop: 圆角大小.
// nRightBottom: 圆角大小.
// nLeftBottom: 圆角大小.
func XDraw_FillRoundRectEx(hDraw int, pRect *RECT, nLeftTop int, nRightTop int, nRightBottom int, nLeftBottom int) int {
	r, _, _ := xDraw_FillRoundRectEx.Call(uintptr(hDraw), uintptr(unsafe.Pointer(pRect)), uintptr(nLeftTop), uintptr(nRightTop), uintptr(nRightBottom), uintptr(nLeftBottom))
	return int(r)
}

// 绘制_圆角矩形扩展, 绘制圆角矩形边框
// hDraw: 图形绘制句柄.
// pRect: 坐标.
// nLeftTop: 圆角大小.
// nRightTop: 圆角大小.
// nRightBottom: 圆角大小.
// nLeftBottom: 圆角大小.
func XDraw_DrawRoundRectEx(hDraw int, pRect *RECT, nLeftTop int, nRightTop int, nRightBottom int, nLeftBottom int) int {
	r, _, _ := xDraw_DrawRoundRectEx.Call(uintptr(hDraw), uintptr(unsafe.Pointer(pRect)), uintptr(nLeftTop), uintptr(nRightTop), uintptr(nRightBottom), uintptr(nLeftBottom))
	return int(r)
}

// 绘制_矩形, 绘制矩形, 使用当前的画刷和画笔. 如果函数成功, 返回非零值, 如果函数失败, 返回值是零
// hDraw: 图形绘制句柄.
// nLeftRect: 左上角X坐标.
// nTopRect: 左上角Y坐标.
// nRightRect: 右下角X坐标.
// nBottomRect: 右下角Y坐标.
func XDraw_Rectangle(hDraw int, nLeftRect int, nTopRect int, nRightRect int, nBottomRect int) bool {
	r, _, _ := xDraw_Rectangle.Call(uintptr(hDraw), uintptr(nLeftRect), uintptr(nTopRect), uintptr(nRightRect), uintptr(nBottomRect))
	return int(r) != 0
}

// 绘制_组框, 绘制组框, 矩形边框
// hDraw: 图形绘制句柄.
// pRect: 矩形坐标.
// pName: 标题名称.
// textColor: 文本颜色.
// textAlpha: 文本透明度.
// pOffset: 文本偏移
func XDraw_DrawGroupBox_Rect(hDraw int, pRect *RECT, pName string, textColor int, textAlpha uint8, pOffset *POINT) int {
	r, _, _ := xDraw_DrawGroupBox_Rect.Call(uintptr(hDraw), uintptr(unsafe.Pointer(pRect)), strPtr(pName), uintptr(textColor), uintptr(textAlpha), uintptr(unsafe.Pointer(pOffset)))
	return int(r)
}

// 绘制_组框圆角, 绘制组框, 圆角边框
// hDraw: 图形绘制句柄.
// pRect: 矩形坐标.
// pName: 标题名称.
// textColor: 文本颜色.
// textAlpha: 文本透明度.
// pOffset: 文本偏移
// nWidth: 圆角宽度.
// nHeight: 圆角高度.
func XDraw_DrawGroupBox_RoundRect(hDraw int, pRect *RECT, pName string, textColor int, textAlpha uint8, pOffset *POINT, nWidth int, nHeight int) int {
	r, _, _ := xDraw_DrawGroupBox_RoundRect.Call(uintptr(hDraw), uintptr(unsafe.Pointer(pRect)), strPtr(pName), uintptr(textColor), uintptr(textAlpha), uintptr(unsafe.Pointer(pOffset)), uintptr(nWidth), uintptr(nHeight))
	return int(r)
}

// 绘制_渐变填充2, 渐变填充, 从一种颜色过渡到另一种颜色
// hDraw: 图形绘制句柄.
// color1: 开始颜色.
// alpha1: 透明度.
// color2: 结束颜色.
// alpha2: 透明度.
// pRect: 矩形坐标.
// mode: 模式, GRADIENT_FILL_
func XDraw_GradientFill2(hDraw int, color1 int, alpha1 uint8, color2 int, alpha2 uint8, pRect *RECT, mode int) int {
	r, _, _ := xDraw_GradientFill2.Call(uintptr(hDraw), uintptr(color1), uintptr(alpha1), uintptr(color2), uintptr(alpha2), uintptr(unsafe.Pointer(pRect)), uintptr(mode))
	return int(r)
}

// 绘制_渐变填充4, 渐变填充,从一种颜色过渡到另一种颜色
// hDraw: 图形绘制句柄.
// color1: 开始颜色.
// color2: 结束颜色
// color3: 开始颜色
// color4: 结束颜色.
// pRect: 矩形坐标.
// mode: 模式, GRADIENT_FILL_
func XDraw_GradientFill4(hDraw int, color1 int, color2 int, color3 int, color4 int, pRect *RECT, mode int) bool {
	r, _, _ := xDraw_GradientFill4.Call(uintptr(hDraw), uintptr(color1), uintptr(color2), uintptr(color3), uintptr(color4), uintptr(unsafe.Pointer(pRect)), uintptr(mode))
	return int(r) != 0
}

// 绘制_边框区域, 绘制边框, 使用指定的画刷绘制指定的区域的边框. 如果函数成功, 返回非零值, 如果函数失败, 返回值是零
// hDraw: 图形绘制句柄.
// hrgn: 区域句柄.
// hbr: 画刷句柄.
// nWidth: 边框宽度
// nHeight: 边框高度
func XDraw_FrameRgn(hDraw int, hrgn int, hbr int, nWidth int, nHeight int) bool {
	r, _, _ := xDraw_FrameRgn.Call(uintptr(hDraw), uintptr(hrgn), uintptr(hbr), uintptr(nWidth), uintptr(nHeight))
	return int(r) != 0
}

// 绘制_边框矩形, 绘制矩形边框, 使用指定的画刷. 如果函数成功,返回非零值,如果函数失败,返回值是零
// hDraw: 图形绘制句柄.
// pRect: 矩形坐标
func XDraw_FrameRect(hDraw int, pRect *RECT) int {
	r, _, _ := xDraw_FrameRect.Call(uintptr(hDraw), uintptr(unsafe.Pointer(pRect)))
	return int(r)
}

// 绘制_线条
// hDraw: 图形绘制句柄.
// x1: 坐标.
// y1: 坐标.
// x2: 坐标.
// y2: 坐标.
func XDraw_DrawLine(hDraw int, x1 int, y1 int, x2 int, y2 int) int {
	r, _, _ := xDraw_DrawLine.Call(uintptr(hDraw), uintptr(x1), uintptr(y1), uintptr(x2), uintptr(y2))
	return int(r)
}

// 绘制_曲线
// hDraw: 图形绘制句柄.
// points: 坐标点数组
// count: 数组大小
// tension: 大于或等于0.0F的值，指定曲线的张力。
func XDraw_DrawCurve(hDraw int, points int, count int, tension float32) int {
	r, _, _ := xDraw_DrawCurve.Call(uintptr(hDraw), uintptr(points), uintptr(count), uintptr(tension))
	return int(r)
}

// 绘制_焦点矩形
// hDraw: 图形绘制句柄.
// pRect: 矩形坐标.
func XDraw_FocusRect(hDraw int, pRect *RECT) int {
	r, _, _ := xDraw_FocusRect.Call(uintptr(hDraw), uintptr(unsafe.Pointer(pRect)))
	return int(r)
}

// 绘制_移动到起点, 更新当前位置到指定点，并返回以前的位置. 如果函数成功, 返回非零值.
// hDraw: 图形绘制句柄.
// X: 坐标.
// Y: 坐标.
// pPoint: 接收以前的当前位置到一个POINT结构的指针, 如果这个参数是NULL指针, 没有返回原来的位置
func XDraw_MoveToEx(hDraw int, X int, Y int, pPoint *POINT) bool {
	r, _, _ := xDraw_MoveToEx.Call(uintptr(hDraw), uintptr(X), uintptr(Y), uintptr(unsafe.Pointer(pPoint)))
	return int(r) != 0
}

// 绘制_线终点, 函数绘制一条线从当前位置到, 但不包括指定点. 如果函数成功, 返回非零值.
// hDraw: 图形绘制句柄.
// nXEnd: X坐标
// nYEnd: Y坐标
func XDraw_LineTo(hDraw int, nXEnd int, nYEnd int) bool {
	r, _, _ := xDraw_LineTo.Call(uintptr(hDraw), uintptr(nXEnd), uintptr(nYEnd))
	return int(r) != 0
}

// 绘制_折线, Polyline() 参见MSDN.
// hDraw: 图形绘制句柄.
// pArrayPt: 参见MSDN.
// arrayPtSize: 参见MSDN.
func XDraw_Polyline(hDraw int, pArrayPt int, arrayPtSize int) bool {
	r, _, _ := xDraw_Polyline.Call(uintptr(hDraw), uintptr(pArrayPt), uintptr(arrayPtSize))
	return int(r) != 0
}

// 绘制_点线, 绘制水平或垂直虚线.
// hDraw: 图形绘制句柄.
// x1: 起点x坐标.
// y1: 起点y坐标.
// x2: 结束点x坐标.
// y2: 结束点y坐标.
func XDraw_Dottedline(hDraw int, x1 int, y1 int, x2 int, y2 int) int {
	r, _, _ := xDraw_Dottedline.Call(uintptr(hDraw), uintptr(x1), uintptr(y1), uintptr(x2), uintptr(y2))
	return int(r)
}

// 绘制_置像素颜色, 函数设置在指定的坐标到指定的颜色的像素. 如果函数成功返回RGB值, 如果失败返回-1
// hDraw: 图形绘制句柄.
// X: 坐标
// Y: 坐标
// crColor: RGB颜色值
func XDraw_SetPixel(hDraw int, X int, Y int, crColor int) int {
	r, _, _ := xDraw_SetPixel.Call(uintptr(hDraw), uintptr(X), uintptr(Y), uintptr(crColor))
	return int(r)
}

// 绘制_复选框, 绘制复选框.
// hDraw: 图形绘制句柄.
// x: 坐标.
// y: 坐标.
// color: 边框颜色.
// bCheck: 是否选中状态.
func XDraw_Check(hDraw int, x int, y int, color int, bCheck bool) int {
	r, _, _ := xDraw_Check.Call(uintptr(hDraw), uintptr(x), uintptr(y), uintptr(color), boolPtr(bCheck))
	return int(r)
}

// 绘制_图标, 绘制图标, DrawIconEx()参见MSDN.
// hDraw: .
// xLeft: .
// yTop: .
// hIcon: .
// cxWidth: .
// cyWidth: .
// istepIfAniCur: .
// hbrFlickerFreeDraw: .
// diFlags: .
func XDraw_DrawIconEx(hDraw int, xLeft int, yTop int, hIcon int, cxWidth int, cyWidth int, istepIfAniCur int, hbrFlickerFreeDraw int, diFlags int) bool {
	r, _, _ := xDraw_DrawIconEx.Call(uintptr(hDraw), uintptr(xLeft), uintptr(yTop), uintptr(hIcon), uintptr(cxWidth), uintptr(cyWidth), uintptr(istepIfAniCur), uintptr(hbrFlickerFreeDraw), uintptr(diFlags))
	return int(r) != 0
}

// 绘制_复制, BitBlt() 参见MSDN.
// hDrawDest: XX.
// nXDest: XX.
// nYDest: XX.
// nWidth: XX.
// nHeight: XX.
// hdcSrc: XX.
// nXSrc: XX.
// nYSrc: XX.
// dwRop: XX.
func XDraw_BitBlt(hDrawDest int, nXDest int, nYDest int, nWidth int, nHeight int, hdcSrc int, nXSrc int, nYSrc int, dwRop int) bool {
	r, _, _ := xDraw_BitBlt.Call(uintptr(hDrawDest), uintptr(nXDest), uintptr(nYDest), uintptr(nWidth), uintptr(nHeight), uintptr(hdcSrc), uintptr(nXSrc), uintptr(nYSrc), uintptr(dwRop))
	return int(r) != 0
}

// 绘制_复制2, BitBlt() 参见MSDN.
// hDrawDest: XX.
// nXDest: XX.
// nYDest: XX.
// nWidth: XX.
// nHeight: XX.
// hDrawSrc: XX.
// nXSrc: XX.
// nYSrc: XX.
// dwRop: XX.
func XDraw_BitBlt2(hDrawDest int, nXDest int, nYDest int, nWidth int, nHeight int, hDrawSrc int, nXSrc int, nYSrc int, dwRop int) bool {
	r, _, _ := xDraw_BitBlt2.Call(uintptr(hDrawDest), uintptr(nXDest), uintptr(nYDest), uintptr(nWidth), uintptr(nHeight), uintptr(hDrawSrc), uintptr(nXSrc), uintptr(nYSrc), uintptr(dwRop))
	return int(r) != 0
}

// 绘制_带透明复制, AlphaBlend() 参见MSDN.
// hDraw: XX.
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
func XDraw_AlphaBlend(hDraw int, nXOriginDest int, nYOriginDest int, nWidthDest int, nHeightDest int, hdcSrc int, nXOriginSrc int, nYOriginSrc int, nWidthSrc int, nHeightSrc int, alpha int) bool {
	r, _, _ := xDraw_AlphaBlend.Call(uintptr(hDraw), uintptr(nXOriginDest), uintptr(nYOriginDest), uintptr(nWidthDest), uintptr(nHeightDest), uintptr(hdcSrc), uintptr(nXOriginSrc), uintptr(nYOriginSrc), uintptr(nWidthSrc), uintptr(nHeightSrc), uintptr(alpha))
	return int(r) != 0
}

// 绘制_三角箭头, 绘制三角型箭头.
// hDraw: 图形绘制句柄.
// align: 箭头方向, 左1, 上2, 右3, 下4.
// x: 中心点X坐标.
// y: 中心点Y坐标.
// width: 三角形宽度.
// height: 三角形高度.
func XDraw_TriangularArrow(hDraw int, align int, x int, y int, width int, height int) int {
	r, _, _ := xDraw_TriangularArrow.Call(uintptr(hDraw), uintptr(align), uintptr(x), uintptr(y), uintptr(width), uintptr(height))
	return int(r)
}

// 绘制_多边形, 绘制多边形.
// hDraw: 图形绘制句柄.
// points: 顶点坐标数组.
// nCount: 顶点数量.
func XDraw_DrawPolygon(hDraw int, points int, nCount int) int {
	r, _, _ := xDraw_DrawPolygon.Call(uintptr(hDraw), uintptr(points), uintptr(nCount))
	return int(r)
}

// 绘制_多边形F, 绘制多边形.
// hDraw: 图形绘制句柄.
// points: 顶点坐标数组.
// nCount: 顶点数量.
func XDraw_DrawPolygonF(hDraw int, points int, nCount int) int {
	r, _, _ := xDraw_DrawPolygonF.Call(uintptr(hDraw), uintptr(points), uintptr(nCount))
	return int(r)
}

// 绘制_填充多边形, 填充多边形.
// hDraw: 图形绘制句柄.
// points: 顶点坐标数组.
// nCount: 顶点数量.
func XDraw_FillPolygon(hDraw int, points int, nCount int) int {
	r, _, _ := xDraw_FillPolygon.Call(uintptr(hDraw), uintptr(points), uintptr(nCount))
	return int(r)
}

// 绘制_填充多边形F, 填充多边形.
// hDraw: 图形绘制句柄.
// points: 顶点坐标数组.
// nCount: 顶点数量.
func XDraw_FillPolygonF(hDraw int, points int, nCount int) int {
	r, _, _ := xDraw_FillPolygonF.Call(uintptr(hDraw), uintptr(points), uintptr(nCount))
	return int(r)
}

// 绘制_图片
// hDraw: 图形绘制句柄.
// hImageFrame: 图片句柄.
// x: x坐标.
// y: y坐标.
func XDraw_Image(hDraw int, hImageFrame int, x int, y int) int {
	r, _, _ := xDraw_Image.Call(uintptr(hDraw), uintptr(hImageFrame), uintptr(x), uintptr(y))
	return int(r)
}

// 绘制_图片2
// hDraw: 图形绘制句柄.
// hImageFrame: 图片句柄.
// x: x坐标.
// y: y坐标.
// width: 宽度.
// height: 高度.
func XDraw_Image2(hDraw int, hImageFrame int, x int, y int, width int, height int) int {
	r, _, _ := xDraw_Image2.Call(uintptr(hDraw), uintptr(hImageFrame), uintptr(x), uintptr(y), uintptr(width), uintptr(height))
	return int(r)
}

// 绘制_图片缩放
// hDraw: 图形绘制句柄.
// hImageFrame: 图片句柄.
// x: x坐标.
// y: y坐标.
// width: 宽度.
// height: 高度.
func XDraw_ImageStretch(hDraw int, hImageFrame int, x int, y int, width int, height int) int {
	r, _, _ := xDraw_ImageStretch.Call(uintptr(hDraw), uintptr(hImageFrame), uintptr(x), uintptr(y), uintptr(width), uintptr(height))
	return int(r)
}

// 绘制_图片自适应
// hDraw: 图形绘制句柄.
// hImageFrame: 图片句柄.
// pRect: 坐标.
// bOnlyBorder: 是否只绘制边缘区域.
func XDraw_ImageAdaptive(hDraw int, hImageFrame int, pRect *RECT, bOnlyBorder bool) int {
	r, _, _ := xDraw_ImageAdaptive.Call(uintptr(hDraw), uintptr(hImageFrame), uintptr(unsafe.Pointer(pRect)), boolPtr(bOnlyBorder))
	return int(r)
}

// 绘制_图片平铺
// hDraw: 图形绘制句柄.
// hImageFrame: 图片句柄.
// pRect: 坐标.
// flag: 标识, 0:从左上角开始平铺, 1:从左下角开始平铺.
func XDraw_ImageExTile(hDraw int, hImageFrame int, pRect *RECT, flag int) int {
	r, _, _ := xDraw_ImageExTile.Call(uintptr(hDraw), uintptr(hImageFrame), uintptr(unsafe.Pointer(pRect)), uintptr(flag))
	return int(r)
}

// 绘制_图片增强
// hDraw: 图形绘制句柄.
// hImageFrame: 图片句柄.
// pRect: 坐标.
// bClip: 是否裁剪区域.
func XDraw_ImageSuper(hDraw int, hImageFrame int, pRect *RECT, bClip bool) int {
	r, _, _ := xDraw_ImageSuper.Call(uintptr(hDraw), uintptr(hImageFrame), uintptr(unsafe.Pointer(pRect)), boolPtr(bClip))
	return int(r)
}

// 绘制_图片增强2
// hDraw: 图形绘制句柄.
// hImageFrame: 图片句柄.
// prcDest: 目标坐标.
// prcSrc: 源坐标.
func XDraw_ImageSuper2(hDraw int, hImageFrame int, prcDest *RECT, prcSrc *RECT) int {
	r, _, _ := xDraw_ImageSuper2.Call(uintptr(hDraw), uintptr(hImageFrame), uintptr(unsafe.Pointer(prcDest)), uintptr(unsafe.Pointer(prcSrc)))
	return int(r)
}

// 绘制_图片增强遮盖, 绘制带遮盖的图片.
// hDraw: 图形绘制句柄.
// hImageFrame: 图片句柄.
// hImageFrameMask: 图片句柄, 遮盖.
// pRect: 坐标.
// pRectMask: 坐标, 遮盖.
// bClip: 是否裁剪区域.
func XDraw_ImageSuperMask(hDraw int, hImageFrame int, hImageFrameMask int, pRect *RECT, pRectMask *RECT, bClip bool) int {
	r, _, _ := xDraw_ImageSuperMask.Call(uintptr(hDraw), uintptr(hImageFrame), uintptr(hImageFrameMask), uintptr(unsafe.Pointer(pRect)), uintptr(unsafe.Pointer(pRectMask)), boolPtr(bClip))
	return int(r)
}

// 绘制_图片遮盖, 绘制带遮盖的图片.
// hDraw: 图形绘制句柄.
// hImageFrame: 图片句柄.
// hImageFrameMask: 图片句柄
// x: hImageFrame X坐标.
// y: hImageFrame Y坐标.
// x2: hImageFrameMask X坐标.
// y2: hImageFrameMask Y坐标.
func XDraw_ImageMask(hDraw int, hImageFrame int, hImageFrameMask int, x int, y int, x2 int, y2 int) int {
	r, _, _ := xDraw_ImageMask.Call(uintptr(hDraw), uintptr(hImageFrame), uintptr(hImageFrameMask), uintptr(x), uintptr(y), uintptr(x2), uintptr(y2))
	return int(r)
}

// 绘制_文本指定矩形, DrawText() 参见MSDN.
// hDraw: 图形绘制句柄.
// lpString: 字符串.
// nCount: 字符串长度.
// lpRect: 坐标.
func XDraw_DrawText(hDraw int, lpString string, nCount int, lpRect *RECT) int {
	r, _, _ := xDraw_DrawText.Call(uintptr(hDraw), strPtr(lpString), uintptr(nCount), uintptr(unsafe.Pointer(lpRect)))
	return int(r)
}

// 绘制_文本下划线
// hDraw: 图形绘制句柄.
// lpString: 字符串.
// nCount: 字符串长度.
// lpRect: 坐标.
// colorLine: 下划线颜色.
// alphaLine: 下划线透明度.
func XDraw_DrawTextUnderline(hDraw int, lpString string, nCount int, lpRect *RECT, colorLine int, alphaLine uint8) int {
	r, _, _ := xDraw_DrawTextUnderline.Call(uintptr(hDraw), strPtr(lpString), uintptr(nCount), uintptr(unsafe.Pointer(lpRect)), uintptr(colorLine), uintptr(alphaLine))
	return int(r)
}

// 绘制_文本, TextOut() 参见MSDN.
// hDraw: 图形绘制句柄.
// nXStart: XX.
// nYStart: XX.
// lpString: XX.
// cbString: XX.
func XDraw_TextOut(hDraw int, nXStart int, nYStart int, lpString string, cbString int) int {
	r, _, _ := xDraw_TextOut.Call(uintptr(hDraw), uintptr(nXStart), uintptr(nYStart), strPtr(lpString), uintptr(cbString))
	return int(r)
}

// 绘制_文本扩展, TextOut() 参见MSDN.
// hDraw: 图形绘制句柄.
// nXStart: XX.
// nYStart: XX.
// lpString: XX.
func XDraw_TextOutEx(hDraw int, nXStart int, nYStart int, lpString string) int {
	r, _, _ := xDraw_TextOutEx.Call(uintptr(hDraw), uintptr(nXStart), uintptr(nYStart), strPtr(lpString))
	return int(r)
}

// 绘制_文本A, TextOut() 参见MSDN.
// hDraw: 图形绘制句柄.
// nXStart: XX.
// nYStart: XX.
// lpString: XX.
func XDraw_TextOutA(hDraw int, nXStart int, nYStart int, lpString int) int {
	r, _, _ := xDraw_TextOutA.Call(uintptr(hDraw), uintptr(nXStart), uintptr(nYStart), uintptr(lpString))
	return int(r)
}
