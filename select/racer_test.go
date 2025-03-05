package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestWebSiteRacer(t *testing.T) {
	slowServer := makeDelayedSever(20 * time.Millisecond)
	fastServer := makeDelayedSever(0 * time.Millisecond)

	defer slowServer.Close()
	defer fastServer.Close()

	slowURL := slowServer.URL
	fastURL := fastServer.URL

	want := fastURL
	got := Racer(slowURL, fastURL)

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func makeDelayedSever(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
