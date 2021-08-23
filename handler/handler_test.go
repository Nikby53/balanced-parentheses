package handler

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

type serviceMock struct{}

func (s serviceMock) Generate(_ int) (string, error) {
	return "}]", nil
}

func HandlerTest(t *testing.T, status int, parameter string) func(t *testing.T) {
	return func(t *testing.T) {
		t.Helper()
		req := httptest.NewRequest("GET", "localhost:8081/generate"+parameter, nil)
		rec := httptest.NewRecorder()
		h := New(serviceMock{})
		h.GenerationHandler(rec, req)
		res := rec.Result()
		if res.StatusCode != status {
			t.Errorf("expected status %v, but got %v", status, res.Status)
		}
	}
}

func TestGenerationHandler(t *testing.T) {
	t.Run("success", HandlerTest(t, 200, "?n=2"))
	t.Run("number query parameters is required", HandlerTest(t, 400, ""))
	t.Run("query parameter should be number", HandlerTest(t, 400, "hello"))
	t.Run("parameter should be greater than zero", HandlerTest(t, 400, "?n=-2"))
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
