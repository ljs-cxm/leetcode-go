package main

import (
	"fmt"
	//	"math"
)

//const limit = math.MaxInt32

func reverse(x int) int {
	var tmp int
	var sum int
	const (
		MAX_INT = 2147483647
		MIN_INT = -2147483648
	)
	flag := 1
	if x < 0 {
		flag = -1
		x = -1 * x
	}
	for x != 0 {
		//		fmt.Println(x)
		tmp = x % 10
		sum = sum*10 + tmp
		x = x / 10
	}
	res := flag * sum
	if (res > MAX_INT) || (res < MIN_INT) {
		return 0
	}
	return res
}

func main() {
	x := -9294967296
	res := reverse(9294967296)
	fmt.Println(res)

	fmt.Println(x)
}
