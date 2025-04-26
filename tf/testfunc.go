package tf

import (
	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xcc"
	"runtime"
)

// TFunc 测试用程序. 测试时使用的函数.
//
// f: 回调函数.
func TFunc(f func(a *app.App, w *window.Window)) {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	a := app.New(true)
	a.EnableDPI(true)
	a.EnableAutoDPI(true)
	w := window.New(0, 0, 600, 400, "Test", 0, xcc.Window_Style_Default)
	f(a, w)
	w.Show(true)
	a.Run()
	a.Exit()
}
