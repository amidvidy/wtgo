package wtgo

/*
 #cgo LDFLAGS: -lwiredtiger
 #include <wiredtiger.h>
 #include <stdlib.h>

int wt_session_close(WT_SESSION* sess,
                     const char* config) {
    return sess->close(sess, config);
}

int wt_session_open_cursor(WT_SESSION* sess,
                           const char* uri,
                           WT_CURSOR* to_dup,
                           const char* config,
                           struct __wt_cursor** cursorp) {
    return sess->open_cursor(sess, uri, to_dup, config, cursorp);
}

int wt_session_create(WT_SESSION* sess,
                      const char* name,
                      const char* config) {
    return sess->create(sess, name, config);
}

int wt_session_begin_transaction(WT_SESSION* sess,
                                 const char* config) {
    return sess->begin_transaction(sess, config);
}

int wt_session_commit_transaction(WT_SESSION* sess,
                                  const char* config) {
    return sess->commit_transaction(sess, config);
}

int wt_session_rollback_transaction(WT_SESSION* sess,
                                    const char* config) {
    return sess->rollback_transaction(sess, config);
}
*/
import "C"

import (
	"unsafe"
)

type Session struct {
	WTSession *C.WT_SESSION
}

func (s *Session) Close() error {
	wterr := C.wt_session_close(s.WTSession, nil)
	if wterr != 0 {
		return &WTError{wterr}
	}
	return nil
}

func (s *Session) OpenCursor(uri, config string) (*Cursor, error) {
	wturi := C.CString(uri)
	wtconfig := C.CString(config)

	defer C.free(unsafe.Pointer(wturi))
	defer C.free(unsafe.Pointer(wtconfig))

	var wtcursor *C.WT_CURSOR

	wterr := C.wt_session_open_cursor(s.WTSession, wturi, nil, wtconfig, &wtcursor)
	if wterr != 0 {
		return nil, &WTError{wterr}
	}
	return &Cursor{wtcursor}, nil
}

func (s *Session) Create(name, config string) error {
	wtname := C.CString(name)
	wtconfig := C.CString(config)

	defer C.free(unsafe.Pointer(wtname))
	defer C.free(unsafe.Pointer(wtconfig))

	wterr := C.wt_session_create(s.WTSession, wtname, wtconfig)
	if wterr != 0 {
		return &WTError{wterr}
	}
	return nil
}

func (s *Session) BeginTransaction(config string) error {
	wtconfig := C.CString(config)

	defer C.free(unsafe.Pointer(wtconfig))

	wterr := C.wt_session_begin_transaction(s.WTSession, wtconfig)
	if wterr != 0 {
		return &WTError{wterr}
	}
	return nil
}

func (s *Session) CommitTransaction(config string) error {
	wtconfig := C.CString(config)

	defer C.free(unsafe.Pointer(wtconfig))

	wterr := C.wt_session_commit_transaction(s.WTSession, wtconfig)
	if wterr != 0 {
		return &WTError{wterr}
	}
	return nil
}

func (s *Session) RollbackTransaction(config string) error {
	wtconfig := C.CString(config)

	defer C.free(unsafe.Pointer(wtconfig))

	wterr := C.wt_session_rollback_transaction(s.WTSession, wtconfig)
	if wterr != 0 {
		return &WTError{wterr}
	}
	return nil
}
