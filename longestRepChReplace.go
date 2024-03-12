package main

func characterReplacement(s string, k int) int {
	left, maxFrequency, longest := 0, 0, 0
	freqMap := map[rune]int{}
	for index, value := range s {
		freqMap[value]++
		maxFrequency = max(maxFrequency, freqMap[value])

		if (index-left)-maxFrequency >= k {
			freqMap[rune(s[left])]--
			left++
		}
		longest = max(longest, index-left+1)
	}
	return longest
}
