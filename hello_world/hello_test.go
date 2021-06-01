package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Fred")
	want := "Hello, Fred"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
