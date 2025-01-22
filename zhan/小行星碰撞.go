package zhan

import "fmt"

// https://leetcode.cn/problems/asteroid-collision/description/?envType=study-plan-v2&envId=leetcode-75
// 给定一个整数数组 asteroids，表示在同一行的小行星。数组中小行星的索引表示它们在空间中的相对位置。
//
// 对于数组中的每一个元素，其绝对值表示小行星的大小，正负表示小行星的移动方向（正表示向右移动，负表示向左移动）。每一颗小行星以相同的速度移动。
//
// 找出碰撞后剩下的所有小行星。碰撞规则：两个小行星相互碰撞，较小的小行星会爆炸。如果两颗小行星大小相同，则两颗小行星都会爆炸。两颗移动方向相同的小行星，永远不会发生碰撞。
//
// 示例 1：
//
// 输入：asteroids = [5,10,-5]
// 输出：[5,10]
// 解释：10 和 -5 碰撞后只剩下 10 。 5 和 10 永远不会发生碰撞。
// 示例 2：
//
// 输入：asteroids = [8,-8]
// 输出：[]
// 解释：8 和 -8 碰撞后，两者都发生爆炸。
// 示例 3：
//
// 输入：asteroids = [10,2,-5]
// 输出：[10]
// 解释：2 和 -5 发生碰撞后剩下 -5 。10 和 -5 发生碰撞后剩下 10 。
func Answer_asteroidCollision() {
	a1 := []int{5, 10, -5}
	r1 := asteroidCollision(a1)
	fmt.Println(r1)

	a2 := []int{8, -8}
	r2 := asteroidCollision(a2)
	fmt.Println(r2)

	a3 := []int{10, 2, -5}
	r3 := asteroidCollision(a3)
	fmt.Println(r3)
}
func asteroidCollision(asteroids []int) []int {
	var result []int
	for _, as := range asteroids {
		if as > 0 {
			result = append(result, as)
		} else {
			for {
				if len(result) == 0 {
					result = append(result, as)
					break
				}
				if result[len(result)-1] < 0 {
					result = append(result, as)
					break
				}

				if result[len(result)-1]+as > 0 {
					break
				} else if result[len(result)-1]+as == 0 {
					result = result[:len(result)-1]
					break
				} else {
					result = result[:len(result)-1]
				}

			}
		}

	}
	return result
}
