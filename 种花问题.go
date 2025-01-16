package main

import "fmt"

// https://leetcode.cn/problems/can-place-flowers/?envType=study-plan-v2&envId=leetcode-75
// 假设有一个很长的花坛，一部分地块种植了花，另一部分却没有。可是，花不能种植在相邻的地块上，它们会争夺水源，两者都会死去。
//
//给你一个整数数组 flowerbed 表示花坛，由若干 0 和 1 组成，其中 0 表示没种植花，1 表示种植了花。另有一个数 n ，能否在不打破种植规则的情况下种入 n 朵花？能则返回 true ，不能则返回 false 。
//
//
//
//示例 1：
//
//输入：flowerbed = [1,0,0,0,1], n = 1
//输出：true
//示例 2：
//
//输入：flowerbed = [1,0,0,0,1], n = 2
//输出：false
//
//
//提示：
//
//1 <= flowerbed.length <= 2 * 104
//flowerbed[i] 为 0 或 1
//flowerbed 中不存在相邻的两朵花
//0 <= n <= flowerbed.length

func answer_canPlaceFlowers() {
	f1 := []int{1, 0, 0, 0, 1}
	n1 := 1
	r1 := canPlaceFlowers(f1, n1)
	fmt.Println(r1)

	f2 := []int{1, 0, 0, 0, 1}
	n2 := 2
	r2 := canPlaceFlowers(f2, n2)
	fmt.Println(r2)

	f3 := []int{1, 0, 0, 0, 0, 0, 1}
	n3 := 3
	r3 := canPlaceFlowers(f3, n3)
	fmt.Println(r3)

	f4 := []int{1, 0, 0, 0, 0, 0, 1}
	n4 := 1
	r4 := canPlaceFlowers(f4, n4)
	fmt.Println(r4)

	f5 := []int{0, 1, 0}
	n5 := 1
	r5 := canPlaceFlowers(f5, n5)
	fmt.Println(r5)

	f6 := []int{0, 0, 1, 0, 1}
	n6 := 1
	r6 := canPlaceFlowers(f6, n6)
	fmt.Println(r6)
}

func canPlaceFlowers(flowerbed []int, n int) bool {
	flow := 0
	var canFlow bool
	for index, i := range flowerbed {
		canFlow = false
		if i == 1 {
			continue
		} else {
			// 每个位置能不能种的规则就是看前一个是不是1 和后面1个是不是1
			if index == 0 {
				// 不看前面
				canFlow = true
			} else {
				// 看前面
				canFlow = flowerbed[index-1] == 0
			}

			if canFlow {
				if index == len(flowerbed)-1 {
					// 不看后面
					canFlow = true
				} else {
					// 看后面
					canFlow = flowerbed[index+1] == 0
				}
			}

			if canFlow {
				flow++
				flowerbed[index] = 1
			}
			if flow >= n {
				return true
			}
		}
	}
	return flow >= n
}
