hello-world-so
==

Simple test/ sample code for building `.so` files from golang, with some ldflag set data, and then loading in python and, if I have time, scala.

## Building


The following will compile a shared object file, exposing an ABI for libraries to use. It will also expose a header file.

```bash
$ go build -ldflags "-X main.ldHello=$(date +"%s")" -buildmode=c-shared -o strings.so strings.go
```

(In this case the unix time is used as the value of `main.ldHello` to show a distinct difference between versions)

This will generate two files:

1.  strings.so
1.  strings.h

The header file is useful for seeing definitions of symbols.

These symbols can be seen as per the usual:

```bash
strings.so: 000000000004f8c0 T _C_Hello
strings.so:                  U ___stack_chk_fail
strings.so:                  U ___stack_chk_guard
strings.so:                  U ___stderrp
strings.so: 000000000004f910 T __cgo_fe5ce5f2a878_Cfunc__Cmalloc
strings.so: 000000000004fdd0 T __cgo_get_context_function
strings.so: 000000000004e900 T __cgo_panic
strings.so: 000000000004f950 T __cgo_release_context
strings.so: 000000000004fb70 T __cgo_sys_thread_start
strings.so: 000000000004d6a0 T __cgo_topofstack
strings.so: 000000000004fce0 T __cgo_wait_runtime_init_done
strings.so: 00000000000026f0 T __cgoexp_fe5ce5f2a878_C_Hello
strings.so:                  U _abort
strings.so: 000000000004e950 T _crosscall2
strings.so: 000000000004fe91 T _crosscall_amd64
strings.so:                  U _fprintf
strings.so:                  U _fputc
strings.so:                  U _free
strings.so:                  U _fwrite
strings.so:                  U _malloc
strings.so:                  U _pthread_attr_destroy
strings.so:                  U _pthread_attr_getstacksize
strings.so:                  U _pthread_attr_init
strings.so:                  U _pthread_cond_broadcast
strings.so:                  U _pthread_cond_wait
strings.so:                  U _pthread_create
strings.so:                  U _pthread_key_create
strings.so:                  U _pthread_key_delete
strings.so:                  U _pthread_mutex_lock
strings.so:                  U _pthread_mutex_unlock
strings.so:                  U _pthread_setspecific
strings.so:                  U _pthread_sigmask
strings.so:                  U _setenv
strings.so:                  U _strerror
strings.so:                  U _unsetenv
strings.so: 000000000004f980 T _x_cgo_init
strings.so: 000000000004fd60 T _x_cgo_notify_runtime_init_done
strings.so: 000000000004fda0 T _x_cgo_set_context_function
strings.so: 000000000004fe00 T _x_cgo_setenv
strings.so: 000000000004fc80 T _x_cgo_sys_thread_create
strings.so: 000000000004fe30 T _x_cgo_thread_start
strings.so: 000000000004fe20 T _x_cgo_unsetenv
strings.so:                  U dyld_stub_binder
```

## Samples

### C

```bash
$ gcc -o hello c/main.c strings.so
$ ./hello
This is a C Application.
Hello, World!
1486300014
```

### Ruby

```
$ ruby ruby/hello.rb
Running from Ruby
Hello, World!
1486300014
```

### Python

```
$ python python/hello.py
Running from python
Hello, World!
1486300014
```


| who       | what |
|-----------|------|
| dockerhub | https://hub.docker.com/r/jspc/hello-world-so/   |
| circleci  | https://circleci.com/gh/jspc/hello-world-so   |
| licence   | MIT   |


Licence
--

MIT License

Copyright (c) 2017 jspc

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
