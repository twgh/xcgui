package app

import (
	"github.com/twgh/xcgui/imagex"
)

// 图片_加载从图片源, 失败返回 nil.
//
// hImageSrc: 图片源句柄.
func NewImageBySrc(hImageSrc int) *imagex.Image {
	return imagex.NewBySrc(hImageSrc)
}

// 图片_加载从文件, 失败返回 nil.
//
// fileName: 图片文件.
func NewImageByFile(fileName string) *imagex.Image {
	return imagex.NewByFile(fileName)
}

// 图片_加载从文件自适应, 加载图片从文件, 自适应图片, 失败返回 nil.
//
// fileName: 图片文件.
//
// leftSize: 坐标.
//
// topSize: 坐标.
//
// rightSize: 坐标.
//
// bottomSize: 坐标.
func NewImageByFileAdaptive(fileName string, leftSize, topSize, rightSize, bottomSize int32) *imagex.Image {
	return imagex.NewByFileAdaptive(fileName, leftSize, topSize, rightSize, bottomSize)
}

// 图片_加载从文件指定区域, 加载图片, 指定区位置及大小, 失败返回 nil.
//
// fileName: 图片文件.
//
// x: 坐标.
//
// y: 坐标.
//
// cx: 宽度.
//
// cy: 高度.
func NewImageByFileRect(fileName string, x, y, cx, cy int32) *imagex.Image {
	return imagex.NewByFileRect(fileName, x, y, cx, cy)
}

// 图片_加载从资源自适应, 加载图片从资源, 自适应图片, 失败返回 nil.
//
// id: 资源ID.
//
// Type: 资源类型.
//
// leftSize: 坐标.
//
// topSize: 坐标.
//
// rightSize: 坐标.
//
// bottomSize: 坐标.
//
// hModule:	从指定模块加载, 例如:DLL, EXE; 如果为空, 从当前EXE加载.
func NewImageByResAdaptive(id int32, Type string, leftSize, topSize, rightSize, bottomSize int32, hModule uintptr) *imagex.Image {
	return imagex.NewByResAdaptive(id, Type, leftSize, topSize, rightSize, bottomSize, hModule)
}

// 图片_加载从资源, 失败返回 nil.
//
// id: 资源ID.
//
// Type: 资源类型.
//
// bStretch: 是否拉伸图片.
//
// hModule:	从指定模块加载, 例如:DLL, EXE; 如果为空, 从当前EXE加载.
func NewImageByRes(id int32, Type string, bStretch bool, hModule uintptr) *imagex.Image {
	return imagex.NewByRes(id, Type, bStretch, hModule)
}

// 图片_加载从ZIP, 加载图片从ZIP压缩包, 失败返回 nil.
//
// zipFileName: ZIP压缩包文件名.
//
// fileName: 图片文件名.
//
// password: ZIP压缩包密码, 不填默认为空.
func NewImageByZip(zipFileName string, fileName string, password ...string) *imagex.Image {
	return imagex.NewByZip(zipFileName, fileName, password...)
}

// 图片_加载从资源ZIP, 失败返回 nil.
//
// id: RC资源ID.
//
// fileName: 图片文件名.
//
// password: ZIP压缩包密码.
//
// hModule: 模块句柄, 不填默认为0.
func NewImageByZipRes(id int32, fileName string, password string, hModule ...uintptr) *imagex.Image {
	return imagex.NewByZipRes(id, fileName, password, hModule...)
}

// 图片_加载从ZIP自适应, 加载图片从ZIP压缩包, 自适应图片, 失败返回 nil.
//
// zipFileName: ZIP压缩包文件名.
//
// fileName: 图片文件名.
//
// password: ZIP压缩包密码.
//
// x1: 坐标.
//
// x2: 坐标.
//
// y1: 坐标.
//
// y2: 坐标.
func NewImageByZipAdaptive(zipFileName string, fileName string, password string, x1, x2, y1, y2 int32) *imagex.Image {
	return imagex.NewByZipAdaptive(zipFileName, fileName, password, x1, x2, y1, y2)
}

// 图片_加载从ZIP指定区域, 加载ZIP图片, 指定区位置及大小, 失败返回 nil.
//
// zipFileName: ZIP文件.
//
// fileName: 图片名称.
//
// password: 密码.
//
// x: 坐标.
//
// y: 坐标.
//
// cx: 宽度.
//
// cy: 高度.
func NewImageByZipRect(zipFileName string, fileName string, password string, x, y, cx, cy int32) *imagex.Image {
	return imagex.NewByZipRect(zipFileName, fileName, password, x, y, cx, cy)
}

// 图片_加载从内存ZIP, 失败返回 nil.
//
// data: 图片数据.
//
// fileName: 图片名称.
//
// password: zip压缩包密码, 不填默认为空.
func NewImageByZipMem(data []byte, fileName string, password ...string) *imagex.Image {
	return imagex.NewByZipMem(data, fileName, password...)
}

// 图片_加载从内存, 加载流图片, 失败返回 nil.
//
// pBuffer: 图片数据.
func NewImageByMem(pBuffer []byte) *imagex.Image {
	return imagex.NewByMem(pBuffer)
}

// 图片_加载从内存指定区域, 加载流图片, 指定区位置及大小, 失败返回 nil.
//
// pBuffer: 图片数据.
//
// x: 坐标.
//
// y: 坐标.
//
// cx: 宽度.
//
// cy: 高度.
func NewImageByMemRect(pBuffer []byte, x, y, cx, cy int32) *imagex.Image {
	return imagex.NewByMemRect(pBuffer, x, y, cx, cy)
}

// 图片_加载从内存自适应, 加载流图片压缩包, 自适应图片(九宫格), 失败返回 nil.
//
// pBuffer: 图片数据.
//
// leftSize: 坐标.
//
// topSize: 坐标.
//
// rightSize: 坐标.
//
// bottomSize: 坐标.
func NewImageByMemAdaptive(pBuffer []byte, leftSize, topSize, rightSize, bottomSize int32) *imagex.Image {
	return imagex.NewByMemAdaptive(pBuffer, leftSize, topSize, rightSize, bottomSize)
}

// 图片_加载从Image, 加载图片从GDI+的Image对象, 失败返回 nil.
//
// pImage: GDI图片对象指针Image*.
func NewImageByImage(pImage uintptr) *imagex.Image {
	return imagex.NewByImage(pImage)
}

// 图片_加载文件图标, 加载文件图标, 从一个EXE文件或DLL文件或图标文件; 例如:*.exe文件的图标, 失败返回 nil.
//
// fileName: 文件名.
func NewImageByExtractIcon(fileName string) *imagex.Image {
	return imagex.NewByExtractIcon(fileName)
}

// 图片_加载从HICON, 创建一个炫彩图片句柄, 从一个现有的图标句柄HICON, 失败返回 nil.
//
// hIcon: 图标句柄.
func NewImageByHICON(hIcon uintptr) *imagex.Image {
	return imagex.NewByHICON(hIcon)
}

// 图片_加载从HBITMAP, 创建一个炫彩图片句柄, 从一个现有的位图句柄 HBITMAP, 失败返回 nil.
//
// hBitmap: 位图句柄.
func NewImageByHBITMAP(hBitmap uintptr) *imagex.Image {
	return imagex.NewByHBITMAP(hBitmap)
}

// 图片_加载从SVG, 失败返回 nil.
//
// hSvg: SVG句柄.
func NewImageBySvg(hSvg int) *imagex.Image {
	return imagex.NewBySvg(hSvg)
}

// 图片_加载从SVG文件, 失败返回 nil.
//
// fileName: 文件名.
func NewImageBySvgFile(fileName string) *imagex.Image {
	return imagex.NewBySvgFile(fileName)
}

// 图片_加载从SVG字符串, 失败返回 nil.
//
// str: 字符串.
func NewImageBySvgString(str string) *imagex.Image {
	return imagex.NewBySvgString(str)
}

// 图片_加载从SVG字符串D2D, 用**纯 Go(Direct2D)**把一段 SVG 渲染成图片句柄; 失败返回 nil.
//
// svgText: 整段 SVG(`<svg viewBox="…"><path d="…"/></svg>`), 或裸的 `<path d="…"/>` 片段.
//
// size: 图片边长(**物理像素**, 正方形). 非方形 viewBox 按较长边等比缩放(与图标约定一致).
//
// color: xc.RGBA 颜色, 它只决定 SVG 里 **currentColor 的取值**:
//   - 单色图标(lucide / heroicons / feather): path 全是 currentColor → 整体就是 color;
//   - 多色图标(iconfont 彩色图标): path 自带 fill="#xxx" → 用它自己的色, color 不影响它.
//
// 线宽取 SVG 的 stroke-width(没有则按 lucide 惯例 2), 不另给参数.
//
// 与 DLL 版的 NewImageBySvgString 的区别: 解析与渲染都在纯 Go 侧(svgp + d2d), 图标站的
// 紧凑写法("a.25.25 0 00.407.195")与圆弧都能正确解析, fill="none" 的子路径不会被糊成填充块;
// 而且**不依赖 app.New(true) 的 D2D 初始化** —— GDI 渲染模式(app.New(false))甚至"还没建窗口"
// 时也能出图. 环境里没有 D2D(d2d1.dll)或 SVG 解析失败时返回 0.
//
// 用法:
//
//	hImg := app.NewImageBySvgStringD2D(svgText, 32, xc.RGBA(255, 255, 255, 255))
//	if hImg == 0 { /* 解析失败或渲染失败 */ }
//	xc.XEle_SetBkImage(ele, hImg) // 或任何吃图片句柄的 API
func NewImageBySvgStringD2D(svgText string, size int32, color uint32) *imagex.Image {
	return imagex.NewBySvgStringD2D(svgText, size, color)
}

// 图片_加载从SVG文件D2D, 同 NewImageBySvgStringD2D, 从文件读 SVG(内部 os.ReadFile), 失败返回 nil.
//
// path: SVG 文件的路径.
//
// size: 图片边长(物理像素), 正方形.
//
// color: xc.RGBA 颜色, 只决定 SVG 里 currentColor 的取值(见 NewImageBySvgStringD2D).
func NewImageBySvgFileD2D(path string, size int32, color uint32) *imagex.Image {
	return imagex.NewBySvgFileD2D(path, size, color)
}

// 从句柄创建图片对象, 失败返回 nil.
func NewImageByHandle(handle int) *imagex.Image {
	return imagex.NewByHandle(handle)
}

// 根据资源文件中的 name 创建对象, 失败返回 nil.
//
// name: 资源名称.
func NewImageByName(name string) *imagex.Image {
	return imagex.NewByName(name)
}

// 从指定的资源文件中, 根据 name 创建对象, 失败返回 nil.
//
// fileName: 资源文件名.
//
// name: 资源名称.
func NewImageByNameEx(fileName, name string) *imagex.Image {
	return imagex.NewByNameEx(fileName, name)
}
