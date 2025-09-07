package edge

import (
	"syscall"
	"unsafe"
)

// ICoreWebView2Profile2 是 ICoreWebView2Profile 的延续, 提供了浏览数据清除功能。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2profile2
type ICoreWebView2Profile2 struct {
	ICoreWebView2Profile
}

// ClearBrowsingData 清除指定类型的浏览数据。
//
// dataKinds: 可以多个执行“或”运算.
func (i *ICoreWebView2Profile2) ClearBrowsingData(dataKinds COREWEBVIEW2_BROWSING_DATA_KINDS, handler *ICoreWebView2ClearBrowsingDataCompletedHandler) error {
	r, _, _ := i.Vtbl.ClearBrowsingData.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(dataKinds),
		uintptr(unsafe.Pointer(handler)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// ClearBrowsingDataInTimeRange 清除指定时间范围内指定类型的浏览数据。
//
// startTime, endTime: 时间范围, 自 UNIX 纪元以来的秒数, 数据将在 [startTime, endTime) 区间内被清除。
func (i *ICoreWebView2Profile2) ClearBrowsingDataInTimeRange(dataKinds COREWEBVIEW2_BROWSING_DATA_KINDS, startTime, endTime int64, handler *ICoreWebView2ClearBrowsingDataCompletedHandler) error {
	start := float64(startTime)
	end := float64(endTime)
	r, _, _ := i.Vtbl.ClearBrowsingDataInTimeRange.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(dataKinds),
		uintptr(unsafe.Pointer(&start)),
		uintptr(unsafe.Pointer(&end)),
		uintptr(unsafe.Pointer(handler)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// ClearBrowsingDataAll 清除所有类型的浏览数据。
func (i *ICoreWebView2Profile2) ClearBrowsingDataAll(handler *ICoreWebView2ClearBrowsingDataCompletedHandler) error {
	r, _, _ := i.Vtbl.ClearBrowsingDataAll.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(handler)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// ClearBrowsingDataEx 在指定的 WebView 事件实现上下文中清除指定类型的浏览数据。
//
// impl: *WebViewEventImpl。
//
// dataKinds: 可以多个执行"或"运算。
//
// cb: 清除完成后的回调处理程序，可以为 nil。
func (i *ICoreWebView2Profile2) ClearBrowsingDataEx(impl *WebViewEventImpl, dataKinds COREWEBVIEW2_BROWSING_DATA_KINDS, cb func(errorCode syscall.Errno) uintptr) error {
	handler := WvEventHandler.GetHandler(impl, "ClearBrowsingDataCompleted")
	if handler == nil {
		var c interface{}
		if cb == nil {
			c = nil
		} else {
			c = cb
		}
		_, _ = WvEventHandler.AddCallBack(impl, "ClearBrowsingDataCompleted", c, nil)
		handler = WvEventHandler.GetHandler(impl, "ClearBrowsingDataCompleted")
	}
	return i.ClearBrowsingData(dataKinds, (*ICoreWebView2ClearBrowsingDataCompletedHandler)(handler))
}

// ClearBrowsingDataInTimeRangeEx 在指定的 WebView 事件实现上下文中清除指定时间范围内指定类型的浏览数据。
//
// impl: *WebViewEventImpl。
//
// dataKinds: 可以多个执行"或"运算。
//
// startTime, endTime: 时间范围, 自 UNIX 纪元以来的秒数, 数据将在 [startTime, endTime) 区间内被清除。
//
// cb: 清除完成后的回调处理程序，可以为 nil。
func (i *ICoreWebView2Profile2) ClearBrowsingDataInTimeRangeEx(impl *WebViewEventImpl, dataKinds COREWEBVIEW2_BROWSING_DATA_KINDS, startTime, endTime int64, cb func(errorCode syscall.Errno) uintptr) error {
	handler := WvEventHandler.GetHandler(impl, "ClearBrowsingDataCompleted")
	if handler == nil {
		var c interface{}
		if cb == nil {
			c = nil
		} else {
			c = cb
		}
		_, _ = WvEventHandler.AddCallBack(impl, "ClearBrowsingDataCompleted", c, nil)
		handler = WvEventHandler.GetHandler(impl, "ClearBrowsingDataCompleted")
	}
	return i.ClearBrowsingDataInTimeRange(dataKinds, startTime, endTime, (*ICoreWebView2ClearBrowsingDataCompletedHandler)(handler))
}

// ClearBrowsingDataAllEx 在指定的 WebView 事件实现上下文中清除所有类型的浏览数据。
//
// impl: *WebViewEventImpl。
//
// cb: 清除完成后的回调处理程序，可以为 nil。
func (i *ICoreWebView2Profile2) ClearBrowsingDataAllEx(impl *WebViewEventImpl, cb func(errorCode syscall.Errno) uintptr) error {
	handler := WvEventHandler.GetHandler(impl, "ClearBrowsingDataCompleted")
	if handler == nil {
		var c interface{}
		if cb == nil {
			c = nil
		} else {
			c = cb
		}
		_, _ = WvEventHandler.AddCallBack(impl, "ClearBrowsingDataCompleted", c, nil)
		handler = WvEventHandler.GetHandler(impl, "ClearBrowsingDataCompleted")
	}
	return i.ClearBrowsingDataAll((*ICoreWebView2ClearBrowsingDataCompletedHandler)(handler))
}
