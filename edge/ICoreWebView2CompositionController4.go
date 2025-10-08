package edge

import (
	"syscall"
	"unsafe"

	"github.com/twgh/xcgui/xc"
)

// ICoreWebView2CompositionController4 是 ICoreWebView2CompositionController3 的延续，用于支持非客户端命中测试功能。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2compositioncontroller4
type ICoreWebView2CompositionController4 struct {
	ICoreWebView2CompositionController3
}

// QueryNonClientRegion 查询指定类型的非客户端区域矩形集合。
//   - 这将用于 NonClientRegionChanged 事件的回调中，该回调的事件参数对象包含一个类型为 COREWEBVIEW2_NON_CLIENT_REGION_KIND 的 region 属性。
//
// kind: 要查询的非客户端区域类型。
func (i *ICoreWebView2CompositionController4) QueryNonClientRegion(kind COREWEBVIEW2_NON_CLIENT_REGION_KIND) (*ICoreWebView2RegionRectCollectionView, error) {
	var rects *ICoreWebView2RegionRectCollectionView
	r, _, _ := i.Vtbl.QueryNonClientRegion.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(kind),
		uintptr(unsafe.Pointer(&rects)),
	)
	if r != 0 {
		return nil, syscall.Errno(r)
	}
	return rects, nil
}

// Event_NonClientRegionChanged 非客户端区域变更事件.
//   - 当应用 HTML 中标记为非客户端的区域发生变化时触发.
//   - 这种变化包括新区域被标记、区域被取消标记，或者区域类型被改为其他类型。
func (i *ICoreWebView2CompositionController4) Event_NonClientRegionChanged(impl *WebViewEventImpl, cb func(sender *ICoreWebView2CompositionController, args *ICoreWebView2NonClientRegionChangedEventArgs) uintptr, allowAddingMultiple ...bool) (int, error) {
	return WvEventHandler.AddCallBack(impl, "NonClientRegionChanged", cb, i, allowAddingMultiple...)
}

// AddNonClientRegionChanged 添加非客户端区域变更事件处理程序。
//   - 当应用 HTML 中标记为非客户端的区域发生变化时触发.
//   - 这种变化包括新区域被标记、区域被取消标记，或者区域类型被改为其他类型。
func (i *ICoreWebView2CompositionController4) AddNonClientRegionChanged(eventHandler *ICoreWebView2NonClientRegionChangedEventHandler, token *EventRegistrationToken) error {
	r, _, _ := i.Vtbl.AddNonClientRegionChanged.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(token)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// RemoveNonClientRegionChanged 移除非客户端区域变更事件处理程序。
func (i *ICoreWebView2CompositionController4) RemoveNonClientRegionChanged(token EventRegistrationToken) error {
	r, _, _ := i.Vtbl.RemoveNonClientRegionChanged.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(token.Value),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// MustGetNonClientRegionAtPoint 获取指定点的非客户端区域类型。出错时会触发全局错误回调。
//   - 如果您使用 ICoreWebView2CompositionController 托管 WebView2，可以在 Win32 窗口过程（WndProc）中调用此方法，以确定鼠标是否正在移动到 WebView2 网页内容上或在其上点击，而这些内容应被视为非客户区的一部分。
//
// point: 指定要检查的点坐标。
func (i *ICoreWebView2CompositionController4) MustGetNonClientRegionAtPoint(point xc.POINT) COREWEBVIEW2_NON_CLIENT_REGION_KIND {
	value, err := i.GetNonClientRegionAtPoint(point)
	ReportErrorAuto(err)
	return value
}

// MustQueryNonClientRegion 查询指定类型的非客户端区域矩形集合。出错时会触发全局错误回调。
//
// kind: 指定要查询的区域类型。
func (i *ICoreWebView2CompositionController4) MustQueryNonClientRegion(kind COREWEBVIEW2_NON_CLIENT_REGION_KIND) *ICoreWebView2RegionRectCollectionView {
	value, err := i.QueryNonClientRegion(kind)
	ReportErrorAuto(err)
	return value
}
