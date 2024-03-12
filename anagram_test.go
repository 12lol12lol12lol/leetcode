package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsAnagram(t *testing.T) {

	type testData struct {
		a1       string
		a2       string
		expected bool
	}
	tableData := map[*testData]struct{}{
		{
			"anagram",
			"nagaram",
			true,
		}: {},
		{
			"rat",
			"car",
			false,
		}: {},
		{
			"ac",
			"bb",
			false,
		}: {},
	}
	for v := range tableData {
		assert.Equal(t, v.expected, isAnagram(v.a1, v.a2), fmt.Sprintf("Error with args a1=%v a2=%v expected=%v", v.a1, v.a2, v.expected))
	}
}
