package adapter

import (
	xc "github.com/twgh/xcgui"
)

type Adapter struct {
	HAdapter_ int
}

// 数据适配器_增加引用计数
func (a *Adapter) AddRef() int {
	return xc.XAd_AddRef(a.HAdapter_)
}

// 数据适配器_释放引用计数
func (a *Adapter) Release() int {
	return xc.XAd_Release(a.HAdapter_)
}

// 数据适配器_取引用计数
func (a *Adapter) GetRefCount() int {
	return xc.XAd_GetRefCount(a.HAdapter_)
}

// 数据适配器_销毁
func (a *Adapter) Destroy() int {
	return xc.XAd_Destroy(a.HAdapter_)
}

// 数据适配器_启用自动销毁
// bEnable: 是否启用
func (a *Adapter) EnableAutoDestroy(bEnable bool) int {
	return xc.XAd_EnableAutoDestroy(a.HAdapter_, bEnable)
}
