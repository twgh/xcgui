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
