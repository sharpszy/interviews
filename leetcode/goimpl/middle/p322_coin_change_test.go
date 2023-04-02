package middle

import (
	"goimpl/common"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Q322. 零钱兑换
// https://leetcode.cn/problems/coin-change/
// 给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。
// 计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1 。
// 你可以认为每种硬币的数量是无限的。
//
// 示例 1：
// 输入：coins = [1, 2, 5], amount = 11
// 输出：3
// 解释：11 = 5 + 5 + 1
//
// 示例 2：
// 输入：coins = [2], amount = 3
// 输出：-1
//
// 示例 3：
// 输入：coins = [1], amount = 0
// 输出：0

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = math.MaxInt
	}
	dp[0] = 0

	for i := range coins {
		for j := coins[i]; j <= amount; j++ {
			if dp[j-coins[i]] != math.MaxInt {
				dp[j] = common.Min(dp[j-coins[i]]+1, dp[j])
			}
		}
	}

	if dp[amount] == math.MaxInt {
		return -1
	}
	return dp[amount]
}

func Test_(t *testing.T) {
	coins := []int{1, 2, 5}
	amount := 11
	assert.Equal(t, 3, coinChange(coins, amount))

	coins = []int{1, 2, 5}
	amount = 16
	assert.Equal(t, 4, coinChange(coins, amount))
}
