package edge

import (
	"errors"
	"sync"
	"unsafe"
)

// WvEventHandler 是 WebView 事件处理器.
var WvEventHandler = newWebviewEventHandler()

type webviewEventHandler struct {
	sync.RWMutex
	// 事件信息map
	EventInfoMap map[*WebViewEventImpl]map[string]eventInfo
}

type eventInfo struct {
	Cbs                 []interface{}
	EvnetHandlerPointer unsafe.Pointer
	// EventToken 添加事件时返回的事件令牌, 用于移除事件的参数
	EventToken *EventRegistrationToken
}

func newWebviewEventHandler() *webviewEventHandler {
	return &webviewEventHandler{
		EventInfoMap: make(map[*WebViewEventImpl]map[string]eventInfo),
	}
}

// AddCallBack 添加回调函数, 返回回调函数索引.
//
// impl: WebViewEventImpl 对象.
//
// obj: 额外的对象, 有部分事件需要(不是从 ICoreWebView2, ICoreWebView2Controller 注册的事件). 具体哪些需要, 可以看此方法的源码.
//
// eventType: 事件类型, 如: NavigationCompleted.
//
// cb: 回调函数. 如果为 nil, 则仅注册事件, 不添加回调函数, 返回回调函数索引为-1.
//
// allowAddingMultiple: 是否允许添加多个回调函数, 默认为 false.
//   - 如果为 false, 那么无论你添加多少次, 都只会有一个回调函数, 也就是说会覆盖旧的回调函数.
//   - 如果为 true, 当你添加多次时, 会添加多个回调函数, 执行顺序是先执行最后添加的, 倒序执行.
func (h *webviewEventHandler) AddCallBack(impl *WebViewEventImpl, eventType string, cb interface{}, obj interface{}, allowAddingMultiple ...bool) (int, error) {
	h.Lock()
	defer h.Unlock()

	// 获取对象的事件回调函数map
	eventMap := h.EventInfoMap[impl]
	if eventMap == nil {
		eventMap = make(map[string]eventInfo)
	}

	info, ok := eventMap[eventType]
	// 判断没有注册过这个类型的事件, 就注册
	if !ok {
		info.EventToken = &EventRegistrationToken{}
		switch eventType {
		case "CursorChanged":
			obj2, ok := obj.(*ICoreWebView2CompositionController)
			if !ok {
				return -1, errors.New("obj is not *ICoreWebView2CompositionController")
			}
			eventHandler := NewICoreWebView2CursorChangedEventHandler(impl)
			err := obj2.AddCursorChanged(eventHandler, info.EventToken)
			if err != nil {
				eventHandler.Release()
				return -1, err
			}
			info.EvnetHandlerPointer = unsafe.Pointer(eventHandler)
		case "CustomItemSelected":
			menuItem, ok := obj.(*ICoreWebView2ContextMenuItem)
			if !ok {
				return -1, errors.New("obj is not *ICoreWebView2ContextMenuItem")
			}
			eventHandler := NewICoreWebView2CustomItemSelectedEventHandler(impl)
			err := menuItem.AddCustomItemSelected(eventHandler, info.EventToken)
			if err != nil {
				eventHandler.Release()
				return -1, err
			}
			info.EvnetHandlerPointer = unsafe.Pointer(eventHandler)
		case "DevToolsProtocolEventReceived":
			receiver, ok := obj.(*ICoreWebView2DevToolsProtocolEventReceiver)
			if !ok {
				return -1, errors.New("obj is not *ICoreWebView2DevToolsProtocolEventReceiver")
			}
			eventHandler := NewICoreWebView2DevToolsProtocolEventReceivedEventHandler(impl)
			err := receiver.AddDevToolsProtocolEventReceived(eventHandler, info.EventToken)
			if err != nil {
				eventHandler.Release()
				return -1, err
			}
			info.EvnetHandlerPointer = unsafe.Pointer(eventHandler)
		case "BytesReceivedChanged":
			downloadOperation, ok := obj.(*ICoreWebView2DownloadOperation)
			if !ok {
				return -1, errors.New("obj is not *ICoreWebView2DownloadOperation")
			}
			eventHandler := NewICoreWebView2BytesReceivedChangedEventHandler(impl)
			err := downloadOperation.AddBytesReceivedChanged(eventHandler, info.EventToken)
			if err != nil {
				eventHandler.Release()
				return -1, err
			}
			info.EvnetHandlerPointer = unsafe.Pointer(eventHandler)
		case "FrameNameChanged":
			frame, ok := obj.(*ICoreWebView2Frame)
			if !ok {
				return -1, errors.New("obj is not *ICoreWebView2Frame")
			}
			eventHandler := NewICoreWebView2FrameNameChangedEventHandler(impl)
			err := frame.AddNameChanged(eventHandler, info.EventToken)
			if err != nil {
				eventHandler.Release()
				return -1, err
			}
			info.EvnetHandlerPointer = unsafe.Pointer(eventHandler)
		case "FrameDestroyed":
			frame, ok := obj.(*ICoreWebView2Frame)
			if !ok {
				return -1, errors.New("obj is not *ICoreWebView2Frame")
			}
			eventHandler := NewICoreWebView2FrameDestroyedEventHandler(impl)
			err := frame.AddDestroyed(eventHandler, info.EventToken)
			if err != nil {
				eventHandler.Release()
				return -1, err
			}
			info.EvnetHandlerPointer = unsafe.Pointer(eventHandler)
		case "AcceleratorKeyPressed":
			eventHandler := NewICoreWebView2AcceleratorKeyPressedEventHandler(impl)
			err := impl.Controller.AddAcceleratorKeyPressed(eventHandler, info.EventToken)
			if err != nil {
				eventHandler.Release()
				return -1, err
			}
			info.EvnetHandlerPointer = unsafe.Pointer(eventHandler)
		case "PermissionRequested":
			eventHandler := NewICoreWebView2PermissionRequestedEventHandler(impl)
			err := impl.CoreWebView.AddPermissionRequested(eventHandler, info.EventToken)
			if err != nil {
				eventHandler.Release()
				return -1, err
			}
			info.EvnetHandlerPointer = unsafe.Pointer(eventHandler)
		case "WebMessageReceived":
			eventHandler := NewICoreWebView2WebMessageReceivedEventHandler(impl)
			err := impl.CoreWebView.AddWebMessageReceived(eventHandler, info.EventToken)
			if err != nil {
				eventHandler.Release()
				return -1, err
			}
			info.EvnetHandlerPointer = unsafe.Pointer(eventHandler)
		case "DocumentTitleChanged":
			eventHandler := NewICoreWebView2DocumentTitleChangedEventHandler(impl)
			err := impl.CoreWebView.AddDocumentTitleChanged(eventHandler, info.EventToken)
			if err != nil {
				eventHandler.Release()
				return -1, err
			}
			info.EvnetHandlerPointer = unsafe.Pointer(eventHandler)
		case "WindowCloseRequested":
			eventHandler := NewICoreWebView2WindowCloseRequestedEventHandler(impl)
			err := impl.CoreWebView.AddWindowCloseRequested(eventHandler, info.EventToken)
			if err != nil {
				eventHandler.Release()
				return -1, err
			}
			info.EvnetHandlerPointer = unsafe.Pointer(eventHandler)
		case "SourceChanged":
			eventHandler := NewICoreWebView2SourceChangedEventHandler(impl)
			err := impl.CoreWebView.AddSourceChanged(eventHandler, info.EventToken)
			if err != nil {
				eventHandler.Release()
				return -1, err
			}
			info.EvnetHandlerPointer = unsafe.Pointer(eventHandler)
		case "NewWindowRequested":
			eventHandler := NewICoreWebView2NewWindowRequestedEventHandler(impl)
			err := impl.CoreWebView.AddNewWindowRequested(eventHandler, info.EventToken)
			if err != nil {
				eventHandler.Release()
				return -1, err
			}
			info.EvnetHandlerPointer = unsafe.Pointer(eventHandler)
		case "WebResourceRequested":
			eventHandler := NewICoreWebView2WebResourceRequestedEventHandler(impl)
			err := impl.CoreWebView.AddWebResourceRequested(eventHandler, info.EventToken)
			if err != nil {
				eventHandler.Release()
				return -1, err
			}
			info.EvnetHandlerPointer = unsafe.Pointer(eventHandler)
		case "ContentLoading":
			eventHandler := NewICoreWebView2ContentLoadingEventHandler(impl)
			err := impl.CoreWebView.AddContentLoading(eventHandler, info.EventToken)
			if err != nil {
				eventHandler.Release()
				return -1, err
			}
			info.EvnetHandlerPointer = unsafe.Pointer(eventHandler)
		case "ContainsFullScreenElementChanged":
			eventHandler := NewICoreWebView2ContainsFullScreenElementChangedEventHandler(impl)
			err := impl.CoreWebView.AddContainsFullScreenElementChanged(eventHandler, info.EventToken)
			if err != nil {
				eventHandler.Release()
				return -1, err
			}
			info.EvnetHandlerPointer = unsafe.Pointer(eventHandler)
		case "ProcessFailed":
			eventHandler := NewICoreWebView2ProcessFailedEventHandler(impl)
			err := impl.CoreWebView.AddProcessFailed(eventHandler, info.EventToken)
			if err != nil {
				eventHandler.Release()
				return -1, err
			}
			info.EvnetHandlerPointer = unsafe.Pointer(eventHandler)
		case "HistoryChanged":
			eventHandler := NewICoreWebView2HistoryChangedEventHandler(impl)
			err := impl.CoreWebView.AddHistoryChanged(eventHandler, info.EventToken)
			if err != nil {
				eventHandler.Release()
				return -1, err
			}
			info.EvnetHandlerPointer = unsafe.Pointer(eventHandler)
		case "ScriptDialogOpening":
			eventHandler := NewICoreWebView2ScriptDialogOpeningEventHandler(impl)
			err := impl.CoreWebView.AddScriptDialogOpening(eventHandler, info.EventToken)
			if err != nil {
				eventHandler.Release()
				return -1, err
			}
			info.EvnetHandlerPointer = unsafe.Pointer(eventHandler)
		case "NavigationStarting":
			eventHandler := NewICoreWebView2NavigationStartingEventHandler(impl)
			err := impl.CoreWebView.AddNavigationStarting(eventHandler, info.EventToken)
			if err != nil {
				eventHandler.Release()
				return -1, err
			}
			info.EvnetHandlerPointer = unsafe.Pointer(eventHandler)
		case "NavigationCompleted":
			eventHandler := NewICoreWebView2NavigationCompletedEventHandler(impl)
			err := impl.CoreWebView.AddNavigationCompleted(eventHandler, info.EventToken)
			if err != nil {
				eventHandler.Release()
				return -1, err
			}
			info.EvnetHandlerPointer = unsafe.Pointer(eventHandler)
		case "Frame_NavigationStarting":
			eventHandler := NewICoreWebView2_Frame_NavigationStartingEventHandler(impl)
			err := impl.CoreWebView.AddFrameNavigationStarting(eventHandler, info.EventToken)
			if err != nil {
				eventHandler.Release()
				return -1, err
			}
			info.EvnetHandlerPointer = unsafe.Pointer(eventHandler)
		case "Frame_NavigationCompleted":
			eventHandler := NewICoreWebView2_Frame_NavigationCompletedEventHandler(impl)
			err := impl.CoreWebView.AddFrameNavigationCompleted(eventHandler, info.EventToken)
			if err != nil {
				eventHandler.Release()
				return -1, err
			}
			info.EvnetHandlerPointer = unsafe.Pointer(eventHandler)
		case "RasterizationScaleChanged":
			c3, err := impl.Controller.GetICoreWebView2Controller3()
			if err != nil {
				return -1, err
			}
			defer c3.Release()
			eventHandler := NewICoreWebView2RasterizationScaleChangedEventHandler(impl)
			err = c3.AddRasterizationScaleChanged(eventHandler, info.EventToken)
			if err != nil {
				eventHandler.Release()
				return -1, err
			}
			info.EvnetHandlerPointer = unsafe.Pointer(eventHandler)
		case "ContextMenuRequested":
			w2_11, err := impl.CoreWebView.GetICoreWebView2_11()
			if err != nil {
				return -1, err
			}
			defer w2_11.Release()
			eventHandler := NewICoreWebView2ContextMenuRequestedEventHandler(impl)
			err = w2_11.AddContextMenuRequested(eventHandler, info.EventToken)
			if err != nil {
				eventHandler.Release()
				return -1, err
			}
			info.EvnetHandlerPointer = unsafe.Pointer(eventHandler)
		case "DOMContentLoaded":
			w2_2, err := impl.CoreWebView.GetICoreWebView2_2()
			if err != nil {
				return -1, err
			}
			defer w2_2.Release()
			eventHandler := NewICoreWebView2DOMContentLoadedEventHandler(impl)
			err = w2_2.AddDomContentLoaded(eventHandler, info.EventToken)
			if err != nil {
				eventHandler.Release()
				return -1, err
			}
			info.EvnetHandlerPointer = unsafe.Pointer(eventHandler)
		case "WebResourceResponseReceived":
			w2_2, err := impl.CoreWebView.GetICoreWebView2_2()
			if err != nil {
				return -1, err
			}
			defer w2_2.Release()
			eventHandler := NewICoreWebView2WebResourceResponseReceivedEventHandler(impl)
			err = w2_2.AddWebResourceResponseReceived(eventHandler, info.EventToken)
			if err != nil {
				eventHandler.Release()
				return -1, err
			}
			info.EvnetHandlerPointer = unsafe.Pointer(eventHandler)
		case "DownloadStarting":
			w2_4, err := impl.CoreWebView.GetICoreWebView2_4()
			if err != nil {
				return -1, err
			}
			defer w2_4.Release()
			eventHandler := NewICoreWebView2DownloadStartingEventHandler(impl)
			err = w2_4.AddDownloadStarting(eventHandler, info.EventToken)
			if err != nil {
				eventHandler.Release()
				return -1, err
			}
			info.EvnetHandlerPointer = unsafe.Pointer(eventHandler)
		case "FrameCreated":
			w2_4, err := impl.CoreWebView.GetICoreWebView2_4()
			if err != nil {
				return -1, err
			}
			defer w2_4.Release()
			eventHandler := NewICoreWebView2FrameCreatedEventHandler(impl)
			err = w2_4.AddFrameCreated(eventHandler, info.EventToken)
			if err != nil {
				eventHandler.Release()
				return -1, err
			}
			info.EvnetHandlerPointer = unsafe.Pointer(eventHandler)
		case "ClientCertificateRequested":
			w2_5, err := impl.CoreWebView.GetICoreWebView2_5()
			if err != nil {
				return -1, err
			}
			defer w2_5.Release()
			eventHandler := NewICoreWebView2ClientCertificateRequestedEventHandler(impl)
			err = w2_5.AddClientCertificateRequested(eventHandler, info.EventToken)
			if err != nil {
				eventHandler.Release()
				return -1, err
			}
			info.EvnetHandlerPointer = unsafe.Pointer(eventHandler)
		case "IsMutedChanged":
			w2_8, err := impl.CoreWebView.GetICoreWebView2_8()
			if err != nil {
				return -1, err
			}
			defer w2_8.Release()
			eventHandler := NewICoreWebView2IsMutedChangedEventHandler(impl)
			err = w2_8.AddIsMutedChanged(eventHandler, info.EventToken)
			if err != nil {
				eventHandler.Release()
				return -1, err
			}
			info.EvnetHandlerPointer = unsafe.Pointer(eventHandler)
		case "DocumentPlayingAudioChanged":
			w2_8, err := impl.CoreWebView.GetICoreWebView2_8()
			if err != nil {
				return -1, err
			}
			defer w2_8.Release()
			eventHandler := NewICoreWebView2IsDocumentPlayingAudioChangedEventHandler(impl)
			err = w2_8.AddIsDocumentPlayingAudioChanged(eventHandler, info.EventToken)
			if err != nil {
				eventHandler.Release()
				return -1, err
			}
			info.EvnetHandlerPointer = unsafe.Pointer(eventHandler)
		case "GetCookiesCompleted":
			info.EventToken = nil
			info.EvnetHandlerPointer = unsafe.Pointer(NewICoreWebView2GetCookiesCompletedHandler(impl))
		case "TrySuspendCompleted":
			info.EventToken = nil
			info.EvnetHandlerPointer = unsafe.Pointer(NewICoreWebView2TrySuspendCompletedHandler(impl))
		case "ExecuteScriptCompleted":
			info.EventToken = nil
			info.EvnetHandlerPointer = unsafe.Pointer(NewICoreWebView2ExecuteScriptCompletedHandler(impl))
		case "AddScriptToExecuteOnDocumentCreatedCompleted":
			info.EventToken = nil
			info.EvnetHandlerPointer = unsafe.Pointer(NewICoreWebView2AddScriptToExecuteOnDocumentCreatedCompletedHandler(impl))
		case "CapturePreviewCompleted":
			info.EventToken = nil
			info.EvnetHandlerPointer = unsafe.Pointer(NewICoreWebView2CapturePreviewCompletedHandler(impl))
		case "CallDevToolsProtocolMethodCompleted":
			info.EventToken = nil
			info.EvnetHandlerPointer = unsafe.Pointer(NewICoreWebView2CallDevToolsProtocolMethodCompletedHandler(impl))
		case "WebResourceResponseViewGetContentCompleted":
			info.EventToken = nil
			info.EvnetHandlerPointer = unsafe.Pointer(NewICoreWebView2WebResourceResponseViewGetContentCompletedHandler(impl))
		case "CreateCoreWebView2CompositionControllerCompleted":
			info.EventToken = nil
			info.EvnetHandlerPointer = unsafe.Pointer(NewICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandler(impl))
		}
	}

	isAddingMultiple := false
	if len(allowAddingMultiple) > 0 {
		isAddingMultiple = allowAddingMultiple[0]
	}
	index := 0
	if isAddingMultiple {
		if cb == nil {
			index = -1
		} else {
			info.Cbs = append(info.Cbs, cb)
			index = len(info.Cbs) - 1
		}
	} else {
		if cb == nil {
			info.Cbs = []interface{}{}
			index = -1
		} else {
			info.Cbs = []interface{}{cb}
		}
	}
	eventMap[eventType] = info
	h.EventInfoMap[impl] = eventMap
	return index, nil
}

// GetHandler 获取指定对象指定事件的 EventHandler.
func (h *webviewEventHandler) GetHandler(impl *WebViewEventImpl, eventType string) unsafe.Pointer {
	h.RLock()
	defer h.RUnlock()
	return h.EventInfoMap[impl][eventType].EvnetHandlerPointer
}

// GetEventToken 获取指定对象指定事件的 EventToken.
func (h *webviewEventHandler) GetEventToken(impl *WebViewEventImpl, eventType string) *EventRegistrationToken {
	h.RLock()
	defer h.RUnlock()
	return h.EventInfoMap[impl][eventType].EventToken
}

// ReleaseEventHandler 释放指定对象的指定事件的 EventHandler 且移除该事件的 CallBack.
func (h *webviewEventHandler) ReleaseEventHandler(impl *WebViewEventImpl, eventType string) {
	h.Lock()
	defer h.Unlock()
	eventMap := h.EventInfoMap[impl]
	info := eventMap[eventType]
	if info.EvnetHandlerPointer != nil {
		ComRelease(info.EvnetHandlerPointer)
	}
	delete(eventMap, eventType)
}

// ReleaseAllEventHandler 释放指定对象的所有事件的 EventHandler, 且移除该对象的所有事件以及CallBack.
func (h *webviewEventHandler) ReleaseAllEventHandler(impl *WebViewEventImpl) {
	h.Lock()
	defer h.Unlock()
	eventMap := h.EventInfoMap[impl]
	for _, info := range eventMap {
		if info.EvnetHandlerPointer != nil {
			ComRelease(info.EvnetHandlerPointer)
		}
	}
	delete(h.EventInfoMap, impl)
}

// GetCallBacks 获取指定对象指定事件的回调函数数组.
func (h *webviewEventHandler) GetCallBacks(impl *WebViewEventImpl, eventType string) []interface{} {
	h.RLock()
	defer h.RUnlock()
	return h.EventInfoMap[impl][eventType].Cbs
}

// RemoveAllCallBack 从 map 里移除指定对象的所有事件以及CallBack.
func (h *webviewEventHandler) RemoveAllCallBack(impl *WebViewEventImpl) {
	h.Lock()
	delete(h.EventInfoMap, impl)
	h.Unlock()
}

// RemoveCallBack 从 map 里移除指定对象指定事件的指定索引的 CallBack.
func (h *webviewEventHandler) RemoveCallBack(impl *WebViewEventImpl, eventType string, index int) bool {
	h.Lock()
	defer h.Unlock()
	info := h.EventInfoMap[impl][eventType]
	l := len(info.Cbs)
	if l == 0 {
		return true
	}
	if index >= l {
		return false
	}
	info.Cbs = append(info.Cbs[:index], info.Cbs[index+1:]...)
	h.EventInfoMap[impl][eventType] = info
	return true
}

// RemoveEvent 从 map 里移除指定对象的指定事件的所有CallBack.
func (h *webviewEventHandler) RemoveEvent(impl *WebViewEventImpl, eventType string) {
	h.Lock()
	defer h.Unlock()
	delete(h.EventInfoMap[impl], eventType)
}
