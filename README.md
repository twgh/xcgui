<h1 align="center">XCGUI</h1>
<p align="center">
    <a href="https://github.com/twgh/xcgui/releases"><img src="https://img.shields.io/badge/release-1.3.12.1-blue.svg?" alt="release"></a>
    <a href="http://www.xcgui.com"><img src="https://img.shields.io/badge/XCGUI-3.1.2-blue.svg?" alt="XCGUI"></a>
    <a href="http://mall.xcgui.com"><img src="https://img.shields.io/badge/Resource-Mall-blue.svg?" alt="Resource Mall"></a>
   <a href="https://golang.org"> <img src="https://img.shields.io/badge/golang-1.16-brightgreen" alt="golang"></a>
    <a href="https://pkg.go.dev/github.com/twgh/xcgui"><img src="https://img.shields.io/badge/go.dev-reference-007d9c ?logo=go&logoColor=white" alt="GoDoc"></a>
    <a href="https://opensource.org/licenses/MIT"><img src="https://img.shields.io/badge/License-MIT-brightgreen.svg?" alt="License"></a>
    <br><br>
    <a href="https://github.com/twgh/xcgui-example">程序示例</a>&nbsp;&nbsp;
	<a href="https://pkg.go.dev/github.com/twgh/xcgui">项目文档</a>&nbsp;&nbsp;
    <a href="http://www.xcgui.com/doc-ui/">官方文档</a>&nbsp;&nbsp;
	<a href="https://github.com/twgh/FileStorage/blob/main/xcgui/help/%E7%82%AB%E5%BD%A9%E7%95%8C%E9%9D%A2%E5%BA%93-%E5%B8%AE%E5%8A%A9%E6%96%87%E6%A1%A3(v3.1)-(2021-10-07).chm?raw=true">帮助文档</a>&nbsp;&nbsp;
    <a href="http://mall.xcgui.com">资源商城</a>
</p>







## 介绍

[English](./README-en.md) | 简体中文

DirectUI设计思想: 在窗口内没有子窗口，界面元素都是逻辑上的区域(无HWND句柄,安全,灵活), 所有UI元素都是自主开发(不受系统限制),  更加灵活的实现各种程序界面,满足不同用户的需求.

XCGUI官网：[www.xcgui.com](http://www.xcgui.com "xcgui 官方网站")

官方资源商城：[mall.xcgui.com](http://mall.xcgui.com "xcgui 资源商城")

## 可视化UI设计器

使用UI设计器可以快速设计界面，节省大量代码。

[![uidesigner](https://z3.ax1x.com/2021/09/15/4Vmh9S.png)](https://github.com/twgh/xcgui-example/tree/main/uidesigner)

[设计器使用例子](https://github.com/twgh/xcgui-example/tree/main/uidesigner)，只有这么多代码：

```go
package main

import (
	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xcc"
)

func main() {
	a := app.New(true)
	// 添加文件搜索路径, 请使用go run运行程序, 如果你使用go build运行, 那么请把这里改成`res`
	a.AddFileSearchPath(`uidesigner\res`)
	// 从zip中加载资源文件
	a.LoadResourceZip(`qqmusic.zip`, "resource.res", "")
	// 从zip中加载布局文件
	hWindow := a.LoadLayoutZip(`qqmusic.zip`, "main.xml", "", 0)
	if hWindow == 0 {
		panic("LoadLayoutZip Error")
	}

	// 创建窗口对象
	w := window.NewWindowByHandle(hWindow)
	// 调整布局
	w.AdjustLayout()
	// 显示窗口
	w.ShowWindow(xcc.SW_SHOW)
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
| 百度网盘     | [下载](https://pan.baidu.com/s/1rC3unQGaxnRUCMm8z8qzvA)，提取码：1111 |
| OneDrive     | [下载](https://1drv.ms/u/s!ApZP3niad5hpdGuodyU_GvugJ_g?e=yBEKmm) |
| Google Drive | [下载](https://drive.google.com/drive/folders/1MuisSsDIr1rjqTkdFIewOgb89SYdf5s6?usp=sharing) |
| Github       | [下载](https://github.com/twgh/FileStorage/tree/main/xcgui)  |

程序运行时需要把`XCGUI.dll`放到程序运行目录。

在开发时最好是放到C:\Windows\System32目录，这样就不需要频繁把dll放到不同程序的运行目录了。

## 简单窗口（纯代码）

[![SimpleWindow](https://i.niupic.com/images/2021/11/01/9FiK.bmp)](https://github.com/twgh/xcgui-example/blob/main/SimpleWindow)


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

xcc包里都是常量，像这样使用：`xcc.Xc_Window_Style_Default`

需要用到xcc包常量的参数注释都是类似这样的，复制`Xc_Window_Style_`到[xcc包](https://github.com/twgh/xcgui/blob/main/xcc/xcconst.go)里搜索即可看到相关常量注释：[![5eX6pD.jpg](https://z3.ax1x.com/2021/10/12/5eX6pD.jpg)](https://github.com/twgh/xcgui/blob/main/window/window.go#L26)

## 命令介绍

xc包里的函数都是xcgui.dll里原本的函数，有一千多个函数，可以直接使用，封装好的类都在其他包里。

在某些情况下，把xc包里的原生函数与封装好的类混合起来使用会更加方便。

## 封装进度

这些类都是基于xc包里的一千多个函数封装的。

| 中文名称                              | 包名             | 类名             | 是否封装完毕 | 文档                                                         |
| ------------------------------------- | ---------------- | ---------------- | ------------ | ------------------------------------------------------------ |
| 程序                                  | app              | App              | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.12.1/app#App) |
| 窗口                                  | window           | Window           | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.12.1/window#Window) |
| 框架窗口                              | window           | FrameWindow      | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.12.1/window#FrameWindow) |
| 模态窗口                              | window           | ModalWindow      | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.12.1/window#ModalWindow) |
| 形状对象                              | widget           | Shape            | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.12.1/widget#Shape) |
| 圆形形状对象                          | widget           | ShapeEllipse     | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.12.1/widget#ShapeEllipse) |
| 形状对象GIF                           | widget           | ShapeGif         | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.12.1/widget#ShapeGif) |
| 组框形状对象                          | widget           | ShapeGroupBox    | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.12.1/widget#ShapeGroupBox) |
| 直线形状对象                          | widget           | ShapeLine        | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.12.1/widget#ShapeLine) |
| 形状对象图片                          | widget           | ShapePicture     | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.12.1/widget#ShapePicture) |
| 矩形形状对象                          | widget           | ShapeRect        | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.12.1/widget#ShapeRect) |
| 形状对象文本                          | widget           | ShapeText        | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.12.1/widget#ShapeText) |
| 表格                                  | widget           | Table            | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.12.1/widget#Table) |
| 按钮                                  | widget           | Button           | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.12.1/widget#Button) |
| 下拉组合框                            | widget           | ComboBox         | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.12.1/widget#ComboBox) |
| 编辑框(常规, 富文本, 聊天气泡)        | widget           | Edit             | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.12.1/widget#Edit) |
| 代码编辑框                            | widget           | Editor           | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.12.1/widget#Editor) |
| 基础元素                              | widget           | Element          | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.12.1/widget#Element) |
| 列表                                  | widget           | List             | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.12.1/widget#List) |
| 列表框                                | widget           | ListBox          | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.12.1/widget#ListBox) |
| 弹出菜单                              | widget           | Menu             | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.12.1/widget#Menu) |
| 进度条                                | widget           | ProgressBar      | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.12.1/widget#ProgressBar) |
| 静态文本连接按钮                      | widget           | TextLink         | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.12.1/widget#TextLink) |
| 布局元素                              | widget           | LayoutEle        | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.12.1/widget#LayoutEle) |
| 布局框架                              | widget           | LayoutFrame      | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.12.1/widget#LayoutFrame) |
| 列表视图                              | widget           | ListView         | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.12.1/widget#ListView) |
| 菜单条                                | widget           | MenuBar          | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.12.1/widget#MenuBar) |
| Pane元素                              | widget           | Pane             | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.12.1/widget#Pane) |
| 滚动条                                | widget           | ScrollBar        | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.12.1/widget#ScrollBar) |
| 滚动视图                              | widget           | ScrollView       | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.12.1/widget#ScrollView) |
| 滑动条元素                            | widget           | SliderBar        | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.12.1/widget#SliderBar) |
| 标签栏元素                            | widget           | TabBar           | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.12.1/widget#TabBar) |
| 工具条                                | widget           | ToolBar          | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.12.1/widget#ToolBar) |
| 列表树元素                            | widget           | Tree             | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.12.1/widget#Tree) |
| 日期时间                              | widget           | DateTime         | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.12.1/widget#DateTime) |
| 月历卡片                              | widget           | MonthCal         | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.12.1/widget#MonthCal) |
| 数据适配器-列表视元素                 | adapter          | AdapterListView  | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.12.1/adapter#AdapterListView) |
| 数据适配器-单列Map-列表头(listHeader) | adapter          | AdapterMap       | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.12.1/adapter#AdapterMap) |
| 数据适配器-XList-XListBox             | adapter          | AdapterTable     | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.12.1/adapter#AdapterTable) |
| 数据适配器-树元素                     | adapter          | AdapterTree      | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.12.1/adapter#AdapterTree) |
| 背景管理器                            | bkmanager        | BkManager        | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.12.1/bkmanager#BkManager) |
| 字体                                  | font             | Font             | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.12.1/font#Font) |
| 图片操作                              | imagex           | Image            | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.12.1/imagex#Image) |
| 列表项模板                            | listitemtemplate | ListItemTemplate | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.12.1/listitemtemplate#ListItemTemplate) |
| 节点                                  | listitemtemplate | Node             | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.12.1/listitemtemplate#Node) |
| 图形绘制                              | drawx            | Draw             | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.12.1/drawx#Draw) |
| XCGUI所有函数                         | xc               |                  | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.12.1/xc#section-documentation) |
| XCGUI常量                             | xcc              |                  | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.12.1/xcc)  |
| 缓动                                  | ease             |                  | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.12.1/ease) |
| 资源操作                              | res              |                  | √            | [文档](https://pkg.go.dev/github.com/twgh/xcgui@v1.3.12.1/res)  |

