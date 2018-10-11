package main

/*
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
