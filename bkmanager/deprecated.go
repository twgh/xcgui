package bkmanager

// Deprecated
//
// !已废弃, 请使用 bkmanager.New()
func NewBkManager() *BkManager {
	return New()
}

// Deprecated
//
// !已废弃, 请使用 bkmanager.NewByHandle()
func NewBkManagerByHandle(handle int) *BkManager {
	return NewByHandle(handle)
}

// Deprecated
//
// !已废弃, 请使用 bkmanager.NewByName()
func NewBkManagerByName(name string) *BkManager {
	return NewByName(name)
}
