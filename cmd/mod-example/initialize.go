package main

/*
#include <httpd.h>
#include <http_config.h>
*/
import "C"

//export initialize
func initialize(pool *C.apr_pool_t, server *C.server_rec) {

}
