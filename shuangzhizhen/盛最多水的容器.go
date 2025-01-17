package shuangzhizhen

import "fmt"

// https://leetcode.cn/problems/container-with-most-water/description/?envType=study-plan-v2&envId=leetcode-75
// 给定一个长度为 n 的整数数组 height 。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i]) 。
//
// 找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
//
// 返回容器可以储存的最大水量。
//
// 说明：你不能倾斜容器。
func Answer_maxArea() {
	h1 := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	r1 := maxArea(h1)
	fmt.Println(r1)

	h2 := []int{1, 1}
	r2 := maxArea(h2)
	fmt.Println(r2)
}

func maxArea(height []int) int {
	result := 0
	for left, right := 0, len(height)-1; left < right; {
		currentR := 0
		if height[left] < height[right] {
			currentR = height[left] * (right - left)
			left++
		} else {
			currentR = height[right] * (right - left)
			right--
		}
		if currentR > result {
			result = currentR
		}
	}
	return result
}
