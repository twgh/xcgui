package widget_test

import (
	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/tf"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"testing"
)

func TestElement_SetFocus(t *testing.T) {
	tf.TFunc(func(a *app.App, w *window.Window) {
		layout := widget.NewLayoutEle(50, 50, 150, 50, w.Handle)
		edit := widget.NewEdit(0, 0, 100, 30, layout.Handle)
		widget.NewButton(50, 100, 100, 30, "setfocus", w.Handle).Event_BnClick(func(pbHandled *bool) int {
			edit.SetFocus()
			return 0
		})
	})
}

func TestElement_SetLeft(t *testing.T) {
	tf.TFunc(func(a *app.App, w *window.Window) {
		btn := widget.NewButton(10, 40, 100, 30, "setleft+5", w.Handle)
		btn.Event_BnClick(func(pbHandled *bool) int {
			btn.SetLeft(btn.GetLeft()+5, true)
			return 0
		})
	})
}

func TestElement_SetTop(t *testing.T) {
	tf.TFunc(func(a *app.App, w *window.Window) {
		btn2 := widget.NewButton(10, 80, 100, 30, "settop+5", w.Handle)
		btn2.Event_BnClick(func(pbHandled *bool) int {
			btn2.SetTop(btn2.GetTop()+5, true)
			return 0
		})
	})
}
