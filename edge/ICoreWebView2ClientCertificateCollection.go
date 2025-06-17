package edge

// ok

import (
	"errors"
	"github.com/twgh/xcgui/wapi"
	"syscall"
	"unsafe"
)

// ICoreWebView2ClientCertificateCollection 客户端证书集合.
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2clientcertificatecollection
type ICoreWebView2ClientCertificateCollection struct {
	Vtbl *ICoreWebView2ClientCertificateCollectionVtbl
}

type ICoreWebView2ClientCertificateCollectionVtbl struct {
	IUnknownVtbl
	GetCount        ComProc
	GetValueAtIndex ComProc
}

func (i *ICoreWebView2ClientCertificateCollection) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2ClientCertificateCollection) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2ClientCertificateCollection) QueryInterface(refiid, object unsafe.Pointer) error {
	r, _, err := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), uintptr(refiid), uintptr(object))
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return err
	}
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetCount 获取集合中的证书数量.
func (i *ICoreWebView2ClientCertificateCollection) GetCount() (uint32, error) {
	var count uint32
	r, _, err := i.Vtbl.GetCount.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&count)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return 0, err
	}
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return count, nil
}

// GetValueAtIndex 获取指定索引处的证书.
func (i *ICoreWebView2ClientCertificateCollection) GetValueAtIndex(index uint32) (*ICoreWebView2ClientCertificate, error) {
	var cert *ICoreWebView2ClientCertificate
	r, _, err := i.Vtbl.GetValueAtIndex.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(index),
		uintptr(unsafe.Pointer(&cert)),
	)
	if !errors.Is(err, wapi.ERROR_SUCCESS) {
		return nil, err
	}
	if r != 0 {
		return nil, syscall.Errno(r)
	}
	return cert, nil
}

// MustGetCount 获取集合中的证书数量. 出错时会触发全局错误回调.
func (i *ICoreWebView2ClientCertificateCollection) MustGetCount() uint32 {
	count, err := i.GetCount()
	ReportErrorAtuo(err)
	return count
}

// MustGetValueAtIndex 获取指定索引处的证书. 出错时会触发全局错误回调.
func (i *ICoreWebView2ClientCertificateCollection) MustGetValueAtIndex(index uint32) *ICoreWebView2ClientCertificate {
	cert, err := i.GetValueAtIndex(index)
	ReportErrorAtuo(err)
	return cert
}
