package erchashu

import "fmt"

// https://leetcode.cn/problems/count-good-nodes-in-binary-tree/?envType=study-plan-v2&envId=leetcode-75

func Answer_goodNodes() {
	tr11 := &TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}
	tr12 := &TreeNode{
		Val:   1,
		Left:  tr11,
		Right: nil,
	}
	tr13 := &TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}
	tr14 := &TreeNode{
		Val:   5,
		Left:  nil,
		Right: nil,
	}
	tr15 := &TreeNode{
		Val:   4,
		Left:  tr13,
		Right: tr14,
	}
	tr16 := &TreeNode{
		Val:   3,
		Left:  tr12,
		Right: tr15,
	}
	r1 := goodNodes(tr16)
	fmt.Println(r1)
}

func goodNodes(root *TreeNode) int {
	res := goodNodesWithParent(root, root.Val)
	return res
}

func goodNodesWithParent(root *TreeNode, max int) int {
	if root == nil {
		return 0
	}
	var res int

	if root.Left != nil {
		leftMax := max
		if root.Left.Val > leftMax {
			leftMax = root.Left.Val
		}
		leftCount := goodNodesWithParent(root.Left, leftMax)
		res += leftCount
	}

	if root.Right != nil {
		rightMax := max
		if root.Right.Val > rightMax {
			rightMax = root.Right.Val
		}
		rightCount := goodNodesWithParent(root.Right, rightMax)
		res += rightCount
	}

	if root.Val >= max {
		res += 1
	}
	return res
}
