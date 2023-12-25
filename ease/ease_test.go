package ease_test

import (
	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/ease"
	"github.com/twgh/xcgui/tf"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
	"testing"
	"time"
)

func TestEx(t *testing.T) {
	tf.TFunc(func(a *app.App, w *window.Window) {
		time.AfterFunc(time.Millisecond*3, func() {
			// 获取窗口坐标
			var rect xc.RECT
			w.GetRect(&rect)

			for i := 0; i < 30; i++ {
				v := ease.Ex(float32(i)/30.0, xcc.Ease_Flag_Back|xcc.Ease_Flag_Out)
				y := int32(v * float32(rect.Top))
				w.SetPosition(rect.Left, y)
				time.Sleep(time.Millisecond * 10)
			}
		})
	})
}
