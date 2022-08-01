package main

func main() {
	l1 := &ListNode{}
	l2 := &ListNode{}
	ln := addTwoNumbers(l1, l2)
	if len(ln) != 2 {
		panic("add two numbers")
	}

}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	list := &ListNode{0, nil}
	result := list
	tmp := 0
	for l1 != nil || l2 != nil || tmp != 0 {
		if l1 != nil {
			tmp += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			tmp += l2.Val
			l2 = l2.Next
		}
		list.Next = &ListNode{tmp % 10, nil}
		tmp = tmp / 10
		list = list.Next
	}
	return result.Next
}

//定义单向链表
type ListNode struct {
	Val  int
	Next *ListNode
}
