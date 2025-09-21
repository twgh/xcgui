package edge

import (
	"fmt"
	"testing"

	"github.com/twgh/xcgui/xc"
)

func TestNewColor(t *testing.T) {
	color := NewColor(200, 200, 200, 255)
	fmt.Println(color.ToUint32(), xc.RGBA(200, 200, 200, 255))
}
