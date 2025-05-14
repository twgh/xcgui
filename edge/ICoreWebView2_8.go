package edge

// ICoreWebView2_8 是 ICoreWebView2_7 接口的延续，支持基本身份验证处理。
//
// https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/reference/win32/icorewebview2_8
type ICoreWebView2_8 struct {
	Vtbl *ICoreWebView2_8Vtbl
}

type ICoreWebView2_8Vtbl struct {
	ICoreWebView2_7Vtbl
	AddIsMutedChanged                   ComProc
	RemoveIsMutedChanged                ComProc
	GetIsMuted                          ComProc
	PutIsMuted                          ComProc
	AddIsDocumentPlayingAudioChanged    ComProc
	RemoveIsDocumentPlayingAudioChanged ComProc
	GetIsDocumentPlayingAudio           ComProc
}

/*TODO:
AddIsMutedChanged
RemoveIsMutedChanged
GetIsMuted
PutIsMuted
AddIsDocumentPlayingAudioChanged
RemoveIsDocumentPlayingAudioChanged
GetIsDocumentPlayingAudio
*/
