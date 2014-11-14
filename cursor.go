package wtgo

// #cgo LDFLAGS: -lwiredtiger
// #include <wiredtiger.h>
// #include <stdlib.h>
import "C"

type Cursor struct {
	WTCursor *C.WT_CURSOR
}

func (cur *Cursor) SetKey(key []byte) {
}

func (cur *Cursor) SetValue(value []byte) {
}

func (cur *Cursor) GetKey() []byte {
}

func (cur *Cursor) SetKey() []byte {
}
