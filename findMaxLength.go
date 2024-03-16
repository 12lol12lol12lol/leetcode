package main

func findMaxLength(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	prefixSum := map[int]int{0: -1}
	currentSum, maxLength := 0, 0
	for index, value := range nums {
		if value == 0 {
			currentSum--
		} else {
			currentSum++
		}
		if value, ok := prefixSum[currentSum]; ok {
			maxLength = max(maxLength, index-value)
		} else {
			prefixSum[currentSum] = index
		}
	}

	return maxLength
}
