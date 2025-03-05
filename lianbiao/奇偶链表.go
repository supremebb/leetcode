package lianbiao

import "fmt"

//给定单链表的头节点 head ，将所有索引为奇数的节点和索引为偶数的节点分别组合在一起，然后返回重新排序的列表。
//
//第一个节点的索引被认为是 奇数 ， 第二个节点的索引为 偶数 ，以此类推。
//
//请注意，偶数组和奇数组内部的相对顺序应该与输入时保持一致。
//
//你必须在 O(1) 的额外空间复杂度和 O(n) 的时间复杂度下解决这个问题。
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func Answer_oddEvenList() {
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
	res1 := oddEvenList(h11)
	printNode := res1
	for printNode != nil {
		fmt.Println(printNode.Val)
		printNode = printNode.Next
	}
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	evenHead := head.Next
	odd := head
	even := evenHead
	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next
	}
	odd.Next = evenHead
	return head
}
