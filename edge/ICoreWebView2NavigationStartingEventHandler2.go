package edge

import "unsafe"

type _ICoreWebView2NavigationStartingEventHandler2Vtbl struct {
	IUnknownVtbl
	Invoke ComProc
}

// ICoreWebView2NavigationStartingEventHandler2 接收框架导航开始事件。
//
//	https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2navigationstartingeventhandler
type ICoreWebView2NavigationStartingEventHandler2 struct {
	vtbl *_ICoreWebView2NavigationStartingEventHandler2Vtbl
	impl _ICoreWebView2NavigationStartingEventHandler2Impl
}

func (i *ICoreWebView2NavigationStartingEventHandler2) AddRef() uintptr {
	r, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2NavigationStartingEventHandler2) Release() uintptr {
	r, _, _ := i.vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func _ICoreWebView2NavigationStartingEventHandler2IUnknownQueryInterface(
	this *ICoreWebView2NavigationStartingEventHandler2,
	refiid, object unsafe.Pointer,
) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2NavigationStartingEventHandler2IUnknownAddRef(
	this *ICoreWebView2NavigationStartingEventHandler2,
) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2NavigationStartingEventHandler2IUnknownRelease(
	this *ICoreWebView2NavigationStartingEventHandler2,
) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2NavigationStartingEventHandler2Invoke(
	this *ICoreWebView2NavigationStartingEventHandler2,
	sender *ICoreWebView2,
	args *ICoreWebView2NavigationStartingEventArgs,
) uintptr {
	return this.impl.NavigationStarting2(sender, args)
}

type _ICoreWebView2NavigationStartingEventHandler2Impl interface {
	IUnknownImpl
	NavigationStarting2(sender *ICoreWebView2, args *ICoreWebView2NavigationStartingEventArgs) uintptr
}

var _ICoreWebView2NavigationStartingEventHandler2Fn = _ICoreWebView2NavigationStartingEventHandler2Vtbl{
	IUnknownVtbl{
		NewComProc(_ICoreWebView2NavigationStartingEventHandler2IUnknownQueryInterface),
		NewComProc(_ICoreWebView2NavigationStartingEventHandler2IUnknownAddRef),
		NewComProc(_ICoreWebView2NavigationStartingEventHandler2IUnknownRelease),
	},
	NewComProc(_ICoreWebView2NavigationStartingEventHandler2Invoke),
}

func NewICoreWebView2NavigationStartingEventHandler2(
	impl _ICoreWebView2NavigationStartingEventHandler2Impl,
) *ICoreWebView2NavigationStartingEventHandler2 {
	return &ICoreWebView2NavigationStartingEventHandler2{
		vtbl: &_ICoreWebView2NavigationStartingEventHandler2Fn,
		impl: impl,
	}
}
