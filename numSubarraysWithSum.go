package main

func numSubarraysWithSum(nums []int, goal int) int {
	currentSum, subArraysCount := 0, 0
	prefixSumMap := map[int]int{}
	for _, value := range nums {
		currentSum += value
		if currentSum == goal {
			subArraysCount++
		}
		if value, ok := prefixSumMap[currentSum-goal]; ok {
			subArraysCount += value
		}
		prefixSumMap[currentSum] += 1
	}

	return subArraysCount
}
