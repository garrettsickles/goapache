package goapache

/*
#cgo CFLAGS: -I/usr/local/opt/httpd/include/httpd
#cgo CFLAGS: -I/usr/include/apache2
#cgo CFLAGS: -I/usr/include/httpd
#cgo CFLAGS: -I/usr/include/apr-1.0
#cgo CFLAGS: -I/usr/include/apr-1
#cgo LDFLAGS: -Wl,-z,relro,-z,now -L/usr/lib64 -lpthread -ldl

#include <http_protocol.h>
*/
import "C"
import "unsafe"

// Request - Wrapper class for apache request_rec
type Request struct {
	RequestRec uintptr `json:"-"`

	ProtoNum       int    `json:protocol_number`
	Chunked        int    `json:chunked`
	ContentLength  int64  `json:content_length`
	Protocol       string `json:protocol`
	Handler        string `json:handler`
	Method         string `json:method`
	UnparsedURI    string `json:unparsed_uri`
	URI            string `json:uri`
}

func ParseRequest(rec uintptr) *Request {
	var r = (*C.request_rec)(unsafe.Pointer(rec))

	return &Request{
		rec,
		int(r.proto_num),
		int(r.chunked),
		int64(r.clength),
		C.GoString(r.protocol),
		C.GoString(r.handler),
		C.GoString(r.method),
		C.GoString(r.unparsed_uri),
		C.GoString(r.uri),
	}
}

func WriteResponse(request *Request, contentType string, code int, data []byte) {
	var r = (*C.request_rec)(unsafe.Pointer(request.RequestRec))

	contentTypeRaw := C.CString(contentType)
	var contentTypeC *C.char = C.ap_make_content_type(r, contentTypeRaw)
	C.free(unsafe.Pointer(contentTypeRaw))
	C.ap_set_content_type(r, contentTypeC)

	C.puts(r.content_type)
	C.puts(contentTypeC)

	C.ap_set_content_length(r, C.long(len(data)))
	C.ap_rwrite(unsafe.Pointer(&data[0]), C.int(len(data)), r)
	r.status = C.int(code)
}
