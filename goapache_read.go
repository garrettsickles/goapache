package goapache

/*
#include <http_protocol.h>
*/
import "C"

import (
	"unsafe"
)

// ReadBody - Read the body of an apache request
func ReadBody(r Request) []byte {
	rec := (*C.request_rec)(unsafe.Pointer(r.RequestRec))

	body := make([]byte, 0)

	if C.ap_should_client_block(rec) == 0 {
		var remain int64 = int64(rec.remaining)
		buf := make([]byte, 8192)

		var size int64
		var pos int64

		for length := int64(C.ap_get_client_block(rec, (*C.char)(unsafe.Pointer(&buf[0])), 8192)); length > 0; {
			if (pos + length) > remain {
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
