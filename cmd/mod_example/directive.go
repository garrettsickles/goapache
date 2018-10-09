package main

/*
#include <httpd.h>
#include <http_config.h>

//
// "configuration" - directive
//
#define DIRECTIVE_CONFIGURATION "configuration"
extern const char* configuration(cmd_parms*,void*,const char*);

//
// Directive(s) List
//
static const command_rec directives[2] =
{
	AP_INIT_TAKE1(
		DIRECTIVE_CONFIGURATION,
		(const char*(*)())(configuration),
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
func configuration(cmd *C.cmd_parms, cfg unsafe.Pointer, arg *C.char) {
	//
	// Setup the "configuration" directive
	//
}
