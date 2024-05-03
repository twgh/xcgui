package xc_test

import (
	"fmt"
	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/tf"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xc"
	"testing"
)

func TestSetBnClicks(t *testing.T) {
	tf.TFunc(func(a *app.App, w *window.Window) {
		a.ShowLayoutFrame(true)
		widget.NewButton(100, 50, 100, 30, "窗口中按钮", w.Handle)

		lay1 := widget.NewLayoutEle(50, 80, 400, 300, w.Handle)
		widget.NewButton(0, 0, 100, 30, "布局元素1中按钮", lay1.Handle)

		lay2 := widget.NewLayoutEle(0, 0, 300, 200, lay1.Handle)
		widget.NewButton(0, 0, 100, 30, "布局元素2中按钮", lay2.Handle)

		xc.SetBnClicks(w.Handle, func(hEle int, pbHandled *bool) int {
			fmt.Println("被单击按钮:", xc.XBtn_GetText(hEle))
			return 0
		})
	})
}

func TestARGB(t *testing.T) {
	fmt.Println(xc.ARGB(255, 201, 100, 255))
	fmt.Println(xc.ARGB2(xc.RGB(255, 201, 100), 255))
	fmt.Println(xc.RGB(255, 201, 100))
	fmt.Println(xc.HexRGB2RGB("ffc964"))
	fmt.Println(xc.RGB2ARGB(xc.HexRGB2RGB("ffc964"), 255))
}

func TestHexRGB2RGB(t *testing.T) {
	fmt.Println(xc.HexRGB2RGB("#ffc964"))
	fmt.Println(xc.HexRGB2RGB("ffc964"))
}

func TestHexRGB2ARGB(t *testing.T) {
	fmt.Println(xc.HexRGB2ARGB("#ffc964", 255))
	fmt.Println(xc.HexRGB2ARGB("ffc964", 255))
}
