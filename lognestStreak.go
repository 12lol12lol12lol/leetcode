package main

func longestConsecutive(nums []int) int {
	m := make(map[int]struct{})
	for _, num := range nums {
		m[num] = struct{}{}
	}
	maxStreak := 0
	for num := range m {
		currentStreak, currentValue := 1, num
		if _, ok := m[currentValue-1]; ok {
			continue
		}
		for {
			if _, ok := m[currentValue+1]; ok {
				currentStreak++
				currentValue++
				continue
			}
			if currentStreak > maxStreak {
				maxStreak = currentStreak
			}
			break
		}
	}
	return maxStreak
}
