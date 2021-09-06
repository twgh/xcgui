# XCGUI

[English](./README-en.md) | 简体中文

DirectUI设计思想: 在窗口内没有子窗口，界面元素都是逻辑上的区域(无HWND句柄,安全,灵活), 所有UI元素都是自主开发(不受系统限制),  更加灵活的实现各种程序界面,满足不同用户的需求.

官方网站：[www.xcgui.com](http://www.xcgui.com "xcgui 官方网站")

# 获取

```go
go get github.com/twgh/xcgui
```

# 项目文档

[项目文档](https://github.com/twgh/xcgui/blob/main/help/%E7%82%AB%E5%BD%A9%E7%95%8C%E9%9D%A2%E5%BA%93-%E5%B8%AE%E5%8A%A9%E6%96%87%E6%A1%A3(v3.0)-(2021-08-04).chm)

# 动态链接库下载

[xcgui.dll](https://github.com/twgh/xcgui/blob/main/help/XCGUI.dll)

程序运行时需要把"XCGUI.dll"放到程序运行目录。

最好是放到C:\Windows\System32目录，这样就不需要把dll放到程序运行目录了。

# 例子

![example](https://github.com/twgh/xcgui/blob/main/example/1/1.jpg)

```go
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
	// 5.退出程序
	a.Exit()
}
```

# 常量

常量都在xcc包里，像这样使用：`xcc.Xc_Window_Style_Default`

# 命令介绍

文件夹外面的xc开头的文件都是原本的api，可以直接使用。

封装好的类都在文件夹里，使用起来更加方便。

