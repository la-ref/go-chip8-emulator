package utils

func BytesToRGBA(color uint32) (uint8, uint8, uint8, uint8) {
	r := uint8((color >> 24) & 0xFF)
	g := uint8((color >> 16) & 0xFF)
	b := uint8((color >> 8) & 0xFF)
	a := uint8(color & 0xFF)
	return r, g, b, a
}

func RGBAToBytes(r, g, b, a uint8) uint32 {
	return (uint32(r) << 24) | (uint32(g) << 16) | (uint32(b) << 8) | uint32(a)
}

func B2i(b bool) uint8 {
	if b {
		return 1
	}
	return 0
}

func I2b(i uint8) bool {
	if i > 0 {
		return true
	}
	return false
}

func ColorLerp(startColor uint32, endColor uint32, percentage float64) uint32 {
	sr, sb, sg, sa := BytesToRGBA(startColor)
	er, eb, eg, ea := BytesToRGBA(endColor)
	ler_r := uint8(float64(sr) + percentage*(float64(er)-float64(sr)))
	ler_g := uint8(float64(sg) + percentage*(float64(eg)-float64(sg)))
	ler_b := uint8(float64(sb) + percentage*(float64(eb)-float64(sb)))
	ler_a := uint8(float64(sa) + percentage*(float64(ea)-float64(sa)))
	return RGBAToBytes(ler_r, ler_g, ler_b, ler_a)
}
