# Go-Apache
| Documentation |
|-|
| [![GoDoc](https://godoc.org/github.com/garrettsickles/goapache?status.png)](https://godoc.org/github.com/garrettsickles/goapache) |

A library to build Apache Modules with Go. With examples!

Create Apache Portable Runtime module using **Go**
--
`What is an Apache Module?`
> If you have never heard of an Apache Module, I would do a quick web search to familiarize yourself.

`Why would I want to build an Apache Module with Go?`
> Go is amazing at writing performant webservers however there are some features (security, logging, metrics, etc..) for which it makes more sense to use a proxy server.
> In these cases you are likely going to use **Nginx** or **Apache** to sit infront of your Go app.
>
> If you use Apache, you are going to configure the application to use multiple modules, one of which could be your custom module!

`When should you develop an Apache Module (in Go, or at all)?`
> Apache Modules are highly performant and configurable pieces of a larger webserver.
>
> Apache Modules natively support shared resources, memory management, and a multiprocessing strategy for parallel execution of requests (see [Apache MPM Worker](https://httpd.apache.org/docs/2.4/mod/worker.html)).
### Development
The `goapache` package support handling and processing Apache server request objects.

#### `goapache`
- [goapache_request.go](goapache_request.go)
  - Contains the `type Request struct {...}` and pertinent methods
  
#### `mod_goapache_example`
- [configure.go](cmd/mod_goapache_example/configure.go)
  - Contains directive handlers for processing the `module.conf` configuration
- [directives.go](cmd/mod_goapache_example/directives.go)
  - Exports the `directives` symbol to the shared object.
  - Registers the directive handlers.
- [handler.go](cmd/mod_goapache_example/handler.go)
  - Exports the `handler` symbol to the shared object.
  - Handles incoming Apache requests.
- [hooks.go](cmd/mod_goapache_example/hooks.go)
  - Exports the `hooks` symbol to the shared object.
  - Registers [Apache Hooks](https://httpd.apache.org/docs/2.4/developer/hooks.html) on the module.
- [initialize.go](cmd/mod_goapache_example/initialize.go)
  - Exports the `initialize` symbol to the shared object.
- [module.go](cmd/mod_goapache_example/module.go)
  - Exports the `mod_goapache_example` to the shared object.
  - Referenced by the Apache Module configuration ([RHEL](configs/rhel/02-goapache_example.conf))
