package handler

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

type serviceMock struct {
	res string
	err error
}

func (s serviceMock) Generate(_ int) (string, error) {
	return s.res, s.err
}

func TestGenerationHandler(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		req := httptest.NewRequest("GET", "localhost:8081/generate?n=2", nil)
		rec := httptest.NewRecorder()
		resMock := "}]"
		h := New(serviceMock{res: resMock, err: nil})
		h.GenerationHandler(rec, req)
		res := rec.Result()
		if res.StatusCode != http.StatusOK {
			t.Errorf("expected status Ok, but got %v", res.Status)
		}
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatalf("could not read response %v", err)
		}
		if string(body) != resMock {
			t.Fatalf("expected 2, but got %v", body)
		}
	})
	t.Run("no query parameter", func(t *testing.T) {
		req := httptest.NewRequest("GET", "localhost:8081/generate", nil)
		rec := httptest.NewRecorder()
		resMock := "}]"
		h := New(serviceMock{res: resMock, err: nil})
		h.GenerationHandler(rec, req)
		res := rec.Result()
		if res.StatusCode != http.StatusBadRequest {
			t.Errorf("expected status Ok, but got %v", res.Status)
		}
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatalf("could not read response %v", err)
		}
		if string(body) != "number query parametres is required\n" {
			t.Fatalf("expected number query parametres is required instead of %s", body)
		}
	})
	t.Run("less than zero", func(t *testing.T) {
		req := httptest.NewRequest("GET", "localhost:8081/generate?n=-2", nil)
		rec := httptest.NewRecorder()
		resMock := "}]"
		h := New(serviceMock{res: resMock, err: nil})
		h.GenerationHandler(rec, req)
		res := rec.Result()
		if res.StatusCode != http.StatusBadRequest {
			t.Errorf("expected status Ok, but got %v", res.Status)
		}
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatalf("could not read response %v", err)
		}
		if string(body) != "parameter should be greater than zero\n" {
			t.Fatalf("expected parameter should be greater than zero instead of %s", body)
		}
	})
	t.Run("not a number", func(t *testing.T) {
		req := httptest.NewRequest("GET", "localhost:8081/generate?n=qweqweq", nil)
		rec := httptest.NewRecorder()
		resMock := "}]"
		h := New(serviceMock{res: resMock, err: nil})
		h.GenerationHandler(rec, req)
		res := rec.Result()
		if res.StatusCode != http.StatusBadRequest {
			t.Errorf("expected status Ok, but got %v", res.Status)
		}
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatalf("could not read response %v", err)
		}
		if string(body) != "query parameter should be number\n" {
			t.Fatalf("expected query parameter should be number instead of %s", body)
		}
	})
}
