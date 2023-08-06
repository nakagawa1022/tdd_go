package context

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type stubStore struct {
	response string
}

type SpyStore struct {
	response  string
	cancelled bool
}

func (s *stubStore) Fetch() string {
	return s.response
}

func (s *SpyStore) Fetch() string {
	time.Sleep(100 * time.Millisecond)
	return s.response
}

func (s *SpyStore) Cancel() {
	s.cancelled = true
}

func TestHandler(t *testing.T) {
	data := "hello, world"
	svr := Server(&stubStore{data})

	request := httptest.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder()

	svr.ServeHTTP(response, request)

	if response.Body.String() != data {
		t.Errorf(`got "%s", want "%s"`, response.Body.String(), data)
	}
}
