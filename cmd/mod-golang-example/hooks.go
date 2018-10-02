package main

/*
#cgo CFLAGS: -I/usr/local/opt/httpd/include/httpd
#cgo CFLAGS: -I/usr/include/apache2

#include <httpd.h>
#include <http_config.h>
#include <apr_pools.h>

extern void initialize(apr_pool_t*,server_rec*);
extern int handler(request_rec*);

void hooks(apr_pool_t *pool)
{
	// ap_hook_child_init(initialize, NULL, NULL, APR_HOOK_FIRST);
	// ap_hook_handler(handler, NULL, NULL, APR_HOOK_MIDDLE);
}

*/
import "C"
