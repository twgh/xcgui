package widget

import (
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

// ProgressBar 进度条.
type ProgressBar struct {
	Element
}

// 进度条_创建, 创建进度条元素.
//
// x: 元素x坐标.
//
// y: 元素y坐标.
//
// cx: 宽度.
//
// cy: 高度.
//
// hParent: 父是窗口资源句柄或UI元素资源句柄.如果是窗口资源句柄将被添加到窗口.
func NewProgressBar(x, y, cx, cy int32, hParent int) *ProgressBar {
	p := &ProgressBar{}
	p.SetHandle(xc.XProgBar_Create(x, y, cx, cy, hParent))
	return p
}

// 从句柄创建对象.
func NewProgressBarByHandle(handle int) *ProgressBar {
	p := &ProgressBar{}
	p.SetHandle(handle)
	return p
}

// 从name创建对象, 失败返回nil.
func NewProgressBarByName(name string) *ProgressBar {
	handle := xc.XC_GetObjectByName(name)
	if handle > 0 {
		p := &ProgressBar{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 从UID创建对象, 失败返回nil.
func NewProgressBarByUID(nUID int32) *ProgressBar {
	handle := xc.XC_GetObjectByUID(nUID)
	if handle > 0 {
		p := &ProgressBar{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 从UID名称创建对象, 失败返回nil.
func NewProgressBarByUIDName(name string) *ProgressBar {
	handle := xc.XC_GetObjectByUIDName(name)
	if handle > 0 {
		p := &ProgressBar{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 进度条_置范围, 设置范围.
//
// range_: 范围.
func (p *ProgressBar) SetRange(range_ int32) *ProgressBar {
	xc.XProgBar_SetRange(p.Handle, range_)
	return p
}

// 进度条_取范围.
func (p *ProgressBar) GetRange() int32 {
	return xc.XProgBar_GetRange(p.Handle)
}

// 进度条_置进度图片.
//
// hImage: 图片句柄.
func (p *ProgressBar) SetImageLoad(hImage int) *ProgressBar {
	xc.XProgBar_SetImageLoad(p.Handle, hImage)
	return p
}

// 进度条_置进度, 设置位置点.
//
// pos: 位置点.
func (p *ProgressBar) SetPos(pos int32) *ProgressBar {
	xc.XProgBar_SetPos(p.Handle, pos)
	return p
}

// 进度条_取进度, 获取当前位置点.
func (p *ProgressBar) GetPos() int32 {
	return xc.XProgBar_GetPos(p.Handle)
}

// 进度条_置水平, 设置水平或垂直.
//
// bHorizon: 水平或垂直.
func (p *ProgressBar) EnableHorizon(bHorizon bool) *ProgressBar {
	xc.XProgBar_EnableHorizon(p.Handle, bHorizon)
	return p
}

// 进度条_启用缩放, 缩放进度贴图为当前进度区域(当前进度所显示区域), 否则为整体100进度区域.
//
// bStretch: 缩放.
func (p *ProgressBar) EnableStretch(bStretch bool) *ProgressBar {
	xc.XProgBar_EnableStretch(p.Handle, bStretch)
	return p
}

// 进度条_启用进度文本 显示进度值文本.
//
// bShow: 是否启用.
func (p *ProgressBar) EnableShowText(bShow bool) *ProgressBar {
	xc.XProgBar_EnableShowText(p.Handle, bShow)
	return p
}

// 进度条_置进度颜色. 设置进度颜色.
//
// color: xc.RGBA 颜色.
func (p *ProgressBar) SetColorLoad(color int) *ProgressBar {
	xc.XProgBar_SetColorLoad(p.Handle, color)
	return p
}

// ------------------------- AddEvent ------------------------- //

// AddEvent_ProgressBar_Change 添加进度条元素进度改变事件.
//
// pFun: 回调函数.
//
// allowAddingMultiple: 允许添加多个回调函数.
func (p *ProgressBar) AddEvent_ProgressBar_Change(pFun XE_PROGRESSBAR_CHANGE1, allowAddingMultiple ...bool) int {
	return xc.EleEventHandler.AddCallBack(p.Handle, xcc.XE_PROGRESSBAR_CHANGE, onXE_PROGRESSBAR_CHANGE, pFun, allowAddingMultiple...)
}

// onXE_PROGRESSBAR_CHANGE 进度条元素进度改变事件.
func onXE_PROGRESSBAR_CHANGE(hEle int, pos int32, pbHandled *bool) int {
	cbs := xc.EleEventHandler.GetCallBacks(hEle, xcc.XE_PROGRESSBAR_CHANGE)
	var ret int
	for i := len(cbs) - 1; i >= 0; i-- {
		ret = cbs[i].(XE_PROGRESSBAR_CHANGE1)(hEle, pos, pbHandled)
		if *pbHandled {
			break
		}
	}
	return ret
}

// ------------------------- 事件 ------------------------- //

type XE_PROGRESSBAR_CHANGE func(pos int32, pbHandled *bool) int            // 进度条元素,进度改变事件.
type XE_PROGRESSBAR_CHANGE1 func(hEle int, pos int32, pbHandled *bool) int // 进度条元素,进度改变事件.

// 进度条元素,进度改变事件.
func (p *ProgressBar) Event_PROGRESSBAR_CHANGE(pFun XE_PROGRESSBAR_CHANGE) bool {
	return xc.XEle_RegEventC(p.Handle, xcc.XE_PROGRESSBAR_CHANGE, pFun)
}

// 进度条元素,进度改变事件.
func (p *ProgressBar) Event_PROGRESSBAR_CHANGE1(pFun XE_PROGRESSBAR_CHANGE1) bool {
	return xc.XEle_RegEventC1(p.Handle, xcc.XE_PROGRESSBAR_CHANGE, pFun)
}
