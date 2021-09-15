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
	// 一个事件可以注册多个处理函数，执行顺序为先执行最后注册的函数，最后执行第一个注册的函数.
	// 当你想拦截当前事件或不想向后传递，只需要将参数(*pbHnadled=true)即可.
	btn.Event_BnClick(event1)
	btn.Event_BnClick(event2)
	btn.Event_BnClick(event3)

	// 3.显示窗口
	win.ShowWindow(xcc.SW_SHOW)
	// 4.运行程序
	a.Run()
	// 5.释放UI库
	a.Exit()
}

func event1(pbHandled *bool) int {
	println("event1")
	return 0
}

func event2(pbHandled *bool) int {
	println("event2")
	return 0
}

func event3(pbHandled *bool) int {
	println("event3")
	*pbHandled = true
	return 0
}
