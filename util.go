package enet

import (
	"reflect"
	"unsafe"
)

func NewEnetUint8(v uint8) Enet_uint8 {
	return SwigcptrEnet_uint8(unsafe.Pointer(&v))
}

func NewEnetUint32(v uint32) Enet_uint32 {
	return SwigcptrEnet_uint32(unsafe.Pointer(&v))
}

func NewEnetUint16(v uint16) Enet_uint16 {
	return SwigcptrEnet_uint16(unsafe.Pointer(&v))
}

func BytesToUintptr(b []byte) (uintptr, int) {
	header := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	return header.Data, header.Len
}

func Uint32BytesToUintptr(u []uint32) (uintptr, int) {
	header := (*reflect.SliceHeader)(unsafe.Pointer(&u))
	return header.Data, header.Len
}

func UintptrToBytes(ptr uintptr, length int) []byte {
	var b []byte
	header := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	header.Len = length
	header.Cap = length
	header.Data = ptr
	return b
}
