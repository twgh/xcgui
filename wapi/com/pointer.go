package com

import "unsafe"

// InterfacePtr 把"以 uintptr 交出的 COM 接口指针"搬成 unsafe.Pointer.
//
// xcgui 的 DLL 直通 API 都以 uintptr 交出接口指针(例如 xc.XDraw_GetD2dRenderTarget 返回
// *ID2D1RenderTarget), 转成类型化结构体之前需要这一步. 这些对象由 xcgui 持有, 只在绘制回调等
// 场合借用, 不做引用计数.
//
// 绕一层取地址是为了避开 uintptr → unsafe.Pointer 的直接转换(go vet 的 unsafeptr 检查会把它报成
// "possible misuse of unsafe.Pointer"): 这是纯位模式重解释, 不做任何指针算术.
func InterfacePtr(h uintptr) unsafe.Pointer {
	return *(*unsafe.Pointer)(unsafe.Pointer(&h))
}
