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
import "unsafe"

// Request - Wrapper class for apache request_rec
type Request struct {
	RequestRec uintptr

	ProtoNum int `json:protocol_number`
	Protocol string `json:protocol`

	Chunked int `json:chunked`

	ContentLength int64 `json:content_length`

	Handler string `json:handler`
	Method string `json:method`
	UnparsedURI string `json:unparsed_uri`
	URI string `json:uri`

}

func ParseRequest(rec uintptr) *Request {
	var r = (*C.request_rec)(unsafe.Pointer(rec))

	return &Request{
		rec,
		int(r.proto_num),
		C.GoString(r.protocol),
		int(r.chunked),
		int64(r.clength),
		C.GoString(r.handler),
		C.GoString(r.method),
		C.GoString(r.unparsed_uri),
		C.GoString(r.uri),
	}
}

func Respond(request *Request) {

}
