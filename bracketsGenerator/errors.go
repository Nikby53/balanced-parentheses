package bracketsGenerator

import "errors"

var (
	// ErrIncorrectInput checks for incorrect input
	ErrIncorrectInput = errors.New("incorrect input")
	// ErrIncorrectNumber checks that number must be greater than zero
	ErrIncorrectNumber = errors.New("enter a number that is greater than zero")
)
