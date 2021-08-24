package service

import (
	"errors"
	"fmt"
	"sync"

	"github.com/Nikby53/balanced-parentheses/brackets"
)

var errIncorrectInput = errors.New("incorrect input")

type Generation interface {
	Generate(num int) (string, error)
}

func New(g Generation) Calculation {
	return Calculation{Calculate: g}
}

type Calculation struct {
	Calculate Generation
}

// CalculateOfBalanced w.
func (c Calculation) CalculateOfBalanced(length int) (float64, error) {
	var (
		wg              sync.WaitGroup
		mutex           sync.Mutex
		count           float64
		percentBalanced float64
	)
	if length <= 0 {
		return float64(length), errIncorrectInput
	}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			parentheses, err := c.Calculate.Generate(length)
			if err != nil {
				fmt.Errorf("%w", errIncorrectInput)
			}
			if brackets.IsBalanced(parentheses) {
				mutex.Lock()
				count++
				mutex.Unlock()
			}
		}()
	}
	wg.Wait()
	percentBalanced = count * 100.00 / 1000.00
	return percentBalanced, nil
}
