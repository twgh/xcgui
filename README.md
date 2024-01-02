<h1 align="center">XCGUI</h1>
<p align="center">
    <a href="https://github.com/twgh/xcgui/releases"><img src="https://img.shields.io/badge/release-1.3.382-blue" alt="release"></a>
    <a href="http://www.xcgui.com"><img src="https://img.shields.io/badge/XCGUI-3.3.8-blue" alt="XCGUI"></a>
   <a href="https://golang.org"> <img src="https://img.shields.io/badge/golang-1.16-blue" alt="golang"></a>
    <a href="https://pkg.go.dev/github.com/twgh/xcgui"><img src="https://img.shields.io/badge/go.dev-reference-brightgreen" alt="GoDoc"></a>
    <a href="https://opensource.org/licenses/MIT"><img src="https://img.shields.io/badge/License-MIT-brightgreen" alt="License"></a>
    <br><br>
    <a href="https://github.com/twgh/xcgui-example">程序示例</a>&nbsp;&nbsp;
    <a href="http://www.xcgui.com/doc-ui/">官方文档</a>&nbsp;&nbsp;
	<a href="https://pkg.go.dev/github.com/twgh/xcgui">项目文档</a>&nbsp;&nbsp;
	<a href="http://mall.xcgui.com">官方资源</a>
</p>









## 介绍

[English](./README-en.md) | 简体中文

- 本库封装自炫彩界面库，功能丰富(近2000个API接口)，简单易用，轻量级，高度DIY自定义，支持一键换肤。
- 炫彩界面库是由C/C++语言开发：软件运行效率高，不需要第三方库的支持，不依赖MFC，ATL，WINDOWS标准控件等。
- DirectUI设计思想：在窗口内没有子窗口，界面元素都是逻辑上的区域(无HWND句柄，安全，灵活)，所有UI元素都是自主开发(不受系统限制)，更加灵活的实现各种程序界面，满足不同用户的需求。
- 拥有免费的UI设计器：快速开发工具，所见即所得，高度自定义系统(DIY)，让UI开发变的更加简单。
- 支持Direct2D，硬件加速，能更大程度的发挥硬件特性，创建高性能，高质量的2D图形。
- [wiki](https://github.com/twgh/xcgui/wiki) 里有简单的入门教程，有空可以看一下，少走弯路。
- 有完善的 [中文官方文档](http://www.xcgui.com/doc-ui/)，[官方资源社区](http://mall.xcgui.com)。
- 官方QQ群(人多用各种语言的都有)：[2283812](https://jq.qq.com/?_wv=1027&k=AiXY4uMc)，只聊Go语言版的QQ群：[793788011](https://jq.qq.com/?_wv=1027&k=bkKgsYYk)

## 获取

```
go get github.com/twgh/xcgui
```

## 可视化UI设计器

使用UI设计器可以快速设计界面，节省大量代码。

[![uidesigner](https://z3.ax1x.com/2021/09/15/4Vmh9S.png)](https://github.com/twgh/xcgui-example/tree/main/uidesigner)

[设计器使用例子](https://github.com/twgh/xcgui-example/tree/main/uidesigner)，只有这么多代码：

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
	// 从内存zip中加载资源文件
	a.LoadResourceZipMem(qqmusic, "resource.res", "")
	// 从内存zip中加载布局文件, 创建窗口对象
	w := window.NewByLayoutZipMem(qqmusic, "main.xml", "", 0, 0)
    
	// songTitle是在main.xml中给歌曲名(shapeText组件)设置的name属性的值.
	// 通过 GetObjectByName 可以获取布局文件中设置了name属性的组件的句柄.
	// 可简化为: widget.NewShapeTextByName("songTitle").
	song := widget.NewShapeTextByHandle(a.GetObjectByName("songTitle"))
	println(song.GetText()) // 输出: 两只老虎爱跳舞
    
	// 调整布局
	w.AdjustLayout()
	// 显示窗口
	w.Show(true)
	a.Run()
	a.Exit()
}
```

## 动态链接库下载

程序运行时需要把`xcgui.dll`放到程序运行目录。

在开发时最好是放到`C:\Windows\System32`目录，这样就不需要频繁把dll放到不同程序的运行目录了。

#### （1）文件直链

| 64位 | [点击下载](https://pkggo-generic.pkg.coding.net/xcgui/file/xcgui.dll?version=latest) |
| ---- | ------------------------------------------------------------ |
| 32位 | [点击下载](https://pkggo-generic.pkg.coding.net/xcgui/file/xcgui-32.dll?version=latest) |

#### （2）命令行下载

64位

```bash
curl -fL "https://pkggo-generic.pkg.coding.net/xcgui/file/xcgui.dll?version=latest" -o xcgui.dll
```

32位

```bash
curl -fL "https://pkggo-generic.pkg.coding.net/xcgui/file/xcgui-32.dll?version=latest" -o xcgui.dll
```

#### （3）使用getxcgui工具下载

> 请确保 `%GOPATH%\bin` 在环境变量path中

```bash
go install github.com/twgh/getxcgui@latest
getxcgui  
```

如果要把dll直接下载到`C:\Windows\System32`目录里，请使用如下命令：

```bash
getxcgui -o %windir%\system32\xcgui.dll
```

此工具的源码在[这里](https://github.com/twgh/getxcgui)，更多flags可以点[进去](https://github.com/twgh/getxcgui#flags)查看

#### （4）网盘下载

网盘内还包含`界面设计器`和`chm帮助文档`

| 网盘         | 下载地址                                                     |
| ------------ | ------------------------------------------------------------ |
| 百度网盘     | [下载](https://pan.baidu.com/s/1rC3unQGaxnRUCMm8z8qzvA?pwd=1111) |
| 蓝奏云     | [下载](https://wwi.lanzoup.com/b0cqd6nkb) |

## 简单窗口（纯代码）

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
	// 1.初始化UI库
	a := app.New(true)
	a.EnableDPI(true)
	a.EnableAutoDPI(true)
	// 2.创建窗口
	w := window.New(0, 0, 430, 300, "xcgui window", 0, xcc.Window_Style_Default|xcc.Window_Style_Drag_Window)

	// 设置窗口边框大小
	w.SetBorderSize(0, 30, 0, 0)
	// 设置窗口图标
	a.SetWindowIcon(imagex.NewBySvgStringW(svgIcon).Handle)
	// 设置窗口透明类型
	w.SetTransparentType(xcc.Window_Transparent_Shadow)
	// 设置窗口阴影
	w.SetShadowInfo(8, 255, 10, false, 0)

	// 创建按钮
	btn := widget.NewButton(165, 135, 100, 30, "Button", w.Handle)
	// 注册按钮被单击事件
	btn.Event_BnClick(func(pbHandled *bool) int {
		a.MessageBox("提示", btn.GetText(), xcc.MessageBox_Flag_Ok|xcc.MessageBox_Flag_Icon_Info, w.GetHWND(), xcc.Window_Style_Modal)
		return 0
	})

	// 3.显示窗口
	w.Show(true)
	// 4.运行程序
	a.Run()
	// 5.释放UI库
	a.Exit()
}

var svgIcon = `<svg t="1669088647057" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="5490" width="22" height="22"><path d="M517.12 512.8704m-432.3328 0a432.3328 432.3328 0 1 0 864.6656 0 432.3328 432.3328 0 1 0-864.6656 0Z" fill="#51C5FF" p-id="5491"></path><path d="M292.1472 418.7136c-85.0432 0-160.4096 41.3696-207.104 105.0624 4.5568 182.7328 122.368 337.3056 285.952 396.032 103.2192-33.28 177.92-130.048 177.92-244.3776 0-141.7216-114.944-256.7168-256.768-256.7168z" fill="#7BE0FF" p-id="5492"></path><path d="M800.2048 571.6992l-101.888-58.8288 101.888-58.8288c16.896-9.728 22.6816-31.3344 12.9536-48.2304l-55.296-95.744c-9.728-16.896-31.3344-22.6816-48.2304-12.9536l-101.888 58.8288V238.336c0-19.5072-15.8208-35.328-35.328-35.328H461.824c-19.5072 0-35.328 15.8208-35.328 35.328v117.6064L324.608 297.1136c-16.896-9.728-38.5024-3.9424-48.2304 12.9536l-55.296 95.744c-9.728 16.896-3.9424 38.5024 12.9536 48.2304l101.888 58.8288-101.888 58.8288c-16.896 9.728-22.6816 31.3344-12.9536 48.2304l55.296 95.744c9.728 16.896 31.3344 22.6816 48.2304 12.9536l101.888-58.8288v117.6064c0 19.5072 15.8208 35.328 35.328 35.328h110.592c19.5072 0 35.328-15.8208 35.328-35.328v-117.6064l101.888 58.8288c16.896 9.728 38.5024 3.9424 48.2304-12.9536l55.296-95.744c9.728-16.896 3.9424-38.5024-12.9536-48.2304z" fill="#CAF8FF" p-id="5493"></path><path d="M517.12 512.8704m-234.24 0a234.24 234.24 0 1 0 468.48 0 234.24 234.24 0 1 0-468.48 0Z" fill="#FFFFFF" p-id="5494"></path><path d="M517.12 512.8704m-103.5776 0a103.5776 103.5776 0 1 0 207.1552 0 103.5776 103.5776 0 1 0-207.1552 0Z" fill="#51C5FF" p-id="5495"></path></svg>`
```

## 常量

xcc包里都是常量，像这样使用：`xcc.Window_Transparent_Shadow`

需要用到xcc包常量的参数注释都是类似这样的，复制`Window_Transparent_`到[xcc包](https://pkg.go.dev/github.com/twgh/xcgui/xcc#pkg-constants)里搜索即可看到相关常量注释：

[![注释](https://s4.ax1x.com/2021/12/22/TQvQzt.png)](https://github.com/twgh/xcgui/blob/main/window/windowbase.go#L630)

## 命令介绍

xc包里包含xcgui.dll里所有的API，有一千多个函数，可以直接使用，封装好的类都在其他包里。

在某些情况下，把xc包里的原生函数与封装好的类混合起来使用会更加方便。

炫彩所有的结构体也都在xc包里。

建议使用 [Goland](https://www.jetbrains.com/go/?from=xcgui) 进行开发，以获得最好的开发体验。本项目所使用的注释格式在Goland里看起来是最好的。

## 事件

炫彩的全部事件都已经定义好了，都是以Event开头的， 以1结尾的事件是会传进去元素的句柄。

在循环中注册事件时，回调函数尽量不要使用匿名函数，使用匿名函数意味着您每次都在创建1个新的回调，最终您将会遇到因创建过多回调导致程序崩溃的报错（大概在2000个回调时会遇到），事件回调函数不使用匿名函数即可避免此问题，一般程序应该不会用到2000个事件，只要注意在多次循环中注册事件时不使用匿名函数就行。这个是Go语言添加的限制，不是xcgui添加的限制。

[![xc-event.png](https://z3.ax1x.com/2021/11/23/opdyh6.png)](https://z3.ax1x.com/2021/11/23/opdyh6.png)

一个事件可以注册多个处理函数，执行顺序为先执行最后注册的函数，最后执行第一个注册的函数，当你想拦截当前事件或不想向后传递，只需要将参数`*pbHandled=true`即可。

## 关于版本号

以`1.3.330`为例，1仅代表本库是正式版，3.33代表是XCGUI官方3.3.3版本，最后的0代表是基于3.33封装的第1个版本，如果基于3.33还有更新，那么会累加。

## JetBrains 开源证书支持

`xcgui` 项目一直以来都是在 [JetBrains](https://www.jetbrains.com/?from=xcgui) 公司旗下的 [GoLand](https://www.jetbrains.com/go/?from=xcgui) 集成开发环境中进行开发，基于 **free JetBrains Open Source license(s)** 正版免费授权，在此表达我的谢意。

[<img src="https://s1.ax1x.com/2022/05/24/XiFI6x.png" alt="jetbrains.png" />](https://www.jetbrains.com/?from=xcgui)

## 模块继承关系

[xcgui模块关系图](http://www.xcgui.com/doc-ui/page_diagram.html)

## 封装进度

这些类都是基于xc包里的一千多个函数封装的。

| 中文名称                              | 包名       | 类名             | 是否封装完毕 | 文档                                                         |
| ------------------------------------- | ---------- | ---------------- | ------------ | ------------------------------------------------------------ |
| 程序（炫彩全局API）                   | app        | App              | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/app#App)     |
| 窗口                                  | window     | Window           | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/window#Window) |
| 框架窗口                              | window     | FrameWindow      | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/window#FrameWindow) |
| 模态窗口                              | window     | ModalWindow      | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/window#ModalWindow) |
| 形状对象                              | widget     | Shape            | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/widget#Shape) |
| 圆形形状对象                          | widget     | ShapeEllipse     | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/widget#ShapeEllipse) |
| 形状对象GIF                           | widget     | ShapeGif         | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/widget#ShapeGif) |
| 组框形状对象                          | widget     | ShapeGroupBox    | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/widget#ShapeGroupBox) |
| 直线形状对象                          | widget     | ShapeLine        | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/widget#ShapeLine) |
| 形状对象图片                          | widget     | ShapePicture     | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/widget#ShapePicture) |
| 矩形形状对象                          | widget     | ShapeRect        | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/widget#ShapeRect) |
| 形状对象文本                          | widget     | ShapeText        | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/widget#ShapeText) |
| 表格                                  | widget     | Table            | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/widget#Table) |
| 按钮                                  | widget     | Button           | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/widget#Button) |
| 下拉组合框                            | widget     | ComboBox         | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/widget#ComboBox) |
| 编辑框(常规, 富文本, 聊天气泡)        | widget     | Edit             | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/widget#Edit) |
| 代码编辑框                            | widget     | Editor           | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/widget#Editor) |
| 基础元素                              | widget     | Element          | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/widget#Element) |
| 列表                                  | widget     | List             | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/widget#List) |
| 列表框                                | widget     | ListBox          | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/widget#ListBox) |
| 弹出菜单                              | widget     | Menu             | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/widget#Menu) |
| 进度条                                | widget     | ProgressBar      | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/widget#ProgressBar) |
| 静态文本连接按钮                      | widget     | TextLink         | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/widget#TextLink) |
| 布局元素                              | widget     | LayoutEle        | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/widget#LayoutEle) |
| 布局框架                              | widget     | LayoutFrame      | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/widget#LayoutFrame) |
| 列表视图                              | widget     | ListView         | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/widget#ListView) |
| 菜单条                                | widget     | MenuBar          | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/widget#MenuBar) |
| Pane元素                              | widget     | Pane             | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/widget#Pane) |
| 滚动条                                | widget     | ScrollBar        | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/widget#ScrollBar) |
| 滚动视图                              | widget     | ScrollView       | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/widget#ScrollView) |
| 滑动条元素                            | widget     | SliderBar        | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/widget#SliderBar) |
| 标签栏元素                            | widget     | TabBar           | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/widget#TabBar) |
| 工具条                                | widget     | ToolBar          | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/widget#ToolBar) |
| 列表树元素                            | widget     | Tree             | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/widget#Tree) |
| 日期时间                              | widget     | DateTime         | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/widget#DateTime) |
| 月历卡片                              | widget     | MonthCal         | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/widget#MonthCal) |
| 数据适配器-列表视元素                 | adapter    | AdapterListView  | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/adapter#AdapterListView) |
| 数据适配器-单列Map-列表头(listHeader) | adapter    | AdapterMap       | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/adapter#AdapterMap) |
| 数据适配器-XList-XListBox             | adapter    | AdapterTable     | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/adapter#AdapterTable) |
| 数据适配器-树元素                     | adapter    | AdapterTree      | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/adapter#AdapterTree) |
| 背景管理器                            | bkmanager  | BkManager        | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/bkmanager#BkManager) |
| 背景对象                              | bkobj      | BkObj            | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/bkobj#BkObj) |
| 字体                                  | font       | Font             | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/font#Font)   |
| 图片操作                              | imagex     | Image            | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/imagex#Image) |
| SVG矢量图形                           | svg        | Svg              | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/svg#Svg)     |
| 列表项模板                            | tmpl       | ListItemTemplate | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/tmpl#ListItemTemplate) |
| 节点                                  | tmpl       | Node             | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/tmpl#Node)   |
| 图形绘制                              | drawx      | Draw             | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/drawx#Draw)  |
| 动画序列                              | ani        | Anima            | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/ani#Anima)   |
| 动画组                                | ani        | AnimaGroup       | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/ani#AnimaGroup) |
| 动画项                                | ani        | AnimaItem        | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/ani#AnimaItem) |
| 动画旋转项                            | ani        | AnimaRotate      | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/ani#AnimaRotate) |
| 动画缩放项                            | ani        | AnimaScale       | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/ani#AnimaScale) |
| 含有XCGUI所有API和结构体              | xc         |                  | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/xc#section-documentation) |
| XCGUI常量                             | xcc        |                  | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/xcc)         |
| 缓动                                  | ease       |                  | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/ease)        |
| 资源操作                              | res        |                  | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui/res)         |
| Windows系统api                        | wapi       |                  | 持续更新     | [文档](https://pkg.go.dev/github.com/twgh/xcgui/wapi)        |
| 调用wapi封装了对窗口的操作            | wapi/wnd   |                  | 持续更新     | [文档](https://pkg.go.dev/github.com/twgh/xcgui/wapi/wnd)    |
| 调用wapi封装了一些常用函数            | wapi/wutil |                  | 持续更新     | [文档](https://pkg.go.dev/github.com/twgh/xcgui/wapi/wutil)  |
