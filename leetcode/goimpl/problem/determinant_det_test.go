package problem

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// 行列式定义：行数 = 列数 N行 * N列
func Determinant(matrix [][]int) int {
	// your code here
	if len(matrix) == 1 {
		return matrix[0][0]
	}

	var det = 0
	for i, x := range matrix[0] {
		minor := make([][]int, len(matrix)-1)
		// 初始化子矩阵
		for j := 0; j < len(minor); j++ {
			minor[j] = append(minor[j], matrix[j+1][0:i]...)
			minor[j] = append(minor[j], matrix[j+1][i+1:]...)
		}
		if i%2 == 0 {
			det += x * Determinant(minor)
		} else {
			det -= x * Determinant(minor)
		}
	}
	return det
}

func Test_Determinant(t *testing.T) {
	r := Determinant([][]int{{1}})
	assert.Equal(t, 1, r)

	r = Determinant([][]int{{1, 3}, {2, 5}})
	assert.Equal(t, -1, r)

	r = Determinant([][]int{{2, 5, 3}, {1, -2, -1}, {1, 3, 4}})
	assert.Equal(t, -20, r)

	r = Determinant([][]int{{2, 5, 3, 1}, {1, -2, -1, 2}, {1, 3, 4, 4}, {2, 3, 4, 4}})
	assert.Equal(t, 23, r)
}
