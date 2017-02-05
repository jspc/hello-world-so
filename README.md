hello-world-so
==

Simple test/ sample code for building `.so` files from golang, with some ldflag set data, and then loading in python and, if I have time, scala.

Building
--

```bash
$ go build -ldflags "-X main.ldHello=$(date +"%s")" -buildmode=c-shared -o strings.so strings.go
```

**Note:** using the date as the value of `main.ldHello` is just a suggestion- anything will do. In this case the unix time is used to show a distinct difference between versions.


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
