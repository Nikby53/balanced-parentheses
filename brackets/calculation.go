package brackets

import (
	"errors"
	"log"
	"sync"
)

var errIncorrectInput = errors.New("incorrect input, please enter a number from 1")

// CalculateOfBalanced method is for calculating the percent
// of balanced string of a certain length.
func CalculateOfBalanced(length int, quantity int) (float64, error) {
	var (
		wg    sync.WaitGroup
		mutex sync.Mutex
		count float64
	)
	if length <= 0 {
		return 0, errIncorrectInput
	}
	for i := 0; i < quantity; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			parentheses, err := Generator{}.Generate(length)
			if err != nil {
				log.Println(err)
				return
			}
			if IsBalanced(parentheses) {
				mutex.Lock()
				count++
				mutex.Unlock()
			}
		}()
	}
	wg.Wait()
	percentBalanced := count * 100.00 / float64(quantity)
	return percentBalanced, nil
}
