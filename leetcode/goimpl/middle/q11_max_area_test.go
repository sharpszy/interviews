package middle

import (
	"goimpl/common"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 暴力解法
func maxArea(height []int) int {
	l := len(height)
	if l < 2 {
		return 0
	}

	area := 0
	for i := 0; i < l; i++ {
		for j := l - 1; j > i; j-- {
			area = common.Max(area, common.Min(height[i], height[j])*(j-i))
		}
	}
	return area
}

// 最优解法：双指针
func maxArea2(height []int) int {
	l, r := 0, len(height)-1
	area := 0
	for l < r {
		area = common.Max(area, (r-l)*common.Min(height[r], height[l]))
		// 移动高度较小的那一端
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return area
}

func Test_maxArea(t *testing.T) {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	assert.Equal(t, 49, maxArea(height))
	assert.Equal(t, 49, maxArea2(height))

	height = []int{1, 1}
	assert.Equal(t, 1, maxArea(height))
	assert.Equal(t, 1, maxArea2(height))
}
