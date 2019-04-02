package main

import (
	"fmt"
)

func main() {
	s := "isdfiiafoadfdi"

	fmt.Println(lengthOfLongestSubstring(s))
}

func lengthOfLongestSubstring(s string) int {
	var left, res int
	usedChar := make(map[rune]int)
	for i, v := range s {
		if k, ok := usedChar[v]; ok && left <= k {
			left = k + 1
		} else {
			res = max(res, i-left+1)
		}

		usedChar[v] = i
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
