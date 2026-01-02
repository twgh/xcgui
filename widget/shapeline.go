package widget

import (
	"github.com/twgh/xcgui/xc"
)

// ShapeLine 直线(形状对象).
type ShapeLine struct {
	Shape
}

// 形状线_创建, 失败返回 nil.
//
// x1: 坐标.
//
// y1: 坐标.
//
// x2: 坐标.
//
// y2: 坐标.
//
// hParent: 父对象句柄.
func NewShapeLine(x1, y1, x2, y2 int32, hParent int) *ShapeLine {
	return NewShapeLineByHandle(xc.XShapeLine_Create(x1, y1, x2, y2, hParent))
}

// 从句柄创建对象, 失败返回 nil.
func NewShapeLineByHandle(handle int) *ShapeLine {
	if handle <= 0 {
		return nil
	}
	p := &ShapeLine{}
	p.SetHandle(handle)
	return p
}

// 从name创建对象, 失败返回 nil.
func NewShapeLineByName(name string) *ShapeLine {
	return NewShapeLineByHandle(xc.XC_GetObjectByName(name))
}

// 从UID创建对象, 失败返回 nil.
func NewShapeLineByUID(nUID int32) *ShapeLine {
	return NewShapeLineByHandle(xc.XC_GetObjectByUID(nUID))
}

// 从UID名称创建对象, 失败返回 nil.
func NewShapeLineByUIDName(name string) *ShapeLine {
	return NewShapeLineByHandle(xc.XC_GetObjectByUIDName(name))
}

// 形状线_置位置.
//
// x1: 坐标.
//
// y1: 坐标.
//
// x2: 坐标.
//
// y2: 坐标.
func (s *ShapeLine) SetPosition(x1, y1, x2, y2 int32) *ShapeLine {
	xc.XShapeLine_SetPosition(s.Handle, x1, y1, x2, y2)
	return s
}

// 形状线_置颜色, 设置直线颜色.
//
// color: xc.RGBA 颜色值.
func (s *ShapeLine) SetColor(color uint32) *ShapeLine {
	xc.XShapeLine_SetColor(s.Handle, color)
	return s
}
