package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Trent")
	want := "Hello, Trent"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
