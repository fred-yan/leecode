package addtwo

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func New() *ListNode {
	return &ListNode{0,nil}
}

func (head *ListNode) Insert (i int, e int) bool{
	p := head
	j := 1
	for nil != p && j < i {
		p = p.Next
		j++
	}
	if nil == p || j > i {
		fmt.Println("pls check i:", i)
		return false
	}
	s := &ListNode{e, nil}
	s.Next = p.Next
	p.Next = s
	return true
}

func (head *ListNode) Traverse(){
	p := head.Next
	for p != nil {
		fmt.Println(p.Val)
		p = p.Next
	}
	fmt.Println("-----done---------")
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	sum := 0
	l3 := New()
	l4 := l3
	p, q := l1.Next, l2.Next
	for  p != nil {
		if q != nil {
			sum = p.Val + q.Val + carry

		} else {
			sum = p.Val + carry
		}
		if sum >= 10 {
			sum = sum - 10
			carry = 1
		} else {
			carry = 0
		}
		temp := &ListNode{0, nil}
		temp.Val = sum
		l4.Next = temp
		l4 = temp

		if q != nil {
			q = q.Next
		}
		p = p.Next
	}

	for q != nil {
		sum = q.Val + carry
		if sum >= 10 {
			sum = sum - 10
			carry = 1
		} else {
			carry = 0
		}
		temp := &ListNode{0, nil}
		temp.Val = sum
		l4.Next = temp
		l4 = temp
		q = q.Next
	}

	if carry != 0 {
		temp := &ListNode{1, nil}
		l4.Next = temp
		l4 = temp
	}

	return l3
}
