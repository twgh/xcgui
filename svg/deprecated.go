package svg

// Deprecated
//
// !已废弃, 请使用 svg.NewByFile().
func NewSvg_LoadFile(pFileName string) *Svg {
	return NewByFile(pFileName)
}

// Deprecated
//
// !已废弃, 请使用 svg.NewByString().
func NewSvg_LoadString(pString string) *Svg {
	return NewByString(pString)
}

// Deprecated
//
// !已废弃, 请使用 svg.NewByStringW().
func NewSvg_LoadStringW(pString string) *Svg {
	return NewByStringW(pString)
}

// Deprecated
//
// !已废弃, 请使用 svg.NewByStringUtf8().
func NewSvg_LoadStringUtf8(pString string) *Svg {
	return NewByStringUtf8(pString)
}

// Deprecated
//
// !已废弃, 请使用 svg.NewByZip().
func NewSvg_LoadZip(pZipFileName, pFileName, pPassword string) *Svg {
	return NewByZip(pZipFileName, pFileName, pPassword)
}

// Deprecated
//
// !已废弃, 请使用 svg.NewByZipMem().
func NewSvg_LoadZipMem(data []byte, pFileName, pPassword string) *Svg {
	return NewByZipMem(data, pFileName, pPassword)
}

// Deprecated
//
// !已废弃, 请使用 svg.NewByRes().
func NewSvg_LoadRes(id int, pType string, hModule int) *Svg {
	return NewByRes(id, pType, hModule)
}
