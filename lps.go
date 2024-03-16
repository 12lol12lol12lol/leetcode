package main

func lps(s string) []int {
	lpsAr := make([]int, len(s))
	if len(s) < 2 {
		return lpsAr
	}
	for i := 1; i < len(s); i++ {
		prevValue := lpsAr[i-1]
		for prevValue > 0 && s[prevValue] != s[i] {
			prevValue = lpsAr[prevValue-1]
		}
		if s[prevValue] == s[i] {
			prevValue++
		}
		lpsAr[i] = prevValue

	}
	return lpsAr
}
