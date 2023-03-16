package simple

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Q14: 最长公共前缀
// https://leetcode.cn/problems/longest-common-prefix/
// 编写一个函数来查找字符串数组中的最长公共前缀。
// 如果不存在公共前缀，返回空字符串 ""。
//
// 示例 1：
// 输入：strs = ["flower","flow","flight"]
// 输出："fl"
//
// 示例 2：
// 输入：strs = ["dog","racecar","car"]
// 输出：""
// 解释：输入不存在公共前缀。

func longestCommonPrefix(strs []string) string {
	minLen, idx := math.MaxInt, -1
	for i, str := range strs {
		if len(str) < minLen {
			minLen = len(str)
			idx = i
		}
	}
	if idx < 0 {
		return ""
	}

	var (
		minIdx = 0
		ch     byte
		f      bool
	)
	for i := 0; i < len(strs[idx]); i++ {
		ch = strs[idx][i]
		f = true
		for j := range strs {
			if strs[j][i] != ch {
				f = false
				break
			}
		}
		if !f {
			break
		}
		minIdx++
	}
	return strs[idx][:minIdx]
}

func Test_longestCommonPrefix(t *testing.T) {
	var strs = []string{"flower", "flow", "flight"}
	assert.Equal(t, "fl", longestCommonPrefix(strs))

	strs = []string{"dog", "racecar", "car"}
	assert.Equal(t, "", longestCommonPrefix(strs))

	s := "abc"
	assert.Equal(t, "", s[:0])
	assert.Equal(t, "ab", s[:len(s)-1])
	assert.Equal(t, "abc", s[:])
}
