package middle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Q3: 无重复字符的最长子串
// https://leetcode.cn/problems/longest-substring-without-repeating-characters/
// 给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。
//
// 示例 1:
// 输入: s = "abcabcbb"
// 输出: 3
// 解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
//
// 示例 2:
// 输入: s = "bbbbb"
// 输出: 1
// 解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
//
// 示例 3:
// 输入: s = "pwwkew"
// 输出: 3
// 解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
//      请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。

func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]int)
	n := len(s)
	left, maxLen := 0, 0
	for i := 0; i < n; i++ {
		if j, ok := m[s[i]]; ok {
			left = max(left, j+1)
		}
		m[s[i]] = i
		maxLen = max(maxLen, i-left+1)
	}
	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Test_lengthOfLongestSubstring(t *testing.T) {
	str := "abcdabcdeaba"
	maxLen := lengthOfLongestSubstring(str)
	assert.Equal(t, 5, maxLen)
}
