package main

//#include <stdlib.h>
//#include <stdint.h>
import "C"
import "unsafe"

//export Free
func Free(ptr unsafe.Pointer) {
	C.free(ptr)
}

//export StrLen
func StrLen(sp *C.char) C.size_t {
	return C.size_t(len(C.GoString(sp)))
}
