package widget

import (
	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/tf"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
	"testing"
)

func TestButton_AddEvent_BnClick(t *testing.T) {
	tf.TFunc(func(a *app.App, w *window.Window) {
		btn1 := NewButton(20, 40, 100, 30, "按钮1", w.Handle)
		// 即使使用匿名函数作为事件回调, 循环超过2000次也不会报错: too many callbacks
		for i := 0; i < 2100; i++ {
			btn1.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
				t.Log("触发按钮1事件")
				return 0
			})
		}
		btn1.AddEvent_Destroy_End(func(hEle int, pbHandled *bool) int {
			t.Log("触发按钮1销毁事件")
			return 0
		})

		btn2 := NewButton(20, 100, 100, 30, "按钮2", w.Handle)
		btn2.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
			t.Log("触发按钮2事件1")
			return 0
		}, true)

		btn2.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
			t.Log("触发按钮2事件2")
			*pbHandled = true // 拦截, 不会触发按钮2事件1
			return 0
		}, true)

		btn2.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
			t.Log("触发按钮2事件3")
			return 0
		}, true)

		btn2.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
			t.Log("触发按钮2事件4")
			return 0
		}, true)

		btn3 := NewButton(20, 160, 100, 30, "销毁按钮1", w.Handle)
		btn3.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
			if xc.XC_IsHELE(btn1.Handle) {
				t.Logf("销毁按钮1, 句柄: %d, 它的所有事件会被自动移除\n", btn1.Handle)
				t.Log("销毁前map:", EventHandler.EventInfoMap)
				btn1.Destroy()
				w.Redraw(false)
				t.Log("销毁后map:", EventHandler.EventInfoMap)
			} else {
				t.Log("按钮1已不存在")
			}
			return 0
		}, true)

		btnCheck := NewButton(20, 220, 100, 30, "选择框", w.Handle)
		btnCheck.SetTypeEx(xcc.Button_Type_Check)
		btnCheck.EnableBkTransparent(true)
		btnCheck.AddEvent_Button_Check(func(hEle int, bCheck bool, pbHandled *bool) int {
			t.Logf("选择框是否选中: %v\n", bCheck)
			return 0
		})
	})
}
