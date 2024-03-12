package main

import (
	"math/rand"
)

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	keys := []int{}
	for _, v := range nums {
		if _, ok := m[v]; !ok {
			keys = append(keys, v)
		}
		m[v]++
	}
	if len(keys) <= k {
		return keys
	}
	partitionFunc := func(left int, right int, pivot int) int {
		pivotFrequence := m[keys[pivot]]
		keys[right], keys[pivot] = keys[pivot], keys[right]
		storeIndex := left
		for i := left; i < right; i++ {
			if m[keys[i]] < pivotFrequence {
				keys[storeIndex], keys[i] = keys[i], keys[storeIndex]
				storeIndex++
			}
		}
		keys[storeIndex], keys[right] = keys[right], keys[storeIndex]
		return storeIndex
	}
	var quickSelect func(int, int, int)
	quickSelect = func(left, right, requiredPosition int) {
		if left == right {
			return
		}
		pivotIndex := rand.Intn(right-left+1) + left
		pivotIndex = partitionFunc(left, right, pivotIndex)
		if pivotIndex == requiredPosition {
			return
		} else if pivotIndex < requiredPosition {
			quickSelect(pivotIndex+1, right, requiredPosition)
		} else {
			quickSelect(left, pivotIndex-1, requiredPosition)
		}
	}
	requiredPosition := len(keys) - k
	quickSelect(0, len(keys)-1, requiredPosition)
	return keys[requiredPosition:]
}
