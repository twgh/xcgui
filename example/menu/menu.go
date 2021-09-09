package main

import (
	"fmt"

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
	win := window.NewWindow(0, 0, 366, 200, "炫彩窗口", 0, xcc.Xc_Window_Style_Default)

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
	btn_Close := widget.NewButton(326, 10, 30, 30, "X", win.Handle)
	btn_Close.SetTextColor(xc.RGB(255, 255, 255), 255)
	btn_Close.SetType(xcc.Button_Type_Close)
	btn_Close.EnableBkTransparent(true)

	// 创建菜单按钮
	btn_Menu := widget.NewButton(50, 50, 30, 30, "menu", win.Handle)
	selected := true // 控制item3是否选中
	btn_Menu.Event_BnClick(func(pbHandled *bool) int {
		// 创建菜单
		menu := widget.NewMenu()
		// 一级菜单
		menu.AddItem(201, "item1", 0, xcc.Menu_Item_Flag_Normal)
		menu.AddItem(202, "item2", 0, xcc.Menu_Item_Flag_Normal)
		if selected {
			menu.AddItem(203, "item3", 0, xcc.Menu_Item_Flag_Check)
		} else {
			menu.AddItem(203, "item3", 0, xcc.Menu_Item_Flag_Normal)
		}

		menu.AddItem(204, "", 0, xcc.Menu_Item_Flag_Separator)
		menu.AddItem(205, "Disable", 0, xcc.Menu_Item_Flag_Disable)
		// 二级菜单
		menu.AddItem(206, "item1", 201, xcc.Menu_Item_Flag_Normal)
		menu.AddItem(207, "item2", 201, xcc.Menu_Item_Flag_Normal)

		// 获取按钮坐标
		var r xc.RECT
		btn_Menu.GetRect(&r)
		// 转换到屏幕坐标
		pt := xc.POINT{X: r.Left, Y: r.Bottom}
		xc.Client2Screen(win.Handle, &pt)
		// 弹出菜单
		menu.Popup(win.Handle, int(pt.X), int(pt.Y), btn_Menu.Handle, xcc.Menu_Popup_Position_Left_Top)
		return 0
	})

	// 注册菜单被选择事件
	btn_Menu.Event_MENU_SELECT(func(nID int, pbHandled *bool) int {
		fmt.Println(nID)
		if nID == 203 {
			selected = !selected
		}
		return 0
	})
	// 注册菜单弹出事件
	btn_Menu.Event_MENU_POPUP(func(HMENUX int, pbHandled *bool) int {
		fmt.Println("弹出菜单")
		return 0
	})
	// 注册菜单退出事件
	btn_Menu.Event_MENU_EXIT(func(pbHandled *bool) int {
		fmt.Println("菜单退出")
		return 0
	})

	// 3.显示窗口
	win.ShowWindow(xcc.SW_SHOW)
	// 4.运行程序
	a.Run()
	// 5.释放UI库
	a.Exit()
}
