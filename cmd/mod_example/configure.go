package main

/*
#include <httpd.h>
#include <http_config.h>
*/
import "C"
import "unsafe"

//export configure
func configure(cmd *C.cmd_parms, cfg unsafe.Pointer, arg *C.char) {

}
