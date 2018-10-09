package main

/*
#include <httpd.h>
#include <http_config.h>

//
// Configurer functions (defined elsewhere, exported symbols)
//
extern const char* configuration(cmd_parms*,void*,const char*);

//
// Directive(s) List
//
const command_rec directive[2] =
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
