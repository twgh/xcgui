package widget

import (
	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/tf"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xc"
	"strconv"
	"testing"
)

func TestComboBox_AddEvent_COMBOBOX_SELECT_END(t *testing.T) {
	tf.TFunc(func(a *app.App, w *window.Window) {
		cbb := NewComboBox(20, 40, 200, 35, w.Handle)
		cbb.CreateAdapter()
		cbb.EnableEdit(false)
		for i := 0; i < 5; i++ {
			cbb.AddItemText("选项" + strconv.Itoa(i))
		}
		cbb.SetSelItem(0)

		cbb.AddEvent_Combobox_Popup_List(func(hEle int, hWindow int, hListBox int, pbHandled *bool) int {
			t.Logf("下拉列表弹出, 窗口句柄: %d, 列表框句柄: %d\n", hWindow, hListBox)
			xc.XListBox_SetItemHeightDefault(hListBox, 30, 30)
			return 0
		})

		cbb.AddEvent_Combobox_Exit_List(func(hEle int, pbHandled *bool) int {
			t.Log("下拉列表退出")
			return 0
		})

		cbb.AddEvent_Combobox_Select(func(hEle int, iItem int32, pbHandled *bool) int {
			t.Log("选择, 第" + xc.Itoa(iItem) + "项")
			return 0
		})

		cbb.AddEvent_ComboBox_Select_End(func(hEle int, iItem int32, pbHandled *bool) int {
			t.Log("选择完成, 第" + xc.Itoa(iItem) + "项")
			return 0
		})
	})
}
