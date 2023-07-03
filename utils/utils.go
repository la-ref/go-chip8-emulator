package utils

func BytesToRGBA(color uint32) (uint8, uint8, uint8, uint8) {
	r := uint8((color >> 24) & 0xFF)
	g := uint8((color >> 16) & 0xFF)
	b := uint8((color >> 8) & 0xFF)
	a := uint8(color & 0xFF)
	return r, g, b, a
}
