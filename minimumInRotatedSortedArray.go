package main

func min(a1, a2 int) int {
	if a1 < a2 {
		return a1
	}
	return a2
}

func findMin(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	if len(nums) < 2 {
		return nums[0]
	}
	if len(nums) < 3 {
		return min(nums[0], nums[1])
	}
	lastIndex := len(nums) - 1
	isRotated := nums[lastIndex] < nums[0]
	if !isRotated {
		return nums[0]
	}
	left, right, mid := 0, lastIndex, 0
	for left < right {
		mid = (right + left) / 2
		if nums[mid] < nums[right] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return nums[left]
}
