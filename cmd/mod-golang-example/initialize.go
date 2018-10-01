package main

/*
#cgo CFLAGS: -I/usr/local/opt/httpd/include/httpd

#include <httpd.h>
#include <http_config.h>
*/
import "C"

//export initialize
func initialize(pool *C.apr_pool_t, server *C.server_rec) {

}
