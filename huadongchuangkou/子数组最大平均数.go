package huadongchuangkou

import "fmt"

// https://leetcode.cn/problems/maximum-average-subarray-i/description/?envType=study-plan-v2&envId=leetcode-75

// 给你一个由 n 个元素组成的整数数组 nums 和一个整数 k 。
//
//请你找出平均数最大且 长度为 k 的连续子数组，并输出该最大平均数。
//
//任何误差小于 10-5 的答案都将被视为正确答案。
//
//
//
//示例 1：
//
//输入：nums = [1,12,-5,-6,50,3], k = 4
//输出：12.75
//解释：最大平均数 (12-5-6+50)/4 = 51/4 = 12.75
//示例 2：
//
//输入：nums = [5], k = 1
//输出：5.00000
//
//
//提示：
//
//n == nums.length
//1 <= k <= n <= 105
//-104 <= nums[i] <= 104

func Answer_findMaxAverage() {
	nums1 := []int{1, 12, -5, -6, 50, 3}
	k1 := 4
	r1 := findMaxAverage(nums1, k1)
	fmt.Println(r1)

	nums2 := []int{5}
	k2 := 1
	r2 := findMaxAverage(nums2, k2)
	fmt.Println(r2)

	nums3 := []int{-1}
	k3 := 1
	r3 := findMaxAverage(nums3, k3)
	fmt.Println(r3)

	nums4 := []int{8860, -853, 6534, 4477, -4589, 8646, -6155, -5577, -1656, -5779, -2619, -8604, -1358, -8009, 4983, 7063, 3104, -1560, 4080, 2763, 5616, -2375, 2848, 1394, -7173, -5225, -8244, -809, 8025, -4072, -4391, -9579, 1407, 6700, 2421, -6685, 5481, -1732, -8892, -6645, 3077, 3287, -4149, 8701, -4393, -9070, -1777, 2237, -3253, -506, -4931, -7366, -8132, 5406, -6300, -275, -1908, 67, 3569, 1433, -7262, -437, 8303, 4498, -379, 3054, -6285, 4203, 6908, 4433, 3077, 2288, 9733, -8067, 3007, 9725, 9669, 1362, -2561, -4225, 5442, -9006, -429, 160, -9234, -4444, 3586, -5711, -9506, -79, -4418, -4348, -5891}
	k4 := 93
	r4 := findMaxAverage(nums4, k4)
	fmt.Println(r4)

}

func findMaxAverage(nums []int, k int) float64 {
	result := float64(0)
	lastAvg := 0
	for i := 0; i <= len(nums)-k; i++ {
		if i == 0 {
			next := 0
			for next < k {
				lastAvg += nums[i+next]
				next++
			}
		} else {
			lastAvg = lastAvg - nums[i-1]
			lastAvg = lastAvg + nums[i+k-1]
		}

		currentResult := float64(lastAvg) / float64(k)
		if i == 0 {
			result = currentResult
			continue
		}
		if currentResult > result {
			result = currentResult
		}
	}
	return result
}
