package main

import (
	"fmt"
	"math"
)

func longestPalindrome(s string) string {
	//	odd number
	n := len(s)
	if n == 0 {
		return ""
	}
	if n == 1 {
		return s
	}

	var res1, res2 int
	var tmp1, tmp2 int
	var str1, str2 string

	for i := 0; i < n-1; i++ {
		tmp1 = 1
		l1 := i - 1
		r1 := i + 1
		for {
			//			fmt.Println("##i=%d", i)
			if l1 >= 0 && r1 <= n-1 {
				//				fmt.Println("s[l1]=%v, s[i]=%v, s[r1]=%v", string(s[l1]), string(s[i]), string(s[r1]))
				if s[l1] == s[r1] {
					tmp1 += 2
					l1--
					r1++
				} else {
					break
				}
			} else {
				break
			}

		}

		if tmp1 > res1 {
			str1 = s[l1+1 : r1]
			res1 = tmp1
		}
		//		fmt.Printf("##str1=%s\n\n", str1)
	}

	//	even number
	for j := 0; j < n-1; j++ {
		tmp2 = 0
		l2 := j
		r2 := j + 1
		for {
			//			fmt.Println("##j=%d", j)
			if l2 >= 0 && r2 <= n-1 {
				//				fmt.Printf("l2=%d, r2=%d\n", l2, r2)
				//				fmt.Println("s[l2]=%v, s[r2]=%v", string(s[l2]), string(s[r2]))
				if s[l2] == s[r2] {
					tmp2 += 2
					l2--
					r2++
				} else {
					break
				}

			} else {
				break
			}
		}
		if tmp2 > res2 {
			str2 = s[l2+1 : r2]
			res2 = tmp2
		}
		//		fmt.Printf("##str2=%s\n\n", str2)
	}

	//	fmt.Printf("##res1=%d, res2=%d\n\n", res1, res2)
	if res2 > res1 {
		return str2
	}

	return str1
}

func main() {
	s := "aaaddaaaaca"
	fmt.Println(longestPalindrome(s))
	fmt.Println(longestPalindrome1(s))
}

func longestPalindrome1(s string) string {
	if s == "" || len(s) == 0 {
		return ""
	}
	l, r := 0, 0
	strLen := len(s)

	for i := 0; i < strLen; i++ {
		a := getMax(s, i, i+1)
		b := getMax(s, i, i)

		max := int(math.Max(float64(a), float64(b)))

		if r-l < max {
			l = i - (max-1)/2
			r = i + max/2
		}
	}

	return s[l : r+1]
}

func getMax(str string, l int, r int) int {
	strLen := len(str)
	left, right := l, r

	for left >= 0 && right < strLen && str[left] == str[right] {
		left--
		right++
	}
	return right - left - 1
}
