package goapache

/*
#include <http_protocol.h>
#include <apr_strings.h>
*/
import "C"
import (
	"errors"
	"fmt"
	"unsafe"
)

type Query struct {
	// scheme ("http"/"ftp"/...)
	Scheme string `json:"scheme"`

	// combined [user[:password]\@]host[:port]
	HostInfo string `json:"hostInfo"`

	// user name, as in http://user:passwd\@host:port/
	User string `json:"user"`

	// password, as in http://user:passwd\@host:port/
	Password string `json:"password"`

	// hostname from URI (or from Host: header)
	HostName string `json:"hostName"`

	// port string (integer representation is in "port")
	PortLiteral string `json:"portLiteral"`

	// the request path (or NULL if only scheme://host was given)
	Path string `json:"path"`

	// Everything after a '?' in the path, if present
	Query string `json:"query"`

	// Trailing "#fragment" string, if present
	Fragment string `json:"fragment"`

	// The port number, numeric, valid only if port_str != NULL
	PortNumber uint16 `json:"portNumber"`
}

// Request - Wrapper class for apache request_rec
type Request struct {
	rec *C.request_rec `json:"-"`

	// Protocol version number of protocol; 1.1 = 1001
	ProtocolNumber int `json:"protocolNumber"`

	// Protocol string, as given to us, or HTTP/0.9
	Protocol string `json:"protocol"`

	// Host, as set by full URI or Host
	HostName string `json:"hostName"`

	// Time when the request started
	RequestTime int64 `json:"requestTime"`

	// Request method (eg. GET, HEAD, POST, etc.)
	Method string `json:"method"`

	// The "real" content length
	ContentLength int64 `json:"contentLength"`

	// sending chunked transfer-coding
	Chunked int `json:"chunked"`

	// The handler string that we use to call a handler function
	Handler string `json:"handler"`

	// The URI without any parsing performed
	UnparsedURI string `json:"unparsedUri"`

	// The path portion of the URI, or "/" if no path provided
	URI string `json:"uri"`

	// The filename on disk corresponding to this response
	Filename string `json:"filename"`

	// The true filename stored in the filesystem, as in the true alpha case
	// and alias correction, e.g. "Image.jpeg" not "IMAGE$1.JPE" on Windows.
	// The core map_to_storage canonicalizes r->filename when they mismatch
	FilenameCanonical string `json:"filenameCanonical"`

	// The PATH_INFO extracted from this request
	PathInfo string `json:"pathInfo"`

	// The QUERY_ARGS extracted from this request
	QueryArgs string `json:"queryArgs"`

	// The translated QUERY_ARGS from the parsed uri
	ParsedURI *Query `json:"parsedURI"`
}

// NewRequest - Populate a request object from an apache request_rec
func NewRequest(rptr uintptr) *Request {
	rec := (*C.request_rec)(unsafe.Pointer(rptr))

	ret := int(C.apr_uri_parse(rec.pool, rec.uri, &rec.parsed_uri))

	fmt.Println(ret)

	return &Request{
		rec,
		int(rec.proto_num),
		C.GoString(rec.protocol),
		C.GoString(rec.hostname),
		int64(rec.request_time),
		C.GoString(rec.method),
		int64(rec.clength),
		int(rec.chunked),
		C.GoString(rec.handler),
		C.GoString(rec.unparsed_uri),
		C.GoString(rec.uri),
		C.GoString(rec.filename),
		C.GoString(rec.canonical_filename),
		C.GoString(rec.path_info),
		C.GoString(rec.args),
		&Query{
			C.GoString(rec.parsed_uri.scheme),
			C.GoString(rec.parsed_uri.hostinfo),
			C.GoString(rec.parsed_uri.user),
			C.GoString(rec.parsed_uri.password),
			C.GoString(rec.parsed_uri.hostname),
			C.GoString(rec.parsed_uri.port_str),
			C.GoString(rec.parsed_uri.path),
			C.GoString(rec.parsed_uri.query),
			C.GoString(rec.parsed_uri.fragment),
			uint16(rec.parsed_uri.port),
		},
	}
}

// SetupClientBlockNoBody - Message should not have a body
func (r *Request) SetupClientBlockNoBody() error {
	ret := int(C.ap_setup_client_block(r.rec, C.REQUEST_NO_BODY))
	if ret == 0 {
		return nil
	}
	return errors.New("Failed to call 'ap_setup_client_block(...,REQUEST_NO_BODY)'")
}

// SetupClientBlockNotChunked - Body without content length
func (r *Request) SetupClientBlockNotChunked() error {
	ret := int(C.ap_setup_client_block(r.rec, C.REQUEST_CHUNKED_ERROR))
	if ret == 0 {
		return nil
	}
	return errors.New("Failed to call 'ap_setup_client_block(...,REQUEST_CHUNKED_ERROR)'")
}

// SetupClientBlockDechunk - If chunked, remove the chunks for me
func (r *Request) SetupClientBlockDechunk() error {
	ret := int(C.ap_setup_client_block(r.rec, C.REQUEST_CHUNKED_DECHUNK))
	if ret == 0 {
		return nil
	}
	return errors.New("Failed to call 'ap_setup_client_block(...,REQUEST_CHUNKED_DECHUNK)'")
}

// SetContentType - Set the http content type for the response
func (r *Request) SetContentType(contentType string) {
	C.ap_set_content_type(r.rec, getPooledString(contentType, r.rec.pool))
}

// SetStatusCode - Set the http status code for the response
func (r *Request) SetStatusCode(statusCode int) {
	r.rec.status = C.int(statusCode)
}

// Respond - Write a response from onto the apache request object
func (r *Request) Respond(data []byte) (int, error) {
	C.ap_set_content_length(r.rec, C.apr_off_t(len(data)))
	bytesSent := C.ap_rwrite(unsafe.Pointer(&data[0]), C.int(len(data)), r.rec)
	return int(bytesSent), nil
}

// ReadBody - Read the body of an apache request from a request object
func (r *Request) ReadBody(blockSize int) ([]byte, error) {
	if blockSize == 0 {
		blockSize = C.HUGE_STRING_LEN
	}

	body := make([]byte, 0)

	if C.ap_should_client_block(r.rec) == 0 {
		remain := int64(r.rec.remaining)
		buffer := make([]byte, 8192)
		bufCPtr := (*C.char)(unsafe.Pointer(&buffer[0]))

		var size int64
		var pos int64

		for length := int64(C.ap_get_client_block(r.rec, bufCPtr, 8192)); length > 0; {
			if (pos + length) > remain {
				size = remain - pos
			} else {
				size = length
			}

			body = append(body, buffer[:size]...)
			pos = pos + size
		}
	}

	return body, nil
}

func getPooledString(str string, pool *C.apr_pool_t) *C.char {
	raw := C.CString(str)
	var pooled *C.char = C.apr_pstrdup(pool, raw)
	C.free(unsafe.Pointer(raw))
	return pooled
}
