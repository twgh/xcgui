<h1 align="center">XCGUI</h1>
<p align="center">
    <a href="https://github.com/twgh/xcgui/releases"><img src="https://img.shields.io/badge/version-1.3.12-blue.svg?style=flat-square" alt="version"></a>
    <img src="https://img.shields.io/badge/golang-1.16-blue"/>
    <a href="https://pkg.go.dev/github.com/twgh/xcgui"><img src="https://img.shields.io/badge/go.dev-reference-007d9c ?logo=go&logoColor=white" alt="GoDoc"></a>
    <a href="https://opensource.org/licenses/MIT"><img src="https://img.shields.io/badge/License-MIT-brightgreen.svg?style=flat-square" alt="License"></a>
    <br><br>
	<a href="https://pkg.go.dev/github.com/twgh/xcgui">Project Doc</a>&nbsp;&nbsp;
    <a href="http://www.xcgui.com/doc-ui/">Official Doc</a>&nbsp;&nbsp;
	<a href="https://github.com/twgh/FileStorage/blob/main/xcgui/help/%E7%82%AB%E5%BD%A9%E7%95%8C%E9%9D%A2%E5%BA%93-%E5%B8%AE%E5%8A%A9%E6%96%87%E6%A1%A3(v3.1)-(2021-10-07).chm?raw=true">CHM Doc</a>&nbsp;&nbsp;
    <a href="https://github.com/twgh/xcgui-example">Examples</a>
</p>






# Introduction

English | [简体中文](./README.md)

DirectUI design idea: there is no sub-window in the window, the interface elements are logical areas (no HWND handle, security, flexibility), all UI elements are developed independently (not limited by the system), more flexible implementation of a variety of program interfaces to meet the needs of different users.

Official website：[www.xcgui.com](http://www.xcgui.com "xcgui 官方网站")

# Visualization UI Designer

[UI Designer usage example](https://github.com/twgh/xcgui-example/tree/main/uidesigner), Using the UI designer can quickly design the interface and save a lot of code.

[![uidesigner](https://z3.ax1x.com/2021/09/15/4Vmh9S.png)](https://github.com/twgh/xcgui-example/tree/main/uidesigner)

Only so much code：

```go
package main

import (
	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xcc"
)

func main() {
	a := app.New(true)
	// Add the file search path. 
    // Please use go run to run the program, if you use go build to run, then please change this to `res`
	a.AddFileSearchPath(`uidesigner\res`)
	// Load resource files from zip
	a.LoadResourceZip(`qqmusic.zip`, "resource.res", "")
	// Load layout file from zip
	hWindow := a.LoadLayoutZip(`qqmusic.zip`, "main.xml", "", 0)
	if hWindow == 0 {
		panic("LoadLayoutZip Error")
	}

	// Create window object
	w := window.NewWindowByHandle(hWindow)
	// adjust the layout
	w.AdjustLayout()
	// Display window
	w.ShowWindow(xcc.SW_SHOW)
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
| Github       | [download](https://github.com/twgh/FileStorage/tree/main/xcgui) |

When the program is running, you need to put `XCGUI.dll`in the program running directory.

It is best to put it in the C:\Windows\System32 directory during development, so that there is no need to frequently put the dll in the running directory of different programs.

# Simple window(Pure code)

[![SimpleWindow](https://i.niupic.com/images/2021/11/01/9FiK.bmp)](https://github.com/twgh/xcgui-example/blob/main/SimpleWindow)

```go
package main

import (
	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xcc"
)

func main() {
	// 1.Initialize XCGUI
	a := app.New(true)
	// 2.Create window
	w := window.NewWindow(0, 0, 430, 300, "", 0, xcc.Window_Style_Simple|xcc.Window_Style_Btn_Close)
	// Set the window border size
	w.SetBorderSize(0, 30, 0, 0)
	// Set window transparency type
	w.SetTransparentType(xcc.Window_Transparent_Shadow)
	// Set window shadow
	w.SetShadowInfo(8, 254, 10, false, 0)
	// 3.Display window
	w.ShowWindow(xcc.SW_SHOW)
	// 4.Run the program
	a.Run()
	// 5.Exit the program
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
| app              | App              | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.12/app#App) |
| window           | Window           | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.12/window#Window) |
| window           | FrameWindow      | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.12/window#FrameWindow) |
| window           | ModalWindow      | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.12/window#ModalWindow) |
| widget           | Shape            | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.12/widget#Shape) |
| widget           | ShapeEllipse     | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.12/widget#ShapeEllipse) |
| widget           | ShapeGif         | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.12/widget#ShapeGif) |
| widget           | ShapeGroupBox    | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.12/widget#ShapeGroupBox) |
| widget           | ShapeLine        | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.12/widget#ShapeLine) |
| widget           | ShapePicture     | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.12/widget#ShapePicture) |
| widget           | ShapeRect        | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.12/widget#ShapeRect) |
| widget           | ShapeText        | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.12/widget#ShapeText) |
| widget           | Table            | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.12/widget#Table) |
| widget           | Button           | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.12/widget#Button) |
| widget           | ComboBox         | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.12/widget#ComboBox) |
| widget           | Edit             | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.12/widget#Edit) |
| widget           | Editor           | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.12/widget#Editor) |
| widget           | Element          | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.12/widget#Element) |
| widget           | List             | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.12/widget#List) |
| widget           | ListBox          | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.12/widget#ListBox) |
| widget           | Menu             | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.12/widget#Menu) |
| widget           | ProgressBar      | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.12/widget#ProgressBar) |
| widget           | TextLink         | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.12/widget#TextLink) |
| widget           | LayoutEle        | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.12/widget#LayoutEle) |
| widget           | LayoutFrame      | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.12/widget#LayoutFrame) |
| widget           | ListView         | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.12/widget#ListView) |
| widget           | MenuBar          | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.12/widget#MenuBar) |
| widget           | Pane             | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.12/widget#Pane) |
| widget           | ScrollBar        | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.12/widget#ScrollBar) |
| widget           | ScrollView       | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.12/widget#ScrollView) |
| widget           | SliderBar        | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.12/widget#SliderBar) |
| widget           | TabBar           | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.12/widget#TabBar) |
| widget           | ToolBar          | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.12/widget#ToolBar) |
| widget           | Tree             | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.12/widget#Tree) |
| widget           | DateTime         | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.12/widget#DateTime) |
| widget           | MonthCal         | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.12/widget#MonthCal) |
| adapter          | AdapterListView  | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.12/adapter#AdapterListView) |
| adapter          | AdapterMap       | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.12/adapter#AdapterMap) |
| adapter          | AdapterTable     | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.12/adapter#AdapterTable) |
| adapter          | AdapterTree      | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.12/adapter#AdapterTree) |
| bkmanager        | BkManager        | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.12/bkmanager#BkManager) |
| font             | Font             | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.12/font#Font) |
| imagex           | Imagex           | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.12/imagex#Image) |
| listitemtemplate | ListItemTemplate | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.12/listitemtemplate#ListItemTemplate) |
| listitemtemplate | Node             | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.12/listitemtemplate#Node) |
| drawx            | Draw             | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.12/drawx#Draw) |
| xc               |                  | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.12/xc#section-documentation) |
| xcc              |                  | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.12/xcc)   |
| ease             |                  | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.12/ease)  |
| res              |                  | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.12/res)   |

