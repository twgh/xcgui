package wutil_test

import (
	"fmt"
	"testing"

	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/tf"
	"github.com/twgh/xcgui/wapi/wutil"
	"github.com/twgh/xcgui/window"
)

func TestGetDropFiles(t *testing.T) {
	tf.TFunc(func(a *app.App, w *window.Window) {
		w.EnableDragFiles(true)
		w.Event_DROPFILES(func(hDropInfo uintptr, pbHandled *bool) int {
			fmt.Println(wutil.GetDropFiles(hDropInfo))
			return 0
		})
	})
}

func TestGetTaskbarRect(t *testing.T) {
	rc, err := wutil.GetTaskbarRect()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("矩形:", rc)
	fmt.Println("宽度:", rc.Right-rc.Left)
	fmt.Println("高度:", rc.Bottom-rc.Top)
}
