package main

import (
	"fmt"
)

// https://leetcode.cn/problems/greatest-common-divisor-of-strings/description/?envType=study-plan-v2&envId=leetcode-75
// 对于字符串 s 和 t，只有在 s = t + t + t + ... + t + t（t 自身连接 1 次或多次）时，我们才认定 “t 能除尽 s”。
//
//给定两个字符串 str1 和 str2 。返回 最长字符串 x，要求满足 x 能除尽 str1 且 x 能除尽 str2 。
//
//
//
//示例 1：
//
//输入：str1 = "ABCABC", str2 = "ABC"
//输出："ABC"
//示例 2：
//
//输入：str1 = "ABABAB", str2 = "ABAB"
//输出："AB"
//示例 3：
//
//输入：str1 = "LEET", str2 = "CODE"
//输出：""
//
//
//提示：
//
//1 <= str1.length, str2.length <= 1000
//str1 和 str2 由大写英文字母组成

func answer_gcdOfStrings() {
	word11 := "ABCABC"
	word12 := "ABC"
	result1 := gcdOfStrings(word11, word12)
	fmt.Println(result1)

	word21 := "ABABAB"
	word22 := "ABAB"
	result2 := gcdOfStrings(word21, word22)
	fmt.Println(result2)

	word31 := "LEET"
	word32 := "CODE"
	word3 := gcdOfStrings(word31, word32)
	fmt.Println(word3)
}

func gcdOfStrings(str1 string, str2 string) string {
	result := ""
	// 首先通过简单的方法判断二者是否存在最大公约数
	// 这个是我没想到的, 如果正面连起来 跟反着连起来一样 则一定有最大公约数
	if str1+str2 != str2+str1 {
		return result
	}
	// 然后获取俩数的最大公约数的长度
	// 这个就是把2个
	gcd := getGCD(len(str1), len(str2))
	result = str1[0:gcd]
	return result
}

func getGCD(a, b int) int {
	if b == 0 {
		return a
	} else {
		return getGCD(b, a%b)
	}
}
