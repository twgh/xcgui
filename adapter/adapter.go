package adapter

import (
	"github.com/twgh/xcgui/objectbase"
	"github.com/twgh/xcgui/xc"
)

// adapter 数据适配器.
type adapter struct {
	objectbase.ObjectBase
}

// 数据适配器_增加引用计数.
func (a *adapter) AddRef() int32 {
	return xc.XAd_AddRef(a.Handle)
}

// 数据适配器_释放引用计数.
func (a *adapter) Release() int32 {
	return xc.XAd_Release(a.Handle)
}

// 数据适配器_取引用计数.
func (a *adapter) GetRefCount() int32 {
	return xc.XAd_GetRefCount(a.Handle)
}

// 数据适配器_销毁.
func (a *adapter) Destroy() *adapter {
	xc.XAd_Destroy(a.Handle)
	return a
}

// 数据适配器_启用自动销毁.
//
// bEnable: 是否启用, 不填默认为 true.
func (a *adapter) EnableAutoDestroy(bEnable ...bool) *adapter {
	enable := true
	if len(bEnable) > 0 {
		enable = bEnable[0]
	}
	xc.XAd_EnableAutoDestroy(a.Handle, enable)
	return a
}
