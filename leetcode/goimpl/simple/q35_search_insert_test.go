package simple

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Q35: 搜索插入位置
// 给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
// 请必须使用时间复杂度为 O(log n) 的算法。
//
// 示例 1:
// 输入: nums = [1,3,5,6], target = 5
// 输出: 2
//
// 示例 2:
// 输入: nums = [1,3,5,6], target = 2
// 输出: 1
//
// 示例 3:
// 输入: nums = [1,3,5,6], target = 7
// 输出: 4

// 二分法查找
func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	start, end := 0, len(nums)
	mid := (start + end) / 2
	for start < mid {
		if nums[mid] > target {
			end = mid
			mid = (start + end) / 2
		} else if nums[mid] < target {
			start = mid
			mid = (start + end) / 2
		} else {
			break
		}
	}
	if nums[mid] < target {
		mid = mid + 1
	}
	return mid
}

func Test_searchInsert(t *testing.T) {
	var (
		nums   = []int{1, 3, 5, 6}
		target = 5
	)
	assert.Equal(t, 2, searchInsert(nums, target))

	nums = []int{1, 3, 5, 6}
	target = 2
	assert.Equal(t, 1, searchInsert(nums, target))

	nums = []int{1, 3, 5, 6}
	target = 0
	assert.Equal(t, 0, searchInsert(nums, target))

	nums = []int{1, 3, 5, 6}
	target = 7
	assert.Equal(t, 4, searchInsert(nums, target))
}
