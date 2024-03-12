package main

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	result := [][]int{}
	if len(nums) < 3 {
		return result
	}
	sort.Ints(nums)
	searchZero := func(current, lo, hi int) {
		for lo < hi {
			sum := nums[current] + nums[lo] + nums[hi]
			if sum > 0 {
				hi--
			} else if sum < 0 {
				lo++
			} else {
				result = append(result, []int{nums[current], nums[lo], nums[hi]})
				lo++
				hi--
				for lo < hi && nums[lo] == nums[lo-1] {
					lo++
				}
			}
		}
	}
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		searchZero(i, i+1, len(nums)-1)

	}
	return result
}

func threeSumMap(nums []int) [][]int {
	result := map[[3]int]struct{}{}
	seen := make(map[int]int, len(nums))
	dups := map[int]struct{}{}
	if len(nums) < 3 {
		return [][]int{}
	}
	for i := 0; i < len(nums); i++ {
		if _, ok := dups[nums[i]]; ok {
			continue
		}
		dups[nums[i]] = struct{}{}
		for j := i + 1; j < len(nums); j++ {
			comp := nums[i] + nums[j]
			if seendIndex, ok := seen[-comp]; ok && seendIndex == i {
				tmpSlice := []int{-comp, nums[i], nums[j]}
				sort.Ints(tmpSlice)
				result[[3]int{tmpSlice[0], tmpSlice[1], tmpSlice[2]}] = struct{}{}
			}
			seen[nums[j]] = i
		}
	}
	answers := [][]int{}
	for k := range result {
		answers = append(answers, []int{k[0], k[1], k[2]})
	}
	return answers
}
