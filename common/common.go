package common

import (
	"syscall"
	"unicode/utf16"
	"unsafe"
)

// StrPtr 将string转换到uintptr.
//	@param s
//	@return uintptr
//
func StrPtr(s string) uintptr {
	if len(s) == 0 {
		return uintptr(0)
	}
	p, _ := syscall.UTF16PtrFromString(s)
	return uintptr(unsafe.Pointer(p))
}

type sliceHeader struct {
	Data uintptr
	Len  int
	Cap  int
}

// UintPtrToString 将uintptr转换到string.
//	@param ptr
//	@return string
//
func UintPtrToString(ptr uintptr) string {
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
//	@param p
//	@return uintptr
//
func Uint16SliceDataPtr(p *[]uint16) uintptr {
	if len(*p) == 0 {
		return uintptr(0)
	}
	return uintptr(unsafe.Pointer(&(*p)[0]))
}

// BoolPtr 将bool转换到uintptr.
//	@param b
//	@return uintptr
//
func BoolPtr(b bool) uintptr {
	if b {
		return uintptr(1)
	}
	return uintptr(0)
}

// Float32Ptr 将float32转换到uintptr.
//	@param f
//	@return uintptr
//
func Float32Ptr(f float32) uintptr {
	return uintptr(*(*uint32)(unsafe.Pointer(&f)))
}

// UintPtrToFloat32 将uintptr转换到float32.
//	@param ptr
//	@return float32
//
func UintPtrToFloat32(ptr uintptr) float32 {
	u := uint32(ptr)
	return *(*float32)(unsafe.Pointer(&u))
}

// ByteSliceDataPtr 将byte[0]指针转换到uintptr.
//	@param b
//	@return uintptr
//
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

// StringToUint16Ptr 返回指向 UTF-8 字符串 s 的 UTF-16 编码的指针，与 syscall.UTF16PtrFromString 不同的是末尾没有添加终止 NUL。
func StringToUint16Ptr(s string) *uint16 {
	return &utf16.Encode([]rune(s))[0]
}

// Uint16SliceToStringSlice 按null字符分割, 把 []uint16 转换到 []string.
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
