package hello

import "testing"

func TestHello(t *testing.T) {
	a := Hello()
	if a != "hello" {
		t.Errorf(`Hello() = %s, wanted "hello"`, a)
	}
}
