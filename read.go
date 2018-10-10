package goapache

/*
#cgo CFLAGS: -I/usr/local/opt/httpd/include/httpd
#cgo CFLAGS: -I/usr/include/apache2
#cgo CFLAGS: -I/usr/include/httpd
#cgo CFLAGS: -I/usr/include/apr-1.0
#cgo CFLAGS: -I/usr/include/apr-1
#cgo LDFLAGS: -shared -Wl,-z,relro,-z,now -L/usr/lib64 -lpthread -ldl

#include "http_protocol.h"
*/
import "C"
import (
	"unsafe"
)

func ReadBody(r *C.request_rec) []byte {
	if (C.ap_should_client_block(r) == 0) {
		var remain C.apr_off_t = r.remaining
		buf := make([]byte, C.HUGE_STRING_LEN)

		var length int64 = 0
		var total int64 = 0
		var pos int64 = 0

		length = C.ap_get_client_block(r, (*C.char)(unsafe.Pointer(&buf[0])), C.HUGE_STRING_LEN)
		for (length > 0) {

		}
		return buf
	}
	return make([]byte, 10)
}