package ani

import (
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

// Anima 动画序列.
type Anima struct {
	animaBase
}

// 动画_创建动画序列, 按顺序执行的动画列表, 返回动画序列对象.
//
// hObjectUI: 绑定的UI对象. UI对象句柄: 窗口句柄, 元素句柄, 形状句柄, SVG句柄.
//
// nLoopCount: 动画循环次数, 0: 无限循环.
func NewAnima(hObjectUI int, nLoopCount int) *Anima {
	p := &Anima{}
	p.SetHandle(xc.XAnima_Create(hObjectUI, nLoopCount))
	return p
}

// 动画_移动, 移动到目标位置, 默认以UI对象中心点为操作方式, 避免出现坐标错位, 返回动画项对象.
//
// duration: 持续时间.
//
// x: 终点位置X(对象左上角坐标).
//
// y: 终点位置Y(对象左上角坐标).
//
// nLoopCount: 动画循环次数, 0:无限循环.
//
// ease_flag: 缓动标识, Ease_Flag_.
//
// bGoBack: 是否返回. 当启用后: 往返到起点, 起点->终点->起点.
func (a *Anima) Move(duration int, x float32, y float32, nLoopCount int, ease_flag xcc.Ease_Flag_, bGoBack bool) *AnimaItem {
	p := &AnimaItem{}
	p.SetHandle(xc.XAnima_Move(a.Handle, duration, x, y, nLoopCount, ease_flag, bGoBack))
	return p
}

// 动画_移动扩展, 从指定位置移动到目标位置, 默认以UI对象中心点为操作方式, 避免出现坐标错位, 返回动画项对象.
//
// duration: 持续时间.
//
// from_x: 起点位置X(对象左上角坐标).
//
// from_y: 起点位置Y(对象左上角坐标).
//
// to_x: 终点位置X(对象左上角坐标).
//
// to_y: 终点位置Y(对象左上角坐标).
//
// nLoopCount: 动画循环次数, 0:无限循环.
//
// ease_flag: 缓动标识, Ease_Flag_.
//
// bGoBack: 是否返回. 当启用后: 往返到起点, 起点->终点->起点.
func (a *Anima) MoveEx(duration int, from_x float32, from_y float32, to_x float32, to_y float32, nLoopCount int, ease_flag xcc.Ease_Flag_, bGoBack bool) *AnimaItem {
	p := &AnimaItem{}
	p.SetHandle(xc.XAnima_MoveEx(a.Handle, duration, from_x, from_y, to_x, to_y, nLoopCount, ease_flag, bGoBack))
	return p
}

// 动画_旋转, 旋转角度支持负数值, 因为负数可以控制反向旋转, 返回动画旋转项对象.
//
// duration: 持续时间.
//
// angle: 角度.
//
// nLoopCount: 动画循环次数, 0: 无限循环.
//
// ease_flag: 缓动标识, Ease_Flag_.
//
// bGoBack: 是否返回. 当启用后: 往返到起点, 起点->终点->起点.
func (a *Anima) Rotate(duration int, angle float32, nLoopCount int, ease_flag xcc.Ease_Flag_, bGoBack bool) *AnimaRotate {
	p := &AnimaRotate{}
	p.SetHandle(xc.XAnima_Rotate(a.Handle, duration, angle, nLoopCount, ease_flag, bGoBack))
	return p
}

// 动画_旋转扩展, 指定起点和终点, 返回动画旋转项对象.
//
// duration: 持续时间.
//
// from: 起点角度.
//
// to: 终点角度.
//
// nLoopCount: 动画循环次数, 0: 无限循环.
//
// ease_flag: 缓动标识, Ease_Flag_.
//
// bGoBack: 是否返回. 当启用后: 往返到起点, 起点->终点->起点.
func (a *Anima) RotateEx(duration int, from float32, to float32, nLoopCount int, ease_flag xcc.Ease_Flag_, bGoBack bool) *AnimaRotate {
	p := &AnimaRotate{}
	p.SetHandle(xc.XAnima_RotateEx(a.Handle, duration, from, to, nLoopCount, ease_flag, bGoBack))
	return p
}

// 动画_缩放, 缩放对象, 默认以自身为中心缩放, 返回动画缩放项对象.
//
// duration: 持续时间.
//
// scaleX: X轴缩放比例.
//
// scaleY: Y轴缩放比例.
//
// nLoopCount: 动画循环次数, 0: 无限循环.
//
// ease_flag: 缓动标识, Ease_Flag_.
//
// bGoBack: 是否返回. 当启用后: 往返到起点, 起点->终点->起点.
func (a *Anima) Scale(duration int, scaleX float32, scaleY float32, nLoopCount int, ease_flag xcc.Ease_Flag_, bGoBack bool) *AnimaScale {
	p := &AnimaScale{}
	p.SetHandle(xc.XAnima_Scale(a.Handle, duration, scaleX, scaleY, nLoopCount, ease_flag, bGoBack))
	return p
}

// 动画_缩放大小, 修改UI对象大小, 默认向右延伸, 返回动画缩放项对象.
//
// duration: 持续时间.
//
// width: 宽度.
//
// height: 高度.
//
// nLoopCount: 动画循环次数, 0: 无限循环.
//
// ease_flag: 缓动标识, Ease_Flag_.
//
// bGoBack: 是否返回. 当启用后: 往返到起点, 起点->终点->起点.
func (a *Anima) ScaleSize(duration int, width float32, height float32, nLoopCount int, ease_flag xcc.Ease_Flag_, bGoBack bool) *AnimaScale {
	p := &AnimaScale{}
	p.SetHandle(xc.XAnima_ScaleSize(a.Handle, duration, width, height, nLoopCount, ease_flag, bGoBack))
	return p
}

// 动画_透明度, 返回动画项对象.
//
// duration: 持续时间.
//
// alpha: 透明度.
//
// nLoopCount: 动画循环次数, 0: 无限循环.
//
// ease_flag: 缓动标识, Ease_Flag_.
//
// bGoBack: 是否返回. 当启用后: 往返到起点, 起点->终点->起点.
func (a *Anima) Alpha(duration int, alpha uint8, nLoopCount int, ease_flag xcc.Ease_Flag_, bGoBack bool) *AnimaItem {
	p := &AnimaItem{}
	p.SetHandle(xc.XAnima_Alpha(a.Handle, duration, alpha, nLoopCount, ease_flag, bGoBack))
	return p
}

// 动画_透明度扩展, 从指定透明度到目标透明度, 返回动画项对象.
//
// duration: 持续时间.
//
// from_alpha: 起始透明度.
//
// to_alpha: 终止透明度.
//
// nLoopCount: 动画循环次数, 0: 无限循环.
//
// ease_flag: 缓动标识, Ease_Flag_.
//
// bGoBack: 是否返回. 当启用后: 往返到起点, 起点->终点->起点.
func (a *Anima) AlphaEx(duration int, from_alpha uint8, to_alpha uint8, nLoopCount int, ease_flag xcc.Ease_Flag_, bGoBack bool) *AnimaItem {
	p := &AnimaItem{}
	p.SetHandle(xc.XAnima_AlphaEx(a.Handle, duration, from_alpha, to_alpha, nLoopCount, ease_flag, bGoBack))
	return p
}

// 动画_颜色, 返回动画项对象.
//
// duration: 持续时间.
//
// color: ABGR 颜色.
//
// nLoopCount: 动画循环次数, 0: 无限循环.
//
// ease_flag: 缓动标识, Ease_Flag_.
//
// bGoBack: 是否返回. 当启用后: 往返到起点, 起点->终点->起点.
func (a *Anima) Color(duration int, color int, nLoopCount int, ease_flag xcc.Ease_Flag_, bGoBack bool) *AnimaItem {
	p := &AnimaItem{}
	p.SetHandle(xc.XAnima_Color(a.Handle, duration, color, nLoopCount, ease_flag, bGoBack))
	return p
}

// 动画_颜色扩展, 从指定颜色到目标颜色, 返回动画项对象.
//
// duration: 持续时间.
//
// from: 起点颜色, ABGR 颜色.
//
// to: 终点颜色, ABGR 颜色.
//
// nLoopCount: 动画循环次数, 0: 无限循环.
//
// ease_flag: 缓动标识, Ease_Flag_.
//
// bGoBack: 是否返回. 当启用后: 往返到起点, 起点->终点->起点.
func (a *Anima) ColorEx(duration int, from int, to int, nLoopCount int, ease_flag xcc.Ease_Flag_, bGoBack bool) *AnimaItem {
	p := &AnimaItem{}
	p.SetHandle(xc.XAnima_ColorEx(a.Handle, duration, from, to, nLoopCount, ease_flag, bGoBack))
	return p
}

// 动画_布局宽度, 修改布局宽度属性, 返回动画项对象.
//
// duration: 持续时间.
//
// nType: 布局宽度类型: xcc.Layout_Size_.
//
// width: 布局宽度.
//
// nLoopCount: 动画循环次数, 0: 无限循环.
//
// ease_flag: 缓动标识, Ease_Flag_.
//
// bGoBack: 是否返回. 当启用后: 往返到起点, 起点->终点->起点.
func (a *Anima) LayoutWidth(duration int, nType xcc.Layout_Size_, width float32, nLoopCount int, ease_flag xcc.Ease_Flag_, bGoBack bool) *AnimaItem {
	p := &AnimaItem{}
	p.SetHandle(xc.XAnima_LayoutWidth(a.Handle, duration, nType, width, nLoopCount, ease_flag, bGoBack))
	return p
}

// 动画_布局高度, 修改布局高度属性, 返回动画项对象.
//
// duration: 持续时间.
//
// nType: 布局高度类型: xcc.Layout_Size_.
//
// height: 布局高度.
//
// nLoopCount: 动画循环次数, 0: 无限循环.
//
// ease_flag: 缓动标识, Ease_Flag_.
//
// bGoBack: 是否返回. 当启用后: 往返到起点, 起点->终点->起点.
func (a *Anima) LayoutHeight(duration int, nType xcc.Layout_Size_, height float32, nLoopCount int, ease_flag xcc.Ease_Flag_, bGoBack bool) *AnimaItem {
	p := &AnimaItem{}
	p.SetHandle(xc.XAnima_LayoutHeight(a.Handle, duration, nType, height, nLoopCount, ease_flag, bGoBack))
	return p
}

// 动画_布局大小, 修改布局宽度和高度, 返回动画项对象.
//
// duration: 持续时间.
//
// nWidthType: 布局宽度类型: xcc.Layout_Size_.
//
// width: 布局宽度.
//
// nHeightType: 布局高度类型: xcc.Layout_Size_.
//
// height: 布局高度.
//
// nLoopCount: 动画循环次数, 0: 无限循环.
//
// ease_flag: 缓动标识, Ease_Flag_.
//
// bGoBack: 是否返回. 当启用后: 往返到起点, 起点->终点->起点.
func (a *Anima) LayoutSize(duration int, nWidthType xcc.Layout_Size_, width float32, nHeightType xcc.Layout_Size_, height float32, nLoopCount int, ease_flag xcc.Ease_Flag_, bGoBack bool) *AnimaItem {
	p := &AnimaItem{}
	p.SetHandle(xc.XAnima_LayoutSize(a.Handle, duration, nWidthType, width, nHeightType, height, nLoopCount, ease_flag, bGoBack))
	return p
}

// 动画_延迟, 返回动画项对象.
//
// duration: 持续时间.
func (a *Anima) Delay(duration float32) *AnimaItem {
	p := &AnimaItem{}
	p.SetHandle(xc.XAnima_Delay(a.Handle, duration))
	return p
}

// 动画_显示, 显示或隐藏UI对象, 返回动画项对象.
//
// duration: 持续时间.
//
// bShow: 显示或隐藏.
func (a *Anima) Show(duration float32, bShow bool) *AnimaItem {
	p := &AnimaItem{}
	p.SetHandle(xc.XAnima_Show(a.Handle, duration, bShow))
	return p
}

// 动画_销毁UI对象, 返回动画项对象.
//
// duration: 持续时间.
func (a *Anima) DestroyObjectUI(duration float32) *AnimaItem {
	p := &AnimaItem{}
	p.SetHandle(xc.XAnima_DestroyObjectUI(a.Handle, duration))
	return p
}

// 动画_延迟扩展, 可以作为一个空动画, 然后在回调里处理自己的算法, 返回动画项句柄.
//
// duration: 持续时间.
//
// nLoopCount: 动画循环次数, 0:无限循环.
//
// ease_flag: 缓动标识, Ease_Flag_.
//
// bGoBack: 是否返回. 当启用后: 往返到起点, 起点->终点->起点.
func (a *Anima) DelayEx(duration float32, nLoopCount int, ease_flag xcc.Ease_Flag_, bGoBack bool) int {
	return xc.XAnima_DelayEx(a.Handle, duration, nLoopCount, ease_flag, bGoBack)
}
