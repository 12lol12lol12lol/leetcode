package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTopKFrequent(t *testing.T) {
	type testData struct {
		nums     []int
		k        int
		expected []int
	}
	tableData := map[*testData]struct{}{
		{
			nums:     []int{1, 1, 1, 2, 2, 3},
			k:        2,
			expected: []int{2, 1},
		}: {},
		{
			nums:     []int{1},
			k:        1,
			expected: []int{1},
		}: {},
	}
	for v := range tableData {
		assert.Equal(t, v.expected, topKFrequent(v.nums, v.k), fmt.Sprintf("error with args nums=%v k=%v", v.nums, v.k))
	}
}
