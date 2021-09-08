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
	// 给类窗口句柄赋值
	win.SetHWindow(hWindow)
	// 调整布局
	win.AdjustLayout()
	// 显示窗口
	win.ShowWindow(xcc.SW_SHOW)

	a.Run()
	a.Exit()
}
