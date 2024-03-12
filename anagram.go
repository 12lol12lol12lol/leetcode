package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	if len(s) == 0 {
		return false
	}
	m := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		m[s[i]] += 1
		if m[s[i]] == 0 {
			delete(m, s[i])
		}
		m[t[i]] -= 1
		if m[t[i]] == 0 {
			delete(m, t[i])
		}
	}
	return len(m) == 0
}
