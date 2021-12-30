<h1 align="center">XCGUI</h1>
<p align="center">
    <a href="https://github.com/twgh/xcgui/releases"><img src="https://img.shields.io/badge/release-1.3.320-blue.svg?" alt="release"></a>
    <a href="http://www.xcgui.com"><img src="https://img.shields.io/badge/XCGUI-3.3.2-blue.svg?" alt="XCGUI"></a>
   <a href="https://golang.org"> <img src="https://img.shields.io/badge/golang-1.16-brightgreen" alt="golang"></a>
    <a href="https://pkg.go.dev/github.com/twgh/xcgui@v1.3.320"><img src="https://img.shields.io/badge/go.dev-reference-brightgreen" alt="GoDoc"></a>
    <a href="https://opensource.org/licenses/MIT"><img src="https://img.shields.io/badge/License-MIT-brightgreen.svg?" alt="License"></a>
    <br><br>
    <a href="https://github.com/twgh/xcgui-example">Examples</a>&nbsp;&nbsp;
	<a href="https://pkg.go.dev/github.com/twgh/xcgui@v1.3.320">Project Doc</a>&nbsp;&nbsp;
    <a href="http://www.xcgui.com/doc-ui/">Official Doc</a>&nbsp;&nbsp;
	<a href="http://www.xcgui.com">Official Website</a>
</p>










## Introduction

English | [简体中文](./README.md)

DirectUI design idea: there is no sub-window in the window, the interface elements are logical areas (no HWND handle, security, flexibility), all UI elements are developed independently (not limited by the system), more flexible implementation of a variety of program interfaces to meet the needs of different users.

Official website：[www.xcgui.com](http://www.xcgui.com "xcgui Official website")

Resource Mall：[mall.xcgui.com](http://mall.xcgui.com "xcgui Resource Mall")

## Get

```go
go get github.com/twgh/xcgui
```

## Visualization UI Designer

Using the UI designer can quickly design the interface and save a lot of code.

[![uidesigner](https://z3.ax1x.com/2021/09/15/4Vmh9S.png)](https://github.com/twgh/xcgui-example/tree/main/uidesigner)

[UI Designer usage example](https://github.com/twgh/xcgui-example/tree/main/uidesigner), Only so much code：

```go
package main

import (
	_ "embed"

	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xcc"
)

//go:embed res/qqmusic.zip
var qqmusic []byte

func main() {
	a := app.New(true)
	// Load resource files from memory zip
	a.LoadResourceZipMem(qqmusic, "resource.res", "")
	// Load layout file from memory zip
	hWindow := a.LoadLayoutZipMem(qqmusic, "main.xml", "", 0, 0)
	// Create window object
	w := window.NewWindowByHandle(hWindow)
	// Adjust the layout
	w.AdjustLayout()
	// Display window
	w.ShowWindow(xcc.SW_SHOW)
	a.Run()
	a.Exit()
}
```

## Dynamic link library download

| NetDisc      | Link                                                         |
| ------------ | ------------------------------------------------------------ |
| OneDrive     | [download](https://1drv.ms/u/s!ApZP3niad5hpdGuodyU_GvugJ_g?e=yBEKmm) |
| Google Drive | [download](https://drive.google.com/drive/folders/1MuisSsDIr1rjqTkdFIewOgb89SYdf5s6?usp=sharing) |
| Github       | [download](https://github.com/twgh/FileStorage/tree/main/xcgui) |

When the program is running, you need to put `XCGUI.dll`in the program running directory.

It is best to put it in the C:\Windows\System32 directory during development, so that there is no need to frequently put the dll in the running directory of different programs.

## Simple window(Pure code)

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

## Const

The constants are all in the xcc package and used like this: `xcc.Window_Style_Default`

## Command introduction

The xc package contains all the APIs in xcgui.dll. There are more than a thousand functions that can be used directly. The encapsulated classes are in other packages.

In some cases, it is more convenient to mix the native functions in the xc package with the encapsulated classes.

All the structures of xcgui are also in the xc package.

## Event

All events of Dazzling have been defined, all start with Event, and events ending with 1 are the handles of the elements that will be passed in.

Try not to use anonymous functions for callback functions. Using anonymous functions means that you are creating a new callback every time. Eventually you will encounter an error that the program crashes due to creating too many callbacks.

[![xc-event.png](https://z3.ax1x.com/2021/11/23/opdyh6.png)](https://z3.ax1x.com/2021/11/23/opdyh6.png)

Multiple processing functions can be registered for an event. The execution order is to execute the last registered function first, and finally execute the first registered function. When you want to intercept the current event or don’t want to pass it backward, you only need to parameter *pbHandled=true. 

## Module relationship

[![xcgui模块关系图.png](https://z3.ax1x.com/2021/11/18/ITAXJf.png)](https://z3.ax1x.com/2021/11/18/ITAXJf.png)

## Schedule

These classes are encapsulated based on more than a thousand functions in the xc package. 

| Package Name     | Class Name       | Finish | Doc                                                          |
| ---------------- | ---------------- | ------ | ------------------------------------------------------------ |
| app              | App              | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.320/app#App) |
| window           | Window           | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.320/window#Window) |
| window           | FrameWindow      | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.320/window#FrameWindow) |
| window           | ModalWindow      | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.320/window#ModalWindow) |
| widget           | Shape            | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.320/widget#Shape) |
| widget           | ShapeEllipse     | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.320/widget#ShapeEllipse) |
| widget           | ShapeGif         | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.320/widget#ShapeGif) |
| widget           | ShapeGroupBox    | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.320/widget#ShapeGroupBox) |
| widget           | ShapeLine        | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.320/widget#ShapeLine) |
| widget           | ShapePicture     | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.320/widget#ShapePicture) |
| widget           | ShapeRect        | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.320/widget#ShapeRect) |
| widget           | ShapeText        | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.320/widget#ShapeText) |
| widget           | Table            | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.320/widget#Table) |
| widget           | Button           | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.320/widget#Button) |
| widget           | ComboBox         | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.320/widget#ComboBox) |
| widget           | Edit             | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.320/widget#Edit) |
| widget           | Editor           | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.320/widget#Editor) |
| widget           | Element          | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.320/widget#Element) |
| widget           | List             | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.320/widget#List) |
| widget           | ListBox          | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.320/widget#ListBox) |
| widget           | Menu             | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.320/widget#Menu) |
| widget           | ProgressBar      | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.320/widget#ProgressBar) |
| widget           | TextLink         | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.320/widget#TextLink) |
| widget           | LayoutEle        | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.320/widget#LayoutEle) |
| widget           | LayoutFrame      | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.320/widget#LayoutFrame) |
| widget           | ListView         | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.320/widget#ListView) |
| widget           | MenuBar          | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.320/widget#MenuBar) |
| widget           | Pane             | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.320/widget#Pane) |
| widget           | ScrollBar        | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.320/widget#ScrollBar) |
| widget           | ScrollView       | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.320/widget#ScrollView) |
| widget           | SliderBar        | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.320/widget#SliderBar) |
| widget           | TabBar           | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.320/widget#TabBar) |
| widget           | ToolBar          | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.320/widget#ToolBar) |
| widget           | Tree             | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.320/widget#Tree) |
| widget           | DateTime         | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.320/widget#DateTime) |
| widget           | MonthCal         | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.320/widget#MonthCal) |
| adapter          | AdapterListView  | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.320/adapter#AdapterListView) |
| adapter          | AdapterMap       | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.320/adapter#AdapterMap) |
| adapter          | AdapterTable     | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.320/adapter#AdapterTable) |
| adapter          | AdapterTree      | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.320/adapter#AdapterTree) |
| bkmanager        | BkManager        | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.320/bkmanager#BkManager) |
| font             | Font             | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.320/font#Font) |
| imagex           | Imagex           | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.320/imagex#Image) |
| svg              | Svg              | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.320/svg#Svg) |
| listitemtemplate | ListItemTemplate | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.320/listitemtemplate#ListItemTemplate) |
| listitemtemplate | Node             | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.320/listitemtemplate#Node) |
| drawx            | Draw             | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.320/drawx#Draw) |
| ani              | Anima            | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.320/ani#Anima) |
| ani              | AnimaGroup       | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.320/ani#AnimaGroup) |
| ani              | AnimaItem        | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.320/ani#AnimaItem) |
| ani              | AnimaRotate      | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.320/ani#AnimaRotate) |
| ani              | AnimaScale       | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.320/ani#AnimaScale) |
| xc               |                  | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.320/xc#section-documentation) |
| xcc              |                  | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.320/xcc) |
| ease             |                  | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.320/ease) |
| res              |                  | √      | [Doc](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.320/res) |

