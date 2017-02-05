package main

type H struct {
	S string
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
