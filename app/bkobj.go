package app

import (
	"github.com/twgh/xcgui/bkobj"
)

// NewBkObjByHandle 从 BkObj 句柄创建 bkobj.BkObj 对象.
func NewBkObjByHandle(handle int) *bkobj.BkObj {
	return bkobj.NewByHandle(handle)
}

// NewBkObjByBkmHandle 从 BkManager 句柄创建 bkobj.BkObj 对象, 失败返回 nil.
//
// hBkm: 背景管理器句柄.
//
// id: 背景对象ID.
func NewBkObjByBkmHandle(hBkm int, id int32) *bkobj.BkObj {
	return bkobj.NewByBkmHandle(hBkm, id)
}
