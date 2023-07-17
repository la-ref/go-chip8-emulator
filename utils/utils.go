package utils

import (
	"os"
	"reflect"
)

func BytesToRGBA(color uint32) (uint8, uint8, uint8, uint8) {
	r := uint8((color >> 24) & 0xFF)
	g := uint8((color >> 16) & 0xFF)
	b := uint8((color >> 8) & 0xFF)
	a := uint8(color & 0xFF)
	return r, g, b, a
}

func ReadFile(filePath string) []byte {
	data, err := os.ReadFile(filePath)
	if err != nil {
		panic("Error opening file")
	}
	return data
}

func Sizeof(s any) int {
	v := reflect.ValueOf(s)
	return v.Len() * int(v.Type().Elem().Size())
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
