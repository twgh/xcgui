package common

import (
	"syscall"
	"unicode/utf16"
	"unsafe"
)

const (
	surr1 = 0xd800
	surr2 = 0xdc00
	surr3 = 0xe000
	surrSelf = 0x10000
	replacementChar = '\uFFFD'
)


// StrPtr 将string转换到uintptr.
//	@param s
//	@return uintptr
//
func StrPtr(s string) uintptr {
	if s == "" {
		return uintptr(0)
	}
	p, _ := syscall.UTF16PtrFromString(s)
	return uintptr(unsafe.Pointer(p))
}

func decodeRune(r1, r2 rune) rune {
	if surr1 <= r1 && r1 < surr2 && surr2 <= r2 && r2 < surr3 {
		return (r1-surr1)<<10 | (r2 - surr2) + surrSelf
	}
	return replacementChar
}

func SafeUTF16ToString(s []uint16)string  {
	var sLen int
	for i, ch := range s {
		if ch == 0 {
			sLen = i
			break
		}
	}
	a := make([]rune, sLen)
	n := 0
	for i := 0; i < sLen; i++ {
		switch r := s[i]; {
		case r < surr1, surr3 <= r:
			// normal rune
			a[n] = rune(r)
		case surr1 <= r && r < surr2 && i+1 < sLen &&
			surr2 <= s[i+1] && s[i+1] < surr3:
			// valid surrogate sequence
			a[n] = decodeRune(rune(r), rune(s[i+1]))
			i++
		default:
			// invalid surrogate sequence
			a[n] = replacementChar
		}
		n++
	}
	return string(a[:n])
}

// UintPtrToString 将uintptr转换到string.
//	@param ptr
//	@return string
//
func UintPtrToString(ptr uintptr) string {
	return SafeUTF16ToString(*(*[]uint16)(unsafe.Pointer(&ptr)))
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
