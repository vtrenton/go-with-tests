package context

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	response  string
	cancelled bool
	t         *testing.T
}

func (s *SpyStore) Fetch() string {
	return s.response
}

func (s *SpyStore) AssertWasCancelled() {
	s.t.Helper()
	if !s.cancelled {
		s.t.Error("store was not told to cancel")
	}
}

func (s *SpyStore) AssertWasNotCancelled() {
	s.t.Helper()
	if s.cancelled {
		s.t.Error("store was told to cancel")
	}
}

func (s *SpyStore) Cancel() {
	s.cancelled = true
}

func TestServer(t *testing.T) {

	t.Run("returns data from the store", func(t *testing.T) {
		data := "hello, world"
		store := &SpyStore{response: data}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf("got %s want %s", response.Body.String(), data)
		}

		store.AssertWasNotCancelled()
	})

	t.Run("tells store to cancel work if cancel is requested", func(t *testing.T) {
		data := "hello, world"
		store := &SpyStore{response: data}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)

		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		store.AssertWasCancelled()

	})
}
