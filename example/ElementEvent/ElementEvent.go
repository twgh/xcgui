package main

import (
	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

func main() {
	// 1.初始化UI库
	a := app.New("")
	// 2.创建窗口
	win := window.NewWindow(0, 0, 466, 300, "炫彩窗口", 0, xcc.Xc_Window_Style_Default)

	// 设置窗口边框大小
	win.SetBorderSize(1, 30, 1, 1)
	// 设置窗口透明类型
	win.SetTransparentType(xcc.Window_Transparent_Shadow)
	// 设置窗口阴影
	win.SetShadowInfo(10, 255, 10, false, 0)
	// 窗口置顶
	win.SetTop()
	// 窗口居中
	win.Center()
	// 创建标签_窗口标题
	lbl_Title := widget.NewShapeText(15, 15, 56, 20, "Title", win.Handle)
	lbl_Title.SetTextColor(xc.RGB(255, 255, 255), 255)

	// 创建结束按钮
	btn_Close := widget.NewButton(426, 10, 30, 30, "X", win.Handle)
	btn_Close.SetTextColor(xc.RGB(255, 255, 255), 255)
	btn_Close.SetType(xcc.Button_Type_Close)
	btn_Close.EnableBkTransparent(true)

	// 创建按钮
	btn := widget.NewButton(50, 50, 70, 30, "button", win.Handle)
	// 按钮被单击事件
	btn.Event_BnClick(func(pbHandled *bool) int {
		println("Event_BnClick")
		return 0
	})

	// 鼠标移动事件
	// btn.Event_MOUSEMOVE(func(nFlags int, pPt *xc.POINT, pbHandled *bool) int {
	// 	println("Event_MOUSEMOVE", nFlags, pPt.X, pPt.Y)
	// 	return 0
	// })

	// 鼠标进入事件
	btn.Event_MOUSESTAY(func(pbHandled *bool) int {
		println("Event_MOUSESTAY")
		return 0
	})

	// 鼠标离开事件
	btn.Event_MOUSELEAVE(func(hEleStay int, pbHandled *bool) int {
		println("Event_MOUSELEAVE", hEleStay)
		return 0
	})

	// 鼠标滚轮滚动事件
	btn.EnableEvent_XE_MOUSEWHEEL(true)
	btn.Event_MOUSEWHEEL(func(nFlags int, pPt *xc.POINT, pbHandled *bool) int {
		println("Event_MOUSEWHEEL", nFlags, pPt.X, pPt.Y)
		return 0
	})

	// 3.显示窗口
	win.ShowWindow(xcc.SW_SHOW)
	// 4.运行程序
	a.Run()
	// 5.释放UI库
	a.Exit()
}
