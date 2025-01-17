package shuangzhizhen

import "fmt"

// https://leetcode.cn/problems/move-zeroes/description/?envType=study-plan-v2&envId=leetcode-75
// 给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
//
//请注意 ，必须在不复制数组的情况下原地对数组进行操作。
//
//
//
//示例 1:
//
//输入: nums = [0,1,0,3,12]
//输出: [1,3,12,0,0]
//示例 2:
//
//输入: nums = [0]
//输出: [0]
//
//
//提示:
//
//1 <= nums.length <= 104
//-231 <= nums[i] <= 231 - 1
//
//
//进阶：你能尽量减少完成的操作次数吗？

func Answer_moveZeroes() {
	n1 := []int{0, 1, 0, 3, 12}
	moveZeroes(n1)
	fmt.Println(n1)

	n2 := []int{0}
	moveZeroes(n2)
	fmt.Println(n2)

	n3 := []int{0, 1}
	moveZeroes(n3)
	fmt.Println(n3)
}

func moveZeroes(nums []int) {
	read := 0
	write := read + 1
	for read < len(nums)-1 && write <= len(nums)-1 {
		if nums[read] == 0 {
			if nums[write] != 0 {
				nums[read], nums[write] = nums[write], nums[read]
				read++
				write = read + 1
			} else {
				// donothing?
				write++
			}
		} else {
			read++
			write = read + 1
		}
	}
}
