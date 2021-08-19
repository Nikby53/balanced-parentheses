// Package brackets implements function for brackets parentheses.
package brackets

import (
	"errors"
	"math/rand"
)

var errIncorrectNumber = errors.New("enter a number that is greater than zero")

// Generator is an empty struct that is
// also a receiver for Generate.
type Generator struct{}

// Generate method creates a random string of brackets of the entered length.
func (g Generator) Generate(num int) (string, error) {
	var parentheses = []rune("(){}[]")
	if num <= 0 {
		return "", errIncorrectNumber
	}
	generated := make([]rune, num)
	for i := range generated {
		generated[i] = parentheses[rand.Intn(len(parentheses))]
	}
	return string(generated), nil
}
