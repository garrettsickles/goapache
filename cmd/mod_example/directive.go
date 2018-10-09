package main

/*
#include <httpd.h>
#include <http_config.h>

typedef const char* const_char_ptr;
#define nullptr NULL

//
// "configuration" - directive
//
#define DIRECTIVE_CONFIGURATION "configuration"
extern const_char_ptr configuration(cmd_parms*,void*,const_char_ptr);

//
// Directive(s) List
//
static const command_rec directives[2] =
{
	AP_INIT_TAKE1(
		DIRECTIVE_CONFIGURATION,
		(const_char_ptr(*)())(configuration),
		NULL,
		RSRC_CONF,
		"The \"configuration\" directive"
	),
	{
		NULL
	}
};
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
