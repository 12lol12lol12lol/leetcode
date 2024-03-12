package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainsDuplicate(t *testing.T) {
	type TestValue struct {
		arr    []int
		isTrue bool
	}
	tableData := map[*TestValue]struct{}{
		{
			arr:    []int{1, 2, 3, 1},
			isTrue: true,
		}: {},
		{
			arr:    []int{1, 2, 3, 4},
			isTrue: false,
		}: {},
		{
			arr:    []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
			isTrue: true,
		}: {},
	}
	for v := range tableData {
		assert.Equal(t, v.isTrue, containsDuplicate(v.arr))
	}
}
