package simple

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Q9: 回文数
// https://leetcode.cn/problems/palindrome-number/
// 给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。
// 回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
//
// 例如，121 是回文，而 123 不是。
// 示例 1：
// 输入：x = 121
// 输出：true
// 示例 2：
//
// 输入：x = -121
// 输出：false
// 解释：从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
// 示例 3：
//
// 输入：x = 10
// 输出：false
// 解释：从右向左读, 为 01 。因此它不是一个回文数。

// 空间复杂度O(1)，时间复杂度O(N)
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	} else if x < 10 {
		return true
	}

	reverse, tmp := 0, x
	for tmp >= 10 {
		reverse = reverse*10 + tmp%10
		tmp = tmp / 10
	}
	// if tmp > 0 {
	reverse = reverse*10 + tmp
	// }
	return reverse == x
}

// 空间复杂度O(1)，时间复杂度O(logN)
func isPalindrome2(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	revertHalfNum := 0
	for x > revertHalfNum {
		revertHalfNum = revertHalfNum*10 + x%10
		x = x / 10
	}
	return x == revertHalfNum || x == revertHalfNum/10
}

func Test_isPalindrome(t *testing.T) {
	nums := []int{121, 1221, 1, 10001}
	for _, n := range nums {
		assert.True(t, isPalindrome(n))
		assert.True(t, isPalindrome2(n))
	}

	nums = []int{120, 122, 10}
	for _, n := range nums {
		assert.False(t, isPalindrome(n))
		assert.False(t, isPalindrome2(n))
	}
}
