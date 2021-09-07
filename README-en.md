# XCGUI

English | [简体中文](./README.md)

DirectUI design idea: there is no sub-window in the window, the interface elements are logical areas (no HWND handle, security, flexibility), all UI elements are developed independently (not limited by the system), more flexible implementation of a variety of program interfaces to meet the needs of different users.

Official website：[www.xcgui.com](http://www.xcgui.com "xcgui 官方网站")

# Get

```go
go get github.com/twgh/xcgui
```

# Documentation

[Project Documentation](https://github.com/twgh/xcgui/blob/main/help/%E7%82%AB%E5%BD%A9%E7%95%8C%E9%9D%A2%E5%BA%93-%E5%B8%AE%E5%8A%A9%E6%96%87%E6%A1%A3(v3.0)-(2021-08-04).chm)

# Dynamic link library download

[xcgui.dll](https://github.com/twgh/xcgui/blob/main/help/XCGUI.dll)

When the program is running, you need to put "XCGUI.dll" in the program running directory.

It is best to put it in the C:\Windows\System32 directory, so that there is no need to put the dll in the program running directory.

# Example

![example](https://github.com/twgh/xcgui/blob/main/example/1/1.jpg)

```go
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
	// 1.Initialize XCGUI
	a := app.New("")
	// 2.Create window
	win := window.NewWindow(0, 0, 466, 300, "XC Window", 0, xcc.Xc_Window_Style_Default)

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
	lbl_Title := shape.NewShapeText(15, 15, 56, 20, "Title", win.HWindow)
	lbl_Title.SetTextColor(xc.RGB(255, 255, 255), 255)

	// Create a minimize button
	btn_Min := widget.NewButton(396, 10, 30, 30, "-", win.HWindow)
	btn_Min.SetTextColor(xc.RGB(255, 255, 255), 255)
	btn_Min.SetType(xcc.Button_Type_Min)
	btn_Min.EnableBkTransparent(true)
	// Create an end button
	btn_Close := widget.NewButton(426, 10, 30, 30, "X", win.HWindow)
	btn_Close.SetTextColor(xc.RGB(255, 255, 255), 255)
	btn_Close.SetType(xcc.Button_Type_Close)
	btn_Close.EnableBkTransparent(true)

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

The functions in the xc package are the original functions in xcgui.dll and can be used directly.

The encapsulated classes are in other folders.

