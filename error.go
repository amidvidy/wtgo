package wtgo

// #cgo LDFLAGS: -lwiredtiger
// #include <wiredtiger.h>
// #include <stdlib.h>
import "C"

type WTError struct {
	code C.int
}

func (e *WTError) Error() string {
	errstr := C.wiredtiger_strerror(e.code)
	defer C.free(unsafe.Pointer(errstr))
	return C.GoString(errstr)
}
