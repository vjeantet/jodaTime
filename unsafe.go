package jodaTime

import "unsafe"

// UnsafeString returns the string under byte buffer
func UnsafeString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
