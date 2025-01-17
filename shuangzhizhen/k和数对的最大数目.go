package shuangzhizhen

import (
	"fmt"
	"sort"
)

// https://leetcode.cn/problems/max-number-of-k-sum-pairs/description/?envType=study-plan-v2&envId=leetcode-75
// 给你一个整数数组 nums 和一个整数 k 。
//
// 每一步操作中，你需要从数组中选出和为 k 的两个整数，并将它们移出数组。
//
// 返回你可以对数组执行的最大操作数。
//
// 示例 1：
//
// 输入：nums = [1,2,3,4], k = 5
// 输出：2
// 解释：开始时 nums = [1,2,3,4]：
// - 移出 1 和 4 ，之后 nums = [2,3]
// - 移出 2 和 3 ，之后 nums = []
// 不再有和为 5 的数对，因此最多执行 2 次操作。
// 示例 2：
//
// 输入：nums = [3,1,3,4,3], k = 6
// 输出：1
// 解释：开始时 nums = [3,1,3,4,3]：
// - 移出前两个 3 ，之后nums = [1,4,3]
// 不再有和为 6 的数对，因此最多执行 1 次操作。
//
// =
// 提示：
//
// 1 <= nums.length <= 105
// 1 <= nums[i] <= 109
// 1 <= k <= 109
func Answer_maxOperations() {
	nums1 := []int{1, 2, 3, 4}
	k1 := 5
	result := maxOperations(nums1, k1)
	fmt.Println(result)

	nums2 := []int{3, 1, 3, 4, 3}
	k2 := 6
	result2 := maxOperations(nums2, k2)
	fmt.Println(result2)

	nums3 := []int{2, 2, 2, 3, 1, 1, 4, 1}
	k3 := 4
	result3 := maxOperations(nums3, k3)
	fmt.Println(result3)
}

func maxOperations(nums []int, k int) int {
	sort.Ints(nums)
	left, right := 0, len(nums)-1
	result := 0
	for left < right {
		temp := nums[left] + nums[right]
		if temp < k {
			left++
		} else if temp > k {
			right--
		} else {
			result++
			left++
			right--
		}
	}
	return result
}
