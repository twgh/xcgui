package widget

import (
	"math/rand"
	"testing"
	"time"

	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/tf"
	"github.com/twgh/xcgui/window"
)

func TestEdit_AddEvent_EditSet(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	tf.TFunc(func(a *app.App, w *window.Window) {
		edit := NewEdit(20, 40, 200, 200, w.Handle)
		edit.EnableMultiLine(true)
		edit.ShowSBarV(true)
		edit.EnableAutoWrap(true)

		edit.AddEvent_Edit_Changed(func(hEle int, pbHandled *bool) int {
			t.Log("编辑框内容改变:", edit.GetText())
			return 0
		})

		edit.AddEvent_Edit_Pos_Changed(func(hEle int, iPos int32, pbHandled *bool) int {
			t.Log("编辑框光标位置改变:", iPos)
			return 0
		})

		edit.Event_EDIT_SET1(func(hEle int, pbHandled *bool) int {
			t.Log("编辑框设置") // 这个事件是干什么的?
			return 0
		})

		edit.AddEvent_Edit_Color_Change(func(hEle int, color uint32, pbHandled *bool) int {
			t.Log("编辑框颜色改变:", color) // 这个是什么颜色被改变?
			return 0
		})

		btn := NewButton(240, 40, 200, 30, "按钮", w.Handle)
		btn.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
			edit.SetDefaultTextColor(rand.Uint32())
			edit.SetCaretColor(rand.Uint32())
			edit.SetTextColor(rand.Uint32())
			edit.SetSelectBkColor(rand.Uint32())
			edit.SetFocusBorderColor(rand.Uint32())
			edit.AddTextUser("addtext")
			edit.MoveEnd()
			edit.Redraw(false)
			return 0
		})
	})
}
