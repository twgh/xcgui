package window_test

import (
	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xcc"
	"math/rand"
	"testing"
	"time"
)

func Test_windowBase_SetSize(t *testing.T) {
	a := app.New(true)
	defer a.Exit()
	rand.Seed(time.Now().Unix())
	w := window.New(0, 0, 500, 500, "SetSize", 0, xcc.Window_Style_Default)

	widget.NewButton(50, 100, 100, 30, "SetSize", w.Handle).Event_BnClick(func(pbHandled *bool) int {
		width := rand.Int31n(400) + 200
		height := rand.Int31n(400) + 200
		w.SetSize(int(width), int(height))
		return 0
	})

	w.Show(true)
	a.Run()
}

func Test_windowBase_SetWidth(t *testing.T) {
	a := app.New(true)
	defer a.Exit()
	rand.Seed(time.Now().Unix())
	w := window.New(0, 0, 500, 500, "SetWidth", 0, xcc.Window_Style_Default)

	widget.NewButton(50, 100, 100, 30, "SetWidth", w.Handle).Event_BnClick(func(pbHandled *bool) int {
		width := rand.Int31n(400) + 200
		w.SetWidth(int(width))
		return 0
	})

	w.Show(true)
	a.Run()
}

func Test_windowBase_SetHeight(t *testing.T) {
	a := app.New(true)
	defer a.Exit()
	rand.Seed(time.Now().Unix())
	w := window.New(0, 0, 500, 500, "SetHeight", 0, xcc.Window_Style_Default)

	widget.NewButton(50, 100, 100, 30, "SetHeight", w.Handle).Event_BnClick(func(pbHandled *bool) int {
		height := rand.Int31n(400) + 200
		w.SetHeight(int(height))
		return 0
	})

	w.Show(true)
	a.Run()
}
