package xc

import "github.com/twgh/xcgui/common"

// 数据适配器_增加引用计数.
//
// hAdapter: 数据适配器句柄.
func XAd_AddRef(hAdapter int) int {
	r, _, _ := xAd_AddRef.Call(uintptr(hAdapter))
	return int(r)
}

// 数据适配器_释放引用计数.
//
// hAdapter: 数据适配器句柄.
func XAd_Release(hAdapter int) int {
	r, _, _ := xAd_Release.Call(uintptr(hAdapter))
	return int(r)
}

// 数据适配器_取引用计数.
//
// hAdapter: 数据适配器句柄.
func XAd_GetRefCount(hAdapter int) int {
	r, _, _ := xAd_GetRefCount.Call(uintptr(hAdapter))
	return int(r)
}

// 数据适配器_销毁.
//
// hAdapter: 数据适配器句柄.
func XAd_Destroy(hAdapter int) int {
	r, _, _ := xAd_Destroy.Call(uintptr(hAdapter))
	return int(r)
}

// 数据适配器_启用自动销毁.
//
// hAdapter: 数据适配器句柄.
//
// bEnable: 是否启用.
func XAd_EnableAutoDestroy(hAdapter int, bEnable bool) int {
	r, _, _ := xAd_EnableAutoDestroy.Call(uintptr(hAdapter), common.BoolPtr(bEnable))
	return int(r)
}
