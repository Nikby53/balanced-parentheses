package brackets

import (
	"errors"
	"log"
)

var errIncorrectInput = errors.New("incorrect input, please enter a number from 1")

// CalculateOfBalanced method is for calculating the percent
// of balanced brackets of certain length.
func CalculateOfBalanced(length, quantity int) (float64, error) {
	cs := make(chan bool, quantity)
	if length <= 0 {
		return 0, errIncorrectInput
	}
	for i := 0; i < quantity; i++ {
		go func(ch chan<- bool) {
			parentheses, err := Generator{}.Generate(length)
			if err != nil {
				log.Fatal(err)
				return
			}
			ch <- IsBalanced(parentheses)
		}(cs)
	}
	var count int
	for i := 0; i < quantity; i++ {
		if <-cs {
			count++
		}
	}
	percentBalanced := float64(count) * 100.00 / float64(quantity)
	return percentBalanced, nil
}
