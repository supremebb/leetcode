package huadongchuangkou

import "fmt"

// https://leetcode.cn/problems/longest-subarray-of-1s-after-deleting-one-element/description/?envType=study-plan-v2&envId=leetcode-75
// 给你一个二进制数组 nums ，你需要从中删掉一个元素。
//
//请你在删掉元素的结果数组中，返回最长的且只包含 1 的非空子数组的长度。
//
//如果不存在这样的子数组，请返回 0 。
//
//
//
//提示 1：
//
//输入：nums = [1,1,0,1]
//输出：3
//解释：删掉位置 2 的数后，[1,1,1] 包含 3 个 1 。
//示例 2：
//
//输入：nums = [0,1,1,1,0,1,1,0,1]
//输出：5
//解释：删掉位置 4 的数字后，[0,1,1,1,1,1,0,1] 的最长全 1 子数组为 [1,1,1,1,1] 。
//示例 3：
//
//输入：nums = [1,1,1]
//输出：2
//解释：你必须要删除一个元素。

func Answer_longestSubarray() {
	//nums1 := []int{1, 1, 0, 1}
	//r1 := longestSubarray(nums1)
	//fmt.Println(r1)
	//
	//nums2 := []int{0, 1, 1, 1, 0, 1, 1, 0, 1}
	//r2 := longestSubarray(nums2)
	//fmt.Println(r2)
	//
	//nums3 := []int{1, 1, 1}
	//r3 := longestSubarray(nums3)
	//fmt.Println(r3)

	nums4 := []int{0, 0, 1, 1}
	r4 := longestSubarray(nums4)
	fmt.Println(r4)
}

func longestSubarray(nums []int) int {
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

	fmt.Println(countNums)
	result := 0
	if len(countNums) == 1 {
		result = countNums[0] - 1
	} else {
		for index, value := range countNums {
			if value > 0 {
				continue
			} else {
				if value == -1 {
					// 那就得看左右连起来是多少了
					right := 0
					if index != len(countNums)-1 {
						right = countNums[index+1]
					}
					left := 0
					if index-1 >= 0 {
						left = countNums[index-1]
					}
					if left+1+right > result {
						result = left + right
					}

				} else {
					// 那就看分别和左右连起来是多少了
					right := 0
					if index != len(countNums)-1 {
						right = countNums[index+1]
					}
					if right > result {
						result = right
					}
					left := 0
					if index-1 >= 0 {
						left = countNums[index-1]
					}
					if left > result {
						result = left
					}
				}

			}

		}
	}

	if result < 0 {
		result = 0
	}
	return result
}
