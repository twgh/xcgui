package widget

import (
	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/tf"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
	"testing"
)

func TestList_AddEvent_List_Select(t *testing.T) {
	tf.TFunc(func(a *app.App, w *window.Window) {
		list := NewList(20, 40, 500, 350, w.Handle)
		list.CreateAdapters(3)
		// 添加列
		list.AddColumnText(147, "name1", "Column1")
		list.AddColumnText(147, "name2", "Column2")
		list.AddColumnText(147, "name3", "Column3")
		// 添加项边框颜色
		list.AddBkBorder(xcc.List_State_Flag_Item_Select_No, xc.RGBA(0, 255, 255, 255), 1)

		for i := int32(0); i < 10; i++ {
			index := list.InsertRowText(i, "item"+xc.Itoa(i)+"-subitem1")
			list.SetItemText(index, 1, "item"+xc.Itoa(i)+"-subitem2")
			list.SetItemText(index, 2, "item"+xc.Itoa(i)+"-subitem3")
		}

		list.AddEvent_List_Header_Click(func(hEle int, iItem int32, pbHandled *bool) int {
			t.Log("列表头被点击:", iItem)
			return 0
		})

		list.AddEvent_List_Select(func(hEle int, iItem int32, pbHandled *bool) int {
			t.Log("列表项被选中:", iItem)
			return 0
		})

		list.AddEvent_List_Header_Width_Change(func(hEle int, iItem int32, nWidth int32, pbHandled *bool) int {
			t.Log("列表头宽度改变:", iItem, nWidth)
			return 0
		})

		list.AddEvent_List_Temp_Create_End(func(hEle int, pItem *xc.List_Item_, nFlag int32, pbHandled *bool) int {
			if nFlag == 1 { // 新模板实例
				// 取列数
				n := list.GetColumnCount()

				// 给布局元素添加边框颜色
				for i := int32(0); i < n; i++ {
					hLayout := list.GetTemplateObject(pItem.Index, i, 0)
					t.Log("项模板里第0个元素是布局元素:", xc.XC_GetObjectType(hLayout) == xcc.XC_ELE_LAYOUT)
					xc.XEle_AddBkBorder(hLayout, xcc.Layout_State_Flag_Full, xc.RGBA(255, 0, 0, 255), 1)
				}

				// 形状文本水平居中
				for i := int32(0); i < n; i++ {
					hShapeText := list.GetTemplateObject(pItem.Index, i, 1)
					t.Log("项模板里第1个元素是形状文本:", xc.XC_GetObjectType(hShapeText) == xcc.XC_SHAPE_TEXT)
					xc.XShapeText_SetTextAlign(hShapeText, xcc.TextAlignFlag_Center)
				}
			}
			return 0
		})
	})
}
