package handler

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

type serviceMock struct {
	err error
}

func (s serviceMock) Generate(_ int) (string, error) {
	return "}]", s.err
}

func TestGenerationHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "localhost:8081/generate?n=2", nil)
	rec := httptest.NewRecorder()
	h := New(serviceMock{
		err: nil,
	})
	h.GenerationHandler(rec, req)
	res := rec.Result()
	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status Ok, but got %v", res.Status)
	}
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("could not read response %v", err)
	}
	if string(b) != "}]" {
		t.Fatalf("expected 2, but got %v", b)
	}
}
