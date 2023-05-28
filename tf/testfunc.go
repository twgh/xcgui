package tf

import (
	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xcc"
)

// TFunc 测试用程序.
//
//	@Description: 测试时使用的函数.
//	@param f
func TFunc(f func(a *app.App, w *window.Window)) {
	a := app.New(true)
	w := window.New(0, 0, 600, 400, "Test", 0, xcc.Window_Style_Default)
	f(a, w)
	w.Show(true)
	a.Run()
	a.Exit()
}
