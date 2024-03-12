package main

func binSearch(left, right, target int, nums []int) int {
	if len(nums) < 1 {
		return -1
	}
	if len(nums) < 2 {
		if nums[0] == target {
			return 0
		}
		return -1
	}
	var mid int
	for left <= right {
		mid = left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func findMinInSorted(nums []int) int {
	if len(nums) < 1 {
		return -1
	}
	if len(nums) < 2 {
		return min(nums[0], nums[1])
	}
	lastIndex := len(nums) - 1
	isRotated := nums[lastIndex] < nums[0]
	if !isRotated {
		return 0
	}
	left, right := 0, lastIndex
	var mid int
	for left <= right {
		mid = left + (right-left)/2
		if nums[mid] > nums[mid+1] {
			return mid + 1
		}
		if nums[mid-1] > nums[mid] {
			return mid
		}
		if nums[mid] > nums[0] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}

func search(nums []int, target int) int {
	if len(nums) < 1 {
		return -1
	}
	if len(nums) < 2 {
		if nums[0] == target {
			return 0
		}
		return -1
	}
	lastIndex := len(nums) - 1
	isRotated := nums[0] > nums[lastIndex]
	if !isRotated {
		if target < nums[0] {
			return -1
		}
		if target > nums[lastIndex] {
			return -1
		}
		return binSearch(0, lastIndex, target, nums)
	} else {
		minIndex := findMinInSorted(nums)
		res1, res2 := binSearch(minIndex, lastIndex, target, nums), binSearch(0, minIndex-1, target, nums)
		if res1*res2 > 0 {
			return -1
		}
		return res1 * res2 * -1
	}

}
