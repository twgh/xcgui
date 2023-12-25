package widget

import (
	"github.com/twgh/xcgui/objectbase"
	"github.com/twgh/xcgui/xc"
)

// Shape 形状对象基类.
type Shape struct {
	objectbase.Widget
}

// 从句柄创建对象.
func NewShapeByHandle(handle int) *Shape {
	p := &Shape{}
	p.SetHandle(handle)
	return p
}

// 从name创建对象, 失败返回nil.
func NewShapeByName(name string) *Shape {
	handle := xc.XC_GetObjectByName(name)
	if handle > 0 {
		p := &Shape{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 从UID创建对象, 失败返回nil.
func NewShapeByUID(nUID int) *Shape {
	handle := xc.XC_GetObjectByUID(nUID)
	if handle > 0 {
		p := &Shape{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 从UID名称创建对象, 失败返回nil.
func NewShapeByUIDName(name string) *Shape {
	handle := xc.XC_GetObjectByUIDName(name)
	if handle > 0 {
		p := &Shape{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 形状_移除, 从父UI元素或窗口,和父布局对象中移除.
func (s *Shape) RemoveShape() int {
	return xc.XShape_RemoveShape(s.Handle)
}

// 形状_取Z序, 获取形状对象Z序, 成功返回索引值, 否则返回 XC_ID_ERROR.
func (s *Shape) GetZOrder() int {
	return xc.XShape_GetZOrder(s.Handle)
}

// 形状_重绘, 重绘形状对象.
func (s *Shape) Redraw() int {
	return xc.XShape_Redraw(s.Handle)
}

// 形状_取宽度, 获取内容宽度.
func (s *Shape) GetWidth() int32 {
	return xc.XShape_GetWidth(s.Handle)
}

// 形状_取高度, 获取内容高度.
func (s *Shape) GetHeight() int32 {
	return xc.XShape_GetHeight(s.Handle)
}

// 形状_移动位置.
//
// x: x坐标.
//
// y: y坐标.
func (s *Shape) SetPosition(x, y int32) int {
	return xc.XShape_SetPosition(s.Handle, x, y)
}

// 形状_取坐标.
//
// pRect: 接收返回坐标.
func (s *Shape) GetRect(pRect *xc.RECT) int {
	return xc.XShape_GetRect(s.Handle, pRect)
}

// 形状_置坐标.
//
// pRect: 坐标.
func (s *Shape) SetRect(pRect *xc.RECT) int {
	return xc.XShape_SetRect(s.Handle, pRect)
}

// 形状_置逻辑坐标, 设置元素坐标, 逻辑坐标, 包含滚动视图偏移.
//
// pRect: 坐标.
//
// bRedraw: 是否重绘.
func (s *Shape) SetRectLogic(pRect *xc.RECT, bRedraw bool) bool {
	return xc.XShape_SetRectLogic(s.Handle, pRect, bRedraw)
}

// 形状_取逻辑坐标, 获取元素坐标, 逻辑坐标, 包含滚动视图偏移.
//
// pRect: 坐标.
func (s *Shape) GetRectLogic(pRect *xc.RECT) int {
	return xc.XShape_GetRectLogic(s.Handle, pRect)
}

// 形状_取基于窗口客户区坐标, 基于窗口客户区坐标.
//
// pRect: 坐标.
func (s *Shape) GetWndClientRect(pRect *xc.RECT) int {
	return xc.XShape_GetWndClientRect(s.Handle, pRect)
}

// 形状_取内容大小 ,仅计算有效内容, 填充父, 权重依赖父级所以无法计算.
//
// pSize: 接收返回内容大小值.
func (s *Shape) GetContentSize(pSize *xc.SIZE) int {
	return xc.XShape_GetContentSize(s.Handle, pSize)
}

// 形状_显示布局边界, 是否显示布局边界.
//
// bShow: 是否显示.
func (s *Shape) ShowLayout(bShow bool) int {
	return xc.XShape_ShowLayout(s.Handle, bShow)
}

// 形状_调整布局.
func (s *Shape) AdjustLayout() int {
	return xc.XShape_AdjustLayout(s.Handle)
}

// 形状_销毁, 销毁形状对象.
func (s *Shape) Destroy() int {
	return xc.XShape_Destroy(s.Handle)
}

// 形状_取位置.
//
// pOutX: 返回X坐标.
//
// pOutY: 返回Y坐标.
func (s *Shape) GetPosition(pOutX, pOutY *int32) int {
	return xc.XShape_GetPosition(s.Handle, pOutX, pOutY)
}

// 形状_置大小.
//
// nWidth: 宽度.
//
// nHeight: 高度.
func (s *Shape) SetSize(nWidth, nHeight int32) int {
	return xc.XShape_SetSize(s.Handle, nWidth, nHeight)
}

// 形状_取大小.
//
// pOutWidth: 返回宽度.
//
// pOutHeight: 返回高度.
func (s *Shape) GetSize(pOutWidth, pOutHeight *int32) int {
	return xc.XShape_GetSize(s.Handle, pOutWidth, pOutHeight)
}

// 形状_置透明度.
//
// alpha: 透明度.
func (s *Shape) SetAlpha(alpha uint8) int {
	return xc.XShape_SetAlpha(s.Handle, alpha)
}

// 形状_取透明度, 返回透明度.
func (s *Shape) GetAlpha() int {
	return xc.XShape_GetAlpha(s.Handle)
}
