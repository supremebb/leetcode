package erchashu

import "fmt"

// https://leetcode.cn/problems/leaf-similar-trees/description/?envType=study-plan-v2&envId=leetcode-75

func Answer_leafSimilar() {
	tr11 := &TreeNode{
		Val:   6,
		Left:  nil,
		Right: nil,
	}
	tr12 := &TreeNode{
		Val:   7,
		Left:  nil,
		Right: nil,
	}
	tr13 := &TreeNode{
		Val:   5,
		Left:  tr11,
		Right: tr12,
	}
	tr14 := &TreeNode{
		Val:   4,
		Left:  nil,
		Right: nil,
	}
	tr15 := &TreeNode{
		Val:   9,
		Left:  nil,
		Right: nil,
	}
	tr16 := &TreeNode{
		Val:   8,
		Left:  nil,
		Right: nil,
	}
	tr17 := &TreeNode{
		Val:   10,
		Left:  nil,
		Right: nil,
	}
	tr18 := &TreeNode{
		Val:   11,
		Left:  tr16,
		Right: tr17,
	}
	tr19 := &TreeNode{
		Val:   2,
		Left:  tr15,
		Right: tr18,
	}
	tr20 := &TreeNode{
		Val:   1,
		Left:  tr14,
		Right: tr19,
	}
	trRoot1 := &TreeNode{
		Val:   3,
		Left:  tr13,
		Right: tr20,
	}

	tr21 := &TreeNode{
		Val:   6,
		Left:  nil,
		Right: nil,
	}
	tr22 := &TreeNode{
		Val:   7,
		Left:  nil,
		Right: nil,
	}
	tr23 := &TreeNode{
		Val:   4,
		Left:  tr11,
		Right: tr12,
	}
	tr24 := &TreeNode{
		Val:   2,
		Left:  tr22,
		Right: tr23,
	}
	tr25 := &TreeNode{
		Val:   5,
		Left:  tr21,
		Right: tr24,
	}
	tr26 := &TreeNode{
		Val:   9,
		Left:  nil,
		Right: nil,
	}
	tr27 := &TreeNode{
		Val:   8,
		Left:  nil,
		Right: nil,
	}
	tr28 := &TreeNode{
		Val:   1,
		Left:  tr26,
		Right: tr27,
	}
	trRoot2 := &TreeNode{
		Val:   3,
		Left:  tr25,
		Right: tr28,
	}

	r1 := leafSimilar(trRoot1, trRoot2)
	fmt.Println(r1)

}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	// 就是看2个root的 叶节点的 数和顺序是否一样
	// 最笨的方法就是把两边的叶节点 拍成数组 然后挨个对吧
	// 这样的话 时间复杂度就是 On 但是空间复杂度会增加

	root1Arr := leafArr(root1)
	root2Arr := leafArr(root2)
	if len(root1Arr) != len(root2Arr) {
		return false
	}
	for i, v := range root1Arr {
		if v != root2Arr[i] {
			return false
		}
	}
	return true
}

func leafArr(tn *TreeNode) []int {
	if tn == nil {
		return nil
	}
	var res []int
	if tn.Left != nil {
		left := leafArr(tn.Left)
		res = append(res, left...)
	}
	if tn.Right != nil {
		right := leafArr(tn.Right)
		res = append(res, right...)
	}
	if tn.Left == nil && tn.Right == nil {
		res = append(res, tn.Val)
	}

	return res
}
