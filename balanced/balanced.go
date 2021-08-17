// Package balanced implements function for balanced parentheses balanced.
package balanced

// IsBalanced is a function that verifies if the given string is a balanced sequence of brackets.
func IsBalanced(s string) bool {
	var parentheses = map[rune]rune{
		']': '[',
		'}': '{',
		')': '(',
	}
	var stack []rune
	for _, v := range s {
		switch {
		case v == '[' || v == '{' || v == '(':
			stack = append(stack, v)
		case v == ']' || v == '}' || v == ')':
			if !(len(stack) != 0 && parentheses[v] == stack[len(stack)-1]) {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
