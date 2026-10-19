// Package types 存放基础结构体, 有多个包共用的, 比如 xc 包和 wapi 包会用到。
package types

type RECT struct {
	Left, Top, Right, Bottom int32
}

type RECTF struct {
	Left, Top, Right, Bottom float32
}

type POINT struct {
	X, Y int32
}

type POINTF struct {
	X, Y float32
}

type SIZE struct {
	CX, CY int32
}
