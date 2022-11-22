package ani

import (
	"github.com/twgh/xcgui/objectbase"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

// AnimaScale 动画缩放项.
type AnimaScale struct {
	objectbase.ObjectBase
}

// 动画缩放_置延伸位置, 设置缩放起点, 确定延伸方向.
//
// position: 位置, Position_Flag_.
func (a *AnimaScale) SetPosition(position xcc.Position_Flag_) bool {
	return xc.XAnimaScale_SetPosition(a.Handle, position)
}
