package qianzhuihe

import "fmt"

// https://leetcode.cn/problems/find-pivot-index/description/?envType=study-plan-v2&envId=leetcode-75
// 给你一个整数数组 nums ，请计算数组的 中心下标 。
//
// 数组 中心下标 是数组的一个下标，其左侧所有元素相加的和等于右侧所有元素相加的和。
//
// 如果中心下标位于数组最左端，那么左侧数之和视为 0 ，因为在下标的左侧不存在元素。这一点对于中心下标位于数组最右端同样适用。
//
// 如果数组有多个中心下标，应该返回 最靠近左边 的那一个。如果数组不存在中心下标，返回 -1 。
//
// 示例 1：
//
// 输入：nums = [1, 7, 3, 6, 5, 6]
// 输出：3
// 解释：
// 中心下标是 3 。
// 左侧数之和 sum = nums[0] + nums[1] + nums[2] = 1 + 7 + 3 = 11 ，
// 右侧数之和 sum = nums[4] + nums[5] = 5 + 6 = 11 ，二者相等。
// 示例 2：
//
// 输入：nums = [1, 2, 3]
// 输出：-1
// 解释：
// 数组中不存在满足此条件的中心下标。
// 示例 3：
//
// 输入：nums = [2, 1, -1]
// 输出：0
// 解释：
// 中心下标是 0 。
// 左侧数之和 sum = 0 ，（下标 0 左侧不存在元素），
// 右侧数之和 sum = nums[1] + nums[2] = 1 + -1 = 0 。
//
// 提示：
//
// 1 <= nums.length <= 104
// -1000 <= nums[i] <= 1000

func Answer_privotIndex() {
	nums1 := []int{1, 7, 3, 6, 5, 6}
	r1 := pivotIndex(nums1)
	fmt.Println(r1)

	nums2 := []int{1, 2, 3}
	r2 := pivotIndex(nums2)
	fmt.Println(r2)

	nums3 := []int{2, 1, -1}
	r3 := pivotIndex(nums3)
	fmt.Println(r3)
}

func pivotIndex(nums []int) int {
	// 算出来总值
	// 算出来每个位置左面的总和
	// 看看 哪个位置左面 = 总和-自己-左面
	numsSum := 0
	leftNums := make([]int, len(nums))
	for index, value := range nums {
		if index == 0 {
			leftNums[index] = 0
		} else {
			leftNums[index] = leftNums[index-1] + nums[index-1]
		}

		numsSum += value
	}

	for index, value := range leftNums {
		if numsSum-value-nums[index] == value {
			return index
		}
	}

	return -1
}
