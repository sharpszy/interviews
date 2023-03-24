package middle

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Q6: N 字形变换
// https://leetcode.cn/problems/zigzag-conversion/
// 将一个给定字符串 s 根据给定的行数 numRows ，以从上往下、从左到右进行 Z 字形排列。
// 比如输入字符串为 "PAYPALISHIRING" 行数为 3 时，排列如下：
//
// P   A   H   N
// A P L S I I G
// Y   I   R
// 之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："PAHNAPLSIIGYIR"。
//
// 请你实现这个将字符串进行指定行数变换的函数：
//
// 示例 1：
// 输入：s = "PAYPALISHIRING", numRows = 3
// 输出："PAHNAPLSIIGYIR"
//
// 示例 2：
// 输入：s = "PAYPALISHIRING", numRows = 4
// 输出："PINALSIGYAHRPI"
// 解释：
// P     I    N
// A   L S  I G
// Y A   H R
// P     I

// 二维矩阵模拟 时间复杂度O(r⋅n)，空间复杂度O(r⋅n)
func zigzag_conversion(s string, rows int) string {
	l, r := len(s), rows
	if r == 1 || l <= r {
		return s
	}

	t := 2*r - 2             // 每个周期字符数
	pc := r - 1              // 每个周期的列数(1 + r - 2)
	c := pc * l / t          // 总列数
	mat := make([][]byte, r) // n行的二维矩阵
	for i := range mat {     // 初始化矩阵
		mat[i] = make([]byte, c)
	}

	x, y := 0, 0
	for i := range s {
		mat[x][y] = s[i]
		if i%t < r-1 {
			x++ // 向下移动
		} else { // 右上移动
			x--
			y++
		}
	}

	ans := make([]byte, 0, l)
	for _, row := range mat {
		for _, ch := range row {
			if ch > 0 {
				ans = append(ans, ch)
			}
		}
	}
	return string(ans)
}

// 压缩矩阵空间 时间复杂度O(n) 空间复杂度O(n)
func zigzag_conversion2(s string, rows int) string {
	r := rows
	if r == 1 || r >= len(s) {
		return s
	}

	mat := make([][]byte, r)
	t, x := r*2-2, 0
	for i, ch := range s {
		mat[x] = append(mat[x], byte(ch))
		if i%t < r-1 {
			x++
		} else {
			x--
		}
	}
	return string(bytes.Join(mat, nil))
}

// https://leetcode.cn/problems/zigzag-conversion/solution/zzi-xing-bian-huan-by-jyd/
// 天秀解法，flag非常妙
func zigzag_conversion3(s string, rows int) string {
	if rows == 1 || rows >= len(s) {
		return s
	}

	mat := make([][]byte, rows)
	for i := range mat {
		mat[i] = make([]byte, 0)
	}

	x, flag := 0, -1
	for i := range s {
		mat[x] = append(mat[x], s[i])
		if x == 0 || x == (rows-1) {
			flag = -flag
		}
		x += flag
	}
	return string(bytes.Join(mat, nil))
}

func Test_zigzag_conversion(t *testing.T) {
	s := "PAYPALISHIRING"
	rows := 3
	assert.Equal(t, "PAHNAPLSIIGYIR", zigzag_conversion(s, rows))
	assert.Equal(t, "PAHNAPLSIIGYIR", zigzag_conversion2(s, rows))
	assert.Equal(t, "PAHNAPLSIIGYIR", zigzag_conversion3(s, rows))
}
