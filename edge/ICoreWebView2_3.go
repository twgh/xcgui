package edge

// ok

import (
	"errors"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

// ICoreWebView2_3 是 ICoreWebView2_2 接口的延续.
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2_3
type ICoreWebView2_3 struct {
	Vtbl *ICoreWebView2_3Vtbl
}

type ICoreWebView2_3Vtbl struct {
	ICoreWebView2_2Vtbl
	TrySuspend                          ComProc
	Resume                              ComProc
	GetIsSuspended                      ComProc
	SetVirtualHostNameToFolderMapping   ComProc
	ClearVirtualHostNameToFolderMapping ComProc
}

func (i *ICoreWebView2_3) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2_3) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2_3) QueryInterface(refiid, object uintptr) error {
	r, _, err := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), refiid, object)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetIsSuspended 获取 WebView 控件是否已挂起。
func (i *ICoreWebView2_3) GetIsSuspended() (bool, error) {
	var result bool
	r, _, err := i.Vtbl.GetIsSuspended.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&result)),
	)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return result, err
	}
	if r != 0 {
		return result, syscall.Errno(r)
	}
	return result, nil
}

// MustGetIsSuspended 获取 WebView 控件是否已挂起。忽略错误。
func (i *ICoreWebView2_3) MustGetIsSuspended() bool {
	r, _ := i.GetIsSuspended()
	return r
}

// Resume 恢复WebView，以便它恢复网页上的活动。
func (i *ICoreWebView2_3) Resume() error {
	r, _, err := i.Vtbl.Resume.Call(
		uintptr(unsafe.Pointer(i)),
	)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// TrySuspend 尝试挂起 WebView 控件, 以节省内存。
//
// handler: 接收 TrySuspend 方法的结果。
func (i *ICoreWebView2_3) TrySuspend(handler *ICoreWebView2TrySuspendCompletedHandler) error {
	r, _, err := i.Vtbl.TrySuspend.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(handler)),
	)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// ClearVirtualHostNameToFolderMapping 清除 SetVirtualHostNameToFolderMapping 添加的本地文件夹的主机名映射。
//
// hostName: 要清除的主机名。
func (i *ICoreWebView2_3) ClearVirtualHostNameToFolderMapping(hostName string) error {
	_hostName, err := windows.UTF16PtrFromString(hostName)
	if err != nil {
		return err
	}

	r, _, err := i.Vtbl.ClearVirtualHostNameToFolderMapping.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_hostName)),
	)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// SetVirtualHostNameToFolderMapping 设置虚拟主机名和文件夹路径之间的映射，以便通过该主机名对网站可用.
//
// hostName: 要映射的主机名。
//
// folderPath: 要映射的文件夹路径。
//
// accessKind: 资源访问权限。
func (i *ICoreWebView2_3) SetVirtualHostNameToFolderMapping(hostName, folderPath string, accessKind COREWEBVIEW2_HOST_RESOURCE_ACCESS_KIND) error {
	_hostName, err := windows.UTF16PtrFromString(hostName)
	if err != nil {
		return err
	}
	_folderPath, err := windows.UTF16PtrFromString(folderPath)
	if err != nil {
		return err
	}

	r, _, err := i.Vtbl.SetVirtualHostNameToFolderMapping.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_hostName)),
		uintptr(unsafe.Pointer(_folderPath)),
		uintptr(accessKind),
	)
	if !errors.Is(err, windows.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}
