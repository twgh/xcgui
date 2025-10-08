package edge

import _ "embed"

//go:embed webviewloader/sdk/x64/WebView2Loader.dll
var WebView2Loader_Dll []byte

//go:embed webviewloader/sdk/x64/WebView2Helper.dll
var WebView2Helper_Dll []byte
