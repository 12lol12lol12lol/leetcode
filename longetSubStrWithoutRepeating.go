package main

func max(a1, a2 int) int {
	if a1 > a2 {
		return a1
	}
	return a2
}

func lengthOfLongestSubstring(s string) int {
	left, longest := 0, 0
	seen := map[rune]int{}
	for index, value := range s {
		if value, ok := seen[value]; ok && value >= left {
			left = value + 1
		}
		longest = max(longest, index-left+1)
		seen[value] = index
	}
	return longest
}
