package ani

import (
	"github.com/twgh/xcgui/xc"
)

// AnimaGroup 动画组.
type AnimaGroup struct {
	animaBase
}

// 动画组_创建, 动画同步组, 当组中动画序列全部完成后, 重新开始.
//
// 当遇到无限循环项时, 直至其他序列完成后终止循环, 避免出现无法到达终点, 无法返回头部进行同步, 返回动画组句柄.
//
// nLoopCount: 动画循环次数, 0: 无限循环.
func NewAnimaGroup(nLoopCount int) *AnimaGroup {
	p := &AnimaGroup{}
	p.SetHandle(xc.XAnimaGroup_Create(nLoopCount))
	return p
}

// 动画组_添加项, 将动画序列添加到组中.
//
// hSequence: 动画序列句柄.
func (a *AnimaGroup) AddItem(hSequence int) int {
	return xc.XAnimaGroup_AddItem(a.Handle, hSequence)
}
