package main

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		if resIndex, ok := m[target-v]; ok {
			return []int{resIndex, i}
		}
		m[v] = i
	}
	return []int{}
}
