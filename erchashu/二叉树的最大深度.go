package erchashu

import "fmt"

// https://leetcode.cn/problems/maximum-depth-of-binary-tree/description/?envType=study-plan-v2&envId=leetcode-75

func Answer_maxDepth() {
	tr11 := &TreeNode{
		Val:   15,
		Left:  nil,
		Right: nil,
	}
	tr12 := &TreeNode{
		Val:   7,
		Left:  nil,
		Right: nil,
	}
	tr13 := &TreeNode{
		Val:   20,
		Left:  tr11,
		Right: tr12,
	}
	tr14 := &TreeNode{
		Val:   9,
		Left:  nil,
		Right: nil,
	}
	tr15 := &TreeNode{
		Val:   3,
		Left:  tr13,
		Right: tr14,
	}
	r1 := maxDepth(tr15)
	fmt.Println(r1)
}

func maxDepth(root *TreeNode) int {
	// 感觉要用到递归
	// 判断left 有没有
	// 判断right 有没有
	// 如果都没有则返回
	// 如果单边有 则继续传下去 然后+1返回
	if root == nil {
		return 0
	}
	var leftMax, rightMax int
	if root.Left != nil {
		leftMax = maxDepth(root.Left)
	}
	if root.Right != nil {
		rightMax = maxDepth(root.Right)
	}

	if leftMax > rightMax {
		return leftMax + 1
	} else {
		return rightMax + 1
	}
}
