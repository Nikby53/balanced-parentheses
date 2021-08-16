package generator

import (
	"errors"
	"testing"

	"github.com/Nikby53/balanced-parentheses/bracketsGenerator"
)

func TestGeneration(t *testing.T) {
	tests := []struct {
		name          string
		length        int
		wantLength    int
		expectedError error
	}{
		{
			name:          "success",
			length:        1,
			wantLength:    1,
			expectedError: bracketsGenerator.ErrIncorrectNumber,
		},
		{
			name:          "negative number",
			length:        -1,
			wantLength:    0,
			expectedError: bracketsGenerator.ErrIncorrectNumber,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := Generation(tt.length)
			if tt.wantLength != len(res) {
				t.Errorf("expected length %v instead of %v", tt.wantLength, res)
			}
			if !errors.Is(err, tt.expectedError) {
				t.Errorf("expected %v instead of %v", tt.expectedError, err)
			}
		})
	}
}
