package wtgo

/*
 #cgo LDFLAGS: -lwiredtiger
 #include <wiredtiger.h>
 #include <stdlib.h>

int wt_session_close(WT_SESSION* sess,
                     const char* config) {
    return sess->close(conn, config);
}

int wt_session_open_cursor(WT_SESSION* sess,
                           const char* uri,
                           WT_CURSOR* to_dup,
                           const char* config,
                           WT_CURSOR** cursorp) {
    return sesss->open_cursor(session, uri, to_dup, config, cursorp);
}
*/
import "C"

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
	wtconfig := C.Cstring(config)

	defer C.free(unsafe.Pointer(wturi))
	defer C.free(unsafe.Pointer(wtconfig))

	var wtcursor *C.WT_CURSOR

	wterr := C.wt_session_open_cursor(s.WTSession, wturi, nil, wtconfig, &wtcursor)
	if wterr != 0 {
		return nil, &WTError{wterr}
	}
	return &Cursor{wtsess}, nil
}
