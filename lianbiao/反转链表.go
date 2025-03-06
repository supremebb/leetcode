package lianbiao

import "fmt"

// https://leetcode.cn/problems/reverse-linked-list/?envType=study-plan-v2&envId=leetcode-75

func Answer_reverseList() {
	h15 := &ListNode{
		Val:  5,
		Next: nil,
	}
	h14 := &ListNode{
		Val:  4,
		Next: h15,
	}
	h13 := &ListNode{
		Val:  3,
		Next: h14,
	}
	h12 := &ListNode{
		Val:  2,
		Next: h13,
	}
	h11 := &ListNode{
		Val:  1,
		Next: h12,
	}
	res1 := reverseList(h11)
	printNode := res1
	for printNode != nil {
		fmt.Println(printNode.Val)
		printNode = printNode.Next
	}
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	lastNode := head
	i := 0
	for {
		oldNext := head.Next
		if i == 0 {
			head.Next = nil
		} else {
			head.Next = lastNode
			lastNode = head
		}
		head = oldNext
		i++
		if head == nil {
			break
		}
	}
	return lastNode
}
