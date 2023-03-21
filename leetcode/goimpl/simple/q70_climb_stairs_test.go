package simple

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Q70: 爬楼梯
// https://leetcode.cn/problems/climbing-stairs/
// 假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
//
// 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
//
// 示例 1：
// 输入：n = 2
// 输出：2
// 解释：有两种方法可以爬到楼顶。
// 1. 1 阶 + 1 阶
// 2. 2 阶
//
// 示例 2：
// 输入：n = 3
// 输出：3
// 解释：有三种方法可以爬到楼顶。
// 1. 1 阶 + 1 阶 + 1 阶
// 2. 1 阶 + 2 阶
// 3. 2 阶 + 1 阶

// 解题关键，推到公式: f(x) = f(x-1) + f(x-2)
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}

	// n-2 , n-1 项
	pn_2, pn_1 := 1, 2
	pn := 0
	for i := 3; i <= n; i++ {
		pn = pn_1 + pn_2
		pn_2 = pn_1
		pn_1 = pn
	}
	return pn
}

func Test_climbStairs(t *testing.T) {
	n := 2
	assert.Equal(t, 2, climbStairs(n))

	n = 3
	assert.Equal(t, 3, climbStairs(n))

	n = 4
	assert.Equal(t, 5, climbStairs(n))

	n = 5
	assert.Equal(t, 8, climbStairs(n))
}
