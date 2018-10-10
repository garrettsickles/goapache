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
	request := goapache.ParseRequest((uintptr)(unsafe.Pointer(rec)))

	jsonBody, _ := json.Marshal(request)

	fmt.Println(string(jsonBody))

	//body := goapache.ReadBody(request)
	//fmt.Println(body)

	goapache.WriteResponse(request, "application/json", 200, jsonBody)

	return C.OK
}
