package bkobj

import (
	"github.com/twgh/xcgui/bkmanager"
)

// Deprecated
//
// !已废弃, 请使用 bkobj.NewByHandle().
func NewBkObjByHandle(handle int) *BkObj {
	return NewByHandle(handle)
}

// Deprecated
//
// !已废弃, 请使用 bkobj.NewByBkm().
func NewBkObjByBkManager(bkm *bkmanager.BkManager, id int) *BkObj {
	return NewByBkm(bkm, id)
}

// Deprecated
//
// !已废弃, 请使用 bkobj.NewByBkmHandle().
func NewBkObjByBkManagerHandle(hBkm, id int) *BkObj {
	return NewByBkmHandle(hBkm, id)
}
