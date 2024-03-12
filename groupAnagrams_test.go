package main

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroupAnagrams(t *testing.T) {
	type testData struct {
		strs     []string
		expected [][]string
	}
	tableData := map[*testData]struct{}{
		{
			strs:     []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			expected: [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}},
		}: {},
		// {
		// 	strs:     []string{""},
		// 	expected: [][]string{{""}},
		// }: {},
		// {
		// 	strs:     []string{"a"},
		// 	expected: [][]string{{"a"}},
		// }: {},
	}
	for v := range tableData {
		fmt.Println(v.expected)
		fmt.Println(groupAnagrams(v.strs))
		assert.True(t, reflect.DeepEqual(v.expected, groupAnagrams(v.strs)), fmt.Sprintf("error with arg=%v", v.strs))
	}
}
