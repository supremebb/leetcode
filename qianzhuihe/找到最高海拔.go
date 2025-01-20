package qianzhuihe

import (
	"fmt"
	"sort"
)

// https://leetcode.cn/problems/find-the-highest-altitude/description/?envType=study-plan-v2&envId=leetcode-75
// 有一个自行车手打算进行一场公路骑行，这条路线总共由 n + 1 个不同海拔的点组成。自行车手从海拔为 0 的点 0 开始骑行。
//
// 给你一个长度为 n 的整数数组 gain ，其中 gain[i] 是点 i 和点 i + 1 的 净海拔高度差（0 <= i < n）。请你返回 最高点的海拔 。
//
// 示例 1：
//
// 输入：gain = [-5,1,5,0,-7]
// 输出：1
// 解释：海拔高度依次为 [0,-5,-4,1,1,-6] 。最高海拔为 1 。
// 示例 2：
//
// 输入：gain = [-4,-3,-2,-1,4,3,2]
// 输出：0
// 解释：海拔高度依次为 [0,-4,-7,-9,-10,-6,-3,-1] 。最高海拔为 0 。
//
// 提示：
//
// n == gain.length
// 1 <= n <= 100
// -100 <= gain[i] <= 100

func Answer_largestAltitude() {
	nums1 := []int{-5, 1, 5, 0, -7}
	r1 := largestAltitude(nums1)
	fmt.Println(r1)

	nums2 := []int{-4, -3, -2, -1, 4, 3, 2}
	r2 := largestAltitude(nums2)
	fmt.Println(r2)
}

func largestAltitude(gain []int) int {

	var highInt []int
	highInt = append(highInt, 0)
	for index, value := range gain {
		highInt = append(highInt, highInt[index]+value)
	}
	sort.Ints(highInt)
	return highInt[len(highInt)-1]
}
