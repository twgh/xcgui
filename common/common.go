package common

import (
	"bytes"
	"errors"
	"image"
	"image/draw"
	"image/gif"
	"image/png"
	"io"
	"os"
	"path/filepath"
	"strings"
	"syscall"
	"time"
	"unicode/utf16"
	"unsafe"
)

// Bytes2String 转换[]byte到string.
//
// b: []byte.
func Bytes2String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// String2Bytes 转换string到[]byte.
//
// s: 文本.
func String2Bytes(s string) []byte {
	sh := (*stringHeader)(unsafe.Pointer(&s))
	bh := sliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}

// stringHeader is the runtime representation of a string.
// It cannot be used safely or portably and its representation may
// change in a later release.
// Moreover, the Data field is not sufficient to guarantee the data
// it references will not be garbage collected, so programs must keep
// a separate, correctly typed pointer to the underlying data.
type stringHeader struct {
	Data uintptr
	Len  int
}

// sliceHeader is the runtime representation of a slice.
// It cannot be used safely or portably and its representation may
// change in a later release.
// Moreover, the Data field is not sufficient to guarantee the data
// it references will not be garbage collected, so programs must keep
// a separate, correctly typed pointer to the underlying data.
type sliceHeader struct {
	Data uintptr
	Len  int
	Cap  int
}

// StrPtr 将string转换到uintptr.
//
// s: 文本.
func StrPtr(s string) uintptr {
	if len(s) == 0 {
		return uintptr(0)
	}
	p, _ := syscall.UTF16PtrFromString(s)
	return uintptr(unsafe.Pointer(p))
}

// UintPtrToString 将uintptr转换到string.
//
// ptr: uintptr.
func UintPtrToString(ptr uintptr) string {
	if ptr == 0 {
		return ""
	}
	s := *(*[]uint16)(unsafe.Pointer(&ptr)) // uintptr转换到[]uint16
	for i := 0; i < len(s); i++ {
		if s[i] == 0 {
			(*sliceHeader)(unsafe.Pointer(&s)).Cap = i // 修改切片的cap
			s = s[0:i]
			break
		}
	}
	return string(utf16.Decode(s))
}

// UTF16PtrToString 接收一个指向 UTF-16 切片的指针，并返回对应的 UTF-8 编码字符串。
//   - 如果指针为空，则返回空字符串。该函数假设 UTF-16 序列以零字符结尾；
//   - 如果不存在零字符，程序可能会崩溃。
func UTF16PtrToString(p *uint16) string {
	if p == nil {
		return ""
	}
	if *p == 0 {
		return ""
	}

	// Find NUL terminator.
	n := 0
	for ptr := unsafe.Pointer(p); *(*uint16)(ptr) != 0; n++ {
		ptr = unsafe.Pointer(uintptr(ptr) + unsafe.Sizeof(*p))
	}
	return syscall.UTF16ToString(unsafe.Slice(p, n))
}

// UintPtrToSlice 将uintptr转换到[]interface{}.
//
// ptr: uintptr.
func UintPtrToSlice(ptr uintptr) []interface{} {
	if ptr == 0 {
		return nil
	}
	s := *(*[]interface{})(unsafe.Pointer(&ptr)) // uintptr转换到[]interface{}
	for i := 0; i < len(s); i++ {
		if s[i] == nil {
			(*sliceHeader)(unsafe.Pointer(&s)).Cap = i // 修改切片的cap
			s = s[0:i]
			break
		}
	}
	return s
}

// Uint16SliceDataPtr 将uint16[0]指针转换到uintptr.
//
// p: uint16[0]的指针.
func Uint16SliceDataPtr(p *[]uint16) uintptr {
	if len(*p) == 0 {
		return uintptr(0)
	}
	return uintptr(unsafe.Pointer(&(*p)[0]))
}

// BoolPtr 将bool转换到uintptr.
//
// b: bool.
func BoolPtr(b bool) uintptr {
	if b {
		return uintptr(1)
	}
	return uintptr(0)
}

// Float32Ptr 将float32转换到uintptr.
//
// f: float32.
func Float32Ptr(f float32) uintptr {
	return uintptr(*(*uint32)(unsafe.Pointer(&f)))
}

// UintPtrToFloat32 将uintptr转换到float32.
//
// ptr: uintptr.
func UintPtrToFloat32(ptr uintptr) float32 {
	if ptr == 0 {
		return 0
	}
	u := uint32(ptr)
	return *(*float32)(unsafe.Pointer(&u))
}

// ByteSliceDataPtr 将byte[0]的指针转换到uintptr.
//
// b: byte[0]的指针.
func ByteSliceDataPtr(b *[]byte) uintptr {
	if len(*b) == 0 {
		return uintptr(0)
	}
	return uintptr(unsafe.Pointer(&(*b)[0]))
}

/* // byte指针到uintptr.
func bytePtr(p *byte) uintptr {
	return uintptr(unsafe.Pointer(p))
} */

/* func UintPtrToString2(r uintptr) string {
    p := (*uint16)(unsafe.Pointer(r))
    if p == nil {
        return ""
    }

    n, end, add := 0, unsafe.Pointer(p), unsafe.Sizeof(*p)
    for *(*uint16)(end) != 0 {
        end = unsafe.Add(end, add)
        n++
    }
    return string(utf16.Decode(unsafe.Slice(p, n)))
} */

/*// RuneToUint16Ptr 返回指向 UTF-8 字符串 r 的 UTF-16 编码的指针.
func RuneToUint16Ptr(r []rune) *uint16 {
	return &utf16.Encode(r)[0]
}*/

// StringToUint16Ptr 返回指向 UTF-8 字符串 s 的 UTF-16 编码的指针，与 syscall.UTF16PtrFromString 不同的是末尾没有添加终止 NUL.
//
// s: 文本.
func StringToUint16Ptr(s string) *uint16 {
	return &utf16.Encode([]rune(s))[0]
}

// Uint16SliceToStringSlice 按null字符分割, 把 []uint16 转换到 []string.
//
// s: []uint16.
func Uint16SliceToStringSlice(s []uint16) []string {
	strSlice := make([]string, 0)
	start := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 0 {
			strSlice = append(strSlice, string(utf16.Decode(s[start:i])))
			start = i + 1

			// 连续null字符, 所以跳出
			if s[start] == 0 {
				break
			}
		}
	}
	return strSlice
}

// GifFrame 包含 Gif 每帧的二进制数据和延时信息.
type GifFrame struct {
	ImageData []byte        // PNG 格式的二进制数据
	Delay     time.Duration // 帧延时
}

// ExtractGifFrames 提取 GIF 的所有帧.
func ExtractGifFrames(gifReader io.Reader) ([]GifFrame, error) {
	// 解码 GIF
	gifData, err := gif.DecodeAll(gifReader)
	if err != nil {
		return nil, err
	}

	if len(gifData.Image) == 0 {
		return nil, errors.New("no image in gif")
	}

	// 创建背景画布（使用第一帧的尺寸）
	bounds := gifData.Image[0].Bounds()
	baseImage := image.NewRGBA(bounds)

	var frames []GifFrame
	// 遍历每一帧
	for i, srcFrame := range gifData.Image {
		// 将当前帧绘制到画布上
		draw.Draw(baseImage, bounds, srcFrame, image.Point{}, draw.Over)

		// 将画布编码为 PNG 字节
		var buf bytes.Buffer
		if err := png.Encode(&buf, baseImage); err != nil {
			return nil, err
		}

		// 收集帧数据
		frames = append(frames, GifFrame{
			ImageData: buf.Bytes(),
			Delay:     time.Duration(gifData.Delay[i]) * 10 * time.Millisecond,
		})
	}
	return frames, nil
}

// GetProcessName 取当前进程名.
func GetProcessName() string {
	// 获取可执行文件路径
	var exePath string
	if path, err := os.Executable(); err == nil {
		exePath = path
	} else { // Executable 获取失败时使用 os.Args[0] 作为备份
		exePath = os.Args[0]
	}
	// 提取文件名
	return filepath.Base(exePath)
}

// GetProcessNameWithoutExt 取当前进程名（不含扩展名）.
func GetProcessNameWithoutExt() string {
	filename := GetProcessName()
	// 去除扩展名
	extension := filepath.Ext(filename)
	return strings.TrimSuffix(filename, extension)
}

// BoolToInt 将 bool 类型转换为 int 类型.
//   - true 返回 1, false 返回 0.
func BoolToInt(input bool) int {
	if input {
		return 1
	}
	return 0
}

// BoolToString 将 bool 类型转换为字符串 true 或 false.
func BoolToString(input bool) string {
	if input {
		return "true"
	}
	return "false"
}

// StringToBool 将字符串转换为 bool 类型.
//   - true 和 1 字符串都会转换为 true, 'true'字符串不区分大小写, 其他值都转换为 false.
func StringToBool(input string) bool {
	if strings.EqualFold(input, "true") || input == "1" {
		return true
	}
	return false
}

// ErrorToErrno 将错误转换为系统调用错误号.
//
// err: 需要转换的错误对象
//
// 返回值:
//   - syscall.Errno: 转换后的系统错误号
//   - bool: 是否成功转换的标志
func ErrorToErrno(err error) (syscall.Errno, bool) {
	var errno syscall.Errno
	if errors.As(err, &errno) { // 处理嵌套错误
		return errno, true
	}
	return 0, false
}
