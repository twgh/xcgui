package common

import (
	"bytes"
	"errors"
	"image"
	"image/draw"
	"image/gif"
	"image/png"
	"io"
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
