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

func TestABGR(t *testing.T) {
	t.Log(xc.ABGR(255, 201, 100, 255))
}

func TestRGBA(t *testing.T) {
	t.Log(xc.RGBA(255, 201, 100, 255))
}

func TestABGR2(t *testing.T) {
	t.Log(xc.ABGR2(xc.BGR(255, 201, 100), 255))
}

func TestRGB(t *testing.T) {
	t.Log(xc.RGB(255, 201, 100))
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

func TestHexRGB2RGB(t *testing.T) {
	t.Log(xc.HexRGB2RGB("#ffc964"))
	t.Log(xc.HexRGB2RGB("ffc964"))
}

func TestHexRGB2ABGR(t *testing.T) {
	t.Log(xc.HexRGB2ABGR("#ffc964", 255))
	t.Log(xc.HexRGB2ABGR("ffc964", 255))
}

func TestHexRGB2BGR(t *testing.T) {
	t.Log(xc.HexRGB2BGR("#ffc964"))
	t.Log(xc.HexRGB2BGR("ffc964"))
}

func TestRGB2ABGR(t *testing.T) {
	t.Log(xc.RGB2ABGR(xc.HexRGB2RGB("ffc964"), 255))
}

func TestRGB2BGR(t *testing.T) {
	t.Log(xc.RGB2BGR(xc.HexRGB2RGB("ffc964")))
}

func TestBGR(t *testing.T) {
	t.Log(xc.BGR(255, 201, 100))
}
