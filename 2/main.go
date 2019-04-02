package main

import (
	"fmt"
)

func main() {
	l1_array := [3]int{3, 4, 5}
	l2_array := [4]int{6, 7, 8, 9}

	l1 := genList(l1_array[:])
	l2 := genList(l2_array[:])
	printList(l1)
	printList(l2)
	l3 := addTwoNumbers(l1, l2)
	printList(l3)

}

//  Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func genList(array []int) *ListNode {
	var list, cur *ListNode

	for _, v := range array {
		if list != nil {
			cur.Next = initListNode(v)
			cur = cur.Next
		} else {
			list = initListNode(v)
			cur = list
		}
	}
	return list
}

func initListNode(n int) *ListNode {
	node := &ListNode{}
	node.Val = n
	node.Next = nil
	return node
}

func printList(list *ListNode) {
	fmt.Printf("Begin to print list\n")
	for {
		if list != nil {
			fmt.Printf("%d", list.Val)
		} else {
			fmt.Printf("\n")
			return
		}

		if list.Next != nil {
			fmt.Printf("->")
		}
		list = list.Next
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var value, flag int
	var l3, cur *ListNode
	if l1 == nil && l2 == nil {
		return initListNode(0)
	}

	for {
		sum := 0
		if l1 == nil && l2 == nil {
			if flag != 0 {
				cur.Next = initListNode(flag)
			}
			return l3
		}

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		if flag != 0 {
			sum += flag
		}

		value = sum % 10
		flag = sum / 10

		if l3 != nil {
			cur.Next = initListNode(value)
			cur = cur.Next
		} else {
			l3 = initListNode(value)
			cur = l3
		}
	}
}
