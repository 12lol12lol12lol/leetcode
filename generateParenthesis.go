package main

func generateParenthesis(n int) []string {
	answer := []string{}
	var backtracking func(current []rune, leftCount, rightCount int)
	backtracking = func(current []rune, leftCount, rightCount int) {
		if len(current) == 2*n {
			answer = append(answer, string(current))
			return
		}
		
	}
}
