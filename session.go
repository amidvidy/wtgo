package wtgo

// #cgo LDFLAGS: -lwiredtiger
// #include <wiredtiger.h>
// #include <stdlib.h>
import "C"

type Session struct {
	WTSession *C.WT_SESSION
}
