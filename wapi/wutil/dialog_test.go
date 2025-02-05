package wutil_test

import (
	"fmt"
	"github.com/twgh/xcgui/wapi"
	"github.com/twgh/xcgui/wapi/wutil"
	"github.com/twgh/xcgui/xc"
	"testing"
)

func TestOpenFiles(t *testing.T) {
	fmt.Println(wutil.OpenFiles(0, []string{"Text Files(*txt)", "*.txt", "All Files(*.*)", "*.*"}, ""))
}

func TestOpenDir(t *testing.T) {
	fmt.Println(wutil.OpenDir(0))
}

func TestOpenFile(t *testing.T) {
	fmt.Println(wutil.OpenFile(0, []string{"Text Files(*txt)", "*.txt", "All Files(*.*)", "*.*"}, ""))
}

func TestOpenFileEx(t *testing.T) {
	arr := wutil.OpenFileEx(wutil.OpenFileOption{
		Title:        "打开文件 (最多选择2个)",
		Filters:      []string{"All Files(*.*)", "*.*", "Text Files(*txt)", "*.txt"},
		MaxOpenFiles: 2,
		// 打开多个文件时, 需要填这个
		Flags: wapi.OFN_ALLOWMULTISELECT | wapi.OFN_EXPLORER | wapi.OFN_PATHMUTEXIST,
	})

	if arr == nil && wapi.CommDlgExtendedError() == wapi.FNERR_BUFFERTOOSMALL {
		fmt.Println("提示", "最多只能选择2个文件")
		return
	}

	for i, s := range arr {
		fmt.Printf("第%d个文件: %s\n", i+1, s)
	}
}

func TestSaveFile(t *testing.T) {
	fmt.Println(wutil.SaveFile(0, []string{"Text Files(*txt)", "*.txt", "All Files(*.*)", "*.*"}, "", "默认文件名.txt"))
}

func TestChooseColor(t *testing.T) {
	rgb := wutil.ChooseColor(0)
	fmt.Println("rgb颜色", rgb)
	fmt.Println("炫彩使用的颜色", xc.RGB2RGBA(rgb, 255))
}
