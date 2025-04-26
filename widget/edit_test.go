package widget

import (
	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/tf"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xc"
	"math/rand"
	"testing"
	"time"
)

func TestEdit_AddEvent_EditSet(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	tf.TFunc(func(a *app.App, w *window.Window) {
		edit := NewEdit(20, 40, 200, 200, w.Handle)
		edit.EnableMultiLine(true)
		edit.ShowSBarV(true)
		edit.EnableAutoWrap(true)

		edit.AddEvent_Edit_Changed(func(hEle int, pbHandled *bool) int {
			t.Log("编辑框内容改变:", edit.GetTextEx())
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

		edit.AddEvent_Edit_Color_Change(func(hEle int, color int, pbHandled *bool) int {
			t.Log("编辑框颜色改变:", color) // 这个是什么颜色被改变?
			return 0
		})

		btn := NewButton(240, 40, 200, 30, "按钮", w.Handle)
		btn.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
			rgbMax := xc.RGB(255, 255, 255)
			edit.SetDefaultTextColor(xc.RGB2RGBA(rand.Intn(rgbMax), 255))
			edit.SetCaretColor(xc.RGB2RGBA(rand.Intn(rgbMax), 255))
			edit.SetTextColor(xc.RGB2RGBA(rand.Intn(rgbMax), 255))
			edit.SetSelectBkColor(xc.RGB2RGBA(rand.Intn(rgbMax), 255))
			edit.SetFocusBorderColor(xc.RGB2RGBA(rand.Intn(rgbMax), 255))
			edit.AddTextUser("addtext")
			edit.MoveEnd()
			edit.Redraw(false)
			return 0
		})
	})
}
