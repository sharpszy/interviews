package simple

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Q26: 删除有序数组中的重复项
// https://leetcode.cn/problems/remove-duplicates-from-sorted-array/
// 给你一个 升序排列 的数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。元素的 相对顺序 应该保持 一致 。
//
// 由于在某些语言中不能改变数组的长度，所以必须将结果放在数组nums的第一部分。更规范地说，如果在删除重复项之后有 k 个元素，那么 nums 的前 k 个元素应该保存最终结果。
//
// 示例 1：
// 输入：nums = [1,1,2]
// 输出：2, nums = [1,2,_]
// 解释：函数应该返回新的长度 2 ，并且原数组 nums 的前两个元素被修改为 1, 2 。不需要考虑数组中超出新长度后面的元素
//
// 示例 2：
// 输入：nums = [0,0,1,1,1,2,2,3,3,4]
// 输出：5, nums = [0,1,2,3,4]
// 解释：函数应该返回新的长度 5 ， 并且原数组 nums 的前五个元素被修改为 0, 1, 2, 3, 4 。不需要考虑数组中超出新长度后面的元素。

func removeDuplicates(nums []int) int {
	var (
		left  = 0
		right = 1
		l     = len(nums)
	)
	if l < 2 {
		return l
	}

	for right < l {
		if nums[left] != nums[right] {
			// 去掉注释仍然成立
			// if right-left > 1 { 中间有连续超过2个以上重复的元素，挪动右侧的元素至left+1的位置
			nums[left+1] = nums[right]
			// }
			left++
		}
		right++
	}
	return left + 1
}

func Test_removeDuplicates(t *testing.T) {
	nums := []int{1, 1, 2}
	r := removeDuplicates(nums)
	assert.Equal(t, 2, r)
	assert.Equal(t, []int{1, 2}, nums[:r])

	nums = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	r = removeDuplicates(nums)
	assert.Equal(t, 5, r)
	assert.Equal(t, []int{0, 1, 2, 3, 4}, nums[:r])
}
