package hash

import (
	"fmt"
	"sort"
)

// https://leetcode.cn/problems/unique-number-of-occurrences/description/?envType=study-plan-v2&envId=leetcode-75
// 给你一个整数数组 arr，如果每个数的出现次数都是独一无二的，就返回 true；否则返回 false。
//
//
//
//示例 1：
//
//输入：arr = [1,2,2,1,1,3]
//输出：true
//解释：在该数组中，1 出现了 3 次，2 出现了 2 次，3 只出现了 1 次。没有两个数的出现次数相同。
//示例 2：
//
//输入：arr = [1,2]
//输出：false
//示例 3：
//
//输入：arr = [-3,0,1,-3,1,1,1,-3,10,0]
//输出：true
//
//
//提示：
//
//1 <= arr.length <= 1000
//-1000 <= arr[i] <= 1000

func Answer_uniqueOccurrences() {
	nums1 := []int{1, 2, 2, 1, 1, 3}
	r1 := uniqueOccurrences(nums1)
	fmt.Println(r1)

	nums2 := []int{1, 2}
	r2 := uniqueOccurrences(nums2)
	fmt.Println(r2)

	nums3 := []int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}
	r3 := uniqueOccurrences(nums3)
	fmt.Println(r3)
}

func uniqueOccurrences(arr []int) bool {

	currentCount := 0
	current := 0
	sort.Ints(arr)
	countMap := make(map[int]int)
	for index, i := range arr {
		if index == 0 {
			currentCount = 1
			current = i
		} else {
			if i == current {
				currentCount++
			} else {
				if _, ok := countMap[currentCount]; ok {
					return false
				} else {
					countMap[currentCount] = 1
					current = i
					currentCount = 1
				}
			}

			if index == len(arr)-1 {
				//结算
				if _, ok := countMap[currentCount]; ok {
					return false
				}
			}
		}
	}
	return true
}
