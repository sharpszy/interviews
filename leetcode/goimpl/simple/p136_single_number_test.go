package simple

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Q136. 只出现一次的数字
// https://leetcode.cn/problems/single-number/
// 给你一个 非空 整数数组 nums ，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
//
// 你必须设计并实现线性时间复杂度的算法来解决此问题，且该算法只使用常量额外空间。
//
// 示例 1 ：
// 输入：nums = [2,2,1]
// 输出：1
//
// 示例 2 ：
// 输入：nums = [4,1,2,1,2]
// 输出：4

// 异或运算，异或满足交换律和结合律 a ^ b ^ a = b
func singleNumber(nums []int) int {
	single := 0
	for _, x := range nums {
		single ^= x
	}
	return single
}

func Test_singleNumber(t *testing.T) {
	nums := []int{2, 2, 1}
	assert.Equal(t, 1, singleNumber(nums))

	nums = []int{4, 1, 2, 1, 2}
	assert.Equal(t, 4, singleNumber(nums))
}
