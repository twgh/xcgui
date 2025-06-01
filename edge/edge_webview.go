package edge

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/twgh/xcgui/wapi"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
	"reflect"
	"strings"
	"sync"
	"syscall"
	"time"
)

var (
	hwndContext = newWindowContext() // 原生窗口上下文
	xcContext   = newWindowContext() // 炫彩窗口或元素上下文
)

// 窗口上下文
type windowContext struct {
	m    map[uintptr]*WebView
	lock sync.RWMutex
}

// 创建一个窗口上下文
func newWindowContext() *windowContext {
	return &windowContext{
		m: map[uintptr]*WebView{},
	}
}

func (w *windowContext) GetWindowContext(handle uintptr) *WebView {
	w.lock.RLock()
	defer w.lock.RUnlock()
	return w.m[handle]
}

func (w *windowContext) SetWindowContext(handle uintptr, data *WebView) {
	w.lock.Lock()
	defer w.lock.Unlock()
	w.m[handle] = data
}

func (w *windowContext) DeleteWindowContext(handle uintptr) {
	w.lock.Lock()
	defer w.lock.Unlock()
	delete(w.m, handle)
}

// GetHEles 获取已记录的所有属于指定窗口的元素句柄.
func (w *windowContext) GetHEles(handle int) []int {
	var hEles []int
	for hEle, _ := range w.m {
		hEle2 := int(hEle)
		if xc.XC_IsHELE(hEle2) && xc.XWidget_GetHWINDOW(hEle2) == handle {
			hEles = append(hEles, hEle2)
		}
	}
	return hEles
}

type rpcMessage struct {
	ID     int               `json:"id"`
	Method string            `json:"method"`
	Params []json.RawMessage `json:"params"`
}

func jsString(v interface{}) string { b, _ := json.Marshal(v); return string(b) }

func (w *WebView) callbinding(d *rpcMessage) (interface{}, error) {
	w.rwxBindings.RLock()
	f, ok := w.bindings[d.Method]
	w.rwxBindings.RUnlock()
	if !ok {
		return nil, nil
	}

	v := reflect.ValueOf(f)
	isVariadic := v.Type().IsVariadic()
	numIn := v.Type().NumIn()
	if (isVariadic && len(d.Params) < numIn-1) || (!isVariadic && len(d.Params) != numIn) {
		return nil, errors.New("function arguments mismatch")
	}
	args := make([]reflect.Value, 0)
	for i := range d.Params {
		var arg reflect.Value
		if isVariadic && i >= numIn-1 {
			arg = reflect.New(v.Type().In(numIn - 1).Elem())
		} else {
			arg = reflect.New(v.Type().In(i))
		}
		if err := json.Unmarshal(d.Params[i], arg.Interface()); err != nil {
			return nil, err
		}
		args = append(args, arg.Elem())
	}

	errorType := reflect.TypeOf((*error)(nil)).Elem()
	res := v.Call(args)
	switch len(res) {
	case 0:
		// No results from the function, just return nil
		return nil, nil

	case 1:
		// One result may be a value, or an error
		if res[0].Type().Implements(errorType) {
			if res[0].Interface() != nil {
				return nil, res[0].Interface().(error)
			}
			return nil, nil
		}
		return res[0].Interface(), nil

	case 2:
		// Two results: first one is value, second is error
		if !res[1].Type().Implements(errorType) {
			return nil, errors.New("second return value must be an error")
		}
		if res[1].Interface() == nil {
			return res[0].Interface(), nil
		}
		return res[0].Interface(), res[1].Interface().(error)

	default:
		return nil, errors.New("unexpected number of return values")
	}
}

// 非客户区鼠标光标句柄
var cursorNoClient uintptr

// 设置非客户区的鼠标光标.
func setNoClientCursor() {
	if cursorNoClient == 0 {
		cursorNoClient = wapi.LoadImageW(0, uintptr(wapi.IDC_ARROW), wapi.IMAGE_CURSOR, 0, 0, wapi.LR_DEFAULTSIZE|wapi.LR_SHARED)
	}
	if cursorNoClient != 0 {
		wapi.SetCursor(cursorNoClient)
	}
}

func wndproc(hwnd uintptr, message uint32, wParam, lParam uintptr) uintptr {
	switch message {
	case wapi.WM_SETCURSOR:
		setNoClientCursor()
		return 0
	case wapi.WM_MOVE, wapi.WM_MOVING:
		if w := hwndContext.GetWindowContext(hwnd); w != nil {
			_ = w.NotifyParentWindowPositionChanged()
			return 0
		}
	case wapi.WM_SIZE:
		if w := hwndContext.GetWindowContext(hwnd); w != nil {
			_ = w.Resize()
			return 0
		}
	case wapi.WM_ACTIVATE:
		if wParam == wapi.WA_INACTIVE {
			break
		}
		if w := hwndContext.GetWindowContext(hwnd); w != nil && w.autofocus {
			_ = w.Focus()
			return 0
		}
	case wapi.WM_CLOSE:
		if w := hwndContext.GetWindowContext(hwnd); w != nil {
			ReportErrorAtuo(w.Close())
			wapi.DestroyWindow(hwnd)
			return 0
		}
	case wapi.WM_NCDESTROY: // 窗口非客户区销毁, 在 WM_DESTROY 之后
		if w := hwndContext.GetWindowContext(hwnd); w != nil {
			// 移除事件
			if xc.XC_IsHWINDOW(w.hParent) { // 原生窗口宿主是炫彩窗口
				xc.XWnd_RemoveEventC(w.hParent, xcc.XWM_WINDPROC, onWndProc)
				xcContext.DeleteWindowContext(uintptr(w.hParent))
			} else if xc.XC_IsHELE(w.hParent) { // 原生窗口宿主是炫彩元素
				xc.XEle_RemoveEventC(w.hParent, xcc.XE_SIZE, onEleSize)
				xc.XEle_RemoveEventC(w.hParent, xcc.XE_SHOW, onEleShow)
				xc.XEle_RemoveEventC(w.hParent, xcc.XE_DESTROY, onEleDestroy)
				xcContext.DeleteWindowContext(uintptr(w.hParent))
			}
			hwndContext.DeleteWindowContext(w.hwnd)
			return 0
		}
	}
	return wapi.DefWindowProc(hwnd, message, wParam, lParam)
}

/*// SetSize 更新原生窗口大小。
//   - webview 是创建在一个用 wapi 创建的原生窗口里的, 然后原生窗口是被嵌入到炫彩窗口或元素里的.
func (w *WebView) SetSize(width, height int32) {
	wapi.SetWindowPos(w.hwnd, 0, 0, 0, width, height, wapi.SWP_NOMOVE|wapi.SWP_NOZORDER|wapi.SWP_NOACTIVATE)
}*/

// Bind 绑定一个 Go 函数，使其以给定的名称作为全局 JavaScript 函数出现。必须在UI线程执行.
//
// f 必须是一个函数:
//   - 函数参数可使用基本类型, 其它的自行测试
//   - 函数可以没有返回值
//   - 函数返回值可以是一个值或一个error
//   - 函数返回值可以是一个值和一个error
func (w *WebView) Bind(name string, f interface{}) error {
	v := reflect.ValueOf(f)
	if v.Kind() != reflect.Func {
		return errors.New("only functions can be bound")
	}
	if n := v.Type().NumOut(); n > 2 {
		return errors.New("function may only return a value or a value+error")
	}
	w.rwxBindings.Lock()
	w.bindings[name] = f
	w.rwxBindings.Unlock()

	initCode := "(function() { var name = " + jsString(name) + ";" + `
		var RPC = window._rpc = (window._rpc || {nextSeq: 1});
		window[name] = function() {
		  var seq = RPC.nextSeq++;
		  var promise = new Promise(function(resolve, reject) {
			RPC[seq] = {
			  resolve: resolve,
			  reject: reject,
			};
		  });
		  window.external.invoke(JSON.stringify({
			id: seq,
			method: name,
			params: Array.prototype.slice.call(arguments),
		  }));
		  return promise;
		}
	})()`

	err := w.Eval(initCode)
	if err != nil {
		return err
	}

	return w.AddScriptToExecuteOnDocumentCreated(initCode, func(errorCode syscall.Errno, id string) uintptr {
		if !errors.Is(errorCode, wapi.S_OK) {
			return 0
		}
		// 存储id
		w.rwxBindings.Lock()
		w.bindingsid[name] = id
		w.rwxBindings.Unlock()
		return 0
	})
}

// Unbind 删除绑定的函数.
func (w *WebView) Unbind(name string) error {
	w.rwxBindings.Lock()
	delete(w.bindings, name)
	id := w.bindingsid[name]
	delete(w.bindingsid, name)
	w.rwxBindings.Unlock()
	return w.RemoveScriptToExecuteOnDocumentCreated(id)
}

// UnbindAll 删除所有绑定的函数.
func (w *WebView) UnbindAll() {
	w.rwxBindings.Lock()
	w.bindings = map[string]interface{}{}

	for _, id := range w.bindingsid {
		_ = w.RemoveScriptToExecuteOnDocumentCreated(id)
	}
	w.bindingsid = map[string]string{}
	w.rwxBindings.Unlock()
}

// GetHWND 返回 webview 所在的原生窗口句柄.
//   - webview 是创建在一个用 wapi 创建的原生窗口里的, 然后原生窗口是被嵌入到炫彩窗口或元素里的.
func (w *WebView) GetHWND() uintptr {
	return w.hwnd
}

// GetHWindow 返回 webview 所在的炫彩窗口句柄.
//   - webview 是创建在一个用 wapi 创建的原生窗口里的, 然后原生窗口是被嵌入到炫彩窗口或元素里的.
func (w *WebView) GetHWindow() int {
	return w.hWindow
}

// GetHParernt 返回 webview 所在的炫彩元素或窗口句柄.
//   - webview 是创建在一个用 wapi 创建的原生窗口里的, 然后原生窗口是被嵌入到炫彩窗口或元素里的.
func (w *WebView) GetHParent() int {
	return w.hParent
}

// EvalAsync 执行 js 代码, 可在回调函数中异步获取结果. 必须在 UI 线程执行.
//
// js: javascript 代码.
//
// cb: 回调函数, 在回调函数里可获取 js 代码返回值.
func (w *WebView) EvalAsync(js string, cb func(errorCode syscall.Errno, result string) uintptr) error {
	return w.ExecuteScript(js, cb)
}

var (
	// ErrEvalTimeout 是执行 js 代码超时.
	ErrEvalTimeout = errors.New("执行超时")
)

// EvalSync 执行 js 代码, 同步取回返回值. 必须在 UI 线程执行.
//
// js: javascript 代码.
//
// timeout: 超时时间, 为空则一直等待.
func (w *WebView) EvalSync(js string, timeout ...time.Duration) (string, error) {
	var err error
	var ret string
	var isDone bool
	var limitT time.Duration
	if len(timeout) > 0 {
		limitT = timeout[0]
	}
	var startT time.Time
	if limitT > 0 {
		startT = time.Now()
	}

	err = w.ExecuteScript(js, func(errorCode syscall.Errno, result string) uintptr {
		if !errors.Is(errorCode, wapi.S_OK) {
			err = errorCode
			return 0
		}
		ret = result
		isDone = true
		return 0
	})
	if err != nil {
		return "", err
	}

	// 等待结果
	// 启动消息循环处理
	var msg wapi.MSG
	for !isDone {
		// 等待超时
		if limitT > 0 {
			if time.Since(startT) > limitT {
				return "", ErrEvalTimeout
			}
		}

		hasMessage := wapi.PeekMessage(&msg, 0, 0, 0, wapi.PM_REMOVE)
		if hasMessage {
			wapi.TranslateMessage(&msg)
			wapi.DispatchMessage(&msg)
		} else {
			wapi.Sleep(1) // 避免 CPU 空转
		}
	}
	return ret, err
}

// BindLog 绑定一个日志输出函数, 参数不限个数, 在js代码中调用, 会在go控制台中输出.
//
// funcName: 自定义函数名, 为空默认为glog.
func (w *WebView) BindLog(funcName ...string) error {
	name := "glog"
	if len(funcName) > 0 {
		name = funcName[0]
		// 名字中不能有空格
		if strings.Contains(name, " ") {
			return errors.New("funcName 中不能有空格")
		}
	}
	return w.Bind(name, func(msg ...interface{}) {
		fmt.Printf("%v\n", msg...)
	})
}

// 创建 WebView2 控制器
func (w *WebView) newWebView2Controller() error {
	var isDone bool
	var err2 error
	// WebView2 控制器创建完成回调
	w.Edge._cbCreateCoreWebView2ControllerCompleted = func(result uintptr, controller *ICoreWebView2Controller) {
		if result != 0 { // 失败
			err2 = fmt.Errorf("CreateCoreWebView2Controller failed: 0x%08x", result)
			isDone = true
			return
		}
		controller.AddRef()
		w.Controller = controller

		var err error
		w.CoreWebView, err = controller.GetCoreWebView2()
		if err != nil {
			err2 = errors.New("GetCoreWebView2 failed: " + err.Error())
			isDone = true
			return
		}
		w.CoreWebView.AddRef()

		// 添加 web 消息接收事件处理程序
		err = w.CoreWebView.AddWebMessageReceived(w.handlerWebMessageReceivedEvent, w.EventRegistrationToken)
		ReportErrorAtuo(err)
		// 添加权限请求事件处理程序
		err = w.CoreWebView.AddPermissionRequested(w.handlerPermissionRequestedEvent, w.EventRegistrationToken)
		ReportErrorAtuo(err)
		// 添加在创建文档时要执行的脚本
		err = w.CoreWebView.AddScriptToExecuteOnDocumentCreated("window.external={invoke:s=>window.chrome.webview.postMessage(s)}", nil)
		ReportErrorAtuo(err)

		if w.focusOnInit {
			err = w.Focus()
			ReportErrorAtuo(err)
		}
		isDone = true
	}

	// 创建 WebView2 控制器
	err := w.Edge.Environment.CreateCoreWebView2Controller(w.hwnd, w.Edge.handler_CreateCoreWebView2ControllerCompleted)
	if err != nil {
		return fmt.Errorf("creating WebView2 controller return: %v", err)
	}

	// 等待 webview2 控制器创建完成
	var msg wapi.MSG
	for !isDone {
		if wapi.GetMessage(&msg, 0, 0, 0) == 0 {
			break
		}
		wapi.TranslateMessage(&msg)
		wapi.DispatchMessage(&msg)
	}
	return err2
}
