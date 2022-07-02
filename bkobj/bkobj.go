package bkobj

import (
	"github.com/twgh/xcgui/bkmanager"
	"github.com/twgh/xcgui/objectbase"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

// BkObj 背景对象.
type BkObj struct {
	objectbase.ObjectBase
}

// NewByHandle 从BkObj句柄创建BkObj对象.
func NewByHandle(handle int) *BkObj {
	p := &BkObj{}
	p.SetHandle(handle)
	return p
}

// NewByBkm 从BkManager对象创建BkObj对象, 失败返回nil.
//
// bkm: 背景管理器对象.
//
// id: 背景对象ID.
func NewByBkm(bkm *bkmanager.BkManager, id int) *BkObj {
	handle := bkm.GetObject(id)
	if handle == 0 {
		return nil
	}
	p := &BkObj{}
	p.SetHandle(handle)
	return p
}

// NewByBkmHandle 从BkManager句柄创建BkObj对象, 失败返回nil.
//
// hBkm: 背景管理器句柄.
//
// id: 背景对象ID.
func NewByBkmHandle(hBkm, id int) *BkObj {
	handle := xc.XBkM_GetObject(hBkm, id)
	if handle == 0 {
		return nil
	}
	p := &BkObj{}
	p.SetHandle(handle)
	return p
}

// 背景对象_置外间距, 外间距与对齐标识(BkObject_Align_Flag_)互相依赖.
//
// left: 左边间距.
//
// top: 顶部间距.
//
// right: 右边间距.
//
// bottom: 底部间距.
func (b *BkObj) SetMargin(left int, top int, right int, bottom int) int {
	return xc.XBkObj_SetMargin(b.Handle, left, top, right, bottom)
}

// 背景对象_置对齐.
//
// nFlags: 对齐方式: xcc.BkObject_Align_Flag_.
func (b *BkObj) SetAlign(nFlags xcc.BkObject_Align_Flag_) int {
	return xc.XBkObj_SetAlign(b.Handle, nFlags)
}

// 背景对象_置图片.
//
// hImage: 图片句柄.
func (b *BkObj) SetImage(hImage int) int {
	return xc.XBkObj_SetImage(b.Handle, hImage)
}

// 背景对象_置旋转.
//
// angle: 旋转角度.
func (b *BkObj) SetRotate(angle float32) int {
	return xc.XBkObj_SetRotate(b.Handle, angle)
}

// 背景对象_置填充颜色.
//
// color: ABGR 颜色值.
func (b *BkObj) SetFillColor(color int) int {
	return xc.XBkObj_SetFillColor(b.Handle, color)
}

// 背景对象_置边框宽度.
//
// width: 宽度.
func (b *BkObj) SetBorderWidth(width int) int {
	return xc.XBkObj_SetBorderWidth(b.Handle, width)
}

// 背景对象_置边框颜色.
//
// color: ABGR 颜色值.
func (b *BkObj) SetBorderColor(color int) int {
	return xc.XBkObj_SetBorderColor(b.Handle, color)
}

// 背景对象_置矩形圆角.
//
// leftTop: 左上角.
//
// leftBottom: 左下角.
//
// rightTop: 右上角.
//
// rightBottom: 右下角.
func (b *BkObj) SetRectRoundAngle(leftTop int, leftBottom int, rightTop int, rightBottom int) int {
	return xc.XBkObj_SetRectRoundAngle(b.Handle, leftTop, leftBottom, rightTop, rightBottom)
}

// 背景对象_启用填充, 启用绘制填充.
//
// bEnable: 是否启用.
func (b *BkObj) EnableFill(bEnable bool) int {
	return xc.XBkObj_EnableFill(b.Handle, bEnable)
}

// 背景对象_启用边框, 启用绘制边框.
//
// bEnable: 是否启用.
func (b *BkObj) EnableBorder(bEnable bool) int {
	return xc.XBkObj_EnableBorder(b.Handle, bEnable)
}

// 背景对象_置文本.
//
// pText: 文本字符串.
func (b *BkObj) SetText(pText string) int {
	return xc.XBkObj_SetText(b.Handle, pText)
}

// 背景对象_置字体.
//
// hFont: 字体句柄.
func (b *BkObj) SetFont(hFont int) int {
	return xc.XBkObj_SetFont(b.Handle, hFont)
}

// 背景对象_置文本对齐.
//
// nAlign: 文本对齐方式: xcc.TextFormatFlag_, xcc.TextAlignFlag_, xcc.TextTrimming_.
func (b *BkObj) SetTextAlign(nAlign xcc.TextFormatFlag_) int {
	return xc.XBkObj_SetTextAlign(b.Handle, nAlign)
}

// 背景对象_取外间距.
//
// pMargin: 接收返回外间距.
func (b *BkObj) GetMargin(pMargin *xc.RECT) int {
	return xc.XBkObj_GetMargin(b.Handle, pMargin)
}

// 背景对象_取对齐, 返回对齐标识: xcc.BkObject_Align_Flag_.
func (b *BkObj) GetAlign() xcc.BkObject_Align_Flag_ {
	return xc.XBkObj_GetAlign(b.Handle)
}

// 背景对象_取图片, 返回图片句柄.
func (b *BkObj) GetImage() int {
	return xc.XBkObj_GetImage(b.Handle)
}

// 背景对象_取旋转角度, 返回旋转角度.
func (b *BkObj) GetRotate() int {
	return xc.XBkObj_GetRotate(b.Handle)
}

// 背景对象_取填充色, 返回ABGR填充色.
func (b *BkObj) GetFillColor() int {
	return xc.XBkObj_GetFillColor(b.Handle)
}

// 背景对象_取边框色, 返回ABGR边框色.
func (b *BkObj) GetBorderColor() int {
	return xc.XBkObj_GetBorderColor(b.Handle)
}

// 背景对象_取边框宽度, 返回边框宽度.
func (b *BkObj) GetBorderWidth() int {
	return xc.XBkObj_GetBorderWidth(b.Handle)
}

// 背景对象_取矩形圆角.
//
// pRect: 接收返回圆角大小.
func (b *BkObj) GetRectRoundAngle(pRect *xc.RECT) int {
	return xc.XBkObj_GetRectRoundAngle(b.Handle, pRect)
}

// 背景对象_是否填充.
func (b *BkObj) IsFill() bool {
	return xc.XBkObj_IsFill(b.Handle)
}

// 背景对象_是否边框.
func (b *BkObj) IsBorder() bool {
	return xc.XBkObj_IsBorder(b.Handle)
}

// 背景对象_取文本.
func (b *BkObj) GetText() string {
	return xc.XBkObj_GetText(b.Handle)
}

// 背景对象_取字体, 返回字体句柄.
func (b *BkObj) GetFont() int {
	return xc.XBkObj_GetFont(b.Handle)
}

// 背景对象_取文本对齐, 返回文本对齐方式: xcc.TextFormatFlag_.
func (b *BkObj) GetTextAlign() xcc.TextFormatFlag_ {
	return xc.XBkObj_GetTextAlign(b.Handle)
}
