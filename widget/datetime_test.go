package widget

import (
	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/tf"
	"github.com/twgh/xcgui/window"
	"testing"
)

func TestDateTime_AddEvent_DateTimeChange(t *testing.T) {
	tf.TFunc(func(a *app.App, w *window.Window) {
		date := NewDateTime(20, 40, 150, 35, w.Handle)
		date.AddEvent_DateTime_Change(func(hEle int, pbHandled *bool) int {
			y, m, d := date.GetDateEx()
			h, m, s := date.GetTimeEx()
			t.Log("日期时间改变:", y, m, d, h, m, s)
			return 0
		})

		date.AddEvent_DateTime_Popup_MonthCal(func(hEle int, hMonthCalWnd int, hMonthCal int, pbHandled *bool) int {
			t.Logf("日期时间弹出月历, 窗口句柄: %d, 月历句柄:%d\n", hMonthCalWnd, hMonthCal)
			return 0
		})

		date.AddEvent_DateTime_Exit_MonthCal(func(hEle int, hMonthCalWnd int, hMonthCal int, pbHandled *bool) int {
			t.Logf("日期时间退出月历, 窗口句柄: %d, 月历句柄:%d\n", hMonthCalWnd, hMonthCal)
			return 0
		})

		Time := NewDateTime(240, 40, 150, 35, w.Handle)
		Time.SetStyle(1)
	})
}
