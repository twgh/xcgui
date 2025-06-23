package edge

import (
	"github.com/twgh/xcgui/common"
	"github.com/twgh/xcgui/wapi"
	"syscall"
	"unsafe"
)

// ICoreWebView2ContextMenuTarget 上下文菜单目标的信息。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2contextmenutarget
type ICoreWebView2ContextMenuTarget struct {
	Vtbl *ICoreWebView2ContextMenuTargetVtbl
}

type ICoreWebView2ContextMenuTargetVtbl struct {
	IUnknownVtbl
	GetKind                    ComProc
	GetIsEditable              ComProc
	GetIsRequestedForMainFrame ComProc
	GetPageUri                 ComProc
	GetFrameUri                ComProc
	GetHasLinkUri              ComProc
	GetLinkUri                 ComProc
	GetHasLinkText             ComProc
	GetLinkText                ComProc
	GetHasSourceUri            ComProc
	GetSourceUri               ComProc
	GetHasSelection            ComProc
	GetSelectionText           ComProc
}

func (i *ICoreWebView2ContextMenuTarget) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2ContextMenuTarget) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2ContextMenuTarget) QueryInterface(refiid, object unsafe.Pointer) error {
	r, _, _ := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), uintptr(refiid), uintptr(object))
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetKind 获取用户选择的上下文菜单目标的类型.
func (i *ICoreWebView2ContextMenuTarget) GetKind() (COREWEBVIEW2_CONTEXT_MENU_TARGET_KIND, error) {
	var kind COREWEBVIEW2_CONTEXT_MENU_TARGET_KIND
	r, _, _ := i.Vtbl.GetKind.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&kind)),
	)
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return kind, nil
}

// GetIsEditable 如果在可编辑组件上请求上下文菜单，则返回 TRUE。
func (i *ICoreWebView2ContextMenuTarget) GetIsEditable() (bool, error) {
	var editable bool
	r, _, _ := i.Vtbl.GetIsEditable.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&editable)),
	)
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return editable, nil
}

// GetIsRequestedForMainFrame 如果在主框架上请求上下文菜单，则返回 TRUE；如果在其他框架上调用，则返回 FALSE。
func (i *ICoreWebView2ContextMenuTarget) GetIsRequestedForMainFrame() (bool, error) {
	var isMainFrame bool
	r, _, _ := i.Vtbl.GetIsRequestedForMainFrame.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&isMainFrame)),
	)
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return isMainFrame, nil
}

// GetPageUri 获取页面的 URI。
func (i *ICoreWebView2ContextMenuTarget) GetPageUri() (string, error) {
	var uri *uint16
	r, _, _ := i.Vtbl.GetPageUri.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&uri)),
	)
	if r != 0 {
		return "", syscall.Errno(r)
	}
	result := common.UTF16PtrToString(uri)
	wapi.CoTaskMemFree(unsafe.Pointer(uri))
	return result, nil
}

// GetFrameUri 获取框架的 URI。
//   - 如果 IsRequestedForMainFrame 为 TRUE，则与 PageUri 匹配。
func (i *ICoreWebView2ContextMenuTarget) GetFrameUri() (string, error) {
	var uri *uint16
	r, _, _ := i.Vtbl.GetFrameUri.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&uri)),
	)
	if r != 0 {
		return "", syscall.Errno(r)
	}
	result := common.UTF16PtrToString(uri)
	wapi.CoTaskMemFree(unsafe.Pointer(uri))
	return result, nil
}

// GetHasLinkUri 如果在包含锚点标签的 HTML 上请求上下文菜单，则返回 TRUE。
func (i *ICoreWebView2ContextMenuTarget) GetHasLinkUri() (bool, error) {
	var hasLink bool
	r, _, _ := i.Vtbl.GetHasLinkUri.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&hasLink)),
	)
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return hasLink, nil
}

// GetLinkUri 获取链接URI（如果 HasLinkUri 为 TRUE，则返回 URI；否则返回空）。
func (i *ICoreWebView2ContextMenuTarget) GetLinkUri() (string, error) {
	var uri *uint16
	r, _, _ := i.Vtbl.GetLinkUri.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&uri)),
	)
	if r != 0 {
		return "", syscall.Errno(r)
	}
	result := common.UTF16PtrToString(uri)
	wapi.CoTaskMemFree(unsafe.Pointer(uri))
	return result, nil
}

// GetHasLinkText 如果在包含锚标记的文本元素上请求上下文菜单，则返回 TRUE。
func (i *ICoreWebView2ContextMenuTarget) GetHasLinkText() (bool, error) {
	var hasText bool
	r, _, _ := i.Vtbl.GetHasLinkText.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&hasText)),
	)
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return hasText, nil
}

// GetLinkText 获取链接文本（如果 HasLinkText 为 TRUE，则获取链接文本；否则返回空）。
func (i *ICoreWebView2ContextMenuTarget) GetLinkText() (string, error) {
	var text *uint16
	r, _, _ := i.Vtbl.GetLinkText.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&text)),
	)
	if r != 0 {
		return "", syscall.Errno(r)
	}
	result := common.UTF16PtrToString(text)
	wapi.CoTaskMemFree(unsafe.Pointer(text))
	return result, nil
}

// GetHasSourceUri 如果在包含源 URI 的 HTML 上请求上下文菜单，则返回 TRUE。
func (i *ICoreWebView2ContextMenuTarget) GetHasSourceUri() (bool, error) {
	var hasSource bool
	r, _, _ := i.Vtbl.GetHasSourceUri.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&hasSource)),
	)
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return hasSource, nil
}

// GetSourceUri 获取元素的活动源 URI（如果 HasSourceUri 为 TRUE，则获取；否则返回空）。
func (i *ICoreWebView2ContextMenuTarget) GetSourceUri() (string, error) {
	var uri *uint16
	r, _, _ := i.Vtbl.GetSourceUri.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&uri)),
	)
	if r != 0 {
		return "", syscall.Errno(r)
	}
	result := common.UTF16PtrToString(uri)
	wapi.CoTaskMemFree(unsafe.Pointer(uri))
	return result, nil
}

// GetHasSelection 判断是否有选中文本.
func (i *ICoreWebView2ContextMenuTarget) GetHasSelection() (bool, error) {
	var hasSelection bool
	r, _, _ := i.Vtbl.GetHasSelection.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&hasSelection)),
	)
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return hasSelection, nil
}

// GetSelectionText 获取所选文本（如果 HasSelection 为 TRUE，则返回所选文本；否则返回空）。
func (i *ICoreWebView2ContextMenuTarget) GetSelectionText() (string, error) {
	var text *uint16
	r, _, _ := i.Vtbl.GetSelectionText.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&text)),
	)
	if r != 0 {
		return "", syscall.Errno(r)
	}
	result := common.UTF16PtrToString(text)
	wapi.CoTaskMemFree(unsafe.Pointer(text))
	return result, nil
}

// MustGetKind 获取用户选择的上下文菜单目标的类型。出错时会触发全局错误回调。
func (i *ICoreWebView2ContextMenuTarget) MustGetKind() COREWEBVIEW2_CONTEXT_MENU_TARGET_KIND {
	kind, err := i.GetKind()
	ReportErrorAtuo(err)
	return kind
}

// MustGetIsEditable 如果在可编辑组件上请求上下文菜单，则返回 TRUE。出错时会触发全局错误回调。
func (i *ICoreWebView2ContextMenuTarget) MustGetIsEditable() bool {
	editable, err := i.GetIsEditable()
	ReportErrorAtuo(err)
	return editable
}

// MustGetIsRequestedForMainFrame 如果在主框架上请求上下文菜单，则返回 TRUE；如果在其他框架上调用，则返回 FALSE。出错时会触发全局错误回调。
func (i *ICoreWebView2ContextMenuTarget) MustGetIsRequestedForMainFrame() bool {
	isMainFrame, err := i.GetIsRequestedForMainFrame()
	ReportErrorAtuo(err)
	return isMainFrame
}

// MustGetPageUri 获取页面的 URI。出错时会触发全局错误回调。
func (i *ICoreWebView2ContextMenuTarget) MustGetPageUri() string {
	uri, err := i.GetPageUri()
	ReportErrorAtuo(err)
	return uri
}

// MustGetFrameUri 获取框架的 URI。出错时会触发全局错误回调。
//   - 如果 IsRequestedForMainFrame 为 TRUE，则与 PageUri 匹配。
func (i *ICoreWebView2ContextMenuTarget) MustGetFrameUri() string {
	uri, err := i.GetFrameUri()
	ReportErrorAtuo(err)
	return uri
}

// MustGetHasLinkUri 如果在包含锚点标签的 HTML 上请求上下文菜单，则返回 TRUE。出错时会触发全局错误回调。
func (i *ICoreWebView2ContextMenuTarget) MustGetHasLinkUri() bool {
	hasLink, err := i.GetHasLinkUri()
	ReportErrorAtuo(err)
	return hasLink
}

// MustGetLinkUri 获取链接URI（如果 HasLinkUri 为 TRUE，则返回 URI；否则返回空）。出错时会触发全局错误回调。
func (i *ICoreWebView2ContextMenuTarget) MustGetLinkUri() string {
	uri, err := i.GetLinkUri()
	ReportErrorAtuo(err)
	return uri
}

// MustGetHasLinkText 如果在包含锚标记的文本元素上请求上下文菜单，则返回 TRUE。出错时会触发全局错误回调。
func (i *ICoreWebView2ContextMenuTarget) MustGetHasLinkText() bool {
	hasText, err := i.GetHasLinkText()
	ReportErrorAtuo(err)
	return hasText
}

// MustGetLinkText 获取链接文本（如果 HasLinkText 为 TRUE，则获取链接文本；否则返回空）。出错时会触发全局错误回调。
func (i *ICoreWebView2ContextMenuTarget) MustGetLinkText() string {
	text, err := i.GetLinkText()
	ReportErrorAtuo(err)
	return text
}

// MustGetHasSourceUri 如果在包含源 URI 的 HTML 上请求上下文菜单，则返回 TRUE。出错时会触发全局错误回调。
func (i *ICoreWebView2ContextMenuTarget) MustGetHasSourceUri() bool {
	hasSource, err := i.GetHasSourceUri()
	ReportErrorAtuo(err)
	return hasSource
}

// MustGetSourceUri 获取元素的活动源 URI（如果 HasSourceUri 为 TRUE，则获取；否则返回空）。出错时会触发全局错误回调。
func (i *ICoreWebView2ContextMenuTarget) MustGetSourceUri() string {
	uri, err := i.GetSourceUri()
	ReportErrorAtuo(err)
	return uri
}

// MustGetHasSelection 判断是否有选中文本。出错时会触发全局错误回调。
func (i *ICoreWebView2ContextMenuTarget) MustGetHasSelection() bool {
	hasSelection, err := i.GetHasSelection()
	ReportErrorAtuo(err)
	return hasSelection
}

// MustGetSelectionText 获取所选文本（如果 HasSelection 为 TRUE，则返回所选文本；否则返回空）。出错时会触发全局错误回调。
func (i *ICoreWebView2ContextMenuTarget) MustGetSelectionText() string {
	text, err := i.GetSelectionText()
	ReportErrorAtuo(err)
	return text
}
