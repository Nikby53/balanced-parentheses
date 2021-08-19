package brackets

import (
	"errors"
	"fmt"
	"testing"
)

func TestGeneration(t *testing.T) {
	tests := []struct {
		generator     Generator
		name          string
		length        int
		wantLength    int
		expectedError error
	}{
		{
			name:       "length 10000",
			length:     10000,
			wantLength: 10000,
		},
		{
			name:          "negative number",
			length:        -1,
			wantLength:    0,
			expectedError: errIncorrectNumber,
		},
		{
			name:          "zero value",
			length:        0,
			wantLength:    0,
			expectedError: errIncorrectNumber,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := tt.generator.Generate(tt.length)
			if tt.wantLength != len(res) {
				t.Errorf("expected length %v instead of %v", tt.wantLength, res)
			}
			if !errors.Is(err, tt.expectedError) {
				t.Errorf("expected %v instead of %v", tt.expectedError, err)
			}
		})
	}
}

func ExampleGenerator_Generate() {
	result, _ := Generator{}.Generate(5)
	fmt.Println(len(result))
	// Output:
	// 5
}
