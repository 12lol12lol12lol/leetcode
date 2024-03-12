package main

func maxArea(height []int) int {
	if len(height) < 2 {
		return 0
	}
	lo, hi, maxArea := 0, len(height)-1, 0
	for lo < hi {
		area := hi - lo
		if height[lo] < height[hi] {
			area *= height[lo]
			lo++
		} else {
			area *= height[hi]
			hi--
		}
		if area > maxArea {
			maxArea = area
		}
	}
	return maxArea
}
