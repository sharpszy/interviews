package simple

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Q1480. 一维数组的动态和
// https://leetcode.cn/problems/running-sum-of-1d-array/
// 给你一个数组 nums 。数组「动态和」的计算公式为：runningSum[i] = sum(nums[0]…nums[i]) 。
// 请返回 nums 的动态和。
// 示例 1：
// 输入：nums = [1,2,3,4]
// 输出：[1,3,6,10]
// 解释：动态和计算过程为 [1, 1+2, 1+2+3, 1+2+3+4] 。
//
// 示例 2：
// 输入：nums = [1,1,1,1,1]
// 输出：[1,2,3,4,5]
// 解释：动态和计算过程为 [1, 1+1, 1+1+1, 1+1+1+1, 1+1+1+1+1] 。

func runningSum(nums []int) []int {
	sum_prev := 0
	for i := range nums {
		nums[i] += sum_prev
		sum_prev = nums[i]
	}
	return nums
}

func Test_runningSum(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	assert.Equal(t, []int{1, 3, 6, 10}, runningSum(nums))

	nums = []int{1, 1, 1, 1, 1}
	assert.Equal(t, []int{1, 2, 3, 4, 5}, runningSum(nums))
}
