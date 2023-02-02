package main

import "fmt"

func main() {
	t1 := []int{2, 4, 3}
	t2 := []int{5, 6, 4}
	h1 := &ListNode{Val: t1[0]}
	l1 := h1
	for i := 1; i < len(t1); i++ {
		l1.Next = &ListNode{Val: t1[i]}
		l1 = l1.Next
	}
	h2 := &ListNode{Val: t2[0]}
	l2 := h2
	for i := 1; i < len(t1); i++ {
		l2.Next = &ListNode{Val: t2[i]}
		l2 = l2.Next
	}
	res := addTwoNumbers(h1, h2)
	res.Show()
}

func (l *ListNode) Show() {
	fmt.Println(l.Val)
	for l.Next != nil {
		l = l.Next
		fmt.Println(l.Val)
	}
}

// 输入：l1 = [2,4,3], l2 = [5,6,4]
// 输出：[7,0,8]
// 解释：342 + 465 = 807.
// 输入：l1 = [0], l2 = [0]
// 输出：[0]
// 输入：l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
// 输出：[8,9,9,9,0,0,0,1]
// 每个链表中的节点数在范围 [1, 100] 内
// 0 <= Node.val <= 9
// 题目数据保证列表表示的数字不含前导零
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	list := &ListNode{0, nil}
	ret := list
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
	return ret.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
