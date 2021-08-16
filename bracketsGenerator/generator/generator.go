package generator

import (
	"math/rand"

	"github.com/Nikby53/balanced-parentheses/bracketsGenerator"
)

func Generation(num int) (string, error) {
	var parentheses = []rune("(){}[]")
	if num <= 0 {
		return "", bracketsGenerator.ErrIncorrectNumber
	}
	s := make([]rune, num)
	for i := range s {
		s[i] = parentheses[rand.Intn(len(parentheses))]
	}
	return string(s), nil
}
