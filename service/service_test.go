package service

import (
	"errors"
	"testing"
)

type serviceMock struct {
	res string
	err error
}

func (s serviceMock) Generate(_ int) (string, error) {
	return s.res, s.err
}

func TestCalculateOfBalanced(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		resMock := "[()]"
		h := New(serviceMock{res: resMock, err: nil})
		temp, err := h.CalculateOfBalanced(4)
		if err != nil {
			t.Errorf("no error expected, but got %v", err)
		}
		if temp != 100 {
			t.Errorf("expected 100.00 instead of %v", temp)
		}
	})
	t.Run("0%", func(t *testing.T) {
		resMock := "[()][]["
		h := New(serviceMock{res: resMock, err: nil})
		temp, err := h.CalculateOfBalanced(4)
		if err != nil {
			t.Errorf("no error expected, but got %v", err)
		}
		if temp != 0 {
			t.Errorf("expected 0.00 instead of %v", temp)
		}
	})
	t.Run("error", func(t *testing.T) {
		resMock := "-1"
		h := New(serviceMock{res: resMock, err: errIncorrectInput})
		_, err := h.CalculateOfBalanced(4)
		if !errors.Is(err, errIncorrectInput) {
			t.Errorf("expected incorrect input instead of %v", errIncorrectInput)
		}
	})
}
