package app_test

import (
	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/imagex"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xcc"
	"testing"
)

func TestNew(t *testing.T) {
	// 可自定义xcgui.dll的路径, 如无需自定义, 则删掉这句代码.
	/*	err := xc.SetXcguiPath(`C:\Users\Administrator\Desktop\XCGUI.dll`)
		if err != nil {
			panic(err)
		}*/
	a := app.New(true)
	w := window.New(0, 0, 500, 500, "", 0, xcc.Window_Style_Default)
	w.Show(true)
	a.Run()
	a.Exit()
}

func TestApp_EnableAutoDPI(t *testing.T) {
	a := app.New(true)
	a.EnableAutoDPI(true)
	w := window.New(0, 0, 500, 500, "", 0, xcc.Window_Style_Default)
	w.Show(true)
	a.Run()
	a.Exit()
}

func TestApp_SetWindowIcon(t *testing.T) {
	svgIcon := `<svg t="1669088647057" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="5490" width="22" height="22"><path d="M517.12 512.8704m-432.3328 0a432.3328 432.3328 0 1 0 864.6656 0 432.3328 432.3328 0 1 0-864.6656 0Z" fill="#51C5FF" p-id="5491"></path><path d="M292.1472 418.7136c-85.0432 0-160.4096 41.3696-207.104 105.0624 4.5568 182.7328 122.368 337.3056 285.952 396.032 103.2192-33.28 177.92-130.048 177.92-244.3776 0-141.7216-114.944-256.7168-256.768-256.7168z" fill="#7BE0FF" p-id="5492"></path><path d="M800.2048 571.6992l-101.888-58.8288 101.888-58.8288c16.896-9.728 22.6816-31.3344 12.9536-48.2304l-55.296-95.744c-9.728-16.896-31.3344-22.6816-48.2304-12.9536l-101.888 58.8288V238.336c0-19.5072-15.8208-35.328-35.328-35.328H461.824c-19.5072 0-35.328 15.8208-35.328 35.328v117.6064L324.608 297.1136c-16.896-9.728-38.5024-3.9424-48.2304 12.9536l-55.296 95.744c-9.728 16.896-3.9424 38.5024 12.9536 48.2304l101.888 58.8288-101.888 58.8288c-16.896 9.728-22.6816 31.3344-12.9536 48.2304l55.296 95.744c9.728 16.896 31.3344 22.6816 48.2304 12.9536l101.888-58.8288v117.6064c0 19.5072 15.8208 35.328 35.328 35.328h110.592c19.5072 0 35.328-15.8208 35.328-35.328v-117.6064l101.888 58.8288c16.896 9.728 38.5024 3.9424 48.2304-12.9536l55.296-95.744c9.728-16.896 3.9424-38.5024-12.9536-48.2304z" fill="#CAF8FF" p-id="5493"></path><path d="M517.12 512.8704m-234.24 0a234.24 234.24 0 1 0 468.48 0 234.24 234.24 0 1 0-468.48 0Z" fill="#FFFFFF" p-id="5494"></path><path d="M517.12 512.8704m-103.5776 0a103.5776 103.5776 0 1 0 207.1552 0 103.5776 103.5776 0 1 0-207.1552 0Z" fill="#51C5FF" p-id="5495"></path></svg>`

	a := app.New(true)
	a.SetWindowIcon(imagex.NewBySvgStringW(svgIcon).Handle)
	w := window.New(0, 0, 500, 500, "", 0, xcc.Window_Style_Default)
	a.Alert("提示框", "提示框图标随之改变")
	w.MessageBox("信息框", "信息框框图标随之改变", xcc.MessageBox_Flag_Ok, xcc.Window_Style_Modal)
	w.Show(true)
	a.Run()
	a.Exit()
}
