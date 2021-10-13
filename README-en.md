<h1 align="center">XCGUI</h1>
<p align="center">
    <a href="https://github.com/twgh/xcgui/releases"><img src="https://img.shields.io/badge/version-1.0.3-blue.svg?style=flat-square" alt="version"></a>
    <img src="https://img.shields.io/badge/golang-1.16-blue"/>
    <a href="https://pkg.go.dev/github.com/twgh/xcgui"><img src="https://img.shields.io/badge/go.dev-reference-007d9c ?logo=go&logoColor=white" alt="GoDoc"></a>
    <a href="https://opensource.org/licenses/MIT"><img src="https://img.shields.io/badge/License-MIT-brightgreen.svg?style=flat-square" alt="License"></a>
    <br><br>
	<a href="https://pkg.go.dev/github.com/twgh/xcgui">Project Document</a>&nbsp;&nbsp;
	<a href="https://github.com/twgh/FileStorage/raw/main/xcgui/help/%E7%82%AB%E5%BD%A9%E7%95%8C%E9%9D%A2%E5%BA%93-%E5%B8%AE%E5%8A%A9%E6%96%87%E6%A1%A3(v3.0)-(2021-09-08).chm">Help Document</a>&nbsp;&nbsp;
    <a href="https://github.com/twgh/xcgui-example">Examples</a>
</p>


# Introduction

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
	// Add the file search path, you need to change it to your own path when you run it, or you can use a relative path
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

# Dynamic link library download

| NetDisc      | Link                                                         |
| ------------ | ------------------------------------------------------------ |
| OneDrive     | [download](https://1drv.ms/u/s!ApZP3niad5hpdGuodyU_GvugJ_g?e=yBEKmm) |
| Google Drive | [download](https://drive.google.com/drive/folders/1MuisSsDIr1rjqTkdFIewOgb89SYdf5s6?usp=sharing) |
| GIthub       | [download](https://github.com/twgh/FileStorage/tree/main/xcgui) |

When the program is running, you need to put "XCGUI.dll" in the program running directory.

It is best to put it in the C:\Windows\System32 directory during development, so that there is no need to frequently put the dll in the running directory of different programs.

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

# Const

The constants are all in the xcc package and used like this: `xcc.Xc_Window_Style_Default`

# Command introduction

The functions in the xc package are the original functions in xcgui.dll. There are more than a thousand functions that can be used directly. The encapsulated classes are in other packages.

In some cases, it is more convenient to mix the native functions in the xc package with the encapsulated classes.

# Schedule

These classes are encapsulated based on more than a thousand functions in the xc package. 

| Package Name     | Class Name       | Finish | Doc                                                          |
| ---------------- | ---------------- | ------ | ------------------------------------------------------------ |
| app              | App              | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.0.3/app#App) |
| window           | Window           | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.0.3/window#Window) |
| window           | FrameWindow      | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.0.3/window#FrameWindow) |
| window           | ModalWindow      | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.0.3/window#ModalWindow) |
| widget           | Shape            | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.0.3/widget#Shape) |
| widget           | ShapeEllipse     | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.0.3/widget#ShapeEllipse) |
| widget           | ShapeGif         | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.0.3/widget#ShapeGif) |
| widget           | ShapeGroupBox    | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.0.3/widget#ShapeGroupBox) |
| widget           | ShapeLine        | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.0.3/widget#ShapeLine) |
| widget           | ShapePicture     | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.0.3/widget#ShapePicture) |
| widget           | ShapeRect        | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.0.3/widget#ShapeRect) |
| widget           | ShapeText        | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.0.3/widget#ShapeText) |
| widget           | Table            | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.0.3/widget#Table) |
| widget           | Button           | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.0.3/widget#Button) |
| widget           | ComboBox         | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.0.3/widget#ComboBox) |
| widget           | Edit             | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.0.3/widget#Edit) |
| widget           | Editor           | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.0.3/widget#Editor) |
| widget           | Element          | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.0.3/widget#Element) |
| widget           | List             | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.0.3/widget#List) |
| widget           | ListBox          | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.0.3/widget#ListBox) |
| widget           | Menu             | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.0.3/widget#Menu) |
| widget           | ProgressBar      | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.0.3/widget#ProgressBar) |
| widget           | TextLink         | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.0.3/widget#TextLink) |
| widget           | LayoutEle        | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.0.3/widget#LayoutEle) |
| widget           | LayoutFrame      | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.0.3/widget#LayoutFrame) |
| widget           | ListView         | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.0.3/widget#ListView) |
| widget           | MenuBar          | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.0.3/widget#MenuBar) |
| widget           | Pane             | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.0.3/widget#Pane) |
| widget           | ScrollBar        | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.0.3/widget#ScrollBar) |
| widget           | ScrollView       | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.0.3/widget#ScrollView) |
| widget           | SliderBar        | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.0.3/widget#SliderBar) |
| widget           | TabBar           | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.0.3/widget#TabBar) |
| widget           | ToolBar          | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.0.3/widget#ToolBar) |
| widget           | Tree             | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.0.3/widget#Tree) |
| widget           | DateTime         | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.0.3/widget#DateTime) |
| widget           | MonthCal         | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.0.3/widget#MonthCal) |
| adapter          | AdapterListView  | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.0.3/adapter#AdapterListView) |
| adapter          | AdapterMap       | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.0.3/adapter#AdapterMap) |
| adapter          | AdapterTable     | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.0.3/adapter#AdapterTable) |
| adapter          | AdapterTree      | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.0.3/adapter#AdapterTree) |
| bkmanager        | BkManager        | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.0.3/bkmanager#BkManager) |
| fontx            | FontX            | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.0.3/fontx#FontX) |
| image            | Image            | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.0.3/image#Image) |
| listitemtemplate | ListItemTemplate | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.0.3/listitemtemplate#ListItemTemplate) |
| listitemtemplate | Node             | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.0.3/listitemtemplate#Node) |
| draw             | Draw             | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.0.3/draw#Draw) |
| xc               |                  | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.0.3/xc#section-documentation) |
| xcc              |                  | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.0.3/xcc)   |
| ease             |                  | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.0.3/ease)  |
| res              |                  | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.0.3/res)   |

