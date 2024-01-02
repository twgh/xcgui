<h1 align="center">XCGUI</h1>
<p align="center">
    <a href="https://github.com/twgh/xcgui/releases"><img src="https://img.shields.io/badge/release-1.3.382-blue" alt="release"></a>
    <a href="http://www.xcgui.com"><img src="https://img.shields.io/badge/XCGUI-3.3.8-blue" alt="XCGUI"></a>
   <a href="https://golang.org"> <img src="https://img.shields.io/badge/golang-1.16-blue" alt="golang"></a>
    <a href="https://pkg.go.dev/github.com/twgh/xcgui"><img src="https://img.shields.io/badge/go.dev-reference-brightgreen" alt="GoDoc"></a>
    <a href="https://opensource.org/licenses/MIT"><img src="https://img.shields.io/badge/License-MIT-brightgreen" alt="License"></a>
    <br><br>
    <a href="https://github.com/twgh/xcgui-example">Examples</a>&nbsp;&nbsp;
	<a href="https://pkg.go.dev/github.com/twgh/xcgui">Project Doc</a>&nbsp;&nbsp;
    <a href="http://www.xcgui.com/doc-ui/">Official Doc</a>&nbsp;&nbsp;
	<a href="http://mall.xcgui.com">Official Resource</a>
</p>








## Introduction

English | [简体中文](./README.md)

- This library is encapsulated from the colorful interface library, with rich functions (nearly 2000 API interfaces), easy to use, lightweight, highly DIY customization, and supports one-click skinning.
- The colorful interface library is developed by C/C++ language: the software runs efficiently, does not need the support of third-party libraries, and does not depend on MFC, ATL, WINDOWS standard controls, etc.
- DirectUI design idea: there is no sub-window in the window, the interface elements are all logical areas (no HWND handle, safe, flexible), all UI elements are developed independently (not restricted by the system), more flexible to achieve various Program interface to meet the needs of different users.
- Has a free UI designer tool: rapid development tools, what you see is what you get, a highly customizable system (DIY), making UI development easier.
- Support Direct2D, hardware acceleration, can make full use of hardware features to create high-performance, high-quality 2D graphics.
- [wiki](https://github.com/twgh/xcgui/wiki) There is a simple introductory tutorial, you can take a look when you have time.

## Get

```
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
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
)

//go:embed res/qqmusic.zip
var qqmusic []byte

func main() {
	a := app.New(true)
	a.EnableDPI(true)
	a.EnableAutoDPI(true)
	// Load resource files from memory zip
	a.LoadResourceZipMem(qqmusic, "resource.res", "")
	// Load layout file from memory zip, Create window object
	w := window.NewByLayoutZipMem(qqmusic, "main.xml", "", 0, 0)
    
	// SongTitle is the value of the name property set for the song name (shapeText component) in main.xml.
	// Through GetObjectByName, you can get the handle to the component with the name property set in the layout file..
	// Can be simplified to: widget.NewShapeTextByName("songTitle").
	song := widget.NewShapeTextByHandle(a.GetObjectByName("songTitle"))
	println(song.GetText()) // output: 两只老虎爱跳舞
    
	// Adjust the layout
	w.AdjustLayout()
	// Display window
	w.Show(true)
	a.Run()
	a.Exit()
}
```

## Dynamic link library download

When the program is running, you need to put `xcgui.dll` in the program running directory.

It is best to put it in the `C:\Windows\System32` directory during development, so that there is no need to frequently put the dll in the running directory of different programs.

#### （1）Link to download

| 64 bit | [download](https://pkggo-generic.pkg.coding.net/xcgui/file/xcgui.dll?version=latest) |
| ------ | ------------------------------------------------------------ |
| 32 bit | [download](https://pkggo-generic.pkg.coding.net/xcgui/file/xcgui-32.dll?version=latest) |

#### （2）Command line download

64bit

```bash
curl -fL "https://pkggo-generic.pkg.coding.net/xcgui/file/xcgui.dll?version=latest" -o xcgui.dll
```

32bit

```bash
curl -fL "https://pkggo-generic.pkg.coding.net/xcgui/file/xcgui-32.dll?version=latest" -o xcgui.dll
```

#### （3）Download using the getxcgui tool

> Please ensure that `%GOPATH%\bin` is in the environment variable `path`

```bash
go install github.com/twgh/getxcgui@latest
getxcgui
```

If you want to download the dll directly to the `C:\Windows\System32` directory, please use the following command:

```bash
getxcgui -o %windir%\system32\xcgui.dll
```

The source code of this tool is [here](https://github.com/twgh/getxcgui). For more flags, you can click [go in](https://github.com/twgh/getxcgui#flags) to view

#### （4）Network disk download

The network disk also contains `Interface Designer` and `chm help documentation`

| NetDisc      | Link                                                         |
| ------------ | ------------------------------------------------------------ |
| LanzouYun | [download](https://wwi.lanzoup.com/b0cqd6nkb) |
| BaiduPan     | [download](https://pan.baidu.com/s/1rC3unQGaxnRUCMm8z8qzvA?pwd=1111) |

## Simple window(Pure code)

[![SimpleWindow](https://s1.ax1x.com/2022/11/22/z14kAs.jpg)](https://github.com/twgh/xcgui-example/blob/main/SimpleWindow)

```go
package main

import (
	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/imagex"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xcc"
)

func main() {
	// 1.Initialize XCGUI
	a := app.New(true)
	a.EnableDPI(true)
	a.EnableAutoDPI(true)
	// 2.Create window
	w := window.New(0, 0, 430, 300, "xcgui window", 0, xcc.Window_Style_Default|xcc.Window_Style_Drag_Window)
    
	// Set the window border size
	w.SetBorderSize(0, 30, 0, 0)
    // Set window icon
	a.SetWindowIcon(imagex.NewBySvgStringW(svgIcon).Handle)
	// Set window transparency type
	w.SetTransparentType(xcc.Window_Transparent_Shadow)
	// Set window shadow
	w.SetShadowInfo(8, 255, 10, false, 0)
    
    // Create a button
	btn := widget.NewButton(165, 135, 100, 30, "Button", w.Handle)
	// Registration button clicked event
	btn.Event_BnClick(func(pbHandled *bool) int {
		a.MessageBox("tip", btn.GetText(), xcc.MessageBox_Flag_Ok|xcc.MessageBox_Flag_Icon_Info, w.GetHWND(), xcc.Window_Style_Modal)
		return 0
	})
    
	// 3.Display window
	w.Show(true)
	// 4.Run the program
	a.Run()
	// 5.Exit the program
	a.Exit()
}

var svgIcon = `<svg t="1669088647057" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="5490" width="22" height="22"><path d="M517.12 512.8704m-432.3328 0a432.3328 432.3328 0 1 0 864.6656 0 432.3328 432.3328 0 1 0-864.6656 0Z" fill="#51C5FF" p-id="5491"></path><path d="M292.1472 418.7136c-85.0432 0-160.4096 41.3696-207.104 105.0624 4.5568 182.7328 122.368 337.3056 285.952 396.032 103.2192-33.28 177.92-130.048 177.92-244.3776 0-141.7216-114.944-256.7168-256.768-256.7168z" fill="#7BE0FF" p-id="5492"></path><path d="M800.2048 571.6992l-101.888-58.8288 101.888-58.8288c16.896-9.728 22.6816-31.3344 12.9536-48.2304l-55.296-95.744c-9.728-16.896-31.3344-22.6816-48.2304-12.9536l-101.888 58.8288V238.336c0-19.5072-15.8208-35.328-35.328-35.328H461.824c-19.5072 0-35.328 15.8208-35.328 35.328v117.6064L324.608 297.1136c-16.896-9.728-38.5024-3.9424-48.2304 12.9536l-55.296 95.744c-9.728 16.896-3.9424 38.5024 12.9536 48.2304l101.888 58.8288-101.888 58.8288c-16.896 9.728-22.6816 31.3344-12.9536 48.2304l55.296 95.744c9.728 16.896 31.3344 22.6816 48.2304 12.9536l101.888-58.8288v117.6064c0 19.5072 15.8208 35.328 35.328 35.328h110.592c19.5072 0 35.328-15.8208 35.328-35.328v-117.6064l101.888 58.8288c16.896 9.728 38.5024 3.9424 48.2304-12.9536l55.296-95.744c9.728-16.896 3.9424-38.5024-12.9536-48.2304z" fill="#CAF8FF" p-id="5493"></path><path d="M517.12 512.8704m-234.24 0a234.24 234.24 0 1 0 468.48 0 234.24 234.24 0 1 0-468.48 0Z" fill="#FFFFFF" p-id="5494"></path><path d="M517.12 512.8704m-103.5776 0a103.5776 103.5776 0 1 0 207.1552 0 103.5776 103.5776 0 1 0-207.1552 0Z" fill="#51C5FF" p-id="5495"></path></svg>`
```

## Const

The constants are all in the xcc package and used like this: `xcc.Window_Style_Default`

## Command introduction

The xc package contains all the APIs in xcgui.dll. There are more than a thousand functions that can be used directly. The encapsulated classes are in other packages.

In some cases, it is more convenient to mix the native functions in the xc package with the encapsulated classes.

All the structures of xcgui are also in the xc package.

 [Goland](https://www.jetbrains.com/go/?from=xcgui) is recommended for development for the best development experience. The comment format I'm using looks the best in Goland.

## Event

All events of Dazzling have been defined, all start with Event, and events ending with 1 are the handles of the elements that will be passed in.

Try not to use anonymous functions for callback functions. Using anonymous functions means that you are creating a new callback every time. Eventually you will encounter an error that the program crashes due to creating too many callbacks.

[![xc-event.png](https://z3.ax1x.com/2021/11/23/opdyh6.png)](https://z3.ax1x.com/2021/11/23/opdyh6.png)

Multiple processing functions can be registered for an event. The execution order is to execute the last registered function first, and finally execute the first registered function. When you want to intercept the current event or don’t want to pass it backward, you only need to parameter *pbHandled=true. 

## About version numbers

Take `1.3.330` as an example, 1 only means that the library is the official version, 3.33 represents the official 3.3.3 version of XCGUI, the last 0 represents the first version based on the 3.33 package, If there is an update based on 3.33, then it will add up.

## JetBrains Open Source Certificate Support

The `xcgui` project has always been under the [GoLand](https://www.jetbrains.com/go/?from=xcgui) integrated development environment, based on **free JetBrains Open Source license(s)** genuine free license, I would like to express my gratitude here.

[<img src="https://s1.ax1x.com/2022/05/24/XiFI6x.png" alt="jetbrains.png" />](https://www.jetbrains.com/?from=xcgui)

## Schedule

These classes are encapsulated based on more than a thousand functions in the xc package. 

| Package Name | Class Name       | Finish              | Doc                                                          |
| ------------ | ---------------- | ------------------- | ------------------------------------------------------------ |
| app          | App              | √                   | [Doc](https://pkg.go.dev/github.com/twgh/xcgui/app#App)      |
| window       | Window           | √                   | [Doc](https://pkg.go.dev/github.com/twgh/xcgui/window#Window) |
| window       | FrameWindow      | √                   | [Doc](https://pkg.go.dev/github.com/twgh/xcgui/window#FrameWindow) |
| window       | ModalWindow      | √                   | [Doc](https://pkg.go.dev/github.com/twgh/xcgui/window#ModalWindow) |
| widget       | Shape            | √                   | [Doc](https://pkg.go.dev/github.com/twgh/xcgui/widget#Shape) |
| widget       | ShapeEllipse     | √                   | [Doc](https://pkg.go.dev/github.com/twgh/xcgui/widget#ShapeEllipse) |
| widget       | ShapeGif         | √                   | [Doc](https://pkg.go.dev/github.com/twgh/xcgui/widget#ShapeGif) |
| widget       | ShapeGroupBox    | √                   | [Doc](https://pkg.go.dev/github.com/twgh/xcgui/widget#ShapeGroupBox) |
| widget       | ShapeLine        | √                   | [Doc](https://pkg.go.dev/github.com/twgh/xcgui/widget#ShapeLine) |
| widget       | ShapePicture     | √                   | [Doc](https://pkg.go.dev/github.com/twgh/xcgui/widget#ShapePicture) |
| widget       | ShapeRect        | √                   | [Doc](https://pkg.go.dev/github.com/twgh/xcgui/widget#ShapeRect) |
| widget       | ShapeText        | √                   | [Doc](https://pkg.go.dev/github.com/twgh/xcgui/widget#ShapeText) |
| widget       | Table            | √                   | [Doc](https://pkg.go.dev/github.com/twgh/xcgui/widget#Table) |
| widget       | Button           | √                   | [Doc](https://pkg.go.dev/github.com/twgh/xcgui/widget#Button) |
| widget       | ComboBox         | √                   | [Doc](https://pkg.go.dev/github.com/twgh/xcgui/widget#ComboBox) |
| widget       | Edit             | √                   | [Doc](https://pkg.go.dev/github.com/twgh/xcgui/widget#Edit)  |
| widget       | Editor           | √                   | [Doc](https://pkg.go.dev/github.com/twgh/xcgui/widget#Editor) |
| widget       | Element          | √                   | [Doc](https://pkg.go.dev/github.com/twgh/xcgui/widget#Element) |
| widget       | List             | √                   | [Doc](https://pkg.go.dev/github.com/twgh/xcgui/widget#List)  |
| widget       | ListBox          | √                   | [Doc](https://pkg.go.dev/github.com/twgh/xcgui/widget#ListBox) |
| widget       | Menu             | √                   | [Doc](https://pkg.go.dev/github.com/twgh/xcgui/widget#Menu)  |
| widget       | ProgressBar      | √                   | [Doc](https://pkg.go.dev/github.com/twgh/xcgui/widget#ProgressBar) |
| widget       | TextLink         | √                   | [Doc](https://pkg.go.dev/github.com/twgh/xcgui/widget#TextLink) |
| widget       | LayoutEle        | √                   | [Doc](https://pkg.go.dev/github.com/twgh/xcgui/widget#LayoutEle) |
| widget       | LayoutFrame      | √                   | [Doc](https://pkg.go.dev/github.com/twgh/xcgui/widget#LayoutFrame) |
| widget       | ListView         | √                   | [Doc](https://pkg.go.dev/github.com/twgh/xcgui/widget#ListView) |
| widget       | MenuBar          | √                   | [Doc](https://pkg.go.dev/github.com/twgh/xcgui/widget#MenuBar) |
| widget       | Pane             | √                   | [Doc](https://pkg.go.dev/github.com/twgh/xcgui/widget#Pane)  |
| widget       | ScrollBar        | √                   | [Doc](https://pkg.go.dev/github.com/twgh/xcgui/widget#ScrollBar) |
| widget       | ScrollView       | √                   | [Doc](https://pkg.go.dev/github.com/twgh/xcgui/widget#ScrollView) |
| widget       | SliderBar        | √                   | [Doc](https://pkg.go.dev/github.com/twgh/xcgui/widget#SliderBar) |
| widget       | TabBar           | √                   | [Doc](https://pkg.go.dev/github.com/twgh/xcgui/widget#TabBar) |
| widget       | ToolBar          | √                   | [Doc](https://pkg.go.dev/github.com/twgh/xcgui/widget#ToolBar) |
| widget       | Tree             | √                   | [Doc](https://pkg.go.dev/github.com/twgh/xcgui/widget#Tree)  |
| widget       | DateTime         | √                   | [Doc](https://pkg.go.dev/github.com/twgh/xcgui/widget#DateTime) |
| widget       | MonthCal         | √                   | [Doc](https://pkg.go.dev/github.com/twgh/xcgui/widget#MonthCal) |
| adapter      | AdapterListView  | √                   | [Doc](https://pkg.go.dev/github.com/twgh/xcgui/adapter#AdapterListView) |
| adapter      | AdapterMap       | √                   | [Doc](https://pkg.go.dev/github.com/twgh/xcgui/adapter#AdapterMap) |
| adapter      | AdapterTable     | √                   | [Doc](https://pkg.go.dev/github.com/twgh/xcgui/adapter#AdapterTable) |
| adapter      | AdapterTree      | √                   | [Doc](https://pkg.go.dev/github.com/twgh/xcgui/adapter#AdapterTree) |
| bkmanager    | BkManager        | √                   | [Doc](https://pkg.go.dev/github.com/twgh/xcgui/bkmanager#BkManager) |
| bkobj        | BkObj            | √                   | [Doc](https://pkg.go.dev/github.com/twgh/xcgui/bkobj#BkObj)  |
| font         | Font             | √                   | [Doc](https://pkg.go.dev/github.com/twgh/xcgui/font#Font)    |
| imagex       | Imagex           | √                   | [Doc](https://pkg.go.dev/github.com/twgh/xcgui/imagex#Image) |
| svg          | Svg              | √                   | [Doc](https://pkg.go.dev/github.com/twgh/xcgui/svg#Svg)      |
| tmpl         | ListItemTemplate | √                   | [Doc](https://pkg.go.dev/github.com/twgh/xcgui/tmpl#ListItemTemplate) |
| tmpl         | Node             | √                   | [Doc](https://pkg.go.dev/github.com/twgh/xcgui/tmpl#Node)    |
| drawx        | Draw             | √                   | [Doc](https://pkg.go.dev/github.com/twgh/xcgui/drawx#Draw)   |
| ani          | Anima            | √                   | [Doc](https://pkg.go.dev/github.com/twgh/xcgui/ani#Anima)    |
| ani          | AnimaGroup       | √                   | [Doc](https://pkg.go.dev/github.com/twgh/xcgui/ani#AnimaGroup) |
| ani          | AnimaItem        | √                   | [Doc](https://pkg.go.dev/github.com/twgh/xcgui/ani#AnimaItem) |
| ani          | AnimaRotate      | √                   | [Doc](https://pkg.go.dev/github.com/twgh/xcgui/ani#AnimaRotate) |
| ani          | AnimaScale       | √                   | [Doc](https://pkg.go.dev/github.com/twgh/xcgui/ani#AnimaScale) |
| xc           |                  | √                   | [Doc](https://pkg.go.dev/github.com/twgh/xcgui/xc#section-documentation) |
| xcc          |                  | √                   | [Doc](https://pkg.go.dev/github.com/twgh/xcgui/xcc)          |
| ease         |                  | √                   | [Doc](https://pkg.go.dev/github.com/twgh/xcgui/ease)         |
| res          |                  | √                   | [Doc](https://pkg.go.dev/github.com/twgh/xcgui/res)          |
| wapi         |                  | Continually updated | [Doc](https://pkg.go.dev/github.com/twgh/xcgui/wapi)         |
| wapi/wnd     |                  | Continually updated | [Doc](https://pkg.go.dev/github.com/twgh/xcgui/wapi/wnd)     |
| wapi/wutil   |                  | Continually updated | [Doc](https://pkg.go.dev/github.com/twgh/xcgui/wapi/wutil)   |
