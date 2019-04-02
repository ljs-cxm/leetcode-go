package main

import (
	"fmt"
	//	"math"
)

func isPalindrome(x int) bool {
	if x != reverse(x) {
		return false
	}
	return true
}

func reverse(x int) int {
	var tmp int
	var sum int

	if x < 0 || x%10 == 0 {
		return -1
	}
	for x != 0 {
		tmp = x % 10
		sum = sum*10 + tmp
		x = x / 10
	}

	return sum
}

func main() {
	fmt.Println(isPalindrome(123321))
	fmt.Println(isPalindrome(-121))
	fmt.Println(isPalindrome(120))
	fmt.Println(isPalindrome(122))
}
