package edge

import (
	"github.com/twgh/xcgui/xc"
	"syscall"
	"unsafe"
)

// ICoreWebView2_9 是 ICoreWebView2_8 接口的延续，用于默认下载对话框的定位和锚定。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2_9
type ICoreWebView2_9 struct {
	Vtbl *ICoreWebView2_9Vtbl
}

type ICoreWebView2_9Vtbl struct {
	ICoreWebView2_8Vtbl
	AddIsDefaultDownloadDialogOpenChanged    ComProc
	RemoveIsDefaultDownloadDialogOpenChanged ComProc
	GetIsDefaultDownloadDialogOpen           ComProc
	OpenDefaultDownloadDialog                ComProc
	CloseDefaultDownloadDialog               ComProc
	GetDefaultDownloadDialogCornerAlignment  ComProc
	PutDefaultDownloadDialogCornerAlignment  ComProc
	GetDefaultDownloadDialogMargin           ComProc
	PutDefaultDownloadDialogMargin           ComProc
}

func (i *ICoreWebView2_9) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2_9) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2_9) QueryInterface(refiid, object unsafe.Pointer) error {
	r, _, _ := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), uintptr(refiid), uintptr(object))
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetIsDefaultDownloadDialogOpen 获取默认下载对话框是否处于打开状态.
//   - 此属性的值仅在显式打开或关闭默认下载对话框时才会改变。
//   - 隐藏 WebView 会隐式隐藏该对话框，但不会改变此属性的值。
func (i *ICoreWebView2_9) GetIsDefaultDownloadDialogOpen() (bool, error) {
	var isOpen bool
	r, _, _ := i.Vtbl.GetIsDefaultDownloadDialogOpen.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&isOpen)),
	)
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return isOpen, nil
}

// OpenDefaultDownloadDialog 打开默认下载对话框.
//   - 如果在存在最近下载记录之前打开该对话框，对话框将显示当前配置文件的所有过往下载记录。
//   - 否则，对话框仅显示最近的下载记录，并提供一个“查看更多”按钮以查看过往下载记录。
//   - 如果对话框处于关闭状态，调用此方法将引发 IsDefaultDownloadDialogOpenChanged 事件。
//   - 如果对话框已打开，则此方法无效。
func (i *ICoreWebView2_9) OpenDefaultDownloadDialog() error {
	r, _, _ := i.Vtbl.OpenDefaultDownloadDialog.Call(
		uintptr(unsafe.Pointer(i)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// CloseDefaultDownloadDialog 关闭默认下载对话框.
//   - 调用此方法时，如果对话框处于打开状态，将引发 IsDefaultDownloadDialogOpenChanged 事件。
//   - 如果对话框已关闭，则此方法无效。
func (i *ICoreWebView2_9) CloseDefaultDownloadDialog() error {
	r, _, _ := i.Vtbl.CloseDefaultDownloadDialog.Call(
		uintptr(unsafe.Pointer(i)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetDefaultDownloadDialogCornerAlignment 获取默认下载对话框的边角对齐方式。
func (i *ICoreWebView2_9) GetDefaultDownloadDialogCornerAlignment() (COREWEBVIEW2_DEFAULT_DOWNLOAD_DIALOG_CORNER_ALIGNMENT, error) {
	var alignment COREWEBVIEW2_DEFAULT_DOWNLOAD_DIALOG_CORNER_ALIGNMENT
	r, _, _ := i.Vtbl.GetDefaultDownloadDialogCornerAlignment.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&alignment)),
	)
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return alignment, nil
}

// SetDefaultDownloadDialogCornerAlignment 设置默认下载对话框的边角对齐方式。
//   - 该对话框可以对齐到 WebView 的任意一个角（请参阅 COREWEBVIEW2_DEFAULT_DOWNLOAD_DIALOG_CORNER_ALIGNMENT）。
//   - 当 WebView 或对话框大小发生变化时，对话框会保持其相对于该角的位置。
//   - 如果 WebView 足够小，对话框可能会部分或完全位于 WebView 边界之外。
//   - 使用 PutDefaultDownloadDialogMargin 设置相对于该角的边距。
//   - 应在初始化期间设置角对齐方式和边距，以确保在首次计算布局时正确应用它们，否则它们将不会生效，直到下次更新 WebView 的位置或大小。
func (i *ICoreWebView2_9) SetDefaultDownloadDialogCornerAlignment(alignment COREWEBVIEW2_DEFAULT_DOWNLOAD_DIALOG_CORNER_ALIGNMENT) error {
	r, _, _ := i.Vtbl.PutDefaultDownloadDialogCornerAlignment.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(alignment),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetDefaultDownloadDialogMargin 获取默认下载对话框的边距.
func (i *ICoreWebView2_9) GetDefaultDownloadDialogMargin() (xc.POINT, error) {
	var margin xc.POINT
	r, _, _ := i.Vtbl.GetDefaultDownloadDialogMargin.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&margin)),
	)
	if r != 0 {
		return xc.POINT{}, syscall.Errno(r)
	}
	return margin, nil
}

// SetDefaultDownloadDialogMargin 设置默认下载对话框的边距.
//   - 边距是一个点，用于描述所选 WebView 角与距离它最近的默认下载对话框角之间的垂直和水平距离。
//   - 正值会使对话框从所选 WebView 角向 WebView 中心移动，负值则会使对话框远离该角。
//   - 使用(0, 0)可使对话框与 WebView 角对齐且无边距。
//   - 应在初始化期间设置角对齐方式和边距，以确保在首次计算布局时它们能正确应用，否则它们将不会生效，直到下次 WebView 位置或大小更新。
func (i *ICoreWebView2_9) SetDefaultDownloadDialogMargin(margin xc.POINT) error {
	r, _, _ := i.Vtbl.PutDefaultDownloadDialogMargin.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&margin)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

/*TODO:
AddIsDefaultDownloadDialogOpenChanged
RemoveIsDefaultDownloadDialogOpenChanged
*/

// MustGetIsDefaultDownloadDialogOpen 获取默认下载对话框是否处于打开状态。出错时会触发全局错误回调。
//   - 此属性的值仅在显式打开或关闭默认下载对话框时才会改变。
//   - 隐藏 WebView 会隐式隐藏该对话框，但不会改变此属性的值。
func (i *ICoreWebView2_9) MustGetIsDefaultDownloadDialogOpen() bool {
	result, err := i.GetIsDefaultDownloadDialogOpen()
	ReportErrorAtuo(err)
	return result
}

// MustGetDefaultDownloadDialogCornerAlignment 获取默认下载对话框的边角对齐方式。出错时会触发全局错误回调。
func (i *ICoreWebView2_9) MustGetDefaultDownloadDialogCornerAlignment() COREWEBVIEW2_DEFAULT_DOWNLOAD_DIALOG_CORNER_ALIGNMENT {
	alignment, err := i.GetDefaultDownloadDialogCornerAlignment()
	ReportErrorAtuo(err)
	return alignment
}

// MustGetDefaultDownloadDialogMargin 获取默认下载对话框的边距。出错时会触发全局错误回调。
func (i *ICoreWebView2_9) MustGetDefaultDownloadDialogMargin() xc.POINT {
	margin, err := i.GetDefaultDownloadDialogMargin()
	ReportErrorAtuo(err)
	return margin
}
