package xc_test

import (
	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/tf"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xc"
	"testing"
)

func TestCallUTWithAny(t *testing.T) {
	tf.TFunc(func(a *app.App, w *window.Window) {
		type Info struct {
			Name string
			Age  int
		}
		data1 := Info{Name: "张三", Age: 18}
		data2 := Info{Name: "李四", Age: 20}

		btn := widget.NewButton(50, 50, 100, 30, "按钮", w.Handle)
		btn.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
			go func() {
				xc.CallUTAny(func(data ...interface{}) int {
					t.Log("可变参数:", data)
					return 0
				}, data1, "男", "北京")

				xc.AutoAny(func(data ...interface{}) int {
					t.Log("可变参数:", data)
					return 0
				}, data2, "女", "上海")

				xc.CallUTAny(func(data ...interface{}) int {
					t.Log("不传参数:", data)
					return 0
				})
			}()
			return 0
		})
	})
}
