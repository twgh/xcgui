package app

import (
	"github.com/twgh/xcgui/bkobj"
)

// NewBkObjByHandle 从 BkObj 句柄创建 bkobj.BkObj 对象.
func NewBkObjByHandle(handle int) *bkobj.BkObj {
	return bkobj.NewByHandle(handle)
}
