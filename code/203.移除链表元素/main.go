package main

import "fmt"

func main() {
	val := 6
	head := &ListNode{Val: 1,
		Next: &ListNode{Val: 2,
			Next: &ListNode{Val: 6,
				Next: &ListNode{Val: 3,
					Next: &ListNode{Val: 4,
						Next: &ListNode{Val: 5,
							Next: &ListNode{Val: 6,
								Next: nil}},
					}}}}}
	ret := removeElements(head, val)
	fmt.Printf("result = %d\n", ret)
}

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	newHead := &ListNode{Val: 0, Next: head}
	pre := newHead
	cur := head
	for cur != nil {
		if cur.Val == val {
			pre.Next = cur.Next
		} else {
			pre = cur
		}
		cur = cur.Next
	}
	return newHead.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
