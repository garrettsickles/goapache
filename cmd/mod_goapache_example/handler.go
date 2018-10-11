package main

/*
#include <httpd.h>
#include <http_protocol.h>
*/
import "C"

import (
	"encoding/json"
	"fmt"
	"unsafe"

	"github.com/garrettsickles/goapache"
)

//export handler
func handler(rec *C.request_rec) C.int {
	request := goapache.NewRequest((uintptr)(unsafe.Pointer(rec)))

	jsonBody, _ := json.Marshal(request)

	fmt.Println(string(jsonBody))

	//body, _ := request.ReadBody()
	//fmt.Println(body)

	request.SetContentType("application/json")
	request.SetStatusCode(200)
	request.Respond(jsonBody)

	return C.OK
}
