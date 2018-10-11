package main

/*
//
// Mac OS
//

//
//    HTTPD
//      - Should be installed by homebrew
//
#cgo CFLAGS: -I/usr/local/opt/httpd/include/httpd

//
//    Apache2
//      - Should be installed by homebrew
//
#cgo CFLAGS: -I/usr/local/opt/apache2/include/httpd

//
//    (A)pache (P)ortable (R)untime
//      - Should be installed by homebrew
//
#cgo CFLAGS: -I/usr/local/opt/apr/libexec/include/apr-1

//
//    (A)pache (P)ortable (R)untime utilities
//      - Should be installed by homebrew
//
#cgo CFLAGS: -I/usr/local/opt/apr-util/libexec/include/apr-1

//
//    Linker flags
//
#cgo LDFLAGS: -shared -Wl,-undefined,dynamic_lookup,-read_only_relocs,suppress -L/usr/lib -lpthread -ldl

*/
import "C"
