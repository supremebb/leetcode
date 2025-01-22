package zhan

import (
	"fmt"
	"strconv"
)

// https://leetcode.cn/problems/decode-string/?envType=study-plan-v2&envId=leetcode-75
// 给定一个经过编码的字符串，返回它解码后的字符串。
//
// 编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数。
//
// 你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且输入的方括号总是符合格式要求的。
//
// 此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 k ，例如不会出现像 3a 或 2[4] 的输入。
//
// 示例 1：
//
// 输入：s = "3[a]2[bc]"
// 输出："aaabcbc"
// 示例 2：
//
// 输入：s = "3[a2[c]]"
// 输出："accaccacc"
// 示例 3：
//
// 输入：s = "2[abc]3[cd]ef"
// 输出："abcabccdcdcdef"
// 示例 4：
//
// 输入：s = "abc3[cd]xyz"
// 输出："abccdcdcdxyz"
//
// 提示：
//
// 1 <= s.length <= 30
// s 由小写英文字母、数字和方括号 '[]' 组成
// s 保证是一个 有效 的输入。
// s 中所有整数的取值范围为 [1, 300]
func Answer_decodeString() {
	//s1 := "3[a]2[bc]"
	//r1 := decodeString(s1)
	//fmt.Println(r1)

	s2 := "3[a2[c]]"
	r2 := decodeString(s2)
	fmt.Println(r2)

	s3 := "100[leetcode]"
	r3 := decodeString(s3)
	fmt.Println(r3)

	s4 := "a"
	r4 := decodeString(s4)
	fmt.Println(r4)
}
func decodeString(s string) string {
	runes := []rune(s)
	var result string
	var index int
	for index <= len(runes)-1 {
		ir, ii := decodeS(runes[index:])
		result += ir
		index = index + ii
	}

	return result
}

func decodeS(runes []rune) (string, int) {
	result := ""
	length := 0
	for index, r := range runes {
		if index < length {
			continue
		}
		if isDigit(r) {
			digit, digLength := digitEnd(runes[index:])

			inr, inLength := decodeS(runes[index+digLength:])
			for i := 0; i < digit; i++ {
				result += inr
			}
			length += digLength
			length += inLength
		} else if r == ']' {
			length += 1
			break
		} else if r == '[' {
			length += 1
		} else {
			result = result + string(r)
			length += 1
		}
	}
	return result, length
}

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

func digitEnd(runes []rune) (int, int) {
	length := 0
	for index, r := range runes {
		if isDigit(r) {
			length++
		} else {
			length = index
			break
		}
	}

	r, _ := runesToInt(runes[:length])
	return r, length
}

func runesToInt(runes []rune) (int, error) {
	// Convert the []rune to a string
	str := string(runes)

	// Use strconv.Atoi to convert the string to an int
	number, err := strconv.Atoi(str)
	if err != nil {
		return 0, fmt.Errorf("failed to convert runes to int: %v", err)
	}

	return number, nil
}
