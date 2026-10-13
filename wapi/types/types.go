// Package types 存放基础结构体。
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
