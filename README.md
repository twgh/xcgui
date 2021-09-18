# XCGUI

[English](./README-en.md) | 简体中文

DirectUI设计思想: 在窗口内没有子窗口，界面元素都是逻辑上的区域(无HWND句柄,安全,灵活), 所有UI元素都是自主开发(不受系统限制),  更加灵活的实现各种程序界面,满足不同用户的需求.

官方网站：[www.xcgui.com](http://www.xcgui.com "xcgui 官方网站")

# 可视化UI设计器

使用UI设计器可以快速设计界面，节省大量代码。

[![uidesigner](https://z3.ax1x.com/2021/09/15/4Vmh9S.png)](https://github.com/twgh/xcgui-example/blob/main/uidesigner/uidesigner.png)

只有这么多代码：

```go
package main

import (
	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xcc"
)

func main() {
	a := app.New("")
	// 添加文件搜索路径, 你运行时需要改成自己的路径
	a.AddFileSearchPath(`D:\GoProject\src\github.com\twgh\xcgui\example\uidesigner\res`)
	// 从zip中加载资源文件
	a.LoadResourceZip("qqmusic.zip", "resource.res", "")
	// 从zip中加载布局文件
	hWindow := a.LoadLayoutZip("qqmusic.zip", "main.xml", "", 0)

	win := &window.Window{}
	// 给类句柄赋值
	win.SetHandle(hWindow)
	// 调整布局
	win.AdjustLayout()
	// 显示窗口
	win.ShowWindow(xcc.SW_SHOW)

	a.Run()
	a.Exit()
}
```

# 获取

```go
go get github.com/twgh/xcgui
```

# 项目文档

[项目文档](https://pkg.go.dev/github.com/twgh/xcgui)        [chm帮助文档](https://github.com/twgh/xcgui-example/blob/main/help/%E7%82%AB%E5%BD%A9%E7%95%8C%E9%9D%A2%E5%BA%93-%E5%B8%AE%E5%8A%A9%E6%96%87%E6%A1%A3(v3.0)-(2021-09-08).chm)

# 动态链接库下载

[xcgui.dll(x64)](https://github.com/twgh/xcgui-example/blob/main/help/x64/XCGUI.dll)        [xcgui.dll(x86)](https://github.com/twgh/xcgui-example/blob/main/help/x86/XCGUI.dll)

程序运行时需要把"XCGUI.dll"放到程序运行目录。

在开发时最好是放到C:\Windows\System32目录，这样就不需要把dll放到程序运行目录了。

# 简单窗口（纯代码）

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
	btn_Min := widget.NewButton(396, 10, 30, 30, "-", win.HWindow)
	btn_Min.SetTextColor(xc.RGB(255, 255, 255), 255)
	btn_Min.SetType(xcc.Button_Type_Min)
	btn_Min.EnableBkTransparent(true)
	// 创建结束按钮
	btn_Close := widget.NewButton(426, 10, 30, 30, "X", win.HWindow)
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

# 更多例子

[xcgui-example](https://github.com/twgh/xcgui-example)

# 常量

xcc包里都是常量，像这样使用：`xcc.Xc_Window_Style_Default`

# 命令介绍

xc包里的函数都是xcgui.dll里原本的函数，有一千多个函数，可以直接使用。

封装好的类都在其他包里。

# 封装进度

这些类都是基于xc包里的一千多个函数封装的，当然你也可以选择直接使用xc包里的函数。

| 中文名称                              | 包名      | 类名             | 是否封装完毕 |
| ------------------------------------- | --------- | ---------------- | ------------ |
| 程序                                  | app       | App              | √            |
| 窗口                                  | window    | Window           | √            |
| 框架窗口                              | window    | FrameWindow      | √            |
| 模态窗口                              | window    | ModalWindow      | √            |
| 形状对象                              | shape     | Shape            | √            |
| 圆形形状对象                          | shape     | ShapeEllipse     | √            |
| 形状对象GIF                           | shape     | ShapeGif         | √            |
| 组框形状对象                          | shape     | ShapeGroupBox    | √            |
| 直线形状对象                          | shape     | ShapeLine        | √            |
| 形状对象图片                          | shape     | ShapePicture     | √            |
| 矩形形状对象                          | shape     | ShapeRect        | √            |
| 形状对象文本                          | shape     | ShapeText        | √            |
| 表格                                  | widget    | Table            | √            |
| 按钮                                  | widget    | Button           | √            |
| 下拉组合框                            | widget    | ComboBox         | √            |
| 编辑框(常规, 富文本, 聊天气泡)        | widget    | Edit             | √            |
| 代码编辑框                            | widget    | Editor           | √            |
| 基础元素                              | widget    | Element          | √            |
| 列表                                  | widget    | List             | √            |
| 列表框                                | widget    | ListBox          | √            |
| 弹出菜单                              | widget    | Menu             | √            |
| 进度条                                | widget    | ProgressBar      | √            |
| 静态文本连接按钮                      | widget    | TextLink         | √            |
| 布局元素                              | widget    | LayoutEle        | √            |
| 布局框架                              | widget    | LayoutFrame      | √            |
| 列表视图                              | widget    | ListView         | √            |
| 菜单条                                | widget    | MenuBar          | √            |
| Pane元素                              | widget    | Pane             | √            |
| 滚动条                                | widget    | ScrollBar        | √            |
| 滚动视图                              | widget    | ScrollView       | √            |
| 滑动条元素                            | widget    | SliderBar        | √            |
| 标签栏元素                            | widget    | TabBar           | √            |
| 工具条                                | widget    | ToolBar          | √            |
| 列表树元素                            | widget    | Tree             | √            |
| 日期时间                              | widget    | DateTime         | √            |
| 月历卡片                              | widget    | MonthCal         | √            |
| 数据适配器-列表视元素                 | adapter   | AdapterListView  | √            |
| 数据适配器-单列Map-列表头(listHeader) | adapter   | AdapterMap       | √            |
| 数据适配器-XList-XListBox             | adapter   | AdapterTable     | √            |
| 数据适配器-树元素                     | adapter   | AdapterTree      | √            |
| 背景管理器                            | bkmanager | BkManager        | √            |
| 字体                                  | fontx     | FontX            | √            |
| 图片操作                              | image     | Image            | √            |
| 列表项模板                            | template  | ListItemTemplate | √            |
| 节点                                  | template  | Node             | √            |
| 图形绘制                              | draw      | Draw             | √            |
| XCGUI所有函数                         | xc        |                  | √            |
| XCGUI常量                             | xcc       |                  | √            |
| 缓动                                  | ease      |                  | √            |
| 资源操作                              | res       |                  | √            |

