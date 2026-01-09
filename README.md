<h1 align="center">XCGUI</h1>
<p align="center">
    <a href="https://github.com/twgh/xcgui/releases"><img src="https://img.shields.io/badge/release-1.3.399-blue" alt="release"></a>
    <a href="http://www.xcgui.com"><img src="https://img.shields.io/badge/XCGUI-3.3.9.1-blue" alt="XCGUI"></a>
   <a href="https://golang.org"> <img src="https://img.shields.io/badge/golang-≥1.18-blue" alt="golang"></a>
    <a href="https://pkg.go.dev/github.com/twgh/xcgui"><img src="https://img.shields.io/badge/go.dev-reference-brightgreen" alt="GoDoc"></a>
    <a href="https://raw.githubusercontent.com/twgh/xcgui/refs/heads/main/xcgui%20license.txt"><img src="https://img.shields.io/badge/License-MIT-brightgreen" alt="License"></a>
    <br><br>
    <a href="https://mcn1fno5w69l.feishu.cn/wiki/space/7489022357177139219?ccm_open_type=lark_wiki_spaceLink&open_tab_from=wiki_home">使用指南</a>&nbsp;&nbsp;
    <a href="https://github.com/twgh/xcgui-example">程序示例</a>&nbsp;&nbsp;
    <a href="http://www.xcgui.com/doc-ui/">官方文档</a>&nbsp;&nbsp;
	<a href="https://pkg.go.dev/github.com/twgh/xcgui">项目文档</a>&nbsp;&nbsp;
	<a href="http://mall.xcgui.com">官方资源</a>
</p>


## 介绍

[English](./README-en.md) | 简体中文

- 本库封装自炫彩界面库，功能丰富(近 2000 个 API 接口)，简单易用，轻量级，高度 DIY 自定义，支持一键换肤。
- 炫彩界面库是由 C/C++ 语言开发：软件运行效率高，不需要第三方库的支持，不依赖 MFC，ATL，WINDOWS 标准控件等。
- DirectUI 设计思想：在窗口内没有子窗口，界面元素都是逻辑上的区域(无 HWND 句柄，安全，灵活)，所有 UI 元素都是自主开发(不受系统限制)，更加灵活的实现各种程序界面，满足不同用户的需求。
- 拥有免费的 UI 设计器：快速开发工具，所见即所得，高度自定义系统(DIY)，让 UI 开发变的更加简单。
- 支持 Direct2D，硬件加速，能更大程度的发挥硬件特性，创建高性能，高质量的 2D 图形。
- 支持 WebView2，可使用前端技术栈开发界面。
- [使用指南](https://mcn1fno5w69l.feishu.cn/wiki/space/7489022357177139219?ccm_open_type=lark_wiki_spaceLink&open_tab_from=wiki_home) 里有入门教程和常见问题，可以看一下，少走弯路。
- 其它相关项目：[xcgui-example](https://github.com/twgh/xcgui-example)，[xc-elementui](https://github.com/twgh/xc-elementui)
- 官方 QQ 群 (能回答问题的人多)：[2283812](https://jq.qq.com/?_wv=1027&k=AiXY4uMc)，只聊 Go 语言版的 QQ 群：[793788011](https://jq.qq.com/?_wv=1027&k=bkKgsYYk)

## 获取

```bash
go get -u github.com/twgh/xcgui
```

## 可视化 UI 设计器

使用 UI 设计器可以快速设计界面，节省大量代码。

[![uidesigner](https://z3.ax1x.com/2021/09/15/4Vmh9S.png)](https://github.com/twgh/xcgui-example/tree/main/Basic/uidesigner)

[设计器使用例子](https://github.com/twgh/xcgui-example/tree/main/Basic/uidesigner)，只有这么多代码：

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
    app.Init()
	a := app.New(true)
	a.EnableAutoDPI(true).EnableDPI(true)
	// 从内存 zip 中加载资源文件
	a.LoadResourceZipMem(qqmusic, "resource.res", "")
	// 从内存 zip 中加载布局文件, 创建窗口对象
	w := window.NewByLayoutZipMem(qqmusic, "main.xml", "", 0, 0)
    
	// songTitle 是在 main.xml 中给歌曲名(shapeText 组件)设置的 name 属性的值.
	// 通过 GetObjectByName 可以获取布局文件中设置了 name 属性的组件的句柄.
	// 可简化为: widget.NewShapeTextByName("songTitle").
	song := widget.NewShapeTextByHandle(app.GetObjectByName("songTitle"))
	println(song.GetText()) // 输出: 两只老虎爱跳舞
    
	// 调整布局
	w.AdjustLayout()
	// 显示窗口
	w.Show(true)
	a.Run()
	a.Exit()
}
```

## 相关资源下载

[百度网盘](https://pan.baidu.com/s/1rC3unQGaxnRUCMm8z8qzvA?pwd=1111)，[蓝奏云](https://wwi.lanzoup.com/b0cqd6nkb)

网盘内包含 `界面设计器`、 `chm帮助文档`

## 简单窗口（纯代码）

[![SimpleWindow](https://s1.ax1x.com/2022/11/22/z14kAs.jpg)](https://github.com/twgh/xcgui-example/tree/main/Basic/SimpleWindow)

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
	// 1.初始化UI库
    app.Init()
	a := app.New(true)
	// 启用自适应 DPI
	a.EnableAutoDPI(true).EnableDPI(true)
	// 2.创建窗口
	w := window.New(0, 0, 430, 300, "xcgui window", 0, xcc.Window_Style_Default|xcc.Window_Style_Drag_Window)

	// 设置窗口边框大小
	w.SetBorderSize(0, 30, 0, 0)
	// 设置全局窗口图标
	a.SetWindowIcon(imagex.NewBySvgString(svgIcon).Handle)
	// 设置窗口透明类型
	w.SetTransparentType(xcc.Window_Transparent_Shadow)
	// 设置窗口阴影
	w.SetShadowInfo(8, 255, 10, false, 0)

	// 创建按钮
	btn := widget.NewButton(165, 135, 100, 30, "Button", w.Handle)
	// 添加按钮点击事件
	btn.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
		w.MessageBox("提示", btn.GetText(), xcc.MessageBox_Flag_Ok|xcc.MessageBox_Flag_Icon_Info, xcc.Window_Style_Modal)
		return 0
	})

	// 3.显示窗口
	w.Show(true)
	// 4.运行程序
	a.Run()
	// 5.释放UI库
	a.Exit()
}

const svgIcon = `<svg t="1669088647057" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="5490" width="22" height="22"><path d="M517.12 512.8704m-432.3328 0a432.3328 432.3328 0 1 0 864.6656 0 432.3328 432.3328 0 1 0-864.6656 0Z" fill="#51C5FF" p-id="5491"></path><path d="M292.1472 418.7136c-85.0432 0-160.4096 41.3696-207.104 105.0624 4.5568 182.7328 122.368 337.3056 285.952 396.032 103.2192-33.28 177.92-130.048 177.92-244.3776 0-141.7216-114.944-256.7168-256.768-256.7168z" fill="#7BE0FF" p-id="5492"></path><path d="M800.2048 571.6992l-101.888-58.8288 101.888-58.8288c16.896-9.728 22.6816-31.3344 12.9536-48.2304l-55.296-95.744c-9.728-16.896-31.3344-22.6816-48.2304-12.9536l-101.888 58.8288V238.336c0-19.5072-15.8208-35.328-35.328-35.328H461.824c-19.5072 0-35.328 15.8208-35.328 35.328v117.6064L324.608 297.1136c-16.896-9.728-38.5024-3.9424-48.2304 12.9536l-55.296 95.744c-9.728 16.896-3.9424 38.5024 12.9536 48.2304l101.888 58.8288-101.888 58.8288c-16.896 9.728-22.6816 31.3344-12.9536 48.2304l55.296 95.744c9.728 16.896 31.3344 22.6816 48.2304 12.9536l101.888-58.8288v117.6064c0 19.5072 15.8208 35.328 35.328 35.328h110.592c19.5072 0 35.328-15.8208 35.328-35.328v-117.6064l101.888 58.8288c16.896 9.728 38.5024 3.9424 48.2304-12.9536l55.296-95.744c9.728-16.896 3.9424-38.5024-12.9536-48.2304z" fill="#CAF8FF" p-id="5493"></path><path d="M517.12 512.8704m-234.24 0a234.24 234.24 0 1 0 468.48 0 234.24 234.24 0 1 0-468.48 0Z" fill="#FFFFFF" p-id="5494"></path><path d="M517.12 512.8704m-103.5776 0a103.5776 103.5776 0 1 0 207.1552 0 103.5776 103.5776 0 1 0-207.1552 0Z" fill="#51C5FF" p-id="5495"></path></svg>`
```

## 常量

xcc 包里都是常量，需要用到 xcc 包常量的参数注释都是类似图片这样的，复制 `Window_Transparent_` 到 [xcc包](https://pkg.go.dev/github.com/twgh/xcgui/xcc#pkg-constants) 里搜索即可看到相关常量及注释：

[![注释](https://s4.ax1x.com/2021/12/22/TQvQzt.png)](https://github.com/twgh/xcgui/blob/main/window/windowbase.go#L630)

## 命令介绍

xc 包里包含 xcgui.dll 里所有的 API，有近 2000 个函数，不习惯使用类的可以直接使用原版函数，封装好的类都在其他包里。

在某些情况下，把 xc 包里的原生函数与封装好的类混合起来使用会更加方便。

炫彩所有的结构体也都在 xc 包里。

建议使用 [Goland](https://www.jetbrains.com/go/?from=xcgui) 进行开发，以获得最好的开发体验。

## 事件

炫彩的一个事件类型可以注册多个回调处理函数，执行顺序为先执行最后注册的回调函数，最后执行第一个注册的回调函数，当你想拦截当前事件或不想向后传递，只需要在回调函数中设置参数 `*pbHandled=true` 即可。

炫彩的全部事件都已经在各种类里定义好了，有两种形式，分别以两种格式开头： `AddEvent` 或 `Event` ，一般使用 `AddEvent` 类型的函数即可。

区别是：由于使用的是 `syscall.NewCallBack` 创建的事件回调函数，该方法限制只能创建 2000 个左右的回调函数，超过就会 panic。当使用 `Event` 类型的函数来注册事件且回调函数是匿名函数时，每次都会创建 1 个新的回调函数，如果不加以控制，就可能会超过 2000 个。而 `AddEvent` 类型的函数会复用创建好的回调函数，可以任意使用匿名函数作为事件回调函数，无需担心超过 2000 个的限制。

## 关于版本号

以 `1.3.330` 为例，1 仅代表本库是正式版，3.33 代表是 XCGUI 官方 3.3.3 版本，最后的 0 代表是基于 3.33 封装的第 1 个版本，如果基于 3.33 还有更新，那么会累加。

## JetBrains 开源证书支持

`xcgui` 项目一直以来都是在 [JetBrains](https://www.jetbrains.com/?from=xcgui) 公司旗下的 [GoLand](https://www.jetbrains.com/go/?from=xcgui) 集成开发环境中进行开发，基于 **free JetBrains Open Source license(s)** 正版免费授权，在此表达我的谢意。

[<img src="https://s1.ax1x.com/2022/05/24/XiFI6x.png" alt="jetbrains.png" />](https://www.jetbrains.com/?from=xcgui)

## 模块继承关系

[xcgui模块关系图](http://www.xcgui.com/doc-ui/page_diagram.html)

## 封装进度

和炫彩相关的类都是基于 xc 包里的近 2000 个函数封装的。

| 中文名称            | 包名         | 类名              | 是否封装完毕 | 文档                                                                     |
|-----------------|------------|-----------------| ------------ |------------------------------------------------------------------------|
| 程序（炫彩全局API）     | app        | App             | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/app#App)                 |
| 窗口              | window     | Window          | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/window#Window)           |
| 框架窗口            | window     | FrameWindow     | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/window#FrameWindow)      |
| 模态窗口            | window     | ModalWindow     | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/window#ModalWindow)      |
| 托盘图标            | window     | TrayIcon        | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/window#TrayIcon)         |
| WebView 环境 | edge | Edge | √ | [文档](https://pkg.go.dev/github.com/twgh/xcgui/edge#Edge) |
| WebView        | edge    | WebView         | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/edge#WebView)     |
| 形状对象            | widget     | Shape           | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/widget#Shape)            |
| 圆形形状对象          | widget     | ShapeEllipse    | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/widget#ShapeEllipse)     |
| 形状对象GIF         | widget     | ShapeGif        | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/widget#ShapeGif)         |
| 组框形状对象          | widget     | ShapeGroupBox   | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/widget#ShapeGroupBox)    |
| 直线形状对象          | widget     | ShapeLine       | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/widget#ShapeLine)        |
| 形状对象图片          | widget     | ShapePicture    | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/widget#ShapePicture)     |
| 矩形形状对象          | widget     | ShapeRect       | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/widget#ShapeRect)        |
| 形状对象文本          | widget     | ShapeText       | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/widget#ShapeText)        |
| 表格              | widget     | Table           | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/widget#Table)            |
| 按钮              | widget     | Button          | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/widget#Button)           |
| 下拉组合框           | widget     | ComboBox        | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/widget#ComboBox)         |
| 编辑框(常规, 富文本, 聊天气泡) | widget     | Edit            | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/widget#Edit)             |
| 代码编辑框           | widget     | Editor          | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/widget#Editor)           |
| 基础元素            | widget     | Element         | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/widget#Element)          |
| 列表              | widget     | List            | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/widget#List)             |
| 列表框             | widget     | ListBox         | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/widget#ListBox)          |
| 弹出菜单            | widget     | Menu            | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/widget#Menu)             |
| 进度条             | widget     | ProgressBar     | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/widget#ProgressBar)      |
| 静态文本链接按钮        | widget     | TextLink        | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/widget#TextLink)         |
| 布局元素            | widget     | LayoutEle       | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/widget#LayoutEle)        |
| 布局框架            | widget     | LayoutFrame     | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/widget#LayoutFrame)      |
| 列表视图            | widget     | ListView        | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/widget#ListView)         |
| 菜单条             | widget     | MenuBar         | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/widget#MenuBar)          |
| Pane元素          | widget     | Pane            | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/widget#Pane)             |
| 滚动条             | widget     | ScrollBar       | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/widget#ScrollBar)        |
| 滚动视图            | widget     | ScrollView      | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/widget#ScrollView)       |
| 滑动条             | widget     | SliderBar       | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/widget#SliderBar)        |
| 标签栏             | widget     | TabBar          | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/widget#TabBar)           |
| 工具条             | widget     | ToolBar         | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/widget#ToolBar)          |
| 树               | widget     | Tree            | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/widget#Tree)             |
| 日期时间            | widget     | DateTime        | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/widget#DateTime)         |
| 月历卡片            | widget     | MonthCal        | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/widget#MonthCal)         |
| Gif 播放器 | widget | GifPlayer | √ | [文档](https://pkg.go.dev/github.com/twgh/xcgui/widget#GifPlayer) |
| 数据适配器-列表视       | adapter    | AdapterListView | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/adapter#AdapterListView) |
| 数据适配器-单列Map-列表头(listHeader) | adapter    | AdapterMap      | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/adapter#AdapterMap)      |
| 数据适配器-XList-XListBox | adapter    | AdapterTable    | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/adapter#AdapterTable)    |
| 数据适配器-树         | adapter    | AdapterTree     | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/adapter#AdapterTree)     |
| 背景管理器           | bkmanager  | BkManager       | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/bkmanager#BkManager)     |
| 背景对象            | bkobj      | BkObj           | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/bkobj#BkObj)             |
| 字体              | font       | Font            | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/font#Font)               |
| 图片操作            | imagex     | Image           | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/imagex#Image)            |
| SVG矢量图形         | svg        | Svg             | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/svg#Svg)                 |
| 列表项模板           | tmpl       | ListItemTemplate | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/tmpl#ListItemTemplate)   |
| 节点              | tmpl       | Node            | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/tmpl#Node)               |
| 图形绘制            | drawx      | Draw            | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/drawx#Draw)              |
| 动画序列            | ani        | Anima           | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/ani#Anima)               |
| 动画组             | ani        | AnimaGroup      | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/ani#AnimaGroup)          |
| 动画项             | ani        | AnimaItem       | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/ani#AnimaItem)           |
| 动画旋转项           | ani        | AnimaRotate     | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/ani#AnimaRotate)         |
| 动画缩放项           | ani        | AnimaScale      | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/ani#AnimaScale)          |
| 含有XCGUI所有API和结构体 | xc         |                 | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/xc#section-documentation) |
| XCGUI常量         | xcc        |                 | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/xcc)                     |
| 缓动              | ease       |                 | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/ease)                    |
| 资源操作            | res        |                 | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/res)                     |
| Windows系统api    | wapi       |                 | 持续更新     | [文档](https://pkg.go.dev/github.com/twgh/xcgui/wapi)                    |
| 调用wapi封装了对窗口的操作 | wapi/wnd   |                 | 持续更新     | [文档](https://pkg.go.dev/github.com/twgh/xcgui/wapi/wnd)                |
| 调用wapi封装了一些常用函数 | wapi/wutil |                 | 持续更新     | [文档](https://pkg.go.dev/github.com/twgh/xcgui/wapi/wutil)              |
