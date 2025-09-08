package edge

import (
	"unsafe"
)

type _ICoreWebView2ProfileDeletedEventHandlerVtbl struct {
	IUnknownVtbl
	Invoke ComProc
}

// ICoreWebView2ProfileDeletedEventHandler 接收配置文件删除事件。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2profiledeletedeventhandler
type ICoreWebView2ProfileDeletedEventHandler struct {
	vtbl *_ICoreWebView2ProfileDeletedEventHandlerVtbl
	impl _ICoreWebView2ProfileDeletedEventHandlerImpl
}

func (i *ICoreWebView2ProfileDeletedEventHandler) AddRef() uintptr {
	r, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func (i *ICoreWebView2ProfileDeletedEventHandler) Release() uintptr {
	r, _, _ := i.vtbl.Release.Call(uintptr(unsafe.Pointer(i)))
	return r
}

func _ICoreWebView2ProfileDeletedEventHandlerIUnknownQueryInterface(this *ICoreWebView2ProfileDeletedEventHandler, refiid, object unsafe.Pointer) uintptr {
	return this.impl.QueryInterface(refiid, object)
}

func _ICoreWebView2ProfileDeletedEventHandlerIUnknownAddRef(this *ICoreWebView2ProfileDeletedEventHandler) uintptr {
	return this.impl.AddRef()
}

func _ICoreWebView2ProfileDeletedEventHandlerIUnknownRelease(this *ICoreWebView2ProfileDeletedEventHandler) uintptr {
	return this.impl.Release()
}

func _ICoreWebView2ProfileDeletedEventHandlerInvoke(this *ICoreWebView2ProfileDeletedEventHandler, sender *ICoreWebView2Profile, args *IUnknown) uintptr {
	return this.impl.ProfileDeleted(sender, args)
}

type _ICoreWebView2ProfileDeletedEventHandlerImpl interface {
	IUnknownImpl
	ProfileDeleted(sender *ICoreWebView2Profile, args *IUnknown) uintptr
}

var _ICoreWebView2ProfileDeletedEventHandlerFn *_ICoreWebView2ProfileDeletedEventHandlerVtbl

func NewICoreWebView2ProfileDeletedEventHandler(impl _ICoreWebView2ProfileDeletedEventHandlerImpl) *ICoreWebView2ProfileDeletedEventHandler {
	if _ICoreWebView2ProfileDeletedEventHandlerFn == nil {
		_ICoreWebView2ProfileDeletedEventHandlerFn = &_ICoreWebView2ProfileDeletedEventHandlerVtbl{
			IUnknownVtbl{
				NewComProc(_ICoreWebView2ProfileDeletedEventHandlerIUnknownQueryInterface),
				NewComProc(_ICoreWebView2ProfileDeletedEventHandlerIUnknownAddRef),
				NewComProc(_ICoreWebView2ProfileDeletedEventHandlerIUnknownRelease),
			},
			NewComProc(_ICoreWebView2ProfileDeletedEventHandlerInvoke),
		}
	}
	return &ICoreWebView2ProfileDeletedEventHandler{
		vtbl: _ICoreWebView2ProfileDeletedEventHandlerFn,
		impl: impl,
	}
}
