package goapache

/*
#cgo CFLAGS: -I/usr/local/opt/httpd/include/httpd
#cgo CFLAGS: -I/usr/include/apache2
#cgo CFLAGS: -I/usr/include/httpd
#cgo CFLAGS: -I/usr/include/apr-1.0
#cgo CFLAGS: -I/usr/include/apr-1
#cgo LDFLAGS: -Wl,-z,relro,-z,now -L/usr/lib64 -lpthread -ldl

#include "httpd.h"
*/
import "C"

import (
	"http"
)

// Request - Wrapper class for apache request_rec
type Request struct {
	RequestRec uintptr

	Handler string
	Method string
}

func GetRquest(rec uintptr) *Request {
	var r = (*C.request_rec)(unsafe.Pointer(rec))

	return &Request(
		rec,
		C.GoString(r.handler),
		C.GoString(r.method)
	)
}
