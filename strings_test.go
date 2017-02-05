package main

import (
	"testing"
)

var (
	newHStr = "Hello, NewH()!"
)

func TestHello(t *testing.T) {
	t.Run("Hello(true)", func(t *testing.T) {
		t.Run("When ldHello is unset", func(t *testing.T) {
			ldHello = ""

			h := Hello(true)
			if h != "" {
				t.Errorf("received %q, expected ''", h)
			}
		})

		t.Run("When ldHello is set", func(t *testing.T) {
			ldHello = "Hello Chump!"

			h := Hello(true)
			if h != ldHello {
				t.Errorf("received %q, expected %q", h, ldHello)
			}
		})
	})

	t.Run("Hello(false)", func(t *testing.T) {
		h := Hello(false)
		if h != "Hello, World!" {
			t.Errorf("received %q, expected 'Hello, World!'", h)
		}
	})
}

func TestNewH(t *testing.T) {
	h := NewH(newHStr)

	t.Run("NewH()", func(t *testing.T) {
		t.Run("correctly sets S value", func(t *testing.T) {
			if h.S != newHStr {
				t.Errorf("received %q, expected %q", h.S, newHStr)
			}
		})
	})
}

func TestH_String(t *testing.T) {
	h := H{newHStr}
	t.Run("h.String()", func(t *testing.T) {
		t.Run("returns correct value", func(t *testing.T) {
			s := h.String()
			if s != newHStr {
				t.Errorf("received %q, expected %q", s, newHStr)
			}
		})
	})
}
