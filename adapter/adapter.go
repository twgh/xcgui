package adapter

import (
	"github.com/twgh/xcgui/xc"
)

type adapter struct {
	HAdapter int
}

// 给本类的HAdapter赋值
func (a *adapter) SetHAdapter(hAdapter int) {
	a.HAdapter = hAdapter
}

// 数据适配器_增加引用计数
func (a *adapter) AddRef() int {
	return xc.XAd_AddRef(a.HAdapter)
}

// 数据适配器_释放引用计数
func (a *adapter) Release() int {
	return xc.XAd_Release(a.HAdapter)
}

// 数据适配器_取引用计数
func (a *adapter) GetRefCount() int {
	return xc.XAd_GetRefCount(a.HAdapter)
}

// 数据适配器_销毁
func (a *adapter) Destroy() int {
	return xc.XAd_Destroy(a.HAdapter)
}

// 数据适配器_启用自动销毁
// bEnable: 是否启用
func (a *adapter) EnableAutoDestroy(bEnable bool) int {
	return xc.XAd_EnableAutoDestroy(a.HAdapter, bEnable)
}
