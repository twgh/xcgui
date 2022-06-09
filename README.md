<h1 align="center">XCGUI</h1>
<p align="center">
    <a href="https://github.com/twgh/xcgui/releases"><img src="https://img.shields.io/badge/release-1.3.340-blue.svg?" alt="release"></a>
    <a href="http://www.xcgui.com"><img src="https://img.shields.io/badge/XCGUI-3.3.4-blue.svg?" alt="XCGUI"></a>
   <a href="https://golang.org"> <img src="https://img.shields.io/badge/golang-1.17-blue" alt="golang"></a>
    <a href="https://pkg.go.dev/github.com/twgh/xcgui"><img src="https://img.shields.io/badge/go.dev-reference-brightgreen" alt="GoDoc"></a>
    <a href="https://opensource.org/licenses/MIT"><img src="https://img.shields.io/badge/License-MIT-brightgreen.svg?" alt="License"></a>
    <br><br>
    <a href="https://github.com/twgh/xcgui-example">程序示例</a>&nbsp;&nbsp;
    <a href="http://www.xcgui.com/doc-ui/">官方文档</a>&nbsp;&nbsp;
	<a href="https://pkg.go.dev/github.com/twgh/xcgui">项目文档</a>&nbsp;&nbsp;
	<a href="http://mall.xcgui.com">官方资源</a>
</p>




## 介绍

[English](./README-en.md) | 简体中文

- 本库封装自炫彩界面库，功能丰富(1000多个API接口)，简单易用，轻量级，高度DIY自定义，支持一键换肤。
- 炫彩界面库是由C/C++语言开发：软件运行效率高，不需要第三方库的支持，不依赖MFC，ATL，WINDOWS标准控件等。
- DirectUI设计思想：在窗口内没有子窗口，界面元素都是逻辑上的区域(无HWND句柄，安全，灵活)，所有UI元素都是自主开发(不受系统限制)，更加灵活的实现各种程序界面，满足不同用户的需求。
- 拥有免费的UI设计器：快速开发工具，所见即所得，高度自定义系统(DIY)，让UI开发变的更加简单。
- 支持Direct2D，硬件加速，能更大程度的发挥硬件特性，创建高性能，高质量的2D图形。
- 编译后，只需携带一个大小为2.5MB的xcgui.dll，没有其他依赖。
- 有完善的中文官方文档：[中文官方文档](http://www.xcgui.com/doc-ui/)

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
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xcc"
)

//go:embed res/qqmusic.zip
var qqmusic []byte

func main() {
	a := app.New(true)
	// 从内存zip中加载资源文件
	a.LoadResourceZipMem(qqmusic, "resource.res", "")
	// 从内存zip中加载布局文件, 创建窗口对象
	w := window.NewWindowByLayoutZipMem(qqmusic, "main.xml", "", 0, 0)
	// 调整布局
	w.AdjustLayout()
	// 显示窗口
	w.ShowWindow(xcc.SW_SHOW)
	a.Run()
	a.Exit()
}
```

## 动态链接库下载

| 网盘         | 下载地址                                                     |
| ------------ | ------------------------------------------------------------ |
| 百度网盘     | [下载](https://pan.baidu.com/s/1rC3unQGaxnRUCMm8z8qzvA?pwd=1111)，提取码：1111 |
| OneDrive     | [下载](https://1drv.ms/u/s!ApZP3niad5hpdGuodyU_GvugJ_g?e=yBEKmm) |
| Google Drive | [下载](https://drive.google.com/drive/folders/1MuisSsDIr1rjqTkdFIewOgb89SYdf5s6?usp=sharing) |
| Github       | [下载](https://github.com/twgh/FileStorage/tree/main/xcgui)  |

程序运行时需要把`XCGUI.dll`放到程序运行目录。

在开发时最好是放到C:\Windows\System32目录，这样就不需要频繁把dll放到不同程序的运行目录了。

## 简单窗口（纯代码）

[![SimpleWindow](https://s1.ax1x.com/2022/05/24/XiEWtg.png)](https://github.com/twgh/xcgui-example/blob/main/SimpleWindow)


```go
package main

import (
	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xcc"
)

func main() {
	// 1.初始化UI库
	a := app.New(true)
	// 2.创建窗口
	w := window.NewWindow(0, 0, 430, 300, "", 0, xcc.Window_Style_Simple|xcc.Window_Style_Btn_Close)
	// 设置窗口边框大小
	w.SetBorderSize(0, 30, 0, 0)
	// 设置窗口透明类型
	w.SetTransparentType(xcc.Window_Transparent_Shadow)
	// 设置窗口阴影
	w.SetShadowInfo(8, 254, 10, false, 0)
	// 3.显示窗口
	w.ShowWindow(xcc.SW_SHOW)
	// 4.运行程序
	a.Run()
	// 5.释放UI库
	a.Exit()
}
```

## 常量

xcc包里都是常量，像这样使用：`xcc.Window_Transparent_Shadow`

需要用到xcc包常量的参数注释都是类似这样的，复制`Window_Transparent_`到[xcc包](https://pkg.go.dev/github.com/twgh/xcgui/xcc#pkg-constants)里搜索即可看到相关常量注释：[![5eX6pD.jpg](https://s4.ax1x.com/2021/12/22/TQvQzt.png)](https://github.com/twgh/xcgui/blob/main/window/windowbase.go#L630)

## 命令介绍

xc包里包含xcgui.dll里所有的API，有一千多个函数，可以直接使用，封装好的类都在其他包里。

在某些情况下，把xc包里的原生函数与封装好的类混合起来使用会更加方便。

炫彩所有的结构体也都在xc包里。

建议使用 [Goland](https://www.jetbrains.com/go/?from=xcgui) 进行开发，以获得最好的开发体验。本项目所使用的注释格式在Goland里看起来是最好的。

## 事件

炫彩的全部事件都已经定义好了，都是以Event开头的， 以1结尾的事件是会传进去元素的句柄。

事件回调函数尽量不要使用匿名函数，使用匿名函数意味着您每次都在创建1个新的回调，最终您将会遇到因创建过多回调导致程序崩溃的报错（大概在2000个回调时会遇到），事件回调函数不使用匿名函数即可避免此问题。

[![xc-event.png](https://z3.ax1x.com/2021/11/23/opdyh6.png)](https://z3.ax1x.com/2021/11/23/opdyh6.png)

一个事件可以注册多个处理函数，执行顺序为先执行最后注册的函数，最后执行第一个注册的函数，当你想拦截当前事件或不想向后传递，只需要将参数`*pbHandled=true`即可。

## 关于版本号

以`1.3.330`为例，1仅代表本库是正式版，3.33代表是XCGUI官方3.3.3版本，最后的0代表是基于3.33封装的第1个版本，如果基于3.33还有更新，那么会累加。

## JetBrains 开源证书支持

`xcgui` 项目一直以来都是在 [JetBrains](https://www.jetbrains.com/?from=xcgui) 公司旗下的 [GoLand](https://www.jetbrains.com/go/?from=xcgui) 集成开发环境中进行开发，基于 **free JetBrains Open Source license(s)** 正版免费授权，在此表达我的谢意。

[<img src="https://s1.ax1x.com/2022/05/24/XiFI6x.png" alt="jetbrains.png" />](https://www.jetbrains.com/?from=xcgui)

## 模块继承关系

[![xcgui模块关系图.png](https://z3.ax1x.com/2021/11/18/ITAXJf.png)](https://z3.ax1x.com/2021/11/18/ITAXJf.png)

## 封装进度

这些类都是基于xc包里的一千多个函数封装的。

| 中文名称                              | 包名             | 类名             | 是否封装完毕 | 文档                                                         |
| ------------------------------------- | ---------------- | ---------------- | ------------ | ------------------------------------------------------------ |
| 程序（炫彩全局API）                   | app              | App              | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.340/app#App) |
| 窗口                                  | window           | Window           | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.340/window#Window) |
| 框架窗口                              | window           | FrameWindow      | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.340/window#FrameWindow) |
| 模态窗口                              | window           | ModalWindow      | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.340/window#ModalWindow) |
| 形状对象                              | widget           | Shape            | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.340/widget#Shape) |
| 圆形形状对象                          | widget           | ShapeEllipse     | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.340/widget#ShapeEllipse) |
| 形状对象GIF                           | widget           | ShapeGif         | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.340/widget#ShapeGif) |
| 组框形状对象                          | widget           | ShapeGroupBox    | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.340/widget#ShapeGroupBox) |
| 直线形状对象                          | widget           | ShapeLine        | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.340/widget#ShapeLine) |
| 形状对象图片                          | widget           | ShapePicture     | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.340/widget#ShapePicture) |
| 矩形形状对象                          | widget           | ShapeRect        | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.340/widget#ShapeRect) |
| 形状对象文本                          | widget           | ShapeText        | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.340/widget#ShapeText) |
| 表格                                  | widget           | Table            | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.340/widget#Table) |
| 按钮                                  | widget           | Button           | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.340/widget#Button) |
| 下拉组合框                            | widget           | ComboBox         | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.340/widget#ComboBox) |
| 编辑框(常规, 富文本, 聊天气泡)        | widget           | Edit             | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.340/widget#Edit) |
| 代码编辑框                            | widget           | Editor           | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.340/widget#Editor) |
| 基础元素                              | widget           | Element          | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.340/widget#Element) |
| 列表                                  | widget           | List             | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.340/widget#List) |
| 列表框                                | widget           | ListBox          | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.340/widget#ListBox) |
| 弹出菜单                              | widget           | Menu             | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.340/widget#Menu) |
| 进度条                                | widget           | ProgressBar      | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.340/widget#ProgressBar) |
| 静态文本连接按钮                      | widget           | TextLink         | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.340/widget#TextLink) |
| 布局元素                              | widget           | LayoutEle        | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.340/widget#LayoutEle) |
| 布局框架                              | widget           | LayoutFrame      | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.340/widget#LayoutFrame) |
| 列表视图                              | widget           | ListView         | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.340/widget#ListView) |
| 菜单条                                | widget           | MenuBar          | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.340/widget#MenuBar) |
| Pane元素                              | widget           | Pane             | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.340/widget#Pane) |
| 滚动条                                | widget           | ScrollBar        | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.340/widget#ScrollBar) |
| 滚动视图                              | widget           | ScrollView       | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.340/widget#ScrollView) |
| 滑动条元素                            | widget           | SliderBar        | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.340/widget#SliderBar) |
| 标签栏元素                            | widget           | TabBar           | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.340/widget#TabBar) |
| 工具条                                | widget           | ToolBar          | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.340/widget#ToolBar) |
| 列表树元素                            | widget           | Tree             | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.340/widget#Tree) |
| 日期时间                              | widget           | DateTime         | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.340/widget#DateTime) |
| 月历卡片                              | widget           | MonthCal         | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.340/widget#MonthCal) |
| 数据适配器-列表视元素                 | adapter          | AdapterListView  | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.340/adapter#AdapterListView) |
| 数据适配器-单列Map-列表头(listHeader) | adapter          | AdapterMap       | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.340/adapter#AdapterMap) |
| 数据适配器-XList-XListBox             | adapter          | AdapterTable     | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.340/adapter#AdapterTable) |
| 数据适配器-树元素                     | adapter          | AdapterTree      | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.340/adapter#AdapterTree) |
| 背景管理器                            | bkmanager        | BkManager        | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.340/bkmanager#BkManager) |
| 背景对象                              | bkobj            | BkObj            | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.340/bkobj#BkObj) |
| 字体                                  | font             | Font             | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.340/font#Font) |
| 图片操作                              | imagex           | Image            | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.340/imagex#Image) |
| SVG矢量图形                           | svg              | Svg              | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.340/svg#Svg) |
| 列表项模板                            | listitemtemplate | ListItemTemplate | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.340/listitemtemplate#ListItemTemplate) |
| 节点                                  | listitemtemplate | Node             | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.340/listitemtemplate#Node) |
| 图形绘制                              | drawx            | Draw             | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.340/drawx#Draw) |
| 动画序列                              | ani              | Anima            | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.340/ani#Anima) |
| 动画组                                | ani              | AnimaGroup       | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.340/ani#AnimaGroup) |
| 动画项                                | ani              | AnimaItem        | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.340/ani#AnimaItem) |
| 动画旋转项                            | ani              | AnimaRotate      | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.340/ani#AnimaRotate) |
| 动画缩放项                            | ani              | AnimaScale       | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.340/ani#AnimaScale) |
| 含有XCGUI所有API和结构体              | xc               |                  | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.340/xc#section-documentation) |
| XCGUI常量                             | xcc              |                  | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.340/xcc) |
| 缓动                                  | ease             |                  | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.340/ease) |
| 资源操作                              | res              |                  | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.340/res) |
| Windows系统api                        | wapi             |                  | 持续更新     | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.340/wapi) |
| 调用wapi封装了对窗口的操作            | wnd              |                  | 持续更新     | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.340/wnd) |
