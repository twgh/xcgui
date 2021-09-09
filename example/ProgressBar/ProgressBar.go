package main

import (
	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

var (
	bar     *widget.ProgressBar
	btn_Add *widget.Button
	btn_Sub *widget.Button
)

func main() {
	// 1.初始化UI库
	a := app.New("")
	// 2.创建窗口
	win := window.NewWindow(0, 0, 436, 100, "炫彩窗口", 0, xcc.Xc_Window_Style_Default)

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
	btn_Close := widget.NewButton(396, 10, 30, 30, "X", win.Handle)
	btn_Close.SetTextColor(xc.RGB(255, 255, 255), 255)
	btn_Close.SetType(xcc.Button_Type_Close)
	btn_Close.EnableBkTransparent(true)

	// 创建进度条
	bar = widget.NewProgressBar(24, 60, 200, 10, win.Handle)
	// 设置进度条左右两边间隔大小
	bar.SetSpaceTwo(0, 0)
	// 设置进度条最大值
	bar.SetRange(100)
	// 设置进度条进度为0
	bar.SetPos(0)

	// 创建按钮_进度加
	btn_Add = widget.NewButton(238, 50, 70, 30, "+", win.Handle)
	btn_Add.Event_BnClick1(onBtnClick)
	// 创建按钮_进度减
	btn_Sub = widget.NewButton(318, 50, 70, 30, "-", win.Handle)
	btn_Sub.Event_BnClick1(onBtnClick)

	// 3.显示窗口
	win.ShowWindow(xcc.SW_SHOW)
	// 4.运行程序
	a.Run()
	// 5.释放UI库
	a.Exit()
}

// 事件_按钮被单击
func onBtnClick(hEle int, pbHandled *bool) int {
	switch hEle {
	case btn_Add.Handle:
		bar.SetPos(bar.GetPos() + 10)
		bar.Redraw(true)
	case btn_Sub.Handle:
		bar.SetPos(bar.GetPos() - 10)
		bar.Redraw(true)
	}
	return 0
}
