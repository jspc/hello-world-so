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

//export Hello
func Hello(ld bool) (hello string) {
	switch ld {
	case true:
		hello = ldHello
	case false:
		hello = "Hello, World!\000"
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
