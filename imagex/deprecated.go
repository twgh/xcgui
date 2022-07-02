package imagex

// Deprecated
//
// !已废弃, 请使用 imagex.NewBySrc().
func NewImage_LoadSrc(hImageSrc int) *Image {
	return NewBySrc(hImageSrc)
}

// Deprecated
//
// !已废弃, 请使用 imagex.NewByFile().
func NewImage_LoadFile(pFileName string) *Image {
	return NewByFile(pFileName)
}

// Deprecated
//
// !已废弃, 请使用 imagex.NewByFileAdaptive().
func NewImage_LoadFileAdaptive(pFileName string, leftSize int, topSize int, rightSize int, bottomSize int) *Image {
	return NewByFileAdaptive(pFileName, leftSize, topSize, rightSize, bottomSize)
}

// Deprecated
//
// !已废弃, 请使用 imagex.NewByFileRect().
func NewImage_LoadFileRect(pFileName string, x int, y int, cx int, cy int) *Image {
	return NewByFileRect(pFileName, x, y, cx, cy)
}

// Deprecated
//
// !已废弃, 请使用 imagex.NewByResAdaptive().
func NewImage_LoadResAdaptive(id int, pType string, leftSize int, topSize int, rightSize int, bottomSize, hModule int) *Image {
	return NewByResAdaptive(id, pType, leftSize, topSize, rightSize, bottomSize, hModule)
}

// Deprecated
//
// !已废弃, 请使用 imagex.NewByRes().
func NewImage_LoadRes(id int, pType string, bStretch bool, hModule int) *Image {
	return NewByRes(id, pType, bStretch, hModule)
}

// Deprecated
//
// !已废弃, 请使用 imagex.NewByZip().
func NewImage_LoadZip(pZipFileName string, pFileName string, pPassword string) *Image {
	return NewByZip(pZipFileName, pFileName, pPassword)
}

// Deprecated
//
// !已废弃, 请使用 imagex.NewByZipAdaptive().
func NewImage_LoadZipAdaptive(pZipFileName string, pFileName string, pPassword string, x1 int, x2 int, y1 int, y2 int) *Image {
	return NewByZipAdaptive(pZipFileName, pFileName, pPassword, x1, x2, y1, y2)
}

// Deprecated
//
// !已废弃, 请使用 imagex.NewByZipRect().
func NewImage_LoadZipRect(pZipFileName string, pFileName string, pPassword string, x int, y int, cx int, cy int) *Image {
	return NewByZipRect(pZipFileName, pFileName, pPassword, x, y, cx, cy)
}

// Deprecated
//
// !已废弃, 请使用 imagex.NewByZipMem().
func NewImage_LoadZipMem(data []byte, pFileName string, pPassword string) *Image {
	return NewByZipMem(data, pFileName, pPassword)
}

// Deprecated
//
// !已废弃, 请使用 imagex.NewByMem().
func NewImage_LoadMemory(pBuffer []byte) *Image {
	return NewByMem(pBuffer)
}

// Deprecated
//
// !已废弃, 请使用 imagex.NewByMemRect().
func NewImage_LoadMemoryRect(pBuffer []byte, x int, y int, cx int, cy int) *Image {
	return NewByMemRect(pBuffer, x, y, cx, cy)
}

// Deprecated
//
// !已废弃, 请使用 imagex.NewByMemAdaptive().
func NewImage_LoadMemoryAdaptive(pBuffer []byte, leftSize int, topSize int, rightSize int, bottomSize int) *Image {
	return NewByMemAdaptive(pBuffer, leftSize, topSize, rightSize, bottomSize)
}

// Deprecated
//
// !已废弃, 请使用 imagex.NewByImage().
func NewImage_LoadFromImage(pImage int) *Image {
	return NewByImage(pImage)
}

// Deprecated
//
// !已废弃, 请使用 imagex.NewByExtractIcon().
func NewImage_LoadFromExtractIcon(pFileName string) *Image {
	return NewByExtractIcon(pFileName)
}

// Deprecated
//
// !已废弃, 请使用 imagex.NewByHICON().
func NewImage_LoadFromHICON(hIcon int) *Image {
	return NewByHICON(hIcon)
}

// Deprecated
//
// !已废弃, 请使用 imagex.NewByHBITMAP().
func NewImage_LoadFromHBITMAP(hBitmap int) *Image {
	return NewByHBITMAP(hBitmap)
}

// Deprecated
//
// !已废弃, 请使用 imagex.NewBySvg().
func NewImage_LoadSvg(hSvg int) *Image {
	return NewBySvg(hSvg)
}

// Deprecated
//
// !已废弃, 请使用 imagex.NewBySvgFile().
func NewImage_LoadSvgFile(pFileName string) *Image {
	return NewBySvgFile(pFileName)
}

// Deprecated
//
// !已废弃, 请使用 imagex.NewBySvgString().
func NewImage_LoadSvgString(pString string) *Image {
	return NewBySvgString(pString)
}

// Deprecated
//
// !已废弃, 请使用 imagex.NewBySvgStringW().
func NewImage_LoadSvgStringW(pString string) *Image {
	return NewBySvgStringW(pString)
}

// Deprecated
//
// !已废弃, 请使用 imagex.NewBySvgStringUtf8().
func NewImage_LoadSvgStringUtf8(pString string) *Image {
	return NewBySvgStringUtf8(pString)
}

// Deprecated
//
// !已废弃, 请使用 imagex.NewByHandle().
func NewImageByHandle(handle int) *Image {
	return NewByHandle(handle)
}

// Deprecated
//
// !已废弃, 请使用 imagex.NewByName().
func NewImageByName(name string) *Image {
	return NewByName(name)
}

// Deprecated
//
// !已废弃, 请使用 imagex.NewByNameEx().
func NewImageByNameEx(fileName, name string) *Image {
	return NewByNameEx(fileName, name)
}
