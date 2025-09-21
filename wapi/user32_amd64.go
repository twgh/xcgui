//go:build amd64

package wapi

// SetWindowLongPtrW 更改指定窗口的属性。该函数还会在额外的窗口内存中设置指定偏移量的值。返回设置前的属性值。如果函数失败，返回值为 0。在 386 架构中, 内部使用的是 SetWindowLongW.
//
// 详情: https://learn.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-SetWindowLongPtrW.
//
// hWnd: 窗口的句柄，间接地是窗口所属的类。如果拥有由 hWnd 参数指定的进程位于 UIPI 层次结构中的进程高于调用线程所在的进程，则 SetWindowLongPtrW 函数将失败。
//
// nIndex: 要设置的值的从零开始的偏移量。有效值介于零到额外窗口内存的字节数之间，减去 LONG_PTR 的大小。若要设置任何其他值，请指定以下值之一: wapi.GWL_ , wapi.GWLP_ . 当 hWnd 参数标识对话框时，也可以使用以下值: wapi.DWLP_ .
//
// dwNewLong: 要设置的新值。
func SetWindowLongPtrW(hWnd uintptr, nIndex int32, dwNewLong int) int {
	ret, _, _ := procSetWindowLongPtrW.Call(
		hWnd,               // 窗口句柄
		uintptr(nIndex),    // 属性索引
		uintptr(dwNewLong), // 新值
	)
	return int(ret) // 返回设置前的值
}

// GetWindowLongPtrW 检索有关指定窗口的信息。该函数还会在额外的窗口内存中检索指定偏移量处的值。返回指定属性的当前值。如果函数失败，返回值为 0。在 386 架构中, 内部使用的是 GetWindowLongW.
//
// 详情: https://learn.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-GetWindowLongPtrW.
//
// hWnd: 窗口的句柄，间接地是窗口所属的类。
//
// nIndex: 要检索的值的从零开始的偏移量。有效值介于零到额外窗口内存的字节数之间，减去 LONG_PTR 的大小。若要设置任何其他值，请指定以下值之一: wapi.GWL_ , wapi.GWLP_ . 当 hWnd 参数标识对话框时，也可以使用以下值: wapi.DWLP_ .
func GetWindowLongPtrW(hWnd uintptr, nIndex int32) int {
	ret, _, _ := procGetWindowLongPtrW.Call(
		hWnd,            // 窗口句柄
		uintptr(nIndex), // 属性索引
	)
	return int(ret) // 返回属性值
}
