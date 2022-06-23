package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := &ListNode{Val: 1, Next: nil}
	l2 := &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: nil}}
	l3 := addTwoNumbers(l1, l2)
	display(l3)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}

	if l1 != nil {
		if l2 != nil {
			result.Val = l1.Val + l2.Val
			l2 = l2.Next
		} else {
			result.Val = l1.Val
		}
		l1 = l1.Next
		if overflow(result) {
			if l1 == nil {
				l1 = &ListNode{Val: 0, Next: nil}
			}
			l1.Val++
		}
	} else {
		if l2 != nil {
			result.Val = l2.Val
			l2 = l2.Next
		} else {
			return nil
		}
		if overflow(result) {
			if l2 == nil {
				l2 = &ListNode{Val: 0, Next: nil}
			}
			l2.Val++
		}
	}

	result.Next = addTwoNumbers(l1, l2)

	return result
}

func overflow(node *ListNode) bool {
	if node.Val >= 10 {
		node.Val -= 10
		return true
	}
	return false
}

func display(list *ListNode) {
	for list != nil {
		fmt.Printf("%v -> ", list.Val)
		list = list.Next
	}
	fmt.Println()
}
