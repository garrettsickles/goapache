package main

/*
#include <httpd.h>
#include <http_config.h>

//
// Directive functions (see 'configure.go')
//
extern const char* configuration(cmd_parms*,void*,const char*);

//
// Directive(s) List
//
const command_rec directives[2] =
{
	AP_INIT_TAKE1(
		"configuration",
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
