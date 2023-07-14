package utils

import (
	"fmt"
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

func SizeOf(v []any) {
	rv := reflect.ValueOf(v)
	len, cap := sizeof(rv)
	fmt.Printf("%v => len: %d bytes, cap: %d bytes\n", rv.Type(), len, cap)
}

func sizeof(rv reflect.Value) (int, int) {
	rt := rv.Type()

	switch rt.Kind() {
	case reflect.Slice:
		size := int(rt.Size())
		if rv.Len() > 0 {
			l, c := sizeof(rv.Index(0))
			return size + (l * rv.Len()), size + (c * rv.Cap())
		}
	}

	return int(rt.Size()), int(rt.Size())
}
