package main

// #cgo CFLAGS: -I/usr/local/opt/httpd/include/httpd
/*
#include <httpd.h>
*/
import "C"

//export handler
func handler(rec *C.request_rec) C.int {
    return 0
}