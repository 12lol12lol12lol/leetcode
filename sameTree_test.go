package main

import "testing"

func TestIsSameTree(t *testing.T) {
	testData := map[string]struct {
		p       []int
		q       []int
		exected bool
	}{
		"Case 1": {
			p:       []int{1, 2, 3},
			q:       []int{1, 2, 3},
			exected: true,
		},
		"Case 2": {
			p:       []int{},
			q:       []int{},
			exected: false,
		},
		"Case 3": {
			p:       []int{},
			q:       []int{},
			exected: false,
		},
	}
}
