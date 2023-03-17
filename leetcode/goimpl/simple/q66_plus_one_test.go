package simple

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// 66. 加一
// https://leetcode.cn/problems/plus-one/
// 给定一个由 整数 组成的 非空 数组所表示的非负整数，在该数的基础上加一。
// 最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。
// 你可以假设除了整数 0 之外，这个整数不会以零开头。
//
// 示例 1：
// 输入：digits = [1,2,3]
// 输出：[1,2,4]
// 解释：输入数组表示数字 123。
//
// 示例 2：
// 输入：digits = [4,3,2,1]
// 输出：[4,3,2,2]
// 解释：输入数组表示数字 4321。

// num => [0, 9]
func plusOne(digits []int, num int) []int {
	var (
		carry = 0
		mod   = 0
		sum   = 0
	)
	for i := len(digits) - 1; i >= 0; i-- {
		if i == len(digits)-1 {
			sum = digits[i] + num + carry
		} else {
			sum = digits[i] + carry
		}
		carry = sum / 10
		mod = sum % 10
		digits[i] = mod
		if carry == 0 {
			break
		}
	}

	if carry > 0 {
		return append([]int{carry}, digits...)
	} else {
		return digits
	}
}

func Test_plusOne(t *testing.T) {
	nums := []int{1, 2, 3}
	assert.Equal(t, []int{1, 2, 4}, plusOne(nums, 1))

	nums = []int{4, 3, 2, 1}
	assert.Equal(t, []int{4, 3, 2, 2}, plusOne(nums, 1))

	nums = []int{9, 9, 9, 8}
	assert.Equal(t, []int{1, 0, 0, 0, 1}, plusOne(nums, 3))
}
