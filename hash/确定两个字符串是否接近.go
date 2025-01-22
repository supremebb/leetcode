package hash

import (
	"reflect"
	"sort"
)

// https://leetcode.cn/problems/determine-if-two-strings-are-close/description/?envType=study-plan-v2&envId=leetcode-75
// 如果可以使用以下操作从一个字符串得到另一个字符串，则认为两个字符串 接近 ：
//
//操作 1：交换任意两个 现有 字符。
//例如，abcde -> aecdb
//操作 2：将一个 现有 字符的每次出现转换为另一个 现有 字符，并对另一个字符执行相同的操作。
//例如，aacabb -> bbcbaa（所有 a 转化为 b ，而所有的 b 转换为 a ）
//你可以根据需要对任意一个字符串多次使用这两种操作。
//
//给你两个字符串，word1 和 word2 。如果 word1 和 word2 接近 ，就返回 true ；否则，返回 false 。
//
//
//
//示例 1：
//
//输入：word1 = "abc", word2 = "bca"
//输出：true
//解释：2 次操作从 word1 获得 word2 。
//执行操作 1："abc" -> "acb"
//执行操作 1："acb" -> "bca"
//示例 2：
//
//输入：word1 = "a", word2 = "aa"
//输出：false
//解释：不管执行多少次操作，都无法从 word1 得到 word2 ，反之亦然。
//示例 3：
//
//输入：word1 = "cabbba", word2 = "abbccc"
//输出：true
//解释：3 次操作从 word1 获得 word2 。
//执行操作 1："cabbba" -> "caabbb"
//执行操作 2："caabbb" -> "baaccc"
//执行操作 2："baaccc" -> "abbccc"
//提示：
//
//1 <= word1.length, word2.length <= 105
//word1 和 word2 仅包含小写英文字母

func closeStrings(word1 string, word2 string) bool {
	// 长度不一样 肯定没戏
	// 统计每个字母出现的次数
	count1, count2 := make([]int, 26), make([]int, 26)
	for _, c := range word1 {
		count1[c-'a']++
	}
	for _, c := range word2 {
		count2[c-'a']++
	}
	for i := 0; i < 26; i++ {
		if count1[i] > 0 && count2[i] == 0 || count1[i] == 0 && count2[i] > 0 {
			return false
		}
	}
	sort.Ints(count1)
	sort.Ints(count2)
	return reflect.DeepEqual(count1, count2)
}
