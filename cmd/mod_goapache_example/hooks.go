package main

/*
#include <httpd.h>
#include <http_config.h>

extern void
initialize(apr_pool_t*,server_rec*);

extern int
handler(request_rec*);

//
// hooks
//    Register hooks for the apache module.
//    This is quit painful to do in Go compared to C-Go
//    so I am leaving it right here for now.
//
//    The functions listed above are the only boiler-plate
//    needed to make this file compile.
//
void hooks(apr_pool_t *pool)
{
	ap_hook_child_init(initialize, NULL, NULL, APR_HOOK_FIRST);
	ap_hook_handler(handler, NULL, NULL, APR_HOOK_MIDDLE);
}
*/
import "C"
