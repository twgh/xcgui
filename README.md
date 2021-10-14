<h1 align="center">XCGUI</h1>
<p align="center">
    <a href="https://github.com/twgh/xcgui/releases"><img src="https://img.shields.io/badge/version-1.0.3-blue.svg?style=flat-square" alt="version"></a>
    <img src="https://img.shields.io/badge/golang-1.16-blue"/>
    <a href="https://pkg.go.dev/github.com/twgh/xcgui"><img src="https://img.shields.io/badge/go.dev-reference-007d9c ?logo=go&logoColor=white" alt="GoDoc"></a>
    <a href="https://opensource.org/licenses/MIT"><img src="https://img.shields.io/badge/License-MIT-brightgreen.svg?style=flat-square" alt="License"></a>
    <br><br>
	<a href="https://pkg.go.dev/github.com/twgh/xcgui">项目文档</a>&nbsp;&nbsp;
	<a href="https://github.com/twgh/FileStorage/raw/main/xcgui/help/%E7%82%AB%E5%BD%A9%E7%95%8C%E9%9D%A2%E5%BA%93-%E5%B8%AE%E5%8A%A9%E6%96%87%E6%A1%A3(v3.0)-(2021-09-08).chm">帮助文档</a>&nbsp;&nbsp;
    <a href="https://github.com/twgh/xcgui-example">程序示例</a>
</p>


## 介绍

[English](./README-en.md) | 简体中文

DirectUI设计思想: 在窗口内没有子窗口，界面元素都是逻辑上的区域(无HWND句柄,安全,灵活), 所有UI元素都是自主开发(不受系统限制),  更加灵活的实现各种程序界面,满足不同用户的需求.

官方网站：[www.xcgui.com](http://www.xcgui.com "xcgui 官方网站")

## 可视化UI设计器

使用UI设计器可以快速设计界面，节省大量代码。

[![uidesigner](https://z3.ax1x.com/2021/09/15/4Vmh9S.png)](https://github.com/twgh/xcgui-example/blob/main/uidesigner/uidesigner.png)

[设计器使用例子](https://github.com/twgh/xcgui-example/tree/main/uidesigner)，只有这么多代码：

```go
package main

import (
	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xcc"
)

func main() {
	a := app.New("")
	// 添加文件搜索路径, 你运行时需要改成自己的路径，也可以使用相对路径
	a.AddFileSearchPath(`D:\GoProject\src\github.com\twgh\xcgui-example\uidesigner\res`)
	// 从zip中加载资源文件
	a.LoadResourceZip("qqmusic.zip", "resource.res", "")
	// 从zip中加载布局文件
	hWindow := a.LoadLayoutZip("qqmusic.zip", "main.xml", "", 0)
	if hWindow == 0 {
		panic("error")
	}
	// 创建窗口对象
	win := window.NewWindowByHandle(hWindow)

	// 调整布局
	win.AdjustLayout()
	// 显示窗口
	win.ShowWindow(xcc.SW_SHOW)

	a.Run()
	a.Exit()
}
```

## 获取

```go
go get github.com/twgh/xcgui
```

## 动态链接库下载

| 网盘         | 下载地址                                                     |
| ------------ | ------------------------------------------------------------ |
| 百度网盘     | [下载](https://pan.baidu.com/s/17zri2GlDOVUY8nvFTXWLFg)，提取码：wcs7 |
| OneDrive     | [下载](https://1drv.ms/u/s!ApZP3niad5hpdGuodyU_GvugJ_g?e=yBEKmm) |
| Google Drive | [下载](https://drive.google.com/drive/folders/1MuisSsDIr1rjqTkdFIewOgb89SYdf5s6?usp=sharing) |
| GIthub       | [下载](https://github.com/twgh/FileStorage/tree/main/xcgui)  |

程序运行时需要把`XCGUI.dll`放到程序运行目录。

在开发时最好是放到C:\Windows\System32目录，这样就不需要频繁把dll放到不同程序的运行目录了。

## 简单窗口（纯代码）

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
	// 1.初始化UI库
	a := app.New("")
	// 2.创建窗口
	win := window.NewWindow(0, 0, 466, 300, "炫彩窗口", 0, xcc.Xc_Window_Style_Default)

	// 设置窗口边框大小
	win.SetBorderSize(1, 30, 1, 1)
	// 设置窗口透明类型
	win.SetTransparentType(xcc.Window_Transparent_Shadow)
	// 设置窗口阴影
	win.SetShadowInfo(10, 255, 10, false, 0)
	// 窗口置顶
	win.SetTop()
	// 窗口居中
	win.Center()
	// 创建标签_窗口标题
	lbl_Title := widget.NewShapeText(15, 15, 56, 20, "Title", win.Handle)
	lbl_Title.SetTextColor(xc.RGB(255, 255, 255), 255)

	// 创建最小化按钮
	btn_Min := widget.NewButton(396, 10, 30, 30, "-", win.Handle)
	btn_Min.SetTextColor(xc.RGB(255, 255, 255), 255)
	btn_Min.SetType(xcc.Button_Type_Min)
	btn_Min.EnableBkTransparent(true)
	// 创建结束按钮
	btn_Close := widget.NewButton(426, 10, 30, 30, "X", win.Handle)
	btn_Close.SetTextColor(xc.RGB(255, 255, 255), 255)
	btn_Close.SetType(xcc.Button_Type_Close)
	btn_Close.EnableBkTransparent(true)

	// 3.显示窗口
	win.ShowWindow(xcc.SW_SHOW)
	// 4.运行程序
	a.Run()
	// 5.释放UI库
	a.Exit()
}
```

## 常量

xcc包里都是常量，像这样使用：`xcc.Xc_Window_Style_Default`

需要用到xcc包常量的参数注释都是类似这样的，复制`Xc_Window_Style_`到[xcc包](https://github.com/twgh/xcgui/blob/main/xcc/xcconst.go)里搜索即可看到相关常量注释：[![5eX6pD.jpg](https://z3.ax1x.com/2021/10/12/5eX6pD.jpg)](https://github.com/twgh/xcgui/blob/main/window/window.go#L19)

## 命令介绍

xc包里的函数都是xcgui.dll里原本的函数，有一千多个函数，可以直接使用，封装好的类都在其他包里。

在某些情况下，把xc包里的原生函数与封装好的类混合起来使用会更加方便。

## 封装进度

这些类都是基于xc包里的一千多个函数封装的。

| 中文名称                              | 包名             | 类名             | 是否封装完毕 | 文档                                                         |
| ------------------------------------- | ---------------- | ---------------- | ------------ | ------------------------------------------------------------ |
| 程序                                  | app              | App              | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.0.3/app#App) |
| 窗口                                  | window           | Window           | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.0.3/window#Window) |
| 框架窗口                              | window           | FrameWindow      | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.0.3/window#FrameWindow) |
| 模态窗口                              | window           | ModalWindow      | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.0.3/window#ModalWindow) |
| 形状对象                              | widget           | Shape            | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.0.3/widget#Shape) |
| 圆形形状对象                          | widget           | ShapeEllipse     | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.0.3/widget#ShapeEllipse) |
| 形状对象GIF                           | widget           | ShapeGif         | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.0.3/widget#ShapeGif) |
| 组框形状对象                          | widget           | ShapeGroupBox    | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.0.3/widget#ShapeGroupBox) |
| 直线形状对象                          | widget           | ShapeLine        | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.0.3/widget#ShapeLine) |
| 形状对象图片                          | widget           | ShapePicture     | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.0.3/widget#ShapePicture) |
| 矩形形状对象                          | widget           | ShapeRect        | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.0.3/widget#ShapeRect) |
| 形状对象文本                          | widget           | ShapeText        | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.0.3/widget#ShapeText) |
| 表格                                  | widget           | Table            | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.0.3/widget#Table) |
| 按钮                                  | widget           | Button           | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.0.3/widget#Button) |
| 下拉组合框                            | widget           | ComboBox         | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.0.3/widget#ComboBox) |
| 编辑框(常规, 富文本, 聊天气泡)        | widget           | Edit             | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.0.3/widget#Edit) |
| 代码编辑框                            | widget           | Editor           | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.0.3/widget#Editor) |
| 基础元素                              | widget           | Element          | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.0.3/widget#Element) |
| 列表                                  | widget           | List             | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.0.3/widget#List) |
| 列表框                                | widget           | ListBox          | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.0.3/widget#ListBox) |
| 弹出菜单                              | widget           | Menu             | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.0.3/widget#Menu) |
| 进度条                                | widget           | ProgressBar      | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.0.3/widget#ProgressBar) |
| 静态文本连接按钮                      | widget           | TextLink         | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.0.3/widget#TextLink) |
| 布局元素                              | widget           | LayoutEle        | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.0.3/widget#LayoutEle) |
| 布局框架                              | widget           | LayoutFrame      | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.0.3/widget#LayoutFrame) |
| 列表视图                              | widget           | ListView         | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.0.3/widget#ListView) |
| 菜单条                                | widget           | MenuBar          | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.0.3/widget#MenuBar) |
| Pane元素                              | widget           | Pane             | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.0.3/widget#Pane) |
| 滚动条                                | widget           | ScrollBar        | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.0.3/widget#ScrollBar) |
| 滚动视图                              | widget           | ScrollView       | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.0.3/widget#ScrollView) |
| 滑动条元素                            | widget           | SliderBar        | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.0.3/widget#SliderBar) |
| 标签栏元素                            | widget           | TabBar           | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.0.3/widget#TabBar) |
| 工具条                                | widget           | ToolBar          | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.0.3/widget#ToolBar) |
| 列表树元素                            | widget           | Tree             | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.0.3/widget#Tree) |
| 日期时间                              | widget           | DateTime         | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.0.3/widget#DateTime) |
| 月历卡片                              | widget           | MonthCal         | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.0.3/widget#MonthCal) |
| 数据适配器-列表视元素                 | adapter          | AdapterListView  | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.0.3/adapter#AdapterListView) |
| 数据适配器-单列Map-列表头(listHeader) | adapter          | AdapterMap       | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.0.3/adapter#AdapterMap) |
| 数据适配器-XList-XListBox             | adapter          | AdapterTable     | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.0.3/adapter#AdapterTable) |
| 数据适配器-树元素                     | adapter          | AdapterTree      | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.0.3/adapter#AdapterTree) |
| 背景管理器                            | bkmanager        | BkManager        | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.0.3/bkmanager#BkManager) |
| 字体                                  | fontx            | FontX            | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.0.3/fontx#FontX) |
| 图片操作                              | image            | Image            | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.0.3/image#Image) |
| 列表项模板                            | listitemtemplate | ListItemTemplate | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.0.3/listitemtemplate#ListItemTemplate) |
| 节点                                  | listitemtemplate | Node             | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.0.3/listitemtemplate#Node) |
| 图形绘制                              | draw             | Draw             | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.0.3/draw#Draw) |
| XCGUI所有函数                         | xc               |                  | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.0.3/xc#section-documentation) |
| XCGUI常量                             | xcc              |                  | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.0.3/xcc)  |
| 缓动                                  | ease             |                  | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.0.3/ease) |
| 资源操作                              | res              |                  | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.0.3/res)  |

