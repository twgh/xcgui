package main

import (
	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/shape"
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
	lbl_Title := shape.NewShapeText(15, 15, 56, 20, "Title", win.HWindow)
	lbl_Title.SetTextColor(xc.RGB(255, 255, 255), 255)

	// 创建结束按钮
	btn_Close := widget.NewButton(426, 10, 30, 30, "X", win.HWindow)
	btn_Close.SetTextColor(xc.RGB(255, 255, 255), 255)
	btn_Close.SetType(xcc.Button_Type_Close)
	btn_Close.EnableBkTransparent(true)

	// 创建按钮_模态窗口
	btn_Modal := widget.NewButton(30, 50, 75, 30, "ModalWindow", win.HWindow)
	// 给按钮绑定事件
	btn_Modal.Event_BnClick(func(pbHandled *bool) int {
		// 创建模态窗口
		win_Modal := window.NewModalWindow(300, 200, "模态窗口", win.HWND, xcc.Xc_Window_Style_Modal|xcc.Xc_Window_Style_Drag_Window)
		// 设置模态窗口边框
		win_Modal.SetBorderSize(1, 25, 1, 1)
		// 在模态窗口上创建关闭按钮
		widget.NewButton(50, 50, 70, 30, "close", win_Modal.HWindow).SetType(xcc.Button_Type_Close)
		// 显示模态窗口
		win_Modal.DoModal()
		return 0
	})

	// 3.显示窗口
	win.ShowWindow(xcc.SW_SHOW)
	// 4.运行程序
	a.Run()
	// 5.释放UI库
	a.Exit()
}
