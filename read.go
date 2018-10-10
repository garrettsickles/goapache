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
	
	body := make([]byte, 0)

	if (C.ap_should_client_block(r) == 0) {
		var remain int64 = int64(r.remaining)
		buf := make([]byte, 8192)

		var length int64 = 0
		var size int64 = 0
		var pos int64 = 0

		for length = int64(C.ap_get_client_block(r, (*C.char)(unsafe.Pointer(&buf[0])), 8192)); (length > 0); {
			if ((pos + length) > remain) {
				size = remain - pos
			} else {
				size = length
			}

			body = append(body, buf[:size]...)
			pos = pos + size
		}
	}

	return body
}