package handler

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

type serviceMock struct {
	res string
}

func (s serviceMock) Generate(_ int) (string, error) {
	return s.res, nil
}

func Success(t testing.TB, req *http.Request) {
	t.Helper()
	rec := httptest.NewRecorder()
	resMock := "}]"
	h := New(serviceMock{res: resMock})
	h.GenerationHandler(rec, req)
	res := rec.Result()
	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status Ok, but got %v", res.Status)
	}
}

func Error(t testing.TB, req *http.Request) {
	t.Helper()
	rec := httptest.NewRecorder()
	resMock := "}]"
	h := New(serviceMock{res: resMock})
	h.GenerationHandler(rec, req)
	res := rec.Result()
	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("expected status Ok, but got %v", res.Status)
	}
}

func TestGenerationHandler(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		req := httptest.NewRequest("GET", "localhost:8081/generate?n=2", nil)
		Success(t, req)
	})
	t.Run("no query parameter", func(t *testing.T) {
		req := httptest.NewRequest("GET", "localhost:8081/generate", nil)
		Error(t, req)
	})
	t.Run("less than zero", func(t *testing.T) {
		req := httptest.NewRequest("GET", "localhost:8081/generate?n=-2", nil)
		Error(t, req)
	})
	t.Run("not a number", func(t *testing.T) {
		req := httptest.NewRequest("GET", "localhost:8081/generate?n=qweqweq", nil)
		Error(t, req)
	})
}

func TestGenerationRequest_Validate(t *testing.T) {
	tests := []struct {
		generationRequest generationRequest
		name              string
		have              string
		expectedError     error
	}{
		{
			name:          "should be a number",
			have:          "localhost:8081/generate",
			expectedError: errNumberRequired,
		},
		{
			name:          "less than zero",
			have:          "localhost:8081/generate?n=-2",
			expectedError: errLessThanZero,
		},
		{
			name:          "should be number",
			have:          "localhost:8081/generate?n=qweqe",
			expectedError: errShouldBeNumber,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest("GET", tt.have, nil)
			request := generationRequest{}
			err := request.Validate(req)
			if !errors.Is(err, tt.expectedError) {
				t.Errorf("expected %v instead of %v", tt.expectedError, err)
			}
		})
	}
}
