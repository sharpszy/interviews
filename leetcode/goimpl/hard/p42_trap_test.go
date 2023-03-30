package hard

import (
	"goimpl/common"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Q42. 接雨水
// https://leetcode.cn/problems/trapping-rain-water/
// 给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
// 示例 1：
// 输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
// 输出：6
// 解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。
//
// 示例 2：
// 输入：height = [4,2,0,3,2,5]
// 输出：9

func trap(height []int) (ans int) {
	left, right := 0, len(height)-1
	leftMax, rightMax := 0, 0
	for left < right {
		leftMax = common.Max(leftMax, height[left])
		rightMax = common.Max(rightMax, height[right])
		if height[left] < height[right] { // 左边小于右边，左边当前位置可以接到雨水
			ans += leftMax - height[left] // 当前位置可以接到的雨水，即为：高度差*宽度1
			left++
		} else {
			ans += rightMax - height[right]
			right--
		}
	}
	return
}

func Test_trap(t *testing.T) {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	assert.Equal(t, 6, trap(height))

	height = []int{4, 2, 0, 3, 2, 5}
	assert.Equal(t, 9, trap(height))
}
