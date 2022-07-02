package drawx

// Deprecated
//
// !已废弃, 请使用 drawx.New().
func NewDraw(hWindow int) *Draw {
	return New(hWindow)
}

// Deprecated
//
// !已废弃, 请使用 drawx.NewGDI().
func NewDrawGDI(hWindow, hdc int) *Draw {
	return NewGDI(hWindow, hdc)
}

// Deprecated
//
// !已废弃, 请使用 drawx.NewByHandle().
func NewDrawByHandle(handle int) *Draw {
	return NewByHandle(handle)
}
