package MyLinkedList

// MyLinkedList
// @Description: 双向链表
type MyLinkedList struct {
	Val  int
	Next *MyLinkedList
	Pre  *MyLinkedList
}

func Constructor() MyLinkedList {
	var myList = &MyLinkedList{
		Val:  -1,
		Next: nil,
		Pre:  nil,
	}
	myList.Next = myList
	myList.Pre = myList
	return *myList
}

func (this *MyLinkedList) Get(index int) int {
	var head = this.Next
	for head != this && index > 0 {
		index--
		head = head.Next
	}
	if index != 0 {
		return -1
	}
	return head.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	var head = &MyLinkedList{
		Val:  val,
		Next: nil,
		Pre:  nil,
	}
	this.Pre = head.Next
	head.Next = this
	this.Next = head.Pre
	head.Pre = this.Next
}

func (this *MyLinkedList) AddAtTail(val int) {

}

func (this *MyLinkedList) AddAtIndex(index int, val int) {

}

func (this *MyLinkedList) DeleteAtIndex(index int) {

}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
