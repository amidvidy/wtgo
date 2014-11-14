package wtgo

// Much of design inspired by @jmhodges' levigo

/*
 #cgo LDFLAGS: -lwiredtiger
 #include <wiredtiger.h>
 #include <stdlib.h>

// Wrappers for calling WT methods accessed through function pointers
int wt_connection_close(WT_CONNECTION* conn,
                        const char* config) {
    return conn->close(conn, config);
}

int wt_connection_open_session(WT_CONNECTION *conn,
                               WT_EVENT_HANDLER *errhandler,
                               const char* config,
                               WT_SESSION **session) {
    return conn->open_session(conn, errhandler, config, session);
}

*/
import "C"

import (
	"unsafe"
)

type WTError struct {
	code C.int
}

func (e *WTError) Error() string {
	errstr := C.wiredtiger_strerror(e.code)
	defer C.free(unsafe.Pointer(errstr))
	return C.GoString(errstr)
}

type Connection struct {
	WTConnection *C.WT_CONNECTION
}

// TODO: better API for all configuration options

func Open(home, config string) (*Connection, error) {
	wthome := C.CString(home)
	wtconfig := C.CString(config)

	defer C.free(unsafe.Pointer(wthome))
	defer C.free(unsafe.Pointer(wtconfig))

	var wtconn *C.WT_CONNECTION

	wterr := C.wiredtiger_open(wthome, nil, wtconfig, &wtconn)
	if wterr != 0 {
		return nil, &WTError{wterr}
	}
	return &Connection{wtconn}, nil
}

func (conn *Connection) Close() error {
	wterr := C.wt_connection_close(conn.WTConnection, nil)
	if wterr != 0 {
		return &WTError{wterr}
	}
	return nil
}

func (conn *Connection) OpenSession(config string) (*Session, error) {
	wtconfig := C.CString(config)
	defer C.free(unsafe.Pointer(wtconfig))

	var wtsess *C.WT_SESSION

	wterr := C.wt_connection_open_session(conn.WTConnection, nil, wtconfig, &wtsess)
	if wterr != 0 {
		return nil, &WTError{wterr}
	}
	return &Session{wtsess}, nil
}
