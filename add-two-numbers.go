package main

import "fmt"

// ListNode ListNode
type ListNode struct {
	Val  int
	Next *ListNode
}

// addTwoNumbers addTwoNumbers
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		res  = new(ListNode)
		list = res
		val  int
	)

	for l1 != nil || l2 != nil || val > 0 {
		list.Next = new(ListNode)
		list = list.Next
		if l1 != nil {
			val += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			val += l2.Val
			l2 = l2.Next
		}

		list.Val = val % 10
		val /= 10
	}
	return res.Next
}

func main() {
	res := addTwoNumbers(&ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}, &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	})

	for res != nil && (res.Val > 0 || res.Next != nil) {
		fmt.Println(res.Val)
		res = res.Next
	}
}
