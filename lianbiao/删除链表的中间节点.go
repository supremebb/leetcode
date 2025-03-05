package lianbiao

import "fmt"

// https://leetcode.cn/problems/delete-the-middle-node-of-a-linked-list/description/?envType=study-plan-v2&envId=leetcode-75
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func Answer_deleteMiddle() {
	h16 := &ListNode{
		Val:  6,
		Next: nil,
	}
	h15 := &ListNode{
		Val:  2,
		Next: h16,
	}
	h14 := &ListNode{
		Val:  1,
		Next: h15,
	}
	h13 := &ListNode{
		Val:  7,
		Next: h14,
	}
	h12 := &ListNode{
		Val:  4,
		Next: h13,
	}
	h11 := &ListNode{
		Val:  3,
		Next: h12,
	}
	h10 := &ListNode{
		Val:  1,
		Next: h11,
	}
	res1 := deleteMiddle(h10)
	printNode := res1
	for printNode != nil {
		fmt.Println(printNode.Val)
		printNode = printNode.Next
	}

	h24 := &ListNode{
		Val:  4,
		Next: nil,
	}
	h23 := &ListNode{
		Val:  3,
		Next: h24,
	}
	h22 := &ListNode{
		Val:  2,
		Next: h23,
	}
	h21 := &ListNode{
		Val:  1,
		Next: h22,
	}

	res2 := deleteMiddle(h21)
	printNode2 := res2
	for printNode2 != nil {
		fmt.Println(printNode2.Val)
		printNode2 = printNode2.Next
	}
}

func deleteMiddle(head *ListNode) *ListNode {
	// 思路就是 On 解决 ，for 循环往下走 计数，每次计数都把中间节点 和下个节点拿出来
	nextMode := head
	var middleMode, middleLeftMode *ListNode
	i := 1
	for ; nextMode != nil; i++ {
		if i == 1 {
			middleLeftMode = nextMode
			middleMode = nextMode
		} else {
			if i%2 == 0 {
				// 往右进一个
				middleLeftMode = middleMode
				middleMode = middleMode.Next
			} else {

			}
		}
		nextMode = nextMode.Next
	}

	middleLeftMode.Next = middleMode.Next
	if i == 2 {
		head = nil
	}
	return head
}
