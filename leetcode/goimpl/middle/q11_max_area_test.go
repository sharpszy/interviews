package middle

import (
	"goimpl/common"
	"testing"

	"github.com/stretchr/testify/assert"
)

func maxArea(height []int) int {
	l := len(height)
	if l < 2 {
		return 0
	}

	max := 0
	for i := 0; i < l; i++ {
		for j := l - 1; j > i; j-- {
			max = common.Max(max, common.Min(height[i], height[j])*(j-i))
		}
	}
	return max
}

func Test_maxArea(t *testing.T) {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	assert.Equal(t, 49, maxArea(height))

	height = []int{1, 1}
	assert.Equal(t, 1, maxArea(height))
}
