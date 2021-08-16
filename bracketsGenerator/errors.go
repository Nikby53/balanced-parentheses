package bracketsGenerator

import "errors"

var (
	ErrIncorrectInput  = errors.New("incorrect input")
	ErrIncorrectNumber = errors.New("enter a number that is greater than zero")
)
