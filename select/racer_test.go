package main

import "testing"

func TestWebSiteRacer(t *testing.T) {
	slowURL := "https://facebook.com"
	fastURL := "https://quii.dev"

	want := fastURL
	got := Racer(slowURL, fastURL)

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
