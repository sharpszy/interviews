package middle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// 17. 电话号码的字母组合
// https://leetcode.cn/problems/letter-combinations-of-a-phone-number/
// 给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。
//
// 给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
// 示例 1：
// 输入：digits = "23"
// 输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]
//
// 示例 2：
// 输入：digits = ""
// 输出：[]
//
// 示例 3：
// 输入：digits = "2"
// 输出：["a","b","c"]

var phoneMap = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

// 不确定有几层循环，使用回溯算法（递归的一种）
func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}

	var combinations []string
	backtrack(&combinations, digits, 0, "")
	return combinations
}

func backtrack(combinations *[]string, digits string, index int, combination string) {
	if index == len(digits) {
		*combinations = append(*combinations, combination)
	} else {
		digit := string(digits[index])
		letters := phoneMap[digit]
		lettersCount := len(letters)
		for i := 0; i < lettersCount; i++ {
			backtrack(combinations, digits, index+1, combination+string(letters[i]))
		}
	}
}

func Test_letterCombinations(t *testing.T) {
	digits := "23"
	result := []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}
	assert.Equal(t, result, letterCombinations(digits))

	digits = "2"
	result = []string{"a", "b", "c"}
	assert.Equal(t, result, letterCombinations(digits))
}
