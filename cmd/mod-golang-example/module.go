package main

/*
#cgo CFLAGS: -I/usr/local/opt/httpd/include/httpd
#cgo CFLAGS: -I/usr/include/apache2
#cgo CFLAGS: -I/usr/include/apr-1.0
#cgo LDFLAGS: -L/usr/lib -lapr-1 -laprutil-1

#include <httpd.h>
#include <http_config.h>
#include <apr_pools.h>
#include <apr_hooks.h>

extern const char* configure(cmd_parms*,void*,const char*);
extern void hooks(apr_pool_t*);

static const command_rec directives[2] =
{
	AP_INIT_TAKE1(
		"configuration",
		(const char*(*)())(configure),
		NULL,
		RSRC_CONF,
		"parameter"
	),
	{
		NULL
	}
};

module AP_MODULE_DECLARE_DATA mod_golang_example =
{
	STANDARD20_MODULE_STUFF,
	NULL,
	NULL,
	NULL,
	NULL,
	directives,
	hooks
};
*/
import "C"

func main() {}
