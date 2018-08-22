package chart

// #cgo LDFLAGS: -lcurses
// #include <stdlib.h>
// #include "chart.h"
import "C"
import (
	"unsafe"
)

func drawChart() {

	year := C.int(5)

	ptr := C.malloc(C.sizeof_int * 1024)
	defer C.free(unsafe.Pointer(ptr))

	C.execute((*C.int)(ptr), year)

	// size := C.greet(name, year, ()
	//
	// b := C.GoBytes(ptr, size)
	// fmt.Println(string(b))
}
