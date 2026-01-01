package xcc

const (
	// XmlTransparentWindow 是一个透明窗口的窗口布局文件.
	//    - 可用于创建 WebView, 这个窗口是有特殊设计的, 它是透明的, 这是为了规避当窗口类型为[透明窗口阴影]时的一个问题:
	//    - 当窗口最小化或最大化时会有一瞬间漏出 WebView 后面的炫彩窗口, 表现出来是闪烁了一下, 所以设计为透明就看不到后面的窗口了. 只有网页颜色和炫彩窗口颜色相差很大时才会容易看出此问题, 追求细节的话是会注意到的.
	//    - 使用此 xml 创建窗口后, 可调用窗口相关方法修改标题, 宽高等.
	XmlTransparentWindow = `<?xml version="1.0" encoding="UTF-8"?>
<!--炫彩界面库-窗口布局文件-->
<head />
<windowUI bkInfoM="{99:1.9.9;98:1(0);5:2(15)20(1)21(3)26(1)22(33554431)23(1);}" bkInfoM_nameT="0(矩形)" border="0,0,0,0" center="true" className="" content="" rect="20,20,500,500" shadowColor="#80000000" shadowDepth="128" shadowSize="8" showT="true" transparentAlpha="255" transparentFlag="shadow" windowStyle="46"></windowUI>`
)
