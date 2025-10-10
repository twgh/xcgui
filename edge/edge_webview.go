package edge

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/twgh/xcgui/common"
	"github.com/twgh/xcgui/wapi"
	"github.com/twgh/xcgui/wapi/wnd"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
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
	for hEle := range w.m {
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

func callbinding(w *WebView, d *rpcMessage) (interface{}, error) {
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
			if w.roundRadius > 0 {
				_ = wnd.SetWindowRound(hwnd, w.roundRadius)
			}
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
	case wapi.WM_DESTROY:
		if w := hwndContext.GetWindowContext(hwnd); w != nil {
			if w.cbDestroy != nil {
				w.cbDestroy(w)
			}
			w.isClose = true
			ReportErrorAuto(w.Close())
			return 0
		}
	case wapi.WM_NCDESTROY: // 窗口非客户区销毁, 在 WM_DESTROY 之后
		if w := hwndContext.GetWindowContext(hwnd); w != nil {
			if w.cbNCDestroy != nil {
				w.cbNCDestroy(w)
			}
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
// name: 函数名. 如果已有同名函数, 则会覆盖.
//   - 可以只是一个name, 将直接绑定在 window.
//   - 可以是多个层级: a.b.c.name, 调用方式: a.b.c.name().
//   - 每个层级的命名规则都必须和 js 变量命名规则一致.
//
// fun: 必须是一个函数:
//   - 函数参数可使用基本类型, 其它的自行测试, 暂不支持[]byte.
//   - 函数可以没有返回值.
//   - 函数返回值可以是一个值或一个error.
//   - 函数返回值可以是一个值和一个error.
func (w *WebView) Bind(name string, fun interface{}) error {
	if name == "" {
		return errors.New("name is empty")
	}
	if strings.Contains(name, " ") {
		return errors.New("name can't contain spaces")
	}
	if fun == nil {
		return errors.New("fun is nil")
	}
	v := reflect.ValueOf(fun)
	if v.Kind() != reflect.Func {
		return errors.New("only functions can be bound")
	}
	if n := v.Type().NumOut(); n > 2 {
		return errors.New("function may only return a value or a value+error")
	}

	w.rwxBindings.Lock()
	w.bindings[name] = fun
	w.rwxBindings.Unlock()

	initCode := fmt.Sprintf(`(function() {
    var RPC = window._rpc = (window._rpc || {
        nextSeq: 1,
        MAX_SEQ: 9007199254740991,
        cleanup: function() {
            var now = Date.now();
            Object.keys(RPC).forEach(function(key) {
                if (key === 'nextSeq' || key === 'MAX_SEQ' || key === 'cleanup') return;
                var req = RPC[key];
                if (req && req.resolved && (now - req.resolvedTime > 300000)) {
                    delete RPC[key];
                }
            });
        }
    });
    var parts = %q.split('.');
    var base = window;
    for (var i = 0; i < parts.length - 1; i++) {
        if (!base[parts[i]]) {
            base[parts[i]] = {};
        }
        base = base[parts[i]];
    }
    var funcName = parts[parts.length - 1];
    base[funcName] = function() {
        var seq = RPC.nextSeq++;
        if (RPC.nextSeq > RPC.MAX_SEQ - 1000) {
            RPC.nextSeq = 1;
            RPC.cleanup();
        }
        if (seq %% 100 === 0) {
            RPC.cleanup();
        }
        var promise = new Promise(function(resolve, reject) {
            RPC[seq] = {
                resolve: function(value) {
                    this.resolved = true;
                    this.resolvedTime = Date.now();
                    resolve(value);
                },
                reject: function(reason) {
                    this.resolved = true;
                    this.resolvedTime = Date.now();
                    reject(reason);
                },
                resolved: false,
                resolvedTime: 0
            };
        });
        window.external.invoke(JSON.stringify({
            id: seq,
            method: %q,
            params: Array.prototype.slice.call(arguments),
        }));
        return promise;
    };
})()`, name, name)

	err := w.Eval(initCode)
	if err != nil {
		return err
	}

	return w.AddScriptToExecuteOnDocumentCreated(initCode, func(errorCode syscall.Errno, id string) uintptr {
		if !errors.Is(errorCode, wapi.S_OK) {
			return 0
		}
		w.rwxBindings.Lock()
		w.bindingsid[name] = id // 存储id
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
			wapi.Sleep(0) // 避免 CPU 空转
		}
	}
	return ret, err
}

// BindLog 绑定一个控制台输出函数, 参数不限个数, 可使用基本类型, 在 js 代码中调用, 会在 go 控制台中输出.
//
// funcName: 自定义函数名, 为空默认为 glog.
func (w *WebView) BindLog(funcName ...string) error {
	var name string
	if len(funcName) > 0 {
		name = funcName[0]
	} else {
		name = "glog"
	}
	return w.Bind(name, func(msg ...interface{}) {
		fmt.Printf("%v\n", msg...)
	})
}

// SetTitle 更新原生窗口的标题。
//   - webview 是创建在一个用 wapi 创建的原生窗口里的, 然后原生窗口是被嵌入到炫彩窗口或元素里的.
func (w *WebView) SetTitle(title string) {
	wapi.SetWindowText(w.hwnd, title)
}

// Event_Destroy 宿主原生窗口销毁事件.
//   - webview 是创建在一个用 wapi 创建的原生窗口里的, 然后原生窗口是被嵌入到炫彩窗口或元素里的.
func (w *WebView) Event_Destroy(cb func(wv *WebView)) {
	w.cbDestroy = cb
}

// Event_NCDestroy 宿主原生窗口非客户区销毁事件.
//   - webview 是创建在一个用 wapi 创建的原生窗口里的, 然后原生窗口是被嵌入到炫彩窗口或元素里的.
//   - 执行顺序在 Event_Destroy 之后, 这个时候原生窗口句柄已无效.
func (w *WebView) Event_NCDestroy(cb func(wv *WebView)) {
	w.cbNCDestroy = cb
}

// 创建 WebView2 控制器
func newWebView2Controller(w *WebView, opt *WebViewOptions) error {
	var isDone bool
	var err2 error
	// WebView2 控制器创建完成回调
	w.Edge._cbCreateCoreWebView2ControllerCompleted = func(result syscall.Errno, controller *ICoreWebView2Controller) {
		if !errors.Is(result, wapi.S_OK) { // 失败
			err2 = errors.New("CreateCoreWebView2Controller failed: " + result.Error())
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

		// 获取浏览器设置
		settings, err := w.GetSettings()
		if err != nil {
			ReportErrorAuto(err)
		} else {
			// 设置是否可开启开发人员工具
			err = settings.SetAreDevToolsEnabled(opt.Debug)
			ReportErrorAuto(err)

			// 设置是否启用非客户区域支持
			s9, err := settings.GetICoreWebView2Settings9()
			if err != nil {
				ReportErrorAuto(err)
			} else {
				err = s9.SetIsNonClientRegionSupportEnabled(opt.AppDrag)
				ReportErrorAuto(err)
				s9.Release()
			}

			// 是否禁用默认的上下文菜单
			if !opt.DefaultContextMenus {
				err = settings.SetAreDefaultContextMenusEnabled(opt.DefaultContextMenus)
				ReportErrorAuto(err)
			}
			// 是否禁用状态栏
			if !opt.StatusBar {
				err = settings.SetIsStatusBarEnabled(opt.StatusBar)
				ReportErrorAuto(err)
			}
			// 是否禁用缩放控件
			if !opt.ZoomControl {
				err = settings.SetIsZoomControlEnabled(opt.ZoomControl)
				ReportErrorAuto(err)
			}

			s3, err := settings.GetICoreWebView2Settings3()
			if err != nil {
				ReportErrorAuto(err)
			} else {
				// 是否禁用浏览器快捷键
				if !opt.BrowserAcceleratorKeys {
					err = s3.SetAreBrowserAcceleratorKeysEnabled(opt.BrowserAcceleratorKeys)
					ReportErrorAuto(err)
				}
				s3.Release()
			}

			settings.Release()
		}
		err = w.Resize()
		ReportErrorAuto(err)

		// 添加 web 消息接收事件处理程序, 添加空回调, 主要目的是注册一下事件
		_, err = w.Event_WebMessageReceived(nil, true)
		ReportErrorAuto(err)
		// 添加权限请求事件处理程序, 添加空回调, 主要目的是注册一下事件
		_, err = w.Event_PermissionRequested(nil, true)
		ReportErrorAuto(err)
		// 添加在创建文档时要执行的脚本
		err = w.CoreWebView.AddScriptToExecuteOnDocumentCreated("window.external={invoke:s=>window.chrome.webview.postMessage(s)}", nil)
		ReportErrorAuto(err)

		if w.focusOnInit {
			err = w.Focus()
			ReportErrorAuto(err)
		}
		isDone = true
	}

	// 创建 WebView2 控制器
	if opt.ProfileName != "" || opt.PrivateMode || opt.DefaultBackgroundColor != nil || opt.ScriptLocale != "默认不设置" {
		env10, err := w.Edge.Environment.GetICoreWebView2Environment10()
		if err != nil {
			return fmt.Errorf("get ICoreWebView2Environment10 failed: %v", err)
		}
		defer env10.Release()

		options, err := env10.CreateCoreWebView2ControllerOptions()
		if err != nil {
			return fmt.Errorf("creating ICoreWebView2ControllerOptions failed: %v", err)
		}
		defer options.Release()

		if opt.ProfileName != "" {
			err = options.SetProfileName(opt.ProfileName)
			if err != nil {
				return fmt.Errorf("setting profile name failed: %v", err)
			}
		}

		if opt.PrivateMode {
			err = options.SetIsInPrivateModeEnabled(true)
			if err != nil {
				return fmt.Errorf("setting private mode failed: %v", err)
			}
		}

		if opt.ScriptLocale != "默认不设置" {
			opt2, err := options.GetICoreWebView2ControllerOptions2()
			if err != nil {
				ReportErrorAuto(err)
			} else {
				defer opt2.Release()
				err = opt2.SetScriptLocale(opt.ScriptLocale)
				if err != nil {
					return fmt.Errorf("setting script locale failed: %v", err)
				}
			}
		}

		if opt.DefaultBackgroundColor != nil {
			opt3, err := options.GetICoreWebView2ControllerOptions3()
			if err != nil {
				ReportErrorAuto(err)
			} else {
				defer opt3.Release()
				err = opt3.SetDefaultBackgroundColor(opt.DefaultBackgroundColor)
				ReportErrorAuto(err) // 出错不直接返回, 因为影响不大
			}
		}

		err = env10.CreateCoreWebView2ControllerWithOptions(w.hwnd, options, w.Edge.handlerCreateCoreWebView2ControllerCompleted)
		if err != nil {
			return fmt.Errorf("creating WebView2 controller with options failed: %v", err)
		}
	} else {
		err := w.Edge.Environment.CreateCoreWebView2Controller(w.hwnd, w.Edge.handlerCreateCoreWebView2ControllerCompleted)
		if err != nil {
			return fmt.Errorf("creating WebView2 controller failed: %v", err)
		}
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

func (w *WebView) msgcb_xcgui(msg string) {
	d := rpcMessage{}
	if err := json.Unmarshal(common.String2Bytes(msg), &d); err != nil {
		return
	}

	id := strconv.Itoa(d.ID)
	if res, err := callbinding(w, &d); err != nil {
		err = w.Eval("window._rpc[" + id + "].reject(" + jsString(err.Error()) + "); window._rpc[" + id + "] = undefined")
		ReportErrorAuto(err)
	} else if b, err := json.Marshal(res); err != nil {
		err = w.Eval("window._rpc[" + id + "].reject(" + jsString(err.Error()) + "); window._rpc[" + id + "] = undefined")
		ReportErrorAuto(err)
	} else {
		err = w.Eval("window._rpc[" + id + "].resolve(" + string(b) + "); window._rpc[" + id + "] = undefined")
		ReportErrorAuto(err)
	}
}
