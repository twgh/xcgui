package xc

import (
	"github.com/twgh/xcgui/common"
	"syscall"

	"github.com/twgh/xcgui/xcc"
)

// 动画_运行, 并且增加引用计数.
//
// hAnimation: 动画序列或动画组句柄.
//
// hRedrawObjectUI: 当更新UI时重绘的UI层. UI对象句柄: 窗口句柄, 元素句柄, 形状句柄, SVG句柄.
func XAnima_Run(hAnimation int, hRedrawObjectUI int) int {
	r, _, _ := xAnima_Run.Call(uintptr(hAnimation), uintptr(hRedrawObjectUI))
	return int(r)
}

// 动画_释放, 停止动画, 并释放引用, 当引用计数为0时自动销毁.
//
// hAnimation: 动画序列或动画组句柄.
//
// bEnd: 是否立即执行到终点.
func XAnima_Release(hAnimation int, bEnd bool) bool {
	r, _, _ := xAnima_Release.Call(uintptr(hAnimation), common.BoolPtr(bEnd))
	return r != 0
}

// 动画_释放扩展, 释放与指定UI对象关联的所有动画, 返回释放动画数量.
//
// hObjectUI: 指定UI对象句柄.
//
// bEnd: 是否立即执行到终点.
func XAnima_ReleaseEx(hObjectUI int, bEnd bool) int {
	r, _, _ := xAnima_ReleaseEx.Call(uintptr(hObjectUI), common.BoolPtr(bEnd))
	return int(r)
}

// 动画_创建动画序列, 按顺序执行的动画列表, 返回动画序列句柄.
//
// hObjectUI: 绑定的UI对象. UI对象句柄: 窗口句柄, 元素句柄, 形状句柄, SVG句柄.
//
// nLoopCount: 动画循环次数, 0: 无限循环.
func XAnima_Create(hObjectUI int, nLoopCount int) int {
	r, _, _ := xAnima_Create.Call(uintptr(hObjectUI), uintptr(nLoopCount))
	return int(r)
}

// 动画_移动, 移动到目标位置, 默认以UI对象中心点为操作方式, 避免出现坐标错位, 返回动画项句柄.
//
// hSequence: 动画序列句柄.
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
func XAnima_Move(hSequence int, duration int, x float32, y float32, nLoopCount int, ease_flag xcc.Ease_Flag_, bGoBack bool) int {
	r, _, _ := xAnima_Move.Call(uintptr(hSequence), uintptr(duration), common.Float32Ptr(x), common.Float32Ptr(y), uintptr(nLoopCount), uintptr(ease_flag), common.BoolPtr(bGoBack))
	return int(r)
}

// 动画_移动扩展, 从指定位置移动到目标位置, 默认以UI对象中心点为操作方式, 避免出现坐标错位, 返回动画项句柄.
//
// hSequence: 动画序列句柄.
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
func XAnima_MoveEx(hSequence int, duration int, from_x float32, from_y float32, to_x float32, to_y float32, nLoopCount int, ease_flag xcc.Ease_Flag_, bGoBack bool) int {
	r, _, _ := xAnima_MoveEx.Call(uintptr(hSequence), uintptr(duration), common.Float32Ptr(from_x), common.Float32Ptr(from_y), common.Float32Ptr(to_x), common.Float32Ptr(to_y), uintptr(nLoopCount), uintptr(ease_flag), common.BoolPtr(bGoBack))
	return int(r)
}

// 动画_旋转, 旋转角度支持负数值, 因为负数可以控制反向旋转, 返回动画旋转项句柄.
//
// hSequence: 动画序列句柄.
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
func XAnima_Rotate(hSequence int, duration int, angle float32, nLoopCount int, ease_flag xcc.Ease_Flag_, bGoBack bool) int {
	r, _, _ := xAnima_Rotate.Call(uintptr(hSequence), uintptr(duration), common.Float32Ptr(angle), uintptr(nLoopCount), uintptr(ease_flag), common.BoolPtr(bGoBack))
	return int(r)
}

// 动画_旋转扩展, 指定起点和终点, 返回动画旋转项句柄.
//
// hSequence: 动画序列句柄.
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
func XAnima_RotateEx(hSequence int, duration int, from float32, to float32, nLoopCount int, ease_flag xcc.Ease_Flag_, bGoBack bool) int {
	r, _, _ := xAnima_RotateEx.Call(uintptr(hSequence), uintptr(duration), common.Float32Ptr(from), common.Float32Ptr(to), uintptr(nLoopCount), uintptr(ease_flag), common.BoolPtr(bGoBack))
	return int(r)
}

// 动画_缩放, 缩放对象, 默认以自身为中心缩放, 返回动画缩放项句柄.
//
// hSequence: 动画序列句柄.
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
func XAnima_Scale(hSequence int, duration int, scaleX float32, scaleY float32, nLoopCount int, ease_flag xcc.Ease_Flag_, bGoBack bool) int {
	r, _, _ := xAnima_Scale.Call(uintptr(hSequence), uintptr(duration), common.Float32Ptr(scaleX), common.Float32Ptr(scaleY), uintptr(nLoopCount), uintptr(ease_flag), common.BoolPtr(bGoBack))
	return int(r)
}

// 动画_缩放大小, 修改UI对象大小, 默认向右延伸, 返回动画缩放项句柄.
//
// hSequence: 动画序列句柄.
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
func XAnima_ScaleSize(hSequence int, duration int, width float32, height float32, nLoopCount int, ease_flag xcc.Ease_Flag_, bGoBack bool) int {
	r, _, _ := xAnima_ScaleSize.Call(uintptr(hSequence), uintptr(duration), common.Float32Ptr(width), common.Float32Ptr(height), uintptr(nLoopCount), uintptr(ease_flag), common.BoolPtr(bGoBack))
	return int(r)
}

// 动画_透明度, 返回动画项句柄.
//
// hSequence: 动画序列句柄.
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
func XAnima_Alpha(hSequence int, duration int, alpha uint8, nLoopCount int, ease_flag xcc.Ease_Flag_, bGoBack bool) int {
	r, _, _ := xAnima_Alpha.Call(uintptr(hSequence), uintptr(duration), uintptr(alpha), uintptr(nLoopCount), uintptr(ease_flag), common.BoolPtr(bGoBack))
	return int(r)
}

// 动画_透明度扩展, 从指定透明度到目标透明度, 返回动画项句柄.
//
// hSequence: 动画序列句柄.
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
func XAnima_AlphaEx(hSequence int, duration int, from_alpha uint8, to_alpha uint8, nLoopCount int, ease_flag xcc.Ease_Flag_, bGoBack bool) int {
	r, _, _ := xAnima_AlphaEx.Call(uintptr(hSequence), uintptr(duration), uintptr(from_alpha), uintptr(to_alpha), uintptr(nLoopCount), uintptr(ease_flag), common.BoolPtr(bGoBack))
	return int(r)
}

// 动画_颜色, 返回动画项句柄.
//
// hSequence: 动画序列句柄.
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
func XAnima_Color(hSequence int, duration int, color int, nLoopCount int, ease_flag xcc.Ease_Flag_, bGoBack bool) int {
	r, _, _ := xAnima_Color.Call(uintptr(hSequence), uintptr(duration), uintptr(color), uintptr(nLoopCount), uintptr(ease_flag), common.BoolPtr(bGoBack))
	return int(r)
}

// 动画_颜色扩展, 从指定颜色到目标颜色, 返回动画项句柄.
//
// hSequence: 动画序列句柄.
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
func XAnima_ColorEx(hSequence int, duration int, from int, to int, nLoopCount int, ease_flag xcc.Ease_Flag_, bGoBack bool) int {
	r, _, _ := xAnima_ColorEx.Call(uintptr(hSequence), uintptr(duration), uintptr(from), uintptr(to), uintptr(nLoopCount), uintptr(ease_flag), common.BoolPtr(bGoBack))
	return int(r)
}

// 动画_布局宽度, 修改布局宽度属性, 返回动画项句柄.
//
// hSequence: 动画序列句柄.
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
func XAnima_LayoutWidth(hSequence int, duration int, nType xcc.Layout_Size_, width float32, nLoopCount int, ease_flag xcc.Ease_Flag_, bGoBack bool) int {
	r, _, _ := xAnima_LayoutWidth.Call(uintptr(hSequence), uintptr(duration), uintptr(nType), common.Float32Ptr(width), uintptr(nLoopCount), uintptr(ease_flag), common.BoolPtr(bGoBack))
	return int(r)
}

// 动画_布局高度, 修改布局高度属性, 返回动画项句柄.
//
// hSequence: 动画序列句柄.
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
func XAnima_LayoutHeight(hSequence int, duration int, nType xcc.Layout_Size_, height float32, nLoopCount int, ease_flag xcc.Ease_Flag_, bGoBack bool) int {
	r, _, _ := xAnima_LayoutHeight.Call(uintptr(hSequence), uintptr(duration), uintptr(nType), common.Float32Ptr(height), uintptr(nLoopCount), uintptr(ease_flag), common.BoolPtr(bGoBack))
	return int(r)
}

// 动画_布局大小, 修改布局宽度和高度, 返回动画项句柄.
//
// hSequence: 动画序列句柄.
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
func XAnima_LayoutSize(hSequence int, duration int, nWidthType xcc.Layout_Size_, width float32, nHeightType xcc.Layout_Size_, height float32, nLoopCount int, ease_flag xcc.Ease_Flag_, bGoBack bool) int {
	r, _, _ := xAnima_LayoutSize.Call(uintptr(hSequence), uintptr(duration), uintptr(nWidthType), common.Float32Ptr(width), uintptr(nHeightType), common.Float32Ptr(height), uintptr(nLoopCount), uintptr(ease_flag), common.BoolPtr(bGoBack))
	return int(r)
}

// 动画_延迟, 返回动画项句柄.
//
// hSequence: 动画序列句柄.
//
// duration: 持续时间.
func XAnima_Delay(hSequence int, duration float32) int {
	r, _, _ := xAnima_Delay.Call(uintptr(hSequence), common.Float32Ptr(duration))
	return int(r)
}

// 动画_显示, 显示或隐藏UI对象, 返回动画项句柄.
//
// hSequence: 动画序列句柄.
//
// duration: 持续时间.
//
// bShow: 显示或隐藏.
func XAnima_Show(hSequence int, duration float32, bShow bool) int {
	r, _, _ := xAnima_Show.Call(uintptr(hSequence), common.Float32Ptr(duration), common.BoolPtr(bShow))
	return int(r)
}

// 动画组_创建, 动画同步组, 当组中动画序列全部完成后, 重新开始.
//
// 当遇到无限循环项时, 直至其他序列完成后终止循环, 避免出现无法到达终点, 无法返回头部进行同步, 返回动画组句柄.
//
// nLoopCount: 动画循环次数, 0: 无限循环.
func XAnimaGroup_Create(nLoopCount int) int {
	r, _, _ := xAnimaGroup_Create.Call(uintptr(nLoopCount))
	return int(r)
}

// 动画组_添加项, 将动画序列添加到组中.
//
// hGroup: 动画组句柄.
//
// hSequence: 动画序列句柄.
func XAnimaGroup_AddItem(hGroup int, hSequence int) int {
	r, _, _ := xAnimaGroup_AddItem.Call(uintptr(hGroup), uintptr(hSequence))
	return int(r)
}

// 动画旋转_置中心, 设置旋转中心点坐标.
//
// hAnimationRotate: 动画旋转项句柄.
//
// x: 坐标X.
//
// y: 坐标Y.
//
// bOffset: TRUE: 相对于自身中心点偏移, FALSE: 绝对坐标.
func XAnimaRotate_SetCenter(hAnimationRotate int, x float32, y float32, bOffset bool) bool {
	r, _, _ := xAnimaRotate_SetCenter.Call(uintptr(hAnimationRotate), common.Float32Ptr(x), common.Float32Ptr(y), common.BoolPtr(bOffset))
	return r != 0
}

// 动画缩放_置延伸位置, 设置缩放起点, 确定延伸方向.
//
// hAnimationScale: 动画缩放项句柄.
//
// position: 位置, Position_Flag_.
func XAnimaScale_SetPosition(hAnimationScale int, position xcc.Position_Flag_) bool {
	r, _, _ := xAnimaScale_SetPosition.Call(uintptr(hAnimationScale), uintptr(position))
	return r != 0
}

// 动画_取UI对象, 获取动画关联的UI对象, 返回UI对象句柄.
//
// hAnimation: 动画序列或动画组句柄.
func XAnima_GetObjectUI(hAnimation int) int {
	r, _, _ := xAnima_GetObjectUI.Call(uintptr(hAnimation))
	return int(r)
}

// 动画项_启用完成释放, 当动画项完成后自动释放.
//
// 例如对多个动画序列进行渐近式延迟, 在动画序列头标添加延时项(时间差), 当延时项完成时自动释放, 后续动画循环就形成一种时间差(因为对齐的时间差销毁了, 他们永远无法对齐时间).
//
// hAnimationItem: 动画项句柄.
//
// bEnable: 是否启用.
func XAnimaItem_EnableCompleteRelease(hAnimationItem int, bEnable bool) int {
	r, _, _ := xAnimaItem_EnableCompleteRelease.Call(uintptr(hAnimationItem), common.BoolPtr(bEnable))
	return int(r)
}

// 动画_启用自动销毁, TRUE: 当引用计数为0时自动销毁, FALSE: 手动销毁.
//
// hAnimation: 动画项或动画序列或动画组句柄.
//
// bEnable: 是否启用.
func XAnima_EnableAutoDestroy(hAnimation int, bEnable bool) int {
	r, _, _ := xAnima_EnableAutoDestroy.Call(uintptr(hAnimation), common.BoolPtr(bEnable))
	return int(r)
}

// 动画_销毁UI对象, 返回动画项句柄.
//
// hSequence: 动画序列句柄.
//
// duration: 持续时间.
func XAnima_DestroyObjectUI(hSequence int, duration float32) int {
	r, _, _ := xAnima_DestroyObjectUI.Call(uintptr(hSequence), common.Float32Ptr(duration))
	return int(r)
}

// 动画_置回调.
//
// hAnimationEx: 动画序列或动画组句柄.
//
// callback: 回调函数.
func XAnima_SetCallBack(hAnimationEx int, callback interface{}) int {
	r, _, _ := xAnima_SetCallBack.Call(uintptr(hAnimationEx), syscall.NewCallback(callback))
	return int(r)
}

// 动画_置用户数据.
//
// hAnimationEx: 动画序列或动画组句柄.
//
// nUserData: 用户数据.
func XAnima_SetUserData(hAnimationEx, nUserData int) int {
	r, _, _ := xAnima_SetUserData.Call(uintptr(hAnimationEx), uintptr(nUserData))
	return int(r)
}

// 动画_取用户数据, 返回用户数据.
//
// hAnimationEx: 动画序列或动画组句柄.
func XAnima_GetUserData(hAnimationEx int) int {
	r, _, _ := xAnima_GetUserData.Call(uintptr(hAnimationEx))
	return int(r)
}

// 动画_停止.
//
// hAnimationEx: 动画序列或动画组句柄.
func XAnima_Stop(hAnimationEx int) bool {
	r, _, _ := xAnima_Stop.Call(uintptr(hAnimationEx))
	return r != 0
}

// 动画_启动.
//
// hAnimationEx: 动画序列或动画组句柄.
func XAnima_Start(hAnimationEx int) bool {
	r, _, _ := xAnima_Start.Call(uintptr(hAnimationEx))
	return r != 0
}

// 动画_暂停.
//
// hAnimationEx: 动画序列或动画组句柄.
func XAnima_Pause(hAnimationEx int) bool {
	r, _, _ := xAnima_Pause.Call(uintptr(hAnimationEx))
	return r != 0
}

// 动画项_置回调.
//
// hAnimationItem: 动画项句柄.
//
// callback: 回调函数.
func XAnimaItem_SetCallback(hAnimationItem int, callback interface{}) int {
	r, _, _ := xAnimaItem_SetCallback.Call(uintptr(hAnimationItem), syscall.NewCallback(callback))
	return int(r)
}

// 动画项_置用户数据.
//
// hAnimationItem: 动画项句柄.
//
// nUserData: 用户数据.
func XAnimaItem_SetUserData(hAnimationItem, nUserData int) int {
	r, _, _ := xAnimaItem_SetUserData.Call(uintptr(hAnimationItem), uintptr(nUserData))
	return int(r)
}

// 动画项_取用户数据, 返回用户数据.
//
// hAnimationItem: 动画项句柄.
func XAnimaItem_GetUserData(hAnimationItem int) int {
	r, _, _ := xAnimaItem_GetUserData.Call(uintptr(hAnimationItem))
	return int(r)
}

// 动画项_启用自动销毁.
//
// hAnimationItem: 动画项句柄.
//
// bEnable: 是否启用.
func XAnimaItem_EnableAutoDestroy(hAnimationItem int, bEnable bool) int {
	r, _, _ := xAnimaItem_EnableAutoDestroy.Call(uintptr(hAnimationItem), common.BoolPtr(bEnable))
	return int(r)
}

// 动画_延迟扩展, 可以作为一个空动画, 然后在回调里处理自己的算法, 返回动画项句柄.
//
// hSequence: 动画序列句柄.
//
// duration: 持续时间.
//
// nLoopCount: 动画循环次数, 0:无限循环.
//
// ease_flag: 缓动标识, Ease_Flag_.
//
// bGoBack: 是否返回. 当启用后: 往返到起点, 起点->终点->起点.
func XAnima_DelayEx(hSequence int, duration float32, nLoopCount int, ease_flag xcc.Ease_Flag_, bGoBack bool) int {
	r, _, _ := xAnima_DelayEx.Call(uintptr(hSequence), common.Float32Ptr(duration), uintptr(nLoopCount), uintptr(ease_flag), common.BoolPtr(bGoBack))
	return int(r)
}

// 动画移动_置标识, 此接口可独立设置x轴移动或y轴移动.
//
// hAnimationMove: 动画移动项句柄.
//
// flags: 动画移动标识, 可组合使用, Animation_Move_.
//
// TODO: 此函数尚未封装到类中.
func XAnimaMove_SetFlag(hAnimationMove int, flags xcc.Animation_Move_) int {
	r, _, _ := xAnimaMove_SetFlag.Call(uintptr(hAnimationMove), uintptr(flags))
	return int(r)
}
