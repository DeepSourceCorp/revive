// Package main ...
package main

import "unsafe"

// #include <stdlib.h>
import "C"

func main() {
	cStr := C.CString("test")

	C.free(unsafe.Pointer(cStr))
}
