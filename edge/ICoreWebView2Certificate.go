package edge

import (
	"syscall"
	"unsafe"

	"github.com/twgh/xcgui/common"
	"github.com/twgh/xcgui/wapi"
)

// ICoreWebView2Certificate 提供对证书元数据的访问。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2certificate
type ICoreWebView2Certificate struct {
	vtbl *ICoreWebView2CertificateVtbl
}

type ICoreWebView2CertificateVtbl struct {
	IUnknownVtbl
	GetSubject                          ComProc
	GetIssuer                           ComProc
	GetValidFrom                        ComProc
	GetValidTo                          ComProc
	GetDerEncodedSerialNumber           ComProc
	GetDisplayName                      ComProc
	ToPemEncoding                       ComProc
	GetPemEncodedIssuerCertificateChain ComProc
}

// GetSubject 获取证书主题。
func (i *ICoreWebView2Certificate) GetSubject() (string, error) {
	var _value *uint16
	r, _, _ := i.vtbl.GetSubject.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_value)),
	)
	if r != 0 {
		return "", syscall.Errno(r)
	}
	value := common.UTF16PtrToString(_value)
	wapi.CoTaskMemFree(unsafe.Pointer(_value))
	return value, nil
}

// GetIssuer 获取颁发该证书的证书颁发机构的名称。
func (i *ICoreWebView2Certificate) GetIssuer() (string, error) {
	var _value *uint16
	r, _, _ := i.vtbl.GetIssuer.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_value)),
	)
	if r != 0 {
		return "", syscall.Errno(r)
	}
	value := common.UTF16PtrToString(_value)
	wapi.CoTaskMemFree(unsafe.Pointer(_value))
	return value, nil
}

// GetValidFrom 获取证书的有效起始日期和时间，以自 UNIX 纪元以来的秒数表示。
func (i *ICoreWebView2Certificate) GetValidFrom() (float64, error) {
	var _value float64
	r, _, _ := i.vtbl.GetValidFrom.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_value)),
	)
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return _value, nil
}

// GetValidTo 获取证书的有效到期日期和时间，以自UNIX纪元以来的秒数表示。
func (i *ICoreWebView2Certificate) GetValidTo() (float64, error) {
	var _value float64
	r, _, _ := i.vtbl.GetValidTo.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_value)),
	)
	if r != 0 {
		return 0, syscall.Errno(r)
	}
	return _value, nil
}

// GetDerEncodedSerialNumber 获取证书的 DER 编码序列号的 Base64 编码。
func (i *ICoreWebView2Certificate) GetDerEncodedSerialNumber() (string, error) {
	var _value *uint16
	r, _, _ := i.vtbl.GetDerEncodedSerialNumber.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_value)),
	)
	if r != 0 {
		return "", syscall.Errno(r)
	}
	value := common.UTF16PtrToString(_value)
	wapi.CoTaskMemFree(unsafe.Pointer(_value))
	return value, nil
}

// GetDisplayName 获取证书的显示名称。
func (i *ICoreWebView2Certificate) GetDisplayName() (string, error) {
	var _value *uint16
	r, _, _ := i.vtbl.GetDisplayName.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_value)),
	)
	if r != 0 {
		return "", syscall.Errno(r)
	}
	value := common.UTF16PtrToString(_value)
	wapi.CoTaskMemFree(unsafe.Pointer(_value))
	return value, nil
}

// ToPemEncoding 将证书转换为 PEM 编码格式, 返回 DER 编码证书的 Base64 编码。
func (i *ICoreWebView2Certificate) ToPemEncoding() (string, error) {
	var _pemEncodedData *uint16
	r, _, _ := i.vtbl.ToPemEncoding.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_pemEncodedData)),
	)
	if r != 0 {
		return "", syscall.Errno(r)
	}
	pemEncodedData := common.UTF16PtrToString(_pemEncodedData)
	wapi.CoTaskMemFree(unsafe.Pointer(_pemEncodedData))
	return pemEncodedData, nil
}

// GetPemEncodedIssuerCertificateChain 获取 PEM 编码的证书颁发者链集合。
//   - 在此集合中，第一个元素是当前证书，其后是中间证书1、中间证书2……中间证书N-1。根证书是该集合中的最后一个元素。
func (i *ICoreWebView2Certificate) GetPemEncodedIssuerCertificateChain() (*ICoreWebView2StringCollection, error) {
	var _value *ICoreWebView2StringCollection
	r, _, _ := i.vtbl.GetPemEncodedIssuerCertificateChain.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_value)),
	)
	if r != 0 {
		return nil, syscall.Errno(r)
	}
	return _value, nil
}

// MustGetSubject 获取证书主题。出错时会触发全局错误回调。
func (i *ICoreWebView2Certificate) MustGetSubject() string {
	result, err := i.GetSubject()
	ReportErrorAtuo(err)
	return result
}

// MustGetIssuer 获取颁发该证书的证书颁发机构的名称。出错时会触发全局错误回调。
func (i *ICoreWebView2Certificate) MustGetIssuer() string {
	result, err := i.GetIssuer()
	ReportErrorAtuo(err)
	return result
}

// MustGetValidFrom 获取证书的有效起始日期和时间，以自 UNIX 纪元以来的秒数表示。出错时会触发全局错误回调。
func (i *ICoreWebView2Certificate) MustGetValidFrom() float64 {
	result, err := i.GetValidFrom()
	ReportErrorAtuo(err)
	return result
}

// MustGetValidTo 获取证书的有效到期日期和时间，以自 UNIX 纪元以来的秒数表示。出错时会触发全局错误回调。
func (i *ICoreWebView2Certificate) MustGetValidTo() float64 {
	result, err := i.GetValidTo()
	ReportErrorAtuo(err)
	return result
}

// MustGetDerEncodedSerialNumber 获取证书的 DER 编码序列号的 Base64 编码。出错时会触发全局错误回调。
func (i *ICoreWebView2Certificate) MustGetDerEncodedSerialNumber() string {
	result, err := i.GetDerEncodedSerialNumber()
	ReportErrorAtuo(err)
	return result
}

// MustGetDisplayName 获取证书的显示名称。出错时会触发全局错误回调。
func (i *ICoreWebView2Certificate) MustGetDisplayName() string {
	result, err := i.GetDisplayName()
	ReportErrorAtuo(err)
	return result
}

// MustToPemEncoding 将证书转换为 PEM 编码格式, 返回 DER 编码证书的 Base64 编码。出错时会触发全局错误回调。
func (i *ICoreWebView2Certificate) MustToPemEncoding() string {
	result, err := i.ToPemEncoding()
	ReportErrorAtuo(err)
	return result
}

// MustGetPemEncodedIssuerCertificateChain 获取 PEM 编码的证书颁发者链集合。出错时会触发全局错误回调。
//   - 在此集合中，第一个元素是当前证书，其后是中间证书1、中间证书2……中间证书N-1。根证书是该集合中的最后一个元素。
func (i *ICoreWebView2Certificate) MustGetPemEncodedIssuerCertificateChain() *ICoreWebView2StringCollection {
	result, err := i.GetPemEncodedIssuerCertificateChain()
	ReportErrorAtuo(err)
	return result
}
