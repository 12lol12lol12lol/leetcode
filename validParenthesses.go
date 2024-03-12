package main

func isValid(s string) bool {
	stack := []rune{}
	parenthessesMap := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
	for _, value := range s {
		if openParenthess, ok := parenthessesMap[value]; ok {
			if len(stack) == 0 {
				return false
			}
			if stack[len(stack)-1] != openParenthess {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, value)
		}
	}
	return len(stack) == 0
}
