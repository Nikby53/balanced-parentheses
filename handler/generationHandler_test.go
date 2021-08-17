package handler

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

type generationHandlerMock struct {
	name string
	err  error
}

func (g *generationHandlerMock) Generation(num int) (string, error) {
	return "", g.err
}

func TestGenerationHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "localhost:8081/generate?n=2", nil)
	if err != nil {
		t.Fatalf("could not created request: %v", err)
	}
	rec := httptest.NewRecorder()
	GenerationHandler(rec, req)
	res := rec.Result()
	if res.StatusCode != http.StatusOK {
		t.Errorf("expercted status Ok but got %v", res.Status)
	}
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("could not read response %v", err)
	}
	if len(string(b)) != 2 {
		t.Fatalf("expected 2 but got %v", len(b))
	}
}
