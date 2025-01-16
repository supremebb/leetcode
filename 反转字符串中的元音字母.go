package main

import "fmt"

// https://leetcode.cn/problems/reverse-vowels-of-a-string/description/?envType=study-plan-v2&envId=leetcode-75
// 给你一个字符串 s ，仅反转字符串中的所有元音字母，并返回结果字符串。
//
//元音字母包括 'a'、'e'、'i'、'o'、'u'，且可能以大小写两种形式出现不止一次。
//
//
//
//示例 1：
//
//输入：s = "IceCreAm"
//
//输出："AceCreIm"
//
//解释：
//
//s 中的元音是 ['I', 'e', 'e', 'A']。反转这些元音，s 变为 "AceCreIm".
//
//示例 2：
//
//输入：s = "leetcode"
//
//输出："leotcede"
//
//
//
//提示：
//
//1 <= s.length <= 3 * 105
//s 由 可打印的 ASCII 字符组成
// 	// a : 97 e:101 i: 105 o:111 u: 117
//	// A : 65 E:69 I: 73 O:79 U:85

func answer_reverseVowels() {
	s1 := "IceCreAm"
	r1 := reverseVowels(s1)
	fmt.Println(r1)

	s2 := "leetcode"
	r2 := reverseVowels(s2)
	fmt.Println(r2)
}

func reverseVowels(s string) string {
	vowelMap := map[rune]int{
		'a': 97,
		'e': 101,
		'i': 105,
		'o': 111,
		'u': 117,
		'A': 65,
		'E': 69,
		'I': 73,
		'O': 79,
		'U': 85,
	}

	var vowels []rune
	vowelIndexMaps := make(map[int]int)

	for index, ch := range s {
		if _, ok := vowelMap[ch]; ok {
			vowels = append(vowels, ch)
			vowelIndexMaps[index] = len(vowels) - 1
		}
	}

	var result string
	for index, ch := range s {
		if _, ok := vowelMap[ch]; ok {
			// 确定自己的位置
			trueIndex := vowelIndexMaps[index]
			reverseIndex := len(vowels) - trueIndex - 1
			result += string(vowels[reverseIndex])
		} else {
			result += string(ch)
		}
	}
	return result
}
