package d2d

import "github.com/twgh/xcgui/wapi/com"

// COM 基础件统一来自 **wapi/com**(ComProc / IUnknown / IUnknownVtbl): 本包几十个虚表字段的类型名
// 靠下面这组类型别名保持简短(写 ComProc 即可), 而实际类型仍属 com 包 —— 需要跨包传递时不用转换.
//
// 注意别名只影响"类型名怎么读", 不改变语义: ComProc.Call 的 uintptrescapes 注解、IUnknown 的虚表
// 指针布局与 nil 安全语义都在 com 包里(见该包文档).
type (
	// ComProc 是一个 COM 过程(虚表里的函数指针).
	ComProc = com.ComProc
	// IUnknown 是 COM 对象的映射(IUnknown 部分).
	IUnknown = com.IUnknown
	// IUnknownVtbl 是 IUnknown 的虚表.
	IUnknownVtbl = com.IUnknownVtbl
)
