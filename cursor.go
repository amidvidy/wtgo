package wtgo

/*
 #cgo LDFLAGS: -lwiredtiger
 #include <wiredtiger.h>
 #include <stdlib.h>

void wt_cursor_set_key(WT_CURSOR* cursor,
                       WT_ITEM* item) {
    return cursor->set_key(cursor, item);
}

void wt_cursor_set_value(WT_CURSOR* cursor,
                         WT_ITEM* item) {
    return cursor->set_value(cursor, item);
}

int wt_cursor_get_key(WT_CURSOR* cursor,
                      WT_ITEM* item) {
    return cursor->get_key(cursor, item);
}

int wt_cursor_get_value(WT_CURSOR* cursor,
                        WT_ITEM* item) {
    return cursor->get_value(cursor, item);
}

int wt_cursor_insert(WT_CURSOR* cursor) {
    return cursor->insert(cursor);
}

int wt_cursor_next(WT_CURSOR* cursor) {
    return cursor->next(cursor);
}

int wt_cursor_prev(WT_CURSOR* cursor) {
    return cursor->prev(cursor);
}

int wt_cursor_reset(WT_CURSOR* cursor) {
    return cursor->reset(cursor);
}
*/
import "C"

import (
	"unsafe"
)

type Cursor struct {
	WTCursor *C.WT_CURSOR
}

func (cur *Cursor) SetKey(key []byte) {
	//var C.WT_ITEM
	wtkey := unsafe.Pointer(&key[0])
	keylen := C.size_t(len(key))

	item := C.WT_ITEM{}

	item.data = wtkey
	item.size = keylen

	C.wt_cursor_set_key(cur.WTCursor, &item)
}

func (cur *Cursor) SetValue(value []byte) {
	wtval := unsafe.Pointer(&value[0])
	vallen := C.size_t(len(value))

	item := C.WT_ITEM{}

	item.data = wtval
	item.size = vallen

	C.wt_cursor_set_value(cur.WTCursor, &item)
}

func (cur *Cursor) GetKey() ([]byte, error) {
	item := C.WT_ITEM{}
	wterr := C.wt_cursor_get_key(cur.WTCursor, &item)

	if wterr != 0 {
		return nil, &WTError{wterr}
	}
	return C.GoBytes(item.data, C.int(item.size)), nil
}

func (cur *Cursor) GetValue() ([]byte, error) {
	item := C.WT_ITEM{}
	wterr := C.wt_cursor_get_value(cur.WTCursor, &item)

	if wterr != 0 {
		return nil, &WTError{wterr}
	}
	return C.GoBytes(item.data, C.int(item.size)), nil
}

func (cur *Cursor) Insert() error {
	wterr := C.wt_cursor_insert(cur.WTCursor)
	if wterr != 0 {
		return &WTError{wterr}
	}
	return nil
}

func (cur *Cursor) Next() error {
	wterr := C.wt_cursor_next(cur.WTCursor)
	if wterr != 0 {
		return &WTError{wterr}
	}
	return nil
}

func (cur *Cursor) Prev() error {
	wterr := C.wt_cursor_prev(cur.WTCursor)
	if wterr != 0 {
		return &WTError{wterr}
	}
	return nil
}

func (cur *Cursor) Reset() error {
	wterr := C.wt_cursor_reset(cur.WTCursor)
	if wterr != 0 {
		return &WTError{wterr}
	}
	return nil
}
