package main

/*
#include <httpd.h>
#include <http_config.h>
#include <apr_pools.h>

////////////////////////////////////////////////////////////////
//                    Exported Symbols
//
// directives[...] (See 'directives.go' for implementation)
//    Contains a NULL terminated array of directive
//    entry-handler pairs for use by the apache module.
//
//    These directives come from the modul's '*****.conf'
//    file in the httpd configuration directory.
//
extern const command_rec directives[2];
//
// initialize(...) (See 'initialize.go' for implementation)
//    The function to be run first at the start of the
//    module. It will run first because it is configured
//    to do so by the hooks(...) function below
//
extern void initialize(apr_pool_t*,server_rec*);
//
// handler(...) (See 'handler.go' for implementation)
//    The entry point for the request. This processes
//    the request_rec from apache.
//
//    This is where all the http-ish thing should occur.
//    For example setting the status code, the reply body,
//    the content length, and so on...
//
extern int handler(request_rec*);
//
// hooks (See 'hooks.go' for implementation)
//    Register hooks for the apache module.
//    This is quit painful to do in Go compared to C-Go
//    so I am leaving it right here for now.
//
//    The functions listed above are the only boiler-plate
//    needed to make this file compile.
//
extern void hooks(apr_pool_t *pool);

////////////////////////////////////////////////////////////////
//                      Implementation
//
// module 'mod_goapache_example'
//
//
module AP_MODULE_DECLARE_DATA mod_goapache_example =
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
