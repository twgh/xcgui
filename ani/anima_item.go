package ani

import (
	"github.com/twgh/xcgui/objectbase"
	"github.com/twgh/xcgui/xc"
)

// AnimaItem 动画项.
type AnimaItem struct {
	objectbase.ObjectBase
}

// 动画项_启用完成释放, 当动画项完成后自动释放.
//
// 例如对多个动画序列进行渐近式延迟, 在动画序列头标添加延时项(时间差), 当延时项完成时自动释放, 后续动画循环就形成一种时间差(因为对齐的时间差销毁了, 他们永远无法对齐时间).
//
// bEnable: 是否启用.
func (a *AnimaItem) EnableCompleteRelease(bEnable bool) *AnimaItem {
	xc.XAnimaItem_EnableCompleteRelease(a.Handle, bEnable)
	return a
}

// 动画项_置回调. TODO: 有问题用不了, 因为 syscall.NewCallback 创建不了没有返回值的回调函数指针.
//
// callback: 回调函数.
func (a *AnimaItem) SetCallback(callback xc.FunAnimationItem) *AnimaItem {
	xc.XAnimaItem_SetCallback(a.Handle, callback)
	return a
}

// 动画项_置用户数据.
//
// nUserData: 用户数据.
func (a *AnimaItem) SetUserData(nUserData int) *AnimaItem {
	xc.XAnimaItem_SetUserData(a.Handle, nUserData)
	return a
}

// 动画项_取用户数据, 返回用户数据.
func (a *AnimaItem) GetUserData() int {
	return xc.XAnimaItem_GetUserData(a.Handle)
}

// 动画项_启用自动销毁.
//
// bEnable: 是否启用.
func (a *AnimaItem) EnableAutoDestroy(bEnable bool) *AnimaItem {
	xc.XAnimaItem_EnableAutoDestroy(a.Handle, bEnable)
	return a
}
