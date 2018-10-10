# Go-Apache

Create Apache Portable Runtime module using **Go**

### Development

#### `mod_goapache_example`
- [configure.go](cmd/mod_goapache_example/configure.go)
  - Contains directive handlers for processing the `module.conf` configuration
- [directives.go](cmd/mod_goapache_example/directives.go)
  - Exports the `directives` symbol to the shared object. Registers the directive handlers.
- [handler.go](cmd/mod_goapache_example/handler.go)
  - Exports the `handler` symbol to the shared object. This handles incoming requests.
- [hooks.go](cmd/mod_goapache_example/hooks.go)
  - Exports the `hooks` symbol to the shared object. Registers [Apache Hooks](https://httpd.apache.org/docs/2.4/developer/hooks.html) on the module.
- [initialize.go](cmd/mod_goapache_example/initialize.go)
  - Exports the `initialize` symbol to the shared object.
- [module.go](cmd/mod_goapache_example/module.go)
  - Exports the `mod_goapache_example` to the shared object. Referenced by the Apache Module configuration ([RHEL](configs/rhel/02-goapache_example.conf))

#### `package goapache`
