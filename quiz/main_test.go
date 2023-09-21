package main

import "testing"

func TestHello(t *testing.T) {
	want := "Hello world"
	if got := Hello("world"); got != want {
		t.Errorf("Hello = %q, want: %q", got, want)
	}
}
