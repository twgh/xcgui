package app

import (
	"github.com/twgh/xcgui/bkmanager"
	"github.com/twgh/xcgui/bkobj"
)

// NewBkObjByHandle 从BkObj句柄创建BkObj对象.
func NewBkObjByHandle(handle int) *bkobj.BkObj {
	return bkobj.NewByHandle(handle)
}

// NewBkObjByBkm 从BkManager对象创建BkObj对象, 失败返回nil.
//
// bkm: 背景管理器对象.
//
// id: 背景对象ID.
func NewBkObjByBkm(bkm *bkmanager.BkManager, id int32) *bkobj.BkObj {
	return bkobj.NewByBkm(bkm, id)
}

// NewBkObjByBkmHandle 从BkManager句柄创建BkObj对象, 失败返回nil.
//
// hBkm: 背景管理器句柄.
//
// id: 背景对象ID.
func NewBkObjByBkmHandle(hBkm int, id int32) *bkobj.BkObj {
	return bkobj.NewByBkmHandle(hBkm, id)
}
