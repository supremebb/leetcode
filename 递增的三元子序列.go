package main

import "fmt"

// https://leetcode.cn/problems/increasing-triplet-subsequence/?envType=study-plan-v2&envId=leetcode-75

// 给你一个整数数组 nums ，判断这个数组中是否存在长度为 3 的递增子序列。
//
//如果存在这样的三元组下标 (i, j, k) 且满足 i < j < k ，使得 nums[i] < nums[j] < nums[k] ，返回 true ；否则，返回 false 。
//
//
//
//示例 1：
//
//输入：nums = [1,2,3,4,5]
//输出：true
//解释：任何 i < j < k 的三元组都满足题意
//示例 2：
//
//输入：nums = [5,4,3,2,1]
//输出：false
//解释：不存在满足题意的三元组
//示例 3：
//
//输入：nums = [2,1,5,0,4,6]
//输出：true
//解释：三元组 (3, 4, 5) 满足题意，因为 nums[3] == 0 < nums[4] == 4 < nums[5] == 6
//
//
//提示：
//
//1 <= nums.length <= 5 * 105
//-231 <= nums[i] <= 231 - 1
//
//
//进阶：你能实现时间复杂度为 O(n) ，空间复杂度为 O(1) 的解决方案吗？

func answer_increasingTriplet() {
	nums1 := []int{1, 2, 3, 4, 5}
	r1 := increasingTriplet(nums1)
	fmt.Println(r1)

	nums2 := []int{5, 4, 3, 2, 1}
	r2 := increasingTriplet(nums2)
	fmt.Println(r2)

	nums3 := []int{2, 1, 5, 0, 4, 6}
	r3 := increasingTriplet(nums3)
	fmt.Println(r3)

	nums4 := []int{20, 100, 10, 12, 5, 13}
	r4 := increasingTriplet(nums4)
	fmt.Println(r4)

	nums5 := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	r5 := increasingTriplet(nums5)
	fmt.Println(r5)

	nums6 := []int{1, 5, 0, 4, 1, 3}
	r6 := increasingTriplet(nums6)
	fmt.Println(r6)
}

func increasingTriplet(nums []int) bool {
	numLength := len(nums)
	if numLength < 3 {
		return false
	}

	LMin, RMax := make([]int, numLength), make([]int, numLength)
	// 拿每个位置 左侧最小值
	LMin[0] = nums[0]

	for i := 1; i < numLength-1; i++ {
		if nums[i-1] < LMin[i-1] {
			LMin[i] = nums[i-1]
		} else {
			LMin[i] = LMin[i-1]
		}
	}
	// 拿每个位置右侧最大值
	RMax[numLength-1] = nums[numLength-1]
	for i := numLength - 2; i > 0; i-- {
		if nums[i+1] > RMax[i+1] {
			RMax[i] = nums[i+1]
		} else {
			RMax[i] = RMax[i+1]
		}
	}

	for index, value := range nums {
		if value > LMin[index] && value < RMax[index] {
			return true
		}

	}

	return false
}
