package xc_test

import (
	"testing"

	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

func TestXInitXCGUI(t *testing.T) {
	// 可自定义xcgui.dll的路径, 如无需自定义, 则删掉这句代码.
	/*	err := xc.SetXcguiPath(`C:\Users\Administrator\Desktop\XCGUI.dll`)
		if err != nil {
			panic(err)
		}*/
	xc.LoadXCGUI()
	xc.XInitXCGUI(true)
	hWindow := xc.XWnd_Create(0, 0, 500, 500, "", 0, xcc.Window_Style_Default)
	xc.XWnd_ShowWindow(hWindow, xcc.SW_SHOW)
	xc.XRunXCGUI()
	xc.XExitXCGUI()
}

func TestWriteDll(t *testing.T) {
	err := xc.WriteDll([]byte("0"))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(xc.GetXcguiPath())
}

func TestGetXcgui(t *testing.T) {
	xc.LoadXCGUI()
	dll := xc.GetXcgui()
	t.Log(dll.Name, dll.Handle())
}

func TestGetXcgui1(t *testing.T) {
	t.Log(xc.GetXcguiPath())
}

func TestSetXcguiPath(t *testing.T) {
	err := xc.SetXcguiPath("C:\\xcgui.dll")
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(xc.GetXcguiPath())
}
