package middle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// 5. 最长回文子串
// https://leetcode.cn/problems/longest-palindromic-substring/
// 给你一个字符串 s，找到 s 中最长的回文子串。
// 如果字符串的反序与原始字符串相同，则该字符串称为回文字符串。
//
// 示例 1：
// 输入：s = "babad"
// 输出："bab"
// 解释："aba" 同样是符合题意的答案。
//
// 示例 2：
// 输入：s = "cbbd"
// 输出："bb"

func longestPalindrome(s string) string {
	l := len(s)
	res := ""
	// 暴力解法，先从最大的子串开始，若为回文串，则直接返回
	for n := l; n > 0; n-- {
		for start := 0; start+n-1 < l; start++ {
			end := start + n - 1
			if isPalindrome(&s, start, end) {
				return s[start : end+1]
			}
		}
	}
	return res
}

func isPalindrome(s *string, start, end int) bool {
	for start <= end {
		if (*s)[start] != (*s)[end] {
			return false
		}
		start++
		end--
	}
	return true
}

func Test_longestPalindrome(t *testing.T) {

	s := "babad"
	assert.Equal(t, "bab", longestPalindrome(s))

	s = ""
	assert.Equal(t, "", longestPalindrome(s))

	s = "bb"
	assert.Equal(t, "bb", longestPalindrome(s))

	s = "cbbd"
	assert.Equal(t, "bb", longestPalindrome(s))
}
