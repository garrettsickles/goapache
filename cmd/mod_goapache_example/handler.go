package main

/*
#include <httpd.h>
*/
import "C"

import (
	"fmt"
	"unsafe"

	"github.com/garrettsickles/goapache"
)

//export handler
func handler(rec *C.request_rec) C.int {
	var request goapache.Request
	request.RequestRec = (uintptr)(unsafe.Pointer(rec))

	body := goapache.ReadBody(request)

	fmt.Println(body)

	return 0
}
