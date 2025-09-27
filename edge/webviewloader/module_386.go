package webviewloader

import _ "embed"

//go:embed sdk/x86/WebView2Loader.dll
var WebView2Loader_Dll []byte

//go:embed sdk/x86/WebView2Helper.dll
var WebView2Helper_Dll []byte
