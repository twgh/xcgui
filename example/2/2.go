package main

import (
	"fmt"
	"strconv"

	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/shape"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xcc"
)

var (
	a              *app.App
	win            *window.Window
	cbb            *widget.ComboBox
	edit           *widget.Edit
	btn_MessageBox *widget.Button
	btn_Add        *widget.Button
	btn_Sub        *widget.Button
	btn_Close      *widget.Button
	btn_Modal      *widget.Button
	bar            *widget.ProgressBar
	link           *widget.TextLink
	lbl_Title      *shape.ShapeText
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
	// 设置窗口大小
	//win.SetRect(&xc.RECT{Left: 0, Top: 0, Right: 1024, Bottom: 768})

	// 创建标签_窗口标题
	lbl_Title = shape.NewShapeText(15, 15, 56, 20, "窗口标题", win.HWindow)
	lbl_Title.SetTextColor(xc.RGB(255, 0, 0), 255)

	// 创建结束按钮
	btn_Close = widget.NewButton(681, 14, 70, 24, "close", win.HWindow)
	btn_Close.SetType(xcc.Button_Type_Close)

	// 创建按钮_信息框
	btn_MessageBox = widget.NewButton(24, 52, 70, 30, "信息框", win.HWindow)
	// 注册按钮事件
	btn_MessageBox.SetBnClick1(onBtnClick)
	// 创建按钮_模态窗口
	btn_Modal = widget.NewButton(105, 52, 70, 30, "模态窗口", win.HWindow)
	btn_Modal.SetBnClick1(onBtnClick)
	// 给按钮添加背景颜色
	btn_Modal.AddBkFill(xcc.Button_Style_Default, xc.RGB(0, 162, 232), 255)
	// 创建进度条
	bar = widget.NewProgressBar(24, 94, 200, 10, win.HWindow)
	bar.SetSpaceTwo(0, 0)
	bar.SetRange(100)
	bar.SetPos(0)

	// 创建按钮_进度加
	btn_Add = widget.NewButton(238, 84, 70, 30, "+", win.HWindow)
	btn_Add.SetBnClick1(onBtnClick)
	// 创建按钮_进度减
	btn_Sub = widget.NewButton(318, 84, 70, 30, "-", win.HWindow)
	btn_Sub.SetBnClick1(onBtnClick)

	// 创建组合框
	cbb = widget.NewComboBox(24, 126, 100, 30, win.HWindow)
	cbb.CreateAdapter()
	// 组合框加入项
	for i := 1; i <= 5; i++ {
		cbb.AddItemText("第" + strconv.Itoa(i) + "项")
	}
	// 组合框选中项
	cbb.SetSelItem(0)
	// 组合框禁止编辑项
	cbb.EnableEdit(false)
	// 注册组合框被选择事件
	cbb.SetComboBoxSelectEnd1(onCbbSelectEnd)

	// 创建编辑框
	edit = widget.NewEdit(138, 126, 100, 30, win.HWindow)
	edit.SetText("hello")

	// 创建超链接标签
	link = widget.NewTextLink(24, 184, 56, 20, "百度一下", win.HWindow)
	link.RegEventC1(xcc.XE_BNCLICK, onBtnClick)

	// 3.显示窗口
	win.ShowWindow(xcc.SW_SHOW)
	// 4.运行程序
	a.Run()
	// 5.释放UI库
	a.Exit()
}

// 按钮被单击
func onBtnClick(hEle int, pbHandled *bool) int {
	fmt.Println("BtnClick:", hEle)
	switch hEle {
	case btn_MessageBox.HEle:
		win.MessageBox("This Is A MessageBox", "提示", xcc.MessageBox_Flag_Ok)
	case btn_Add.HEle:
		bar.SetPos(bar.GetPos() + 10)
		bar.Redraw(true)
	case link.HEle:
	case btn_Sub.HEle:
		bar.SetPos(bar.GetPos() - 10)
		bar.Redraw(true)
	case btn_Modal.HEle:
		win_Modal := window.NewModalWindow(500, 350, "模态窗口", win.HWND, xcc.Xc_Window_Style_Modal)
		widget.NewButton(50, 50, 70, 30, "close", win_Modal.HWindow).SetType(xcc.Button_Type_Close)
		win_Modal.DoModal()
	}
	return 0
}

// 组合框被选择
func onCbbSelectEnd(hEle int, iItem int, pbHandled *bool) int {
	switch hEle {
	case cbb.HEle:
		edit.SetText(cbb.GetItemText(iItem, 0))
		edit.Redraw(true)
	}
	return 0
}
