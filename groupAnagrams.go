package main

import (
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	anogramMap := make(map[string][]string)
	for _, str := range strs {
		s := strings.Split(str, "")
		sort.Strings(s)
		r := strings.Join(s, "")
		anogramMap[r] = append(anogramMap[r], str)
	}
	res := make([][]string, 0, len(anogramMap))
	for _, v := range anogramMap {
		res = append(res, v)
	}
	return res

}
