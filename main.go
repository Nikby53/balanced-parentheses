package main

import (
	"fmt"

	"github.com/Nikby53/balanced-parentheses/brackets"
	"github.com/Nikby53/balanced-parentheses/service"
)

func main() {
	fmt.Println(service.Calculation{Calculate: brackets.Generator{}}.CalculateOfBalanced(2))
}
