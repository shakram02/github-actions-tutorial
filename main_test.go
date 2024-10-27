package main

import "testing"

func TestMain(t *testing.T) {
	g := Greeter{}

	if g.Greet("test") != "hello: test" {
		t.Fail()
	}
}
