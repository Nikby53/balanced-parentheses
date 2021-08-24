// Package brackets implements function for brackets parentheses.
package brackets

// IsBalanced gets a string and returns true if the parentheses are balanced.
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
