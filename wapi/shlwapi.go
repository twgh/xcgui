package wapi

import (
	"syscall"
	"unsafe"
)

var (
	// Library.
	Shlwapi = syscall.NewLazyDLL("shlwapi")

	// Functions.
	procSHCreateMemStream = Shlwapi.NewProc("SHCreateMemStream")
)

// SHCreateMemStream 创建内存流.
//
// initData: 此缓冲区的内容用于设置内存流的初始内容。 如果此参数为 NULL，则返回的内存流没有任何初始内容。
//
// https://learn.microsoft.com/zh-cn/windows/win32/api/shlwapi/nf-shlwapi-shcreatememstream
func SHCreateMemStream(initData []byte) (uintptr, error) {
	var pInit uintptr
	var cbInit uint32
	if len(initData) > 0 {
		cbInit = uint32(len(initData))
		pInit = uintptr(unsafe.Pointer(&initData[0]))
	}
	ret, _, err := procSHCreateMemStream.Call(
		pInit,
		uintptr(cbInit),
	)
	if ret == 0 {
		return 0, err
	}
	return ret, nil
}
