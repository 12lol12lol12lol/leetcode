package main

func longestPrefix(s string) string {
	p1, p2 := 0, len(s)
	answ := ""
	for p1 < len(s)-1 {
		p1++
		p2--
		if s[:p1] == s[p2:] {
			answ = s[:p1]
		}
	}
	return answ
}
