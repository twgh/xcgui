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

// ReportErrorAuto 报告错误到全局错误回调, 自动获取调用者信息.
func ReportErrorAuto(err error) {
	if err != nil && webviewErrorCallBack != nil {
		webviewErrorCallBack(NewWebViewErrorWithSkip(err, 2))
	}
}

// WebViewError 是 WebView 错误.
type WebViewError struct {
	FullName string // 报错所在的完整函数路径, 例: github.com/twgh/xcgui/edge.(*ICoreWebView2).Reload
	Method   string // 报错所在的函数名, 例: Reload
	File     string // 报错所在文件, 例: D:/GOProject/src/github.com/twgh/xcgui/edge/ICoreWebView2.go
	Line     int    // 报错所在文件的行号
	Err      error  // 错误
}

func (e *WebViewError) Error() string {
	return e.Err.Error()
}

// ErrorWithMethod 返回例: Reload report error: xxx
func (e *WebViewError) ErrorWithMethod() string {
	return fmt.Sprintf("%s report error: %v", e.Method, e.Err)
}

// ErrorWithFullName 返回例: github.com/twgh/xcgui/edge.(*ICoreWebView2).Reload report error: xxx
func (e *WebViewError) ErrorWithFullName() string {
	return fmt.Sprintf("%s report error: %v", e.FullName, e.Err)
}

// ErrorWithFile 返回例: D:/GOProject/src/github.com/twgh/xcgui/edge/ICoreWebView2.go:480, github.com/twgh/xcgui/edge.(*ICoreWebView2).Reload report error: xxx
func (e *WebViewError) ErrorWithFile() string {
	return fmt.Sprintf("%s:%d, %s report error: %v", e.File, e.Line, e.FullName, e.Err)
}

// NewWebViewError 创建 WebViewError.
func NewWebViewError(method string, err error) *WebViewError {
	return &WebViewError{
		Method: method,
		Err:    err,
	}
}

// NewWebViewErrorAuto 创建 WebViewError. 自动获取调用者信息.
func NewWebViewErrorAuto(err error) *WebViewError {
	return NewWebViewErrorWithSkip(err, 2)
}

// NewWebViewErrorWithSkip 创建 WebViewError. 可以指定跳过多少层调用来自动获取调用者信息.
func NewWebViewErrorWithSkip(err error, skip int) *WebViewError {
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
	// 去除类型前缀（如 "main.(*ICoreWebView2_2)"）
	if strings.Contains(lastPart, ")") {
		return strings.Split(lastPart, ")")[1]
	}
	return lastPart
}
