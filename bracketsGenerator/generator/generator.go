// Package generator is for generation random sequence of brackets.
package generator

import (
	"math/rand"

	"github.com/Nikby53/balanced-parentheses/bracketsGenerator"
)

// Store ww.
type Store struct{}

// GenerationRepo ww.
type GenerationRepo interface {
	Generation(num int) (string, error)
}

// Generation function creates a random string of brackets of the entered length.
func (s Store) Generation(num int) (string, error) {
	var parentheses = []rune("(){}[]")
	if num <= 0 {
		return "", bracketsGenerator.ErrIncorrectNumber
	}
	generated := make([]rune, num)
	for i := range generated {
		generated[i] = parentheses[rand.Intn(len(parentheses))]
	}
	return string(generated), nil
}
