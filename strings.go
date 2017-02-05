package main

import (
	"C"
)

var (
	ldHello string
)

type H struct {
	S string
}

//export C_Hello
func C_Hello(ld bool) *C.char {
	return C.CString(Hello(ld))
}

func Hello(ld bool) (hello string) {
	switch ld {
	case true:
		hello = ldHello
	case false:
		hello = "Hello, World!"
	}

	return
}

func NewH(s string) (h H) {
	h.S = s
	return
}

func (h H) String() string {
	return h.S
}

func main() {}
