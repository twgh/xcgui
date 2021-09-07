package xc

// 模板_加载从文件, 列表项模板文件载入, 返回模板句柄
// nType: 模板类型, ListItemTemp_Type_
// pFileName: 文件名.
func XTemp_Load(nType int, pFileName string) int {
	r, _, _ := xTemp_Load.Call(uintptr(nType), strPtr(pFileName))
	return int(r)
}

// 模板_加载从ZIP, 加载列表项模板从zip压缩包中, 返回模板句柄
// nType: 模板类型
// pZipFile: zip文件
// pFileName: 文件名
// pPassword: zip密码
func XTemp_LoadZip(nType int, pZipFile string, pFileName string, pPassword string) int {
	r, _, _ := xTemp_LoadZip.Call(uintptr(nType), strPtr(pZipFile), strPtr(pFileName), strPtr(pPassword))
	return int(r)
}

// 模板_加载从内存ZIP, 加载列表项模板从内存zip压缩包中, 返回模板句柄
// nType: 模板类型, ListItemTemp_Type_
// data: 内存块指针
// length: 内存块大小
// pFileName: 文件名
// pPassword: zip密码
func XTemp_LoadZipMem(nType int, data int, length int, pFileName string, pPassword string) int {
	r, _, _ := xTemp_LoadZipMem.Call(uintptr(nType), uintptr(data), uintptr(length), strPtr(pFileName), strPtr(pPassword))
	return int(r)
}

// 模板_加载从文件扩展, 加载列表项模板从文件
// nType: 模板类型, ListItemTemp_Type_
// pFileName: 文件名
// pOutTemp1: 返回模板句柄1
// pOutTemp2: 返回模板句柄2
func XTemp_LoadEx(nType int, pFileName string, pOutTemp1 int, pOutTemp2 int) bool {
	r, _, _ := xTemp_LoadEx.Call(uintptr(nType), strPtr(pFileName), uintptr(pOutTemp1), uintptr(pOutTemp2))
	return int(r) != 0
}

// 模板_加载从ZIP扩展, 加载列表项模板从zip压缩包中
// nType: 模板类型, ListItemTemp_Type_
// pZipFile: zip文件
// pFileName: 文件名
// pPassword: zip密码
// pOutTemp1: 返回模板句柄1, 项模板
// pOutTemp2: 返回模板句柄2, 列表头模板或列表视组模板
func XTemp_LoadZipEx(nType int, pZipFile string, pFileName string, pPassword string, pOutTemp1 int, pOutTemp2 int) bool {
	r, _, _ := xTemp_LoadZipEx.Call(uintptr(nType), strPtr(pZipFile), strPtr(pFileName), strPtr(pPassword), uintptr(pOutTemp1), uintptr(pOutTemp2))
	return int(r) != 0
}

// 模板_加载从内存ZIP扩展, 加载列表项模板从内存zip压缩包中
// nType: 模板类型, ListItemTemp_Type_
// data: 内存块指针
// length: 内存块大小, 字节为单位
// pFileName: 文件名
// pPassword: zip密码
// pOutTemp1: 返回模板句柄1, 项模板
// pOutTemp2: 返回模板句柄2, 列表头模板或列表视组模板
func XTemp_LoadZipMemEx(nType int, data int, length int, pFileName string, pPassword string, pOutTemp1 int, pOutTemp2 int) bool {
	r, _, _ := xTemp_LoadZipMemEx.Call(uintptr(nType), uintptr(data), uintptr(length), strPtr(pFileName), strPtr(pPassword), uintptr(pOutTemp1), uintptr(pOutTemp2))
	return int(r) != 0
}

// 模板_加载从字符串, 加载列表项模板文件从内存字符串
// nType: 模板类型, ListItemTemp_Type_
// pStringXML: 字符串指针.
func XTemp_LoadFromString(nType int, pStringXML int) int {
	r, _, _ := xTemp_LoadFromString.Call(uintptr(nType), uintptr(pStringXML))
	return int(r)
}

// 模板_加载从字符串扩展, 加载列表项模板文件从内存字符串
// nType: 模板类型, ListItemTemp_Type_
// pStringXML: 字符串内容
// pOutTemp1: 返回模板句柄1
// pOutTemp2: 返回模板句柄2
func XTemp_LoadFromStringEx(nType int, pStringXML int, pOutTemp1 int, pOutTemp2 int) bool {
	r, _, _ := xTemp_LoadFromStringEx.Call(uintptr(nType), uintptr(pStringXML), uintptr(pOutTemp1), uintptr(pOutTemp2))
	return int(r) != 0
}

// 模板_取类型, 获取列表项模板类型, 返回: ListItemTemp_Type_
// hTemp: 列表项模板句柄.
func XTemp_GetType(hTemp int) int {
	r, _, _ := xTemp_GetType.Call(uintptr(hTemp))
	return int(r)
}

// 模板_销毁, 项模板销毁
// hTemp: 项模板句柄.
func XTemp_Destroy(hTemp int) bool {
	r, _, _ := xTemp_Destroy.Call(uintptr(hTemp))
	return int(r) != 0
}

// 模板_创建, 创建项模板, 返回模板句柄
// nType: 模板类型, ListItemTemp_Type_
func XTemp_Create(nType int) int {
	r, _, _ := xTemp_Create.Call(uintptr(nType))
	return int(r)
}

// 模板_添加根节点
// hTemp: 项模板句柄.
// pNode: 节点指针.
func XTemp_AddNodeRoot(hTemp int, pNode int) bool {
	r, _, _ := xTemp_AddNodeRoot.Call(uintptr(hTemp), uintptr(pNode))
	return int(r) != 0
}

// 模板_添加子节点
// pParentNode: 父节点指针.
// pNode: 节点指针.
func XTemp_AddNode(pParentNode int, pNode int) bool {
	r, _, _ := xTemp_AddNode.Call(uintptr(pParentNode), uintptr(pNode))
	return int(r) != 0
}

// 模板_创建节点
// nType: 对象类型.
func XTemp_CreateNode(nType int) int {
	r, _, _ := xTemp_CreateNode.Call(uintptr(nType))
	return int(r)
}

// 模板_置节点属性
// pNode: 节点指针.
// pName: 属性名.
// pAttr: 属性值.
func XTemp_SetNodeAttribute(pNode int, pName string, pAttr string) bool {
	r, _, _ := xTemp_SetNodeAttribute.Call(uintptr(pNode), strPtr(pName), strPtr(pAttr))
	return int(r) != 0
}

// 模板_置节点属性扩展
// pNode: 节点指针
// itemID: 模板项ID
// pName: 属性名
// pAttr: 属性值
func XTemp_SetNodeAttributeEx(pNode int, itemID int, pName string, pAttr string) bool {
	r, _, _ := xTemp_SetNodeAttributeEx.Call(uintptr(pNode), uintptr(itemID), strPtr(pName), strPtr(pAttr))
	return int(r) != 0
}

// 模板_取列表中的节点
// hTemp: 模板句柄
// index: 节点位置索引
func XTemp_List_GetNode(hTemp int, index int) int {
	r, _, _ := xTemp_List_GetNode.Call(uintptr(hTemp), uintptr(index))
	return int(r)
}

// 模板_取节点, 获取节点, 根据itemID
// pNode: 节点指针
// itemID: ID.
func XTemp_GetNode(pNode int, itemID int) int {
	r, _, _ := xTemp_GetNode.Call(uintptr(pNode), uintptr(itemID))
	return int(r)
}

// 模板_克隆节点, 获取列表项模板类型, 返回: ListItemTemp_Type_
// pNode: 节点指针
func XTemp_CloneNode(pNode int) int {
	r, _, _ := xTemp_CloneNode.Call(uintptr(pNode))
	return int(r)
}
