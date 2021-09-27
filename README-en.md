# XCGUI

English | [简体中文](./README.md)

DirectUI design idea: there is no sub-window in the window, the interface elements are logical areas (no HWND handle, security, flexibility), all UI elements are developed independently (not limited by the system), more flexible implementation of a variety of program interfaces to meet the needs of different users.

Official website：[www.xcgui.com](http://www.xcgui.com "xcgui 官方网站")

# Visualization UI Designer

Using the UI designer can quickly design the interface and save a lot of code.

[![uidesigner](https://z3.ax1x.com/2021/09/15/4Vmh9S.png)](https://github.com/twgh/xcgui-example/blob/main/uidesigner/uidesigner.png)

Only so much code：

```go
package main

import (
	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xcc"
)

func main() {
	a := app.New("")
	// Add file search path, you need to change to your own path when you run
	a.AddFileSearchPath(`D:\GoProject\src\github.com\twgh\xcgui\example\uidesigner\res`)
	// Load resource files from zip
	a.LoadResourceZip("qqmusic.zip", "resource.res", "")
	// Load layout file from zip
	hWindow := a.LoadLayoutZip("qqmusic.zip", "main.xml", "", 0)
	if hWindow == 0 {
		panic("error")
	}
	// Create window object
	win := window.NewWindowByHandle(hWindow)
	
	// Assign a value to the class handle
	win.SetHandle(hWindow)
	// Adjust the layout
	win.AdjustLayout()
	// Display window
	win.ShowWindow(xcc.SW_SHOW)

	a.Run()
	a.Exit()
}
```

# Get

```go
go get github.com/twgh/xcgui
```

# Documentation

[Documentation](https://pkg.go.dev/github.com/twgh/xcgui)    [Chm Help Document](https://github.com/twgh/xcgui-example/blob/main/help/%E7%82%AB%E5%BD%A9%E7%95%8C%E9%9D%A2%E5%BA%93-%E5%B8%AE%E5%8A%A9%E6%96%87%E6%A1%A3(v3.0)-(2021-09-08).chm)

# Dynamic link library download

[xcgui.dll(x64)](https://github.com/twgh/xcgui-example/blob/main/help/x64/XCGUI.dll)        [xcgui.dll(x86)](https://github.com/twgh/xcgui-example/blob/main/help/x86/XCGUI.dll)

When the program is running, you need to put "XCGUI.dll" in the program running directory.

It is best to put it in the C:\Windows\System32 directory during development, so that there is no need to put the dll in the program running directory.

# Simple window(Pure code)

[![SimpleWindow](https://z3.ax1x.com/2021/09/15/4VnNuj.jpg)](https://github.com/twgh/xcgui-example/blob/main/SimpleWindow/SimpleWindow.jpg)

```go
package main

import (
	"github.com/twgh/xcgui/app"
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
	lbl_Title := widget.NewShapeText(15, 15, 56, 20, "Title", win.Handle)
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

# More examples

[xcgui-example](https://github.com/twgh/xcgui-example)

# Const

The constants are all in the xcc package and used like this: `xcc.Xc_Window_Style_Default`

# Command introduction

The functions in the xc package are the original functions in xcgui.dll. There are more than a thousand functions that can be used directly.

The encapsulated classes are in other packages.

# Schedule

These classes are encapsulated based on more than a thousand functions in the xc package. Of course, you can also choose to use the functions in the xc package directly.

| Package Name     | Class Name       | Finish |
| ---------------- | ---------------- | ------ |
| app              | App              | √      |
| window           | Window           | √      |
| window           | FrameWindow      | √      |
| window           | ModalWindow      | √      |
| shape            | Shape            | √      |
| shape            | ShapeEllipse     | √      |
| shape            | ShapeGif         | √      |
| shape            | ShapeGroupBox    | √      |
| shape            | ShapeLine        | √      |
| shape            | ShapePicture     | √      |
| shape            | ShapeRect        | √      |
| shape            | ShapeText        | √      |
| widget           | Table            | √      |
| widget           | Button           | √      |
| widget           | ComboBox         | √      |
| widget           | Edit             | √      |
| widget           | Editor           | √      |
| widget           | Element          | √      |
| widget           | List             | √      |
| widget           | ListBox          | √      |
| widget           | Menu             | √      |
| widget           | ProgressBar      | √      |
| widget           | TextLink         | √      |
| widget           | LayoutEle        | √      |
| widget           | LayoutFrame      | √      |
| widget           | ListView         | √      |
| widget           | MenuBar          | √      |
| widget           | Pane             | √      |
| widget           | ScrollBar        | √      |
| widget           | ScrollView       | √      |
| widget           | SliderBar        | √      |
| widget           | TabBar           | √      |
| widget           | ToolBar          | √      |
| widget           | Tree             | √      |
| widget           | DateTime         | √      |
| widget           | MonthCal         | √      |
| adapter          | AdapterListView  | √      |
| adapter          | AdapterMap       | √      |
| adapter          | AdapterTable     | √      |
| adapter          | AdapterTree      | √      |
| bkmanager        | BkManager        | √      |
| fontx            | FontX            | √      |
| image            | Image            | √      |
| listitemtemplate | ListItemTemplate | √      |
| listitemtemplate | Node             | √      |
| draw             | Draw             | √      |
| xc               |                  | √      |
| xcc              |                  | √      |
| ease             |                  | √      |
| res              |                  | √      |

