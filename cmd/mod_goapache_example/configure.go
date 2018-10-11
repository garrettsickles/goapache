package main

/*
#include <httpd.h>
#include <http_config.h>

//
// Defines to make CGo easier to use
//
typedef const char* const_char_ptr;
#define nullptr NULL
*/
import "C"
import "unsafe"

//export configuration
func configuration(cmd *C.cmd_parms, cfg unsafe.Pointer, arg C.const_char_ptr) C.const_char_ptr {
	//
	// Setup the "configuration" directive
	//
	return C.const_char_ptr(C.nullptr)
}
