package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestWebSiteRacer(t *testing.T) {
	t.Run("compares the speed of two urls and returns the url of the faster one", func(t *testing.T) {
		slowServer := makeDelayedSever(20 * time.Millisecond)
		fastServer := makeDelayedSever(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, err := Racer(slowURL, fastURL)

		if err != nil {
			t.Errorf("did not expect an error but got one: %s", err)
		}

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("returns an error if the server doesnt respond in a specified time", func(t *testing.T) {
		server := makeDelayedSever(25 * time.Microsecond)

		defer server.Close()

		_, err := ConfigurableRacer(server.URL, server.URL, 20*time.Millisecond)

		if err != nil {
			t.Errorf("Expected an error bu didn't get one")
		}
	})
}

func makeDelayedSever(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
