package window

import (
	"github.com/twgh/xcgui/xcc"
)

// Deprecated
//
// !已废弃, 请使用 window.New().
func NewWindow(x int, y int, cx int, cy int, pTitle string, hWndParent int, XCStyle xcc.Window_Style_) *Window {
	return New(x, y, cx, cy, pTitle, hWndParent, XCStyle)
}

// Deprecated
//
// !已废弃, 请使用 window.NewEx()
func NewWindowEx(dwExStyle int, dwStyle int, lpClassName string, x int, y int, cx int, cy int, pTitle string, hWndParent int, XCStyle xcc.Window_Style_) *Window {
	return NewEx(dwExStyle, dwStyle, lpClassName, x, y, cx, cy, pTitle, hWndParent, XCStyle)
}

// Deprecated
//
// !已废弃, 请使用 window.NewByHandle()
func NewWindowByHandle(hWindow int) *Window {
	return NewByHandle(hWindow)
}

// Deprecated
//
// !已废弃, 请使用 window.NewByLayout()
func NewWindowByLayout(pFileName string, hParent, hAttachWnd int) *Window {
	return NewByLayout(pFileName, hParent, hAttachWnd)
}

// Deprecated
//
// !已废弃, 请使用 window.NewByLayoutZip()
func NewWindowByLayoutZip(pZipFileName string, pFileName string, pPassword string, hParent, hAttachWnd int) *Window {
	return NewByLayoutZip(pZipFileName, pFileName, pPassword, hParent, hAttachWnd)
}

// Deprecated
//
// !已废弃, 请使用 window.NewByLayoutZipMem()
func NewWindowByLayoutZipMem(data []byte, pFileName string, pPassword string, hParent int, hAttachWnd int) *Window {
	return NewByLayoutZipMem(data, pFileName, pPassword, hParent, hAttachWnd)
}

// Deprecated
//
// !已废弃, 请使用 window.NewByLayoutStringW()
func NewWindowByLayoutStringW(pStringXML string, hParent int, hAttachWnd int) *Window {
	return NewByLayoutStringW(pStringXML, hParent, hAttachWnd)
}

// Deprecated
//
// !已废弃, 请使用 window.NewByName()
func NewWindowByName(name string) *Window {
	return NewByName(name)
}

// Deprecated
//
// !已废弃, 请使用 window.NewByUID()
func NewWindowByUID(nUID int) *Window {
	return NewByUID(nUID)
}

// Deprecated
//
// !已废弃, 请使用 window.NewByUIDName()
func NewWindowByUIDName(name string) *Window {
	return NewByUIDName(name)
}
