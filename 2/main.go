package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	ret := twoSum(nums, target)
	fmt.Printf("answer: %v\n", ret)
}

//  Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	tmp = 0
	var l3 *ListNode
	for {
		if l1.Next == nil && l2.Next == nil {
			if !tmp {
				l3.Val = tmp
			}
			return l3
		}

		if l1.Next == nil {

		}

		if l2.Next == nil {

		}
	}

}
