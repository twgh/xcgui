package webviewloader

import _ "embed"

//go:embed sdk/x64/WebView2Loader.dll
var Dll []byte
