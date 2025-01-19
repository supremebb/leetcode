package huadongchuangkou

import "fmt"

// https://leetcode.cn/problems/maximum-number-of-vowels-in-a-substring-of-given-length/description/?envType=study-plan-v2&envId=leetcode-75

// 请返回字符串 s 中长度为 k 的单个子字符串中可能包含的最大元音字母数。
//
//英文中的 元音字母 为（a, e, i, o, u）。
//
//
//
//示例 1：
//
//输入：s = "abciiidef", k = 3
//输出：3
//解释：子字符串 "iii" 包含 3 个元音字母。
//示例 2：
//
//输入：s = "aeiou", k = 2
//输出：2
//解释：任意长度为 2 的子字符串都包含 2 个元音字母。
//示例 3：
//
//输入：s = "leetcode", k = 3
//输出：2
//解释："lee"、"eet" 和 "ode" 都包含 2 个元音字母。
//示例 4：
//
//输入：s = "rhythms", k = 4
//输出：0
//解释：字符串 s 中不含任何元音字母。
//示例 5：
//
//输入：s = "tryhard", k = 4
//输出：1

func Answer_maxVowels() {
	s1 := "abciiidef"
	k1 := 3
	r1 := maxVowels(s1, k1)
	fmt.Println(r1)

	s2 := "aeiou"
	k2 := 2
	r2 := maxVowels(s2, k2)
	fmt.Println(r2)

	s3 := "rhythms"
	k3 := 4
	r3 := maxVowels(s3, k3)
	fmt.Println(r3)

	s4 := "weallloveyou"
	k4 := 7
	r4 := maxVowels(s4, k4)
	fmt.Println(r4)
}

func maxVowels(s string, k int) int {
	vowelMap := map[rune]int{
		'a': 97,
		'e': 101,
		'i': 105,
		'o': 111,
		'u': 117,
	}
	result := 0
	lastFirstIsYY := false
	lastYYCount := 0
	for index, ch := range s {
		if index > len(s)-k {
			return result
		}
		if index == 0 {
			if _, ok := vowelMap[ch]; ok {
				lastYYCount++
				lastFirstIsYY = true
			}
			for j := 1; j < k; j++ {
				nextS := s[index+j]
				if _, ok := vowelMap[rune(nextS)]; ok {
					lastYYCount++
				}
			}
			result = lastYYCount
		} else {
			if lastFirstIsYY {
				lastYYCount--
			}
			firstS := s[index]
			lastS := s[index+k-1]
			if _, ok := vowelMap[rune(firstS)]; ok {
				lastFirstIsYY = true
			} else {
				lastFirstIsYY = false
			}
			if _, ok := vowelMap[rune(lastS)]; ok {
				lastYYCount++
			}
			if lastYYCount > result {
				result = lastYYCount
			}
		}
	}
	return result
}
