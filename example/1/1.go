package main

import (
	xc "github.com/twgh/xcgui"
	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/shape"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xcc"
)

var (
	a         *app.App
	win       *window.Window
	btn_Close *widget.Button
	lbl_Title *shape.ShapeText
)

func main() {
	// 1.初始化UI库
	a = app.New("")
	// 2.创建窗口
	win = window.NewWindow(0, 0, 766, 518, "炫彩窗口", 0, xcc.Xc_Window_Style_Default)

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
	lbl_Title = shape.NewShapeText(15, 15, 56, 20, "Title", win.HWindow)
	lbl_Title.SetTextColor(xc.RGB(255, 255, 255), 255)

	// 创建最小化按钮
	widget.NewButton(636, 14, 38, 24, "Min", win.HWindow).SetType(xcc.Button_Type_Min)
	// 创建最大化按钮
	widget.NewButton(675, 14, 38, 24, "Max", win.HWindow).SetType(xcc.Button_Type_Max)
	// 创建结束按钮
	btn_Close = widget.NewButton(714, 14, 38, 24, "Close", win.HWindow)
	btn_Close.SetType(xcc.Button_Type_Close)

	// 3.显示窗口
	win.ShowWindow(xcc.SW_SHOW)
	// 4.运行程序
	a.Run()
	// 5.释放UI库
	a.Exit()
}
