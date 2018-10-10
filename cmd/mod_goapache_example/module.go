package main

/*
#cgo CFLAGS: -I/usr/local/opt/httpd/include/httpd
#cgo CFLAGS: -I/usr/include/apache2
#cgo CFLAGS: -I/usr/include/httpd
#cgo CFLAGS: -I/usr/include/apr-1.0
#cgo CFLAGS: -I/usr/include/apr-1
#cgo LDFLAGS: -shared -Wl,-z,relro,-z,now -L/usr/lib64 -lpthread -ldl

#include <httpd.h>
#include <http_config.h>
#include <apr_pools.h>

extern void hooks(apr_pool_t*);
extern const command_rec directives[2];

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
