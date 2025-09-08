package edge

import (
	"syscall"
	"unsafe"
)

// ICoreWebView2Find 提供用于在网页视图中查找和浏览文本的方法及属性的接口。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2find
type ICoreWebView2Find struct {
	Vtbl *ICoreWebView2FindVtbl
}

type ICoreWebView2FindVtbl struct {
	IUnknownVtbl
	GetActiveMatchIndex                       ComProc
	GetMatchCount                             ComProc
	AddActiveMatchIndexChangedEventHandler    ComProc
	RemoveActiveMatchIndexChangedEventHandler ComProc
	AddMatchCountChangedEventHandler          ComProc
	RemoveMatchCountChangedEventHandler       ComProc
	Start                                     ComProc
	FindNext                                  ComProc
	FindPrevious                              ComProc
	Stop                                      ComProc
}

func (i *ICoreWebView2Find) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2Find) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2Find) QueryInterface(refiid, object unsafe.Pointer) error {
	r, _, _ := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), uintptr(refiid), uintptr(object))
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetActiveMatchIndex 获取当前活动匹配项的索引。
//   - 如果没有活动匹配项，则返回-1。第一个匹配项的索引从1开始。
func (i *ICoreWebView2Find) GetActiveMatchIndex() (int32, error) {
	var value int32
	r, _, _ := i.Vtbl.GetActiveMatchIndex.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return value, nil
}

// GetMatchCount 获取匹配项的总数。
func (i *ICoreWebView2Find) GetMatchCount() (int32, error) {
	var value int32
	r, _, _ := i.Vtbl.GetMatchCount.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return value, nil
}

// Event_FindActiveMatchIndexChanged 查找活动匹配项索引更改事件。
//   - 当当前活动的查找匹配项索引发生变化时触发。
//   - 这种情况可能发生在用户导航到不同的匹配项时，或者在通过编程方式更改活动匹配项时。
func (i *ICoreWebView2Find) Event_FindActiveMatchIndexChanged(impl *WebViewEventImpl, cb func(sender *ICoreWebView2Find) uintptr, allowAddingMultiple ...bool) (int, error) {
	return WvEventHandler.AddCallBack(impl, "FindActiveMatchIndexChanged", cb, i, allowAddingMultiple...)
}

// Event_FindMatchCountChanged 查找匹配项数量更改事件。
//   - 当文档中的匹配项总数因新的查找会话或文档更改而发生变化时，会引发此事件。
func (i *ICoreWebView2Find) Event_FindMatchCountChanged(impl *WebViewEventImpl, cb func(sender *ICoreWebView2Find) uintptr, allowAddingMultiple ...bool) (int, error) {
	return WvEventHandler.AddCallBack(impl, "FindMatchCountChanged", cb, i, allowAddingMultiple...)
}

// AddActiveMatchIndexChangedEventHandler 添加活动匹配项索引更改事件处理程序。
//   - 这种情况可能发生在用户导航到不同的匹配项时，或者在通过编程方式更改活动匹配项时。
func (i *ICoreWebView2Find) AddActiveMatchIndexChangedEventHandler(eventHandler *ICoreWebView2FindActiveMatchIndexChangedEventHandler, token *EventRegistrationToken) error {
	r, _, _ := i.Vtbl.AddActiveMatchIndexChangedEventHandler.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(token)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// RemoveActiveMatchIndexChangedEventHandler 移除活动匹配项索引更改事件处理程序。
func (i *ICoreWebView2Find) RemoveActiveMatchIndexChangedEventHandler(token EventRegistrationToken) error {
	r, _, _ := i.Vtbl.RemoveActiveMatchIndexChangedEventHandler.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(token.Value),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// AddMatchCountChangedEventHandler 添加匹配项数量更改事件处理程序。
//   - 当文档中的匹配项总数因新的查找会话或文档更改而发生变化时，会引发此事件。
func (i *ICoreWebView2Find) AddMatchCountChangedEventHandler(eventHandler *ICoreWebView2FindMatchCountChangedEventHandler, token *EventRegistrationToken) error {
	r, _, _ := i.Vtbl.AddMatchCountChangedEventHandler.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(token)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// RemoveMatchCountChangedEventHandler 移除匹配项数量更改事件处理程序。
func (i *ICoreWebView2Find) RemoveMatchCountChangedEventHandler(token EventRegistrationToken) error {
	r, _, _ := i.Vtbl.RemoveMatchCountChangedEventHandler.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(token.Value),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// Start 使用指定的查找选项异步启动查找。
//   - 显示查找栏并启动查找会话。如果已在进行查找会话，它将被停止并替换为这个新实例。
//   - 如果调用时传入空字符串，查找栏会显示，但不会执行查找操作。
//   - 启动后更改 FindOptions 对象不会影响正在进行的查找会话。要更改正在进行的查找会话，必须使用新的或修改后的 FindOptions 对象再次调用 Start。
//   - Start 支持 HTML 和 TXT 文档查询。一般来说，此 API 专为基于文本的查找会话设计。
//   - 如果以编程方式在没有文本字段的其他文件格式上启动查找会话，查找会话会尝试执行，但无法找到任何匹配项（会静默失败）。
//   - 注意：当 UI 显示且查找栏中包含查找词，并且匹配项已填充到查找栏的计数器中时，异步操作完成。UI 显示和计数器中匹配项的填充之间可能存在轻微延迟。
//   - MatchCountChanged 和 ActiveMatchIndexChanged 事件仅在 Start 完成后触发；否则，它们将保持默认值（活动匹配索引为-1，匹配计数为0）。
//   - 要启动新的查找会话（从第一个匹配项开始搜索），请在调用 Start 之前调用 Stop。
//   - 如果在不调用 Stop 的情况下，连续使用相同的选项调用 Start，查找会话将从现有会话的当前位置继续。在不更改参数的情况下调用 Start，其行为将类似于 FindNext 或 FindPrevious，具体取决于最近执行的搜索操作。如果两者都未被调用，Start 默认向前查找。但是，在正在进行的查找会话期间再次调用 Start 不会从当前活动匹配项的位置继续。例如，对于文本“1 1 A 1 1”，先启动查找“A”的会话，然后启动另一个查找“1”的会话，它会从文档开头开始搜索，无论之前的活动匹配项在哪里。此行为表明，更改查找查询会启动一个全新的查找会话，而不是从之前的匹配索引继续。
func (i *ICoreWebView2Find) Start(options *ICoreWebView2FindOptions, handler *ICoreWebView2FindStartCompletedHandler) error {
	r, _, _ := i.Vtbl.Start.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(options)),
		uintptr(unsafe.Pointer(handler)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// StartEx 使用指定的查找选项异步启动查找。
//   - 显示查找栏并启动查找会话。如果已在进行查找会话，它将被停止并替换为这个新实例。
//   - 如果调用时传入空字符串，查找栏会显示，但不会执行查找操作。
//   - 启动后更改 FindOptions 对象不会影响正在进行的查找会话。要更改正在进行的查找会话，必须使用新的或修改后的 FindOptions 对象再次调用 Start。
//   - Start 支持 HTML 和 TXT 文档查询。一般来说，此 API 专为基于文本的查找会话设计。
//   - 如果以编程方式在没有文本字段的其他文件格式上启动查找会话，查找会话会尝试执行，但无法找到任何匹配项（会静默失败）。
//   - 注意：当 UI 显示且查找栏中包含查找词，并且匹配项已填充到查找栏的计数器中时，异步操作完成。UI 显示和计数器中匹配项的填充之间可能存在轻微延迟。
//   - MatchCountChanged 和 ActiveMatchIndexChanged 事件仅在 Start 完成后触发；否则，它们将保持默认值（活动匹配索引为-1，匹配计数为0）。
//   - 要启动新的查找会话（从第一个匹配项开始搜索），请在调用 Start 之前调用 Stop。
//   - 如果在不调用 Stop 的情况下，连续使用相同的选项调用 Start，查找会话将从现有会话的当前位置继续。在不更改参数的情况下调用 Start，其行为将类似于 FindNext 或 FindPrevious，具体取决于最近执行的搜索操作。如果两者都未被调用，Start 默认向前查找。但是，在正在进行的查找会话期间再次调用 Start 不会从当前活动匹配项的位置继续。例如，对于文本“1 1 A 1 1”，先启动查找“A”的会话，然后启动另一个查找“1”的会话，它会从文档开头开始搜索，无论之前的活动匹配项在哪里。此行为表明，更改查找查询会启动一个全新的查找会话，而不是从之前的匹配索引继续。
func (i *ICoreWebView2Find) StartEx(impl *WebViewEventImpl, options *ICoreWebView2FindOptions, cb func(errorCode syscall.Errno) uintptr) error {
	handler := WvEventHandler.GetHandler(impl, "FindStartCompleted")
	if handler == nil {
		var c interface{}
		if cb == nil {
			c = nil
		} else {
			c = cb
		}
		_, _ = WvEventHandler.AddCallBack(impl, "FindStartCompleted", c, nil)
		handler = WvEventHandler.GetHandler(impl, "FindStartCompleted")
	}
	return i.Start(options, (*ICoreWebView2FindStartCompletedHandler)(handler))
}

// FindNext 导航到文档中的下一个匹配项。
//   - 如果没有找到匹配项，FindNext 将回绕到第一个匹配项的索引。
//   - 如果在没有活动的查找会话时调用，FindNext 将无提示地失败。
func (i *ICoreWebView2Find) FindNext() error {
	r, _, _ := i.Vtbl.FindNext.Call(
		uintptr(unsafe.Pointer(i)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// FindPrevious 导航到文档中的上一个匹配项。
//   - 如果没有找到匹配项，FindPrevious 将回绕到最后一个匹配项的索引。
//   - 如果在没有活跃的查找会话时调用，FindPrevious 将无提示地失败。
func (i *ICoreWebView2Find) FindPrevious() error {
	r, _, _ := i.Vtbl.FindPrevious.Call(
		uintptr(unsafe.Pointer(i)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// Stop 停止当前的“查找”会话并隐藏查找栏。
//   - 如果在没有活动的查找会话时调用，它将静默地不执行任何操作。
func (i *ICoreWebView2Find) Stop() error {
	r, _, _ := i.Vtbl.Stop.Call(
		uintptr(unsafe.Pointer(i)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// MustGetActiveMatchIndex 获取当前活动匹配项的索引。出错时会触发全局错误回调。
//   - 如果没有活动匹配项，则返回-1。第一个匹配项的索引从1开始。
func (i *ICoreWebView2Find) MustGetActiveMatchIndex() int32 {
	value, err := i.GetActiveMatchIndex()
	ReportErrorAuto(err)
	return value
}

// MustGetMatchCount 获取匹配项的总数。出错时会触发全局错误回调。
func (i *ICoreWebView2Find) MustGetMatchCount() int32 {
	value, err := i.GetMatchCount()
	ReportErrorAuto(err)
	return value
}
