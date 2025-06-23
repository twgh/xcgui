package edge

import (
	"github.com/twgh/xcgui/common"
	"github.com/twgh/xcgui/wapi"
	"syscall"
	"unsafe"
)

// ICoreWebView2ClientCertificate 提供对客户端证书元数据的访问权限。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2clientcertificate
type ICoreWebView2ClientCertificate struct {
	Vtbl *ICoreWebView2ClientCertificateVtbl
}

type ICoreWebView2ClientCertificateVtbl struct {
	IUnknownVtbl
	GetSubject                          ComProc
	GetIssuer                           ComProc
	GetValidFrom                        ComProc
	GetValidTo                          ComProc
	GetDerEncodedSerialNumber           ComProc
	GetDisplayName                      ComProc
	ToPemEncoding                       ComProc
	GetPemEncodedIssuerCertificateChain ComProc
	GetKind                             ComProc
}

func (i *ICoreWebView2ClientCertificate) AddRef() uintptr {
	r, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2ClientCertificate) Release() uintptr {
	r, _, _ := i.Vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2ClientCertificate) QueryInterface(refiid, object unsafe.Pointer) error {
	r, _, _ := i.Vtbl.QueryInterface.Call(uintptr(unsafe.Pointer(i)), uintptr(refiid), uintptr(object))
	if r != 0 {
		return syscall.Errno(r)
	}
	return nil
}

// GetSubject 获取证书主题.
func (i *ICoreWebView2ClientCertificate) GetSubject() (string, error) {
	var subject *uint16
	r, _, _ := i.Vtbl.GetSubject.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&subject)),
	)
	if r != 0 {
		return "", syscall.Errno(r)
	}
	defer wapi.CoTaskMemFree(unsafe.Pointer(subject))
	return common.UTF16PtrToString(subject), nil
}

// GetIssuer 获取颁发证书的证书颁发机构名称。
func (i *ICoreWebView2ClientCertificate) GetIssuer() (string, error) {
	var issuer *uint16
	r, _, _ := i.Vtbl.GetIssuer.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&issuer)),
	)
	if r != 0 {
		return "", syscall.Errno(r)
	}
	defer wapi.CoTaskMemFree(unsafe.Pointer(issuer))
	return common.UTF16PtrToString(issuer), nil
}

// GetValidFrom 获取证书的有效起始日期和时间，以自 UNIX 纪元以来的秒数表示。
func (i *ICoreWebView2ClientCertificate) GetValidFrom() (float64, error) {
	var validFrom float64
	r, _, _ := i.Vtbl.GetValidFrom.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&validFrom)),
	)
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return validFrom, nil
}

// GetValidTo 获取证书的有效截止日期和时间，以自 UNIX 纪元以来的秒数表示。
func (i *ICoreWebView2ClientCertificate) GetValidTo() (float64, error) {
	var validTo float64
	r, _, _ := i.Vtbl.GetValidTo.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&validTo)),
	)
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return validTo, nil
}

// GetDerEncodedSerialNumber 获取证书的 DER 编码序列号的 Base64 编码。
func (i *ICoreWebView2ClientCertificate) GetDerEncodedSerialNumber() (string, error) {
	var serialNumber *uint16
	r, _, _ := i.Vtbl.GetDerEncodedSerialNumber.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&serialNumber)),
	)
	if r != 0 {
		return "", syscall.Errno(r)
	}
	defer wapi.CoTaskMemFree(unsafe.Pointer(serialNumber))
	return common.UTF16PtrToString(serialNumber), nil
}

// GetDisplayName 获取证书的显示名称.
func (i *ICoreWebView2ClientCertificate) GetDisplayName() (string, error) {
	var displayName *uint16
	r, _, _ := i.Vtbl.GetDisplayName.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&displayName)),
	)
	if r != 0 {
		return "", syscall.Errno(r)
	}
	defer wapi.CoTaskMemFree(unsafe.Pointer(displayName))
	return common.UTF16PtrToString(displayName), nil
}

// ToPemEncoding 将证书转换为 PEM 编码数据。
func (i *ICoreWebView2ClientCertificate) ToPemEncoding() (string, error) {
	var pemData *uint16
	r, _, _ := i.Vtbl.ToPemEncoding.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&pemData)),
	)
	if r != 0 {
		return "", syscall.Errno(r)
	}
	defer wapi.CoTaskMemFree(unsafe.Pointer(pemData))
	return common.UTF16PtrToString(pemData), nil
}

// GetPemEncodedIssuerCertificateChain 获取 PEM 编码的客户端证书颁发者链集合。
//   - 在此集合中，第一个元素是当前证书，后面依次是中间证书1、中间证书2 …… 中间证书N - 1。根证书是集合中的最后一个元素。
func (i *ICoreWebView2ClientCertificate) GetPemEncodedIssuerCertificateChain() (*ICoreWebView2StringCollection, error) {
	var collection *ICoreWebView2StringCollection
	r, _, _ := i.Vtbl.GetPemEncodedIssuerCertificateChain.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&collection)),
	)
	if r != 0 {
		return nil, syscall.Errno(r)
	}
	return collection, nil
}

// GetKind 获取证书类型.
func (i *ICoreWebView2ClientCertificate) GetKind() (COREWEBVIEW2_CLIENT_CERTIFICATE_KIND, error) {
	var kind COREWEBVIEW2_CLIENT_CERTIFICATE_KIND
	r, _, _ := i.Vtbl.GetKind.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&kind)),
	)
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return kind, nil
}

// MustGetSubject 获取证书主题. 出错时会触发全局错误回调.
func (i *ICoreWebView2ClientCertificate) MustGetSubject() string {
	subject, err := i.GetSubject()
	ReportErrorAtuo(err)
	return subject
}

// MustGetIssuer 获取颁发证书的证书颁发机构名称. 出错时会触发全局错误回调.
func (i *ICoreWebView2ClientCertificate) MustGetIssuer() string {
	issuer, err := i.GetIssuer()
	ReportErrorAtuo(err)
	return issuer
}

// MustGetValidFrom 获取证书的有效起始日期和时间，以自 UNIX 纪元以来的秒数表示. 出错时会触发全局错误回调.
func (i *ICoreWebView2ClientCertificate) MustGetValidFrom() float64 {
	validFrom, err := i.GetValidFrom()
	ReportErrorAtuo(err)
	return validFrom
}

// MustGetValidTo 获取证书的有效截止日期和时间，以自 UNIX 纪元以来的秒数表示. 出错时会触发全局错误回调.
func (i *ICoreWebView2ClientCertificate) MustGetValidTo() float64 {
	validTo, err := i.GetValidTo()
	ReportErrorAtuo(err)
	return validTo
}

// MustGetDerEncodedSerialNumber 获取证书的 DER 编码序列号的 Base64 编码. 出错时会触发全局错误回调.
func (i *ICoreWebView2ClientCertificate) MustGetDerEncodedSerialNumber() string {
	serialNumber, err := i.GetDerEncodedSerialNumber()
	ReportErrorAtuo(err)
	return serialNumber
}

// MustGetDisplayName 获取证书的显示名称. 出错时会触发全局错误回调.
func (i *ICoreWebView2ClientCertificate) MustGetDisplayName() string {
	displayName, err := i.GetDisplayName()
	ReportErrorAtuo(err)
	return displayName
}

// MustToPemEncoding 将证书转换为 PEM 编码数据. 出错时会触发全局错误回调.
func (i *ICoreWebView2ClientCertificate) MustToPemEncoding() string {
	pemData, err := i.ToPemEncoding()
	ReportErrorAtuo(err)
	return pemData
}

// MustGetPemEncodedIssuerCertificateChain 获取 PEM 编码的客户端证书颁发者链集合. 出错时会触发全局错误回调.
//   - 在此集合中，第一个元素是当前证书，后面依次是中间证书1、中间证书2 …… 中间证书N - 1。根证书是集合中的最后一个元素。
func (i *ICoreWebView2ClientCertificate) MustGetPemEncodedIssuerCertificateChain() *ICoreWebView2StringCollection {
	collection, err := i.GetPemEncodedIssuerCertificateChain()
	ReportErrorAtuo(err)
	return collection
}

// MustGetKind 获取证书类型. 出错时会触发全局错误回调.
func (i *ICoreWebView2ClientCertificate) MustGetKind() COREWEBVIEW2_CLIENT_CERTIFICATE_KIND {
	kind, err := i.GetKind()
	ReportErrorAtuo(err)
	return kind
}
