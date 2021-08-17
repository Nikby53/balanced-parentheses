package balanced

import (
	"fmt"
	"testing"
)

func TestIsBalanced(t *testing.T) {
	tests := []struct {
		in   string
		name string
		want bool
	}{
		{
			name: "balanced",
			in:   "{[{}]()}",
			want: true,
		},
		{
			name: "all parentheses with numbers are balanced",
			in:   "(((1 + 2) * 3) - 4)/5",
			want: true,
		},
		{
			name: "unbalanced",
			in:   "[(]",
			want: false,
		},
		{
			name: "empty string",
			in:   "",
			want: false,
		},
		{
			name: "starts with closed one",
			in:   "]})",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := IsBalanced(tt.in)
			if tt.want != res {
				t.Errorf("expected %v instead of %v", tt.want, res)
			}
		})
	}
}

func ExampleIsBalanced() {
	result := IsBalanced("()[]{}")
	fmt.Println(result)
	// Output:
	// true
}
