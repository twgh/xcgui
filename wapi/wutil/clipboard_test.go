package wutil_test

import (
	"fmt"
	"github.com/twgh/xcgui/wapi/wutil"
	"testing"
)

func TestSetClipboardText(t *testing.T) {
	err := wutil.SetClipboardText("SetClipboardText")
	if err != nil {
		fmt.Println(err)
	}
	text, err := wutil.GetClipboardText()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(text)
}
