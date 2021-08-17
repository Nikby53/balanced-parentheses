// Package balanced implements function for balanced parentheses balanced.
package balanced

// IsBalanced is a function that verifies if the given string is a balanced sequence of brackets.
func IsBalanced(s string) bool {
	var parentheses = map[rune]rune{
		']': '[',
		'}': '{',
		')': '(',
	}
	if s == "" {
		return false
	}
	var stack []rune
	for _, v := range s {
		n := len(stack) - 1
		if v == '[' || v == '{' || v == '(' {
			stack = append(stack, v)
		} else if v == ']' || v == '}' || v == ')' {
			if (len(stack) != 0 && parentheses[v] != stack[n]) || len(stack) == 0 {
				return false
			}
			stack = stack[:n]
		}
	}
	return len(stack) == 0
}
