package middle

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

//Q7: 整数反转
// https://leetcode.cn/problems/reverse-integer/
// 给你一个 32 位的有符号整数 x ，返回将 x 中的数字部分反转后的结果。
// 如果反转后整数超过 32 位的有符号整数的范围 [−231,  231 − 1] ，就返回 0。
// 假设环境不允许存储 64 位整数（有符号或无符号）。
// 示例 1：
// 输入：x = 123
// 输出：321
//
// 示例 2：
// 输入：x = -123
// 输出：-321

func reverses_int32(x int) int {
	var (
		res   int
		digit int
	)
	for x != 0 {
		if res < math.MinInt32/10 || res > math.MaxInt32/10 { // 注意溢出条件
			return 0
		}
		digit = x % 10
		res = res*10 + digit
		x = x / 10
	}
	return res
}

func Test_reverses_int32(t *testing.T) {
	x := 123
	assert.Equal(t, 321, reverses_int32(x))

	x = -123
	assert.Equal(t, -321, reverses_int32(x))

	x = 123000
	assert.Equal(t, 321, reverses_int32(x))
}
