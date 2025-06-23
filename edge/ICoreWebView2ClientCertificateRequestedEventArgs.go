package edge

import (
	"github.com/twgh/xcgui/common"
	"github.com/twgh/xcgui/wapi"
	"syscall"
	"unsafe"
)

// ICoreWebView2ClientCertificateRequestedEventArgs 客户端证书请求事件的参数.
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2clientcertificaterequestedeventargs
type ICoreWebView2ClientCertificateRequestedEventArgs struct {
	Vtbl *ICoreWebView2ClientCertificateRequestedEventArgsVtbl
}

type ICoreWebView2ClientCertificateRequestedEventArgsVtbl struct {
	IUnknownVtbl
	GetHost                          ComProc
	GetPort                          ComProc
	GetIsProxy                       ComProc
	GetAllowedCertificateAuthorities ComProc
	GetMutuallyTrustedCertificates   ComProc
	GetSelectedCertificate           ComProc
	PutSelectedCertificate           ComProc
	GetCancel                        ComProc
	PutCancel                        ComProc
	GetHandled                       ComProc
	PutHandled                       ComProc
	GetDeferral                      ComProc
}

func (i *ICoreWebView2ClientCertificateRequestedEventArgs) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2ClientCertificateRequestedEventArgs) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2ClientCertificateRequestedEventArgs) QueryInterface(refiid, object unsafe.Pointer) error {
	r, _, _ := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), uintptr(refiid), uintptr(object))
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetHost 获取请求客户端证书认证的服务器主机名。
func (i *ICoreWebView2ClientCertificateRequestedEventArgs) GetHost() (string, error) {
	var host *uint16
	r, _, _ := i.Vtbl.GetHost.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&host)),
	)
	if r != 0 {
		return "", syscall.Errno(r)
	}
	defer wapi.CoTaskMemFree(unsafe.Pointer(host))
	return common.UTF16PtrToString(host), nil
}

// GetPort 获取请求客户端证书认证的服务器端口。
func (i *ICoreWebView2ClientCertificateRequestedEventArgs) GetPort() (int, error) {
	var port int
	r, _, _ := i.Vtbl.GetPort.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&port)),
	)
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return port, nil
}

// GetIsProxy 获取发出此请求的服务器是否为 HTTP 代理.
//   - 如果服务器是源服务器，则返回 false。
func (i *ICoreWebView2ClientCertificateRequestedEventArgs) GetIsProxy() (bool, error) {
	var isProxy bool
	r, _, _ := i.Vtbl.GetIsProxy.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&isProxy)),
	)
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return isProxy, nil
}

// GetSelectedCertificate 获取所选证书。
func (i *ICoreWebView2ClientCertificateRequestedEventArgs) GetSelectedCertificate() (*ICoreWebView2ClientCertificate, error) {
	var cert *ICoreWebView2ClientCertificate
	r, _, _ := i.Vtbl.GetSelectedCertificate.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&cert)),
	)
	if r != 0 {
		return nil, syscall.Errno(r)
	}
	return cert, nil
}

// SetSelectedCertificate 设置用于响应服务器的证书。
func (i *ICoreWebView2ClientCertificateRequestedEventArgs) SetSelectedCertificate(cert *ICoreWebView2ClientCertificate) error {
	r, _, _ := i.Vtbl.PutSelectedCertificate.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(cert)),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetCancel 获取是否取消证书选择。
func (i *ICoreWebView2ClientCertificateRequestedEventArgs) GetCancel() (bool, error) {
	var cancel bool
	r, _, _ := i.Vtbl.GetCancel.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&cancel)),
	)
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return cancel, nil
}

// SetCancel 设置是否取消证书选择。
//   - 如果取消，则无论 Handled 属性如何，请求都会中止。默认值为 false。
func (i *ICoreWebView2ClientCertificateRequestedEventArgs) SetCancel(cancel bool) error {
	r, _, _ := i.Vtbl.PutCancel.Call(
		uintptr(unsafe.Pointer(i)),
		common.BoolPtr(cancel),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetDeferral 获取延迟对象. 使用此操作可在稍后完成该事件。
func (i *ICoreWebView2ClientCertificateRequestedEventArgs) GetDeferral() (*ICoreWebView2Deferral, error) {
	var deferral *ICoreWebView2Deferral
	r, _, _ := i.Vtbl.GetDeferral.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&deferral)),
	)
	if r != 0 {
		return nil, syscall.Errno(r)
	}
	return deferral, nil
}

// GetHandled 获取是否已处理此事件。
func (i *ICoreWebView2ClientCertificateRequestedEventArgs) GetHandled() (bool, error) {
	var handled bool
	r, _, _ := i.Vtbl.GetHandled.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&handled)),
	)
	if r != 0 {
		return false, syscall.Errno(r)
	}
	return handled, nil
}

// SetHandled 设置是否已处理此事件。
//   - 可以将此标志设置为 true，以便在有或没有证书的情况下响应服务器。
//   - 如果此标志为 true 且存在 SelectedCertificate，则使用选定证书响应服务器，否则不使用证书响应服务器。
//   - 默认情况下，Handled 和 Cancel 的值为 false，并显示默认的客户端证书选择对话框提示，允许用户选择证书。除非 Handled 设置为 true，否则将忽略 SelectedCertificate。
func (i *ICoreWebView2ClientCertificateRequestedEventArgs) SetHandled(handled bool) error {
	r, _, _ := i.Vtbl.PutHandled.Call(
		uintptr(unsafe.Pointer(i)),
		common.BoolPtr(handled),
	)
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetAllowedCertificateAuthorities 获取一个集合, 该集合包含服务器允许的证书颁发机构的 DER 编码可分辨名称的 Base64 编码。
func (i *ICoreWebView2ClientCertificateRequestedEventArgs) GetAllowedCertificateAuthorities() (*ICoreWebView2StringCollection, error) {
	var collection *ICoreWebView2StringCollection
	r, _, _ := i.Vtbl.GetAllowedCertificateAuthorities.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&collection)),
	)
	if r != 0 {
		return nil, syscall.Errno(r)
	}
	return collection, nil
}

// GetMutuallyTrustedCertificates 获取相互信任的 CA 证书集合.
func (i *ICoreWebView2ClientCertificateRequestedEventArgs) GetMutuallyTrustedCertificates() (*ICoreWebView2ClientCertificateCollection, error) {
	var collection *ICoreWebView2ClientCertificateCollection
	r, _, _ := i.Vtbl.GetMutuallyTrustedCertificates.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&collection)),
	)
	if r != 0 {
		return nil, syscall.Errno(r)
	}
	return collection, nil
}

// MustGetHost 获取请求客户端证书认证的服务器主机名。出错时会触发全局错误回调。
func (i *ICoreWebView2ClientCertificateRequestedEventArgs) MustGetHost() string {
	host, err := i.GetHost()
	ReportErrorAtuo(err)
	return host
}

// MustGetPort 获取请求客户端证书认证的服务器端口。出错时会触发全局错误回调。
func (i *ICoreWebView2ClientCertificateRequestedEventArgs) MustGetPort() int {
	port, err := i.GetPort()
	ReportErrorAtuo(err)
	return port
}

// MustGetIsProxy 获取发出此请求的服务器是否为 HTTP 代理. 出错时会触发全局错误回调。
//   - 如果服务器是源服务器，则返回 false。
func (i *ICoreWebView2ClientCertificateRequestedEventArgs) MustGetIsProxy() bool {
	isProxy, err := i.GetIsProxy()
	ReportErrorAtuo(err)
	return isProxy
}

// MustGetSelectedCertificate 获取所选证书。出错时会触发全局错误回调。
func (i *ICoreWebView2ClientCertificateRequestedEventArgs) MustGetSelectedCertificate() *ICoreWebView2ClientCertificate {
	cert, err := i.GetSelectedCertificate()
	ReportErrorAtuo(err)
	return cert
}

// MustGetCancel 获取是否取消证书选择。出错时会触发全局错误回调。
func (i *ICoreWebView2ClientCertificateRequestedEventArgs) MustGetCancel() bool {
	cancel, err := i.GetCancel()
	ReportErrorAtuo(err)
	return cancel
}

// MustGetDeferral 获取延迟对象。使用此操作可在稍后完成该事件。出错时会触发全局错误回调。
func (i *ICoreWebView2ClientCertificateRequestedEventArgs) MustGetDeferral() *ICoreWebView2Deferral {
	deferral, err := i.GetDeferral()
	ReportErrorAtuo(err)
	return deferral
}

// MustGetHandled 获取是否已处理此事件。出错时会触发全局错误回调。
func (i *ICoreWebView2ClientCertificateRequestedEventArgs) MustGetHandled() bool {
	handled, err := i.GetHandled()
	ReportErrorAtuo(err)
	return handled
}

// MustGetAllowedCertificateAuthorities 获取一个集合, 该集合包含服务器允许的证书颁发机构的 DER 编码可分辨名称的 Base64 编码。出错时会触发全局错误回调.
func (i *ICoreWebView2ClientCertificateRequestedEventArgs) MustGetAllowedCertificateAuthorities() *ICoreWebView2StringCollection {
	collection, err := i.GetAllowedCertificateAuthorities()
	ReportErrorAtuo(err)
	return collection
}

// MustGetMutuallyTrustedCertificates 获取相互信任的 CA 证书集合. 出错时会触发全局错误回调.
func (i *ICoreWebView2ClientCertificateRequestedEventArgs) MustGetMutuallyTrustedCertificates() *ICoreWebView2ClientCertificateCollection {
	collection, err := i.GetMutuallyTrustedCertificates()
	ReportErrorAtuo(err)
	return collection
}
