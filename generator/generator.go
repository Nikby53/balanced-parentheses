// Package generator is for generation random sequence of brackets.
package generator

import (
	"errors"
	"math/rand"
)

// Generator ww.
type Generator struct{}

var ErrIncorrectNumber = errors.New("enter a number that is greater than zero")

// Generation function creates a random string of brackets of the entered length.
func (g Generator) Generation(num int) (string, error) {
	var parentheses = []rune("(){}[]")
	if num <= 0 {
		return "", ErrIncorrectNumber
	}
	generated := make([]rune, num)
	for i := range generated {
		generated[i] = parentheses[rand.Intn(len(parentheses))]
	}
	return string(generated), nil
}
