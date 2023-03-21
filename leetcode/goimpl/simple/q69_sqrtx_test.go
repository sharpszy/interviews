package simple

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Q69. x 的平方根
// https://leetcode.cn/problems/sqrtx/
// 给你一个非负整数 x ，计算并返回 x 的 算术平方根 。
// 由于返回类型是整数，结果只保留 整数部分 ，小数部分将被 舍去 。
// 注意：不允许使用任何内置指数函数和算符，例如 pow(x, 0.5) 或者 x ** 0.5 。
//
// 示例 1：
// 输入：x = 4
// 输出：2
//
// 示例 2：
// 输入：x = 8
// 输出：2
// 解释：8 的算术平方根是 2.82842..., 由于返回类型是整数，小数部分将被舍去。

// 二分法
func mySqrt(x int) int {
	var (
		l, r     = 0, x
		ans      = -1
		mid  int = l + (r-l)/2
	)
	for l <= r {
		if mid*mid <= x {
			ans = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
		mid = l + (r-l)/2
	}
	return ans
}

func Test_mySqrt(t *testing.T) {
	x := 4
	r := mySqrt(x)
	assert.Equal(t, 2, r)

	x = 0
	r = mySqrt(x)
	assert.Equal(t, 0, r)

	x = 8
	r = mySqrt(x)
	assert.Equal(t, 2, r)

	x = 17
	r = mySqrt(x)
	assert.Equal(t, 4, r)
}
