package xc

import _ "embed"

//go:embed dll/x86/xcgui.dll
var DLL []byte

// xcgui.dll 的 CRC32 值
const CRC32 = "858CE8ED"
