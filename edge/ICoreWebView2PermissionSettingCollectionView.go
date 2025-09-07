package edge

import (
	"syscall"
	"unsafe"
)

// ICoreWebView2PermissionSettingCollectionView 提供对权限设置集合的只读访问。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2permissionsettingcollectionview
type ICoreWebView2PermissionSettingCollectionView struct {
	Vtbl *ICoreWebView2PermissionSettingCollectionViewVtbl
}

type ICoreWebView2PermissionSettingCollectionViewVtbl struct {
	IUnknownVtbl
	GetValueAtIndex ComProc
	GetCount        ComProc
}

func (i *ICoreWebView2PermissionSettingCollectionView) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2PermissionSettingCollectionView) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2PermissionSettingCollectionView) QueryInterface(refiid, object unsafe.Pointer) error {
	r, _, _ := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), uintptr(refiid), uintptr(object))
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetValueAtIndex 获取集合中指定索引处的权限设置。
func (i *ICoreWebView2PermissionSettingCollectionView) GetValueAtIndex(index uint32) (*ICoreWebView2PermissionSetting, error) {
	var permissionSetting *ICoreWebView2PermissionSetting
	r, _, _ := i.Vtbl.GetValueAtIndex.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(index),
		uintptr(unsafe.Pointer(&permissionSetting)),
	)
	if r != 0 {
		return nil, syscall.Errno(r)
	}
	return permissionSetting, nil
}

// GetCount 获取集合中权限设置的数量。
func (i *ICoreWebView2PermissionSettingCollectionView) GetCount() (uint32, error) {
	var value uint32
	r, _, _ := i.Vtbl.GetCount.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return value, nil
}

// MustGetValueAtIndex 获取集合中指定索引处的权限设置。出错时会触发全局错误回调。
func (i *ICoreWebView2PermissionSettingCollectionView) MustGetValueAtIndex(index uint32) *ICoreWebView2PermissionSetting {
	permissionSetting, err := i.GetValueAtIndex(index)
	ReportErrorAtuo(err)
	return permissionSetting
}

// MustGetCount 获取集合中权限设置的数量。出错时会触发全局错误回调。
func (i *ICoreWebView2PermissionSettingCollectionView) MustGetCount() uint32 {
	value, err := i.GetCount()
	ReportErrorAtuo(err)
	return value
}
