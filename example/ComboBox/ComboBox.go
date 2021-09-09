package main

import (
	"strconv"

	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

var edit *widget.Edit // 编辑框

func main() {
	// 1.初始化UI库
	a := app.New("")
	// 2.创建窗口
	win := window.NewWindow(0, 0, 466, 300, "XCGUI", 0, xcc.Xc_Window_Style_Default)

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

	// 创建组合框
	cbb := widget.NewComboBox(24, 50, 100, 30, win.Handle)
	cbb.CreateAdapter()
	// 组合框加入项
	for i := 1; i <= 5; i++ {
		cbb.AddItemText("item" + strconv.Itoa(i))
	}
	// 组合框选中项
	cbb.SetSelItem(0)
	// 组合框禁止编辑项
	cbb.EnableEdit(false)
	// 注册组合框被选择事件
	cbb.Event_ComboBox_Select_End(func(iItem int, pbHandled *bool) int {
		edit.SetText(cbb.GetItemText(iItem, 0))
		edit.Redraw(true)
		return 0
	})

	// 创建编辑框
	edit = widget.NewEdit(138, 50, 100, 30, win.Handle)
	edit.SetText("hello")

	// 3.显示窗口
	win.ShowWindow(xcc.SW_SHOW)
	// 4.运行程序
	a.Run()
	// 5.释放UI库
	a.Exit()
}
