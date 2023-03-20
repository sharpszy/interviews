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

// 将第一个元素座位基准，遍历剩余元素，直到查找到不相等字符为止
// 时间复杂度O(N)，空间复杂度O(1)
func longestCommonPrefix2(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	idx := 0
outter:
	for ; idx < len(strs[0]); idx++ {
		for _, str := range strs {
			if idx >= len(str) || str[idx] != strs[0][idx] {
				break outter
			}
		}
	}
	return strs[0][:idx]
}

func Test_longestCommonPrefix(t *testing.T) {
	var strs = []string{"flower", "flow", "flight"}
	assert.Equal(t, "fl", longestCommonPrefix(strs))
	assert.Equal(t, "fl", longestCommonPrefix2(strs))

	strs = []string{"flower1", "flower23", "flower45", "flower789"}
	assert.Equal(t, "flower", longestCommonPrefix(strs))
	assert.Equal(t, "flower", longestCommonPrefix2(strs))

	strs = []string{"dog", "racecar", "car"}
	assert.Equal(t, "", longestCommonPrefix(strs))
	assert.Equal(t, "", longestCommonPrefix2(strs))

	s := "abc"
	assert.Equal(t, "", s[:0])
	assert.Equal(t, "ab", s[:len(s)-1])
	assert.Equal(t, "abc", s[:])
}
