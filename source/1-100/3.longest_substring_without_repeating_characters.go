package main

import (
	"fmt"
)

/*
3. 无重复字符的最长子串

给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。

示例：
输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
*/

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}

func lengthOfLongestSubstring(s string) int {
	maxLen := 0
	m := map[byte]int{}

	for i, j := 0, 0; i < len(s); i++ {
		if i > 0 {
			delete(m, s[i-1])
		}

		for j < len(s) && m[s[j]] == 0 {
			m[s[j]]++
			j++
		}

		if j-i > maxLen {
			maxLen = j - i
		}
	}

	return maxLen
}
