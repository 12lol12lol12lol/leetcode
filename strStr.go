package main

func strStr(haystack, needle string) int {
	if len(needle) > len(haystack) {
		return -1
	} else if len(needle) == len(haystack) {
		if needle == haystack {
			return 0
		}
		return -1
	}
	if len(needle) < 1 {
		return -1
	}
	prefixTable := make([]int, len(haystack)+len(needle)+1)
	searchStr := needle + "#" + haystack
	for i := 1; i < len(searchStr); i++ {
		prevPrefValue := prefixTable[i-1]
		for prevPrefValue > 0 && searchStr[prevPrefValue] != searchStr[i] {
			prevPrefValue = prefixTable[prevPrefValue-1]
		}
		if searchStr[i] == searchStr[prevPrefValue] {
			prevPrefValue++
		}
		if prevPrefValue == len(needle) {
			return i - len(needle) - prevPrefValue
		}
		prefixTable[i] = prevPrefValue
	}
	return -1
}
