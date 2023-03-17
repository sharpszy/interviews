package simple

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Q27: 移除元素
// https://leetcode.cn/problems/remove-element/
// 给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素，并返回移除后数组的新长度。
// 不要使用额外的数组空间，你必须仅使用 O(1) 额外空间并 原地 修改输入数组。
// 元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。
//
// 示例 1：
// 输入：nums = [3,2,2,3], val = 3
// 输出：2, nums = [2,2]
// 解释：函数应该返回新的长度 2, 并且 nums 中的前两个元素均为 2。你不需要考虑数组中超出新长度后面的元素。例如，函数返回的新长度为 2 ，而 nums = [2,2,3,3] 或 nums = [2,2,0,0]，也会被视作正确答案。
//
// 示例 2：
// 输入：nums = [0,1,2,2,3,0,4,2], val = 2
// 输出：5, nums = [0,1,4,0,3]
// 解释：函数应该返回新的长度 5, 并且 nums 中的前五个元素为 0, 1, 3, 0, 4。注意这五个元素可为任意顺序。你不需要考虑数组中超出新长度后面的元素。

func removeElement(nums []int, val int) int {
	i := -1
	for _, x := range nums {
		if x != val {
			i++
			nums[i] = x
		}
	}
	return i + 1
}

func Test_removeElement(t *testing.T) {
	nums := []int{3, 2, 2, 3}
	val := 3
	r := removeElement(nums, val)
	assert.Equal(t, 2, r)
	assert.Equal(t, []int{2, 2}, nums[:r])

	nums = []int{0, 1, 2, 2, 3, 0, 4, 2}
	val = 2
	r = removeElement(nums, val)
	assert.Equal(t, 5, r)
	assert.Equal(t, []int{0, 1, 3, 0, 4}, nums[:r])
}