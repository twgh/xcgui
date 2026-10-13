package wapi

import (
	"unsafe"

	"github.com/twgh/xcgui/wapi/com"
)

// Deprecated
//
// !这是旧版函数, 请使用 com.NewGUIDPointer
func NewGUIDPointer(guid string) unsafe.Pointer {
	return com.NewGUIDPointer(guid)
}
