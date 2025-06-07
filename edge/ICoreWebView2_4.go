package edge

import (
	"errors"
	"github.com/twgh/xcgui/common"
	"github.com/twgh/xcgui/wapi"
	"syscall"
	"unsafe"
)

// ICoreWebView2_4 是 ICoreWebView2_3 接口的延续，以支持 FrameCreated 和 下载开始 事件。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2_4
type ICoreWebView2_4 struct {
	Vtbl *ICoreWebView2_4Vtbl
}

type ICoreWebView2_4Vtbl struct {
	ICoreWebView2_3Vtbl
	AddFrameCreated        ComProc
	RemoveFrameCreated     ComProc
	AddDownloadStarting    ComProc
	RemoveDownloadStarting ComProc
}

func (i *ICoreWebView2_4) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2_4) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2_4) QueryInterface(refiid, object unsafe.Pointer) error {
	r, _, err := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), uintptr(refiid), uintptr(object))
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// AddFrameCreated 添加框架创建事件处理程序。
//   - 当创建新的 iframe 时触发。
//   - 处理此事件以访问 ICoreWebView2Frame 对象。
func (i *ICoreWebView2_4) AddFrameCreated(eventHandler *ICoreWebView2FrameCreatedEventHandler, token *EventRegistrationToken) error {
	r, _, err := i.Vtbl.AddFrameCreated.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(token)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// RemoveFrameCreated 移除框架创建事件处理程序。
func (i *ICoreWebView2_4) RemoveFrameCreated(token EventRegistrationToken) error {
	r, _, err := i.Vtbl.RemoveFrameCreated.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(token.Value),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

/*TODO:

AddDownloadStarting
RemoveDownloadStarting
*/

// ICoreWebView2Frame 提供对 iframe 信息的直接访问。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2frame
type ICoreWebView2Frame struct {
	Vtbl *ICoreWebView2FrameVtbl
}

type ICoreWebView2FrameVtbl struct {
	IUnknownVtbl
	GetName                          ComProc
	AddNameChanged                   ComProc
	RemoveNameChanged                ComProc
	AddHostObjectToScriptWithOrigins ComProc
	RemoveHostObjectFromScript       ComProc
	AddDestroyed                     ComProc
	RemoveDestroyed                  ComProc
	IsDestroyed                      ComProc
}

// GetName 获取 iframe 的 window.name 属性的值。
//   - 默认值等同于声明它的 iframe HTML 标签。
//   - 即使 iframe 被销毁，你也可以访问此属性。
func (i *ICoreWebView2Frame) GetName() (string, error) {
	var name *uint16
	r, _, err := i.Vtbl.GetName.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&name)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return "", err
	}
	if r != 0 {
		return "", syscall.Errno(r)
	}
	result := common.UTF16PtrToString(name)
	wapi.CoTaskMemFree(unsafe.Pointer(name))
	return result, nil
}

// MustGetName 获取 iframe 的 window.name 属性的值。出错时会触发全局错误回调。
func (i *ICoreWebView2Frame) MustGetName() string {
	name, err := i.GetName()
	ReportErrorAtuo(err)
	return name
}

// AddNameChanged 添加名称改变事件处理程序.
//   - 当 iframe 更改其 window.name 属性时触发。
func (i *ICoreWebView2Frame) AddNameChanged(eventHandler *ICoreWebView2FrameNameChangedEventHandler, token *EventRegistrationToken) error {
	r, _, err := i.Vtbl.AddNameChanged.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(token)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// RemoveNameChanged 移除名称改变事件处理程序.
func (i *ICoreWebView2Frame) RemoveNameChanged(token EventRegistrationToken) error {
	r, _, err := i.Vtbl.RemoveNameChanged.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(token.Value),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// AddHostObjectToScriptWithOrigins 将提供的主机对象添加到在具有指定名称的 iframe 中运行的脚本中，该脚本适用于指定来源的列表。
//   - 只有当访问期间iframe的源与传入的源之一匹配时，该iframe才能访问宿主对象。
//   - 在与文档的源进行比较之前，提供的源将被规范化。因此，方案名称将转换为小写，主机将根据需要进行 Punycode 解码，默认端口值将被移除等等。这意味着源的主机可能是 Punycode 编码的，也可能不是，但无论如何都会匹配。
//   - 如果列表中包含格式错误的源，调用将失败。
//   - 该方法可以连续多次调用，而无需针对同一对象名称调用 RemoveHostObjectFromScript。它将用新对象和新的源列表替换先前的对象。
//
// todo: 这个对象的格式在 go 中还没测试明白. 需结合 go-ole 库, 在对象中嵌入 ole.IDispatch, 并实现其方法.
func (i *ICoreWebView2Frame) AddHostObjectToScriptWithOrigins(name string, object interface{}, origins []string) error {
	_name, err := syscall.UTF16PtrFromString(name)
	if err != nil {
		return err
	}

	count := len(origins)
	originsPtr := make([]*uint16, count)
	for i, origin := range origins {
		originsPtr[i], err = syscall.UTF16PtrFromString(origin)
		if err != nil {
			return err
		}
	}

	r, _, err := i.Vtbl.AddHostObjectToScriptWithOrigins.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_name)),
		uintptr(unsafe.Pointer(&object)),
		uintptr(count),
		uintptr(unsafe.Pointer(&originsPtr[0])),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// RemoveHostObjectFromScript 删除由名称指定的宿主对象，以便在 iframe 中的 JavaScript 代码无法再访问它。
//   - 虽然新的访问尝试会被拒绝，但如果该对象已被 iframe 中的 JavaScript 代码获取，那么 JavaScript 代码将继续有权访问该对象。
//   - 对已被移除或从未添加过的名称调用此方法会失败。如果 iframe 被销毁，此方法也会返回失败。
func (i *ICoreWebView2Frame) RemoveHostObjectFromScript(name string) error {
	_name, err := syscall.UTF16PtrFromString(name)
	if err != nil {
		return err
	}
	r, _, err := i.Vtbl.RemoveHostObjectFromScript.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_name)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// AddDestroyed 添加框架销毁事件处理程序.
//   - 当与此 ICoreWebView2Frame 对象对应的 iframe 被移除或包含该 iframe 的文档被销毁时触发。
func (i *ICoreWebView2Frame) AddDestroyed(eventHandler *ICoreWebView2FrameDestroyedEventHandler, token *EventRegistrationToken) error {
	r, _, err := i.Vtbl.AddDestroyed.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(token)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// RemoveDestroyed 移除框架销毁事件处理程序.
func (i *ICoreWebView2Frame) RemoveDestroyed(token EventRegistrationToken) error {
	r, _, err := i.Vtbl.RemoveDestroyed.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(token.Value),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// IsDestroyed 检查框架是否已被销毁. 在销毁事件期间返回 true。
func (i *ICoreWebView2Frame) IsDestroyed() (bool, error) {
	var destroyed bool
	r, _, err := i.Vtbl.IsDestroyed.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&destroyed)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return false, err
	}
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return destroyed, nil
}

// MustIsDestroyed 检查框架是否已被销毁。在销毁事件期间返回 true。出错时会触发全局错误回调。
func (i *ICoreWebView2Frame) MustIsDestroyed() bool {
	result, err := i.IsDestroyed()
	ReportErrorAtuo(err)
	return result
}
