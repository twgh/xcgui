package app_test

import (
	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xcc"
)

func ExampleNew() {
	// 可自定义xcgui.dll的路径, 如无需自定义, 则删掉这句代码.
	/*	err := xc.SetXcguiPath(`C:\Users\Administrator\Desktop\XCGUI.dll`)
		if err != nil {
			panic(err)
		}*/
	a := app.New(true)
	w := window.New(0, 0, 500, 500, "", 0, xcc.Window_Style_Default)
	w.ShowWindow(xcc.SW_SHOW)
	a.Run()
	a.Exit()
}

func ExampleApp_ShowAndRun() {
	a := app.New(true)
	w := window.New(0, 0, 500, 500, "", 0, xcc.Window_Style_Default)
	a.ShowAndRun(w.Handle)
	a.Exit()
}
