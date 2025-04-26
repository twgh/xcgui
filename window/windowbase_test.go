package window_test

import (
	"fmt"
	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/tf"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"math/rand"
	"testing"
	"time"
)

func Test_windowBase_SetSize(t *testing.T) {
	tf.TFunc(func(a *app.App, w *window.Window) {
		rand.Seed(time.Now().Unix())
		widget.NewButton(50, 100, 100, 30, "SetSize", w.Handle).AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
			width := rand.Int31n(400) + 200
			height := rand.Int31n(400) + 200
			w.SetSize(width, height)
			return 0
		})
	})
}

func Test_windowBase_SetWidth(t *testing.T) {
	tf.TFunc(func(a *app.App, w *window.Window) {
		rand.Seed(time.Now().Unix())
		widget.NewButton(50, 100, 100, 30, "SetWidth", w.Handle).AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
			width := rand.Int31n(400) + 200
			w.SetWidth(width)
			return 0
		})
	})
}

func Test_windowBase_SetHeight(t *testing.T) {
	tf.TFunc(func(a *app.App, w *window.Window) {
		rand.Seed(time.Now().Unix())
		widget.NewButton(50, 100, 100, 30, "SetHeight", w.Handle).AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
			height := rand.Int31n(400) + 200
			w.SetHeight(height)
			return 0
		})
	})
}

func Test_windowBase_SetLeft(t *testing.T) {
	tf.TFunc(func(a *app.App, w *window.Window) {
		widget.NewButton(50, 100, 100, 30, "SetLeft+5", w.Handle).AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
			w.SetLeft(w.GetLeft() + 5)
			return 0
		})
	})
}

func Test_windowBase_SetTopEdge(t *testing.T) {
	tf.TFunc(func(a *app.App, w *window.Window) {
		widget.NewButton(50, 100, 100, 30, "SetTopEdge+5", w.Handle).AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
			w.SetTopEdge(w.GetTop() + 5)
			return 0
		})
	})
}

func Test_windowBase_GetRectDPI(t *testing.T) {
	tf.TFunc(func(a *app.App, w *window.Window) {
		fmt.Printf("系统缩放倍数: %d%%\n", w.GetDPI()*100/96)
		t.Log(w.GetRectEx())
		t.Log(w.GetRectDPI())
	})
}

func Test_windowBase_SetTop(t *testing.T) {
	tf.TFunc(func(a *app.App, w *window.Window) {
		btn := widget.NewButton(50, 100, 100, 30, "置顶", w.Handle)
		btn.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
			if btn.GetText() == "置顶" {
				w.SetTop(true)
				btn.SetText("取消置顶").Redraw(false)
			} else {
				w.SetTop(false)
				btn.SetText("置顶").Redraw(false)
			}
			return 0
		})
	})
}

func Test_windowBase_AddEvent_NCDestroy(t *testing.T) {
	tf.TFunc(func(a *app.App, w *window.Window) {
		w.AddEvent_Close(func(hWindow int, pbHandled *bool) int {
			t.Log("窗口关闭事件")
			return 0
		})

		w.AddEvent_Destroy(func(hWindow int, pbHandled *bool) int {
			t.Log("窗口即将销毁事件")
			return 0
		})

		w.AddEvent_NCDestroy(func(hWindow int, pbHandled *bool) int {
			t.Log("窗口非客户区销毁事件")
			return 0
		})
	})
}
