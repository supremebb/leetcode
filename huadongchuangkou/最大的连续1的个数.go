package huadongchuangkou

import "fmt"

// https://leetcode.cn/problems/max-consecutive-ones-iii/description/?envType=study-plan-v2&envId=leetcode-75
// 给定一个二进制数组 nums 和一个整数 k，假设最多可以翻转 k 个 0 ，则返回执行操作后 数组中连续 1 的最大个数 。
//
// 示例 1：
//
// 输入：nums = [1,1,1,0,0,0,1,1,1,1,0], K = 2
// 输出：6
// 解释：[1,1,1,0,0,1,1,1,1,1,1]
// 粗体数字从 0 翻转到 1，最长的子数组长度为 6。
// 示例 2：
//
// 输入：nums = [0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1], K = 3
// 输出：10
// 解释：[0,0,1,1,1,1,1,1,1,1,1,1,0,0,0,1,1,1,1]
// 粗体数字从 0 翻转到 1，最长的子数组长度为 10。
//
// 提示：
//
// 1 <= nums.length <= 105
// nums[i] 不是 0 就是 1
// 0 <= k <= nums.length
func Answer_longestOnes() {
	nums1 := []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}
	k1 := 2
	r1 := longestOnes(nums1, k1)
	fmt.Println(r1)
}

func longestOnes(nums []int, k int) int {
	// 3 -3 4 -1
	// 把nums转换成1个记录连续的数组
	var countNums []int
	lastCount := 0
	for index, i := range nums {
		if i == 1 {
			if lastCount >= 0 {
				lastCount++
			} else {
				countNums = append(countNums, lastCount)
				lastCount = 1
			}
		} else if i == 0 {
			if lastCount <= 0 {
				lastCount--
			} else {
				countNums = append(countNums, lastCount)
				lastCount = -1
			}
		}

		if index == len(nums)-1 {
			countNums = append(countNums, lastCount)
		}
	}

	// 我做不出来了
	return 0
}
