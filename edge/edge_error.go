package edge

import (
	"fmt"
	"runtime"
	"strings"
)

// 全局错误回调
var webviewErrorCallBack func(err *WebViewError)

// SetErrorCallBack 设置全局 WebView 错误回调. 调用 Must 系列方法时, 出错会触发该回调. 还有一些不方便直接 return 的地方也会把错误报告到该回调.
func SetErrorCallBack(cb func(*WebViewError)) {
	webviewErrorCallBack = cb
}

// ReportError 报告错误到全局错误回调.
func ReportError(method string, err error) {
	if err != nil && webviewErrorCallBack != nil {
		webviewErrorCallBack(NewWebViewError(method, err))
	}
}

// ReportError2 报告错误到全局错误回调, 自动获取调用者函数名.
func ReportError2(err error) {
	if err != nil && webviewErrorCallBack != nil {
		webviewErrorCallBack(NewWebViewError2(err))
	}
}

// WebViewError 是 WebView 错误.
type WebViewError struct {
	FullName string
	Method   string
	File     string
	Line     int
	Err      error
}

func (e *WebViewError) Error() string {
	return fmt.Sprintf("%s failed: %v", e.FullName, e.Err)
}

// NewWebViewError2 创建 WebViewError. 自动获取调用者方法名. 向上查找 2 层调用者.
func NewWebViewError2(err error) *WebViewError {
	return newWebViewErrorWithSkip(err, 2)
}

// NewWebViewError1 创建 WebViewError. 自动获取调用者方法名. 向上查找 1 层调用者.
func NewWebViewError1(err error) *WebViewError {
	return newWebViewErrorWithSkip(err, 1)
}

func newWebViewErrorWithSkip(err error, skip int) *WebViewError {
	pc, _file, line, ok := runtime.Caller(skip)
	if !ok {
		return &WebViewError{Method: "unknown", Err: err}
	}
	fullName := runtime.FuncForPC(pc).Name()
	methodName := extractMethodName(fullName)
	return &WebViewError{FullName: fullName, Method: methodName, File: _file, Line: line, Err: err}
}

// 从完整函数名中提取方法名
func extractMethodName(fullName string) string {
	// 示例输入：main.(*ICoreWebView2_2).MustGetCookieManager
	parts := strings.Split(fullName, ".")
	if len(parts) == 0 {
		return fullName
	}
	lastPart := parts[len(parts)-1]
	// 去除类型前缀（如 "(*ICoreWebView2_2)"）
	if strings.Contains(lastPart, ")") {
		return strings.Split(lastPart, ")")[1]
	}
	return lastPart
}

// NewWebViewError 创建 WebViewError.
func NewWebViewError(method string, err error) *WebViewError {
	return &WebViewError{
		Method: method,
		Err:    err,
	}
}
