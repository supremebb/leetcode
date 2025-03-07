package lianbiao

import "fmt"

// 在一个大小为 n 且 n 为 偶数 的链表中，对于 0 <= i <= (n / 2) - 1 的 i ，第 i 个节点（下标从 0 开始）的孪生节点为第 (n-1-i) 个节点 。
//
// 比方说，n = 4 那么节点 0 是节点 3 的孪生节点，节点 1 是节点 2 的孪生节点。这是长度为 n = 4 的链表中所有的孪生节点。
// 孪生和 定义为一个节点和它孪生节点两者值之和。
//
// 给你一个长度为偶数的链表的头节点 head ，请你返回链表的 最大孪生和 。
// https://leetcode.cn/problems/maximum-twin-sum-of-a-linked-list/description/?envType=study-plan-v2&envId=leetcode-75

func Answer_pairSum() {
	h14 := &ListNode{
		Val:  4,
		Next: nil,
	}
	h13 := &ListNode{
		Val:  7,
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
	res1 := pairSum(h11)
	fmt.Println(res1)

}

func pairSum(head *ListNode) int {
	// 第一反应就是把链表转换成有序的数组存储数据. On 就可以 但是需要额外空间
	nodeArr := make([]int, 0)
	for ; head != nil; head = head.Next {
		nodeArr = append(nodeArr, head.Val)
	}
	var result int
	nodeLength := len(nodeArr)
	for i, v := range nodeArr {
		if i == nodeLength/2 {
			break
		}
		if v+nodeArr[nodeLength-i-1] > result {
			result = v + nodeArr[nodeLength-i-1]
		}
	}
	return result
}
