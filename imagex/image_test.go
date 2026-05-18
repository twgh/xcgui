package imagex_test

import (
	"image"
	"image/color"
	"testing"

	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/imagex"
	"github.com/twgh/xcgui/tf"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xcc"
)

func TestNewByGoImage(t *testing.T) {
	const w, h = 256, 256
	// 创建一张 RGBA 图片
	img := image.NewRGBA(image.Rect(0, 0, w, h))

	// 左半边：红色
	red := color.RGBA{R: 255, G: 0, B: 0, A: 255}
	// 右半边：蓝色
	blue := color.RGBA{R: 0, G: 0, B: 255, A: 255}

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if x < w/2 {
				img.Set(x, y, red)
			} else {
				img.Set(x, y, blue)
			}
		}
	}

	tf.TFunc(func(a *app.App, w *window.Window) {
		img1 := imagex.NewByGoImage(img)
		w.AddBkImage(xcc.Window_State_Flag_Body_Leave, img1.Handle)
	})
}

func TestModifyByGoImage(t *testing.T) {
	const w, h = 256, 256
	// 创建2张 RGBA 图片
	img1 := image.NewRGBA(image.Rect(0, 0, w, h))
	img2 := image.NewRGBA(image.Rect(0, 0, w, h))

	red := color.RGBA{R: 255, G: 0, B: 0, A: 255}
	blue := color.RGBA{R: 0, G: 0, B: 255, A: 255}

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if x < w/2 {
				img1.Set(x, y, red)  // 左半边：红色
				img2.Set(x, y, blue) // 右半边：蓝色
			} else {
				img1.Set(x, y, blue) // 右半边：蓝色
				img2.Set(x, y, red)  // 左半边：红色
			}
		}
	}

	tf.TFunc(func(a *app.App, w *window.Window) {
		image1 := imagex.NewByGoImage(img1)

		pic := widget.NewShapePicture(30, 100, 256, 256, w.Handle)
		pic.SetImage(image1.Handle)

		widget.NewButton(30, 30, 100, 30, "修改图片", w.Handle).AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
			if w.GetProperty("背景图片") == "2" {
				w.SetProperty("背景图片", "1")
				image1.ModifyDataGoImage(img1)
			} else {
				w.SetProperty("背景图片", "2")
				image1.ModifyDataGoImage(img2)
			}

			pic.Redraw()
			return 0
		})
	})
}
