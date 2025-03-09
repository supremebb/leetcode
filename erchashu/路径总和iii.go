package erchashu

import "fmt"

// https://leetcode.cn/problems/path-sum-iii/?envType=study-plan-v2&envId=leetcode-7

func Answer_pathSum() {
	tr11 := &TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}
	tr12 := &TreeNode{
		Val:   -2,
		Left:  nil,
		Right: nil,
	}
	tr13 := &TreeNode{
		Val:   3,
		Left:  tr11,
		Right: tr12,
	}
	tr14 := &TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}
	tr15 := &TreeNode{
		Val:   2,
		Left:  nil,
		Right: tr14,
	}
	tr16 := &TreeNode{
		Val:   5,
		Left:  tr13,
		Right: tr15,
	}
	tr17 := &TreeNode{
		Val:   11,
		Left:  nil,
		Right: nil,
	}
	tr18 := &TreeNode{
		Val:   -3,
		Left:  nil,
		Right: tr17,
	}
	trRoot1 := &TreeNode{
		Val:   10,
		Left:  tr16,
		Right: tr18,
	}
	r1 := pathSum(trRoot1, 8)
	fmt.Println(r1)

	tr21 := &TreeNode{
		Val:   5,
		Left:  nil,
		Right: nil,
	}
	tr22 := &TreeNode{
		Val:   4,
		Left:  nil,
		Right: tr21,
	}
	tr23 := &TreeNode{
		Val:   3,
		Left:  nil,
		Right: tr22,
	}
	tr24 := &TreeNode{
		Val:   2,
		Left:  nil,
		Right: tr23,
	}
	tr2Root := &TreeNode{
		Val:   1,
		Left:  nil,
		Right: tr24,
	}
	r2 := pathSum(tr2Root, 3)
	fmt.Println(r2)
}

func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	// 从1个节点往下纵深 一直到 left 和 right 都是nil 才算
	var result int

	result += pathSumNode(root, 0, targetSum)
	if root.Left != nil {
		result += pathSum(root.Left, targetSum)
	}
	if root.Right != nil {
		result += pathSum(root.Right, targetSum)
	}
	return result
}

func pathSumNode(root *TreeNode, currentSum, targetSum int) int {
	var result int
	currentSum += root.Val
	if currentSum == targetSum {
		result++
	}

	if root.Left != nil {
		leftResult := pathSumNode(root.Left, currentSum, targetSum)
		result += leftResult
	}
	if root.Right != nil {
		rightResult := pathSumNode(root.Right, currentSum, targetSum)
		result += rightResult
	}

	return result
}
