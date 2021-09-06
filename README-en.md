# XCGUI

English | [简体中文](./README.md)

DirectUI design idea: there is no sub-window in the window, the interface elements are logical areas (no HWND handle, security, flexibility), all UI elements are developed independently (not limited by the system), more flexible implementation of a variety of program interfaces to meet the needs of different users.

Official website：[www.xcgui.com]([www.xcgui.com](http://www.xcgui.com/))

# Get

```go
go get github.com/twgh/xcgui
```

# Documentation

[Project Documentation](https://github.com/twgh/xcgui/blob/main/help/%E7%82%AB%E5%BD%A9%E7%95%8C%E9%9D%A2%E5%BA%93-%E5%B8%AE%E5%8A%A9%E6%96%87%E6%A1%A3(v3.0)-(2021-08-04).chm)

# Dynamic link library download

[xcgui.dll](https://github.com/twgh/xcgui/blob/main/help/XCGUI.dll)

When the program is running, you need to put "XCGUI.dll" in the help folder in the program running directory.It is best to put it in the C:\Windows\System32 directory, so that there is no need to put the dll in the program running directory.

# Example

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
	// 1.Initialize XCGUI
	a = app.New("")
	// 2.Create window
	win = window.NewWindow(0, 0, 766, 518, "炫彩窗口", 0, xcc.Xc_Window_Style_Default)

	// Set the window border size
	win.SetBorderSize(1, 30, 1, 1)
	// Set window transparency type
	win.SetTransparentType(xcc.Window_Transparent_Shadow)
	// Set window shadow
	win.SetShadowInfo(10, 255, 10, false, 0)
	// Window on top
	win.SetTop()
	// Window centered
	win.Center()
	// Create label window title
	lbl_Title = shape.NewShapeText(15, 15, 56, 20, "Title", win.HWindow)
	lbl_Title.SetTextColor(xc.RGB(255, 255, 255), 255)

	// Create a minimize button
	widget.NewButton(636, 14, 38, 24, "Min", win.HWindow).SetType(xcc.Button_Type_Min)
	// Create a maximize button
	widget.NewButton(675, 14, 38, 24, "Max", win.HWindow).SetType(xcc.Button_Type_Max)
	// Create an end button
	btn_Close = widget.NewButton(714, 14, 38, 24, "Close", win.HWindow)
	btn_Close.SetType(xcc.Button_Type_Close)

	// 3.Display window
	win.ShowWindow(xcc.SW_SHOW)
	// 4.Run the program
	a.Run()
	// 5.exit the program
	a.Exit()
}
```

# Const

The constants are all in the xcc package and used like this: `xcc.Xc Window Style_Default`

# Command introduction

The files beginning with xc outside the folder are the original api and can be used directly.

The encapsulated classes are all in the folder, which is more convenient to use.

