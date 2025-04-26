package widget

import (
	"fmt"
	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/tf"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
	"testing"
)

func TestListBox_AddEvent_ListBox_Select(t *testing.T) {
	tf.TFunc(func(a *app.App, w *window.Window) {
		// 创建ListBox
		lb := NewListBox(12, 33, 400, 350, w.Handle)
		// 创建数据适配器, 这个必须创建, 存储数据的
		lb.CreateAdapter()
		lb.ShowSBarV(true)
		// 添加项边框颜色
		lb.AddBkBorder(xcc.ListBox_State_Flag_Item_Select_No, xc.RGBA(0, 255, 255, 255), 1)

		for i := 0; i < 15; i++ {
			// 添加行
			lb.AddItemText(fmt.Sprintf("item-%d", i))
		}

		lb.AddEvent_ListBox_Select(func(hEle int, iItem int32, pbHandled *bool) int {
			t.Log("列表框项被选择:", iItem)
			return 0
		})

		lb.AddEvent_ListBox_Temp_Create_End(func(hEle int, pItem *xc.ListBox_Item_, nFlag int32, pbHandled *bool) int {
			if nFlag == 1 { // 新模板实例
				// 给布局元素添加边框颜色
				hLayout := lb.GetTemplateObject(pItem.Index, 0)
				t.Log("项模板里第0个元素是布局元素:", xc.XC_GetObjectType(hLayout) == xcc.XC_ELE_LAYOUT)
				xc.XEle_AddBkBorder(hLayout, xcc.Layout_State_Flag_Full, xc.RGBA(255, 0, 0, 255), 1)

				// 形状文本水平居中
				hShapeText := lb.GetTemplateObject(pItem.Index, 1)
				t.Log("项模板里第1个元素是形状文本:", xc.XC_GetObjectType(hShapeText) == xcc.XC_SHAPE_TEXT)
				xc.XShapeText_SetTextAlign(hShapeText, xcc.TextAlignFlag_Center)
			}
			return 0
		})
	})
}
