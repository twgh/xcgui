package widget_test

import (
	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xcc"
	"testing"
)

func TestElement_SetFocus(t *testing.T) {
	a := app.New(true)
	defer a.Exit()

	w := window.New(0, 0, 500, 500, "", 0, xcc.Window_Style_Default)
	layout := widget.NewLayoutEle(50, 50, 150, 50, w.Handle)
	edit := widget.NewEdit(0, 0, 100, 30, layout.Handle)

	widget.NewButton(50, 100, 100, 30, "setfocus", w.Handle).Event_BnClick(func(pbHandled *bool) int {
		edit.SetFocus()
		return 0
	})

	w.Show(true)
	a.Run()
}
