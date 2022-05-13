package common

import (
	"syscall"
	"unsafe"
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

// UintPtrToString 将uintptr转换到string.
//	@param ptr
//	@return string
//
func UintPtrToString(ptr uintptr) string {
	return syscall.UTF16ToString(*(*[]uint16)(unsafe.Pointer(&ptr)))
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
