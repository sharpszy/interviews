package middle

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Q43. 字符串相乘
// https://leetcode.cn/problems/multiply-strings/
// 给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。
// 注意：不能使用任何内置的 BigInteger 库或直接将输入转换为整数。
//
// 示例 1:
// 输入: num1 = "2", num2 = "3"
// 输出: "6"
//
// 示例 2:
// 输入: num1 = "123", num2 = "456"
// 输出: "56088"

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	if len(num1) > len(num2) {
		num1, num2 = num2, num1
	}

	var (
		str             string
		buf             []int = make([]int, len(num1)+len(num2)) // 需要处理高位为0的情况
		a, b            int
		mul, mod, carry int
	)
	for i := len(num1) - 1; i >= 0; i-- {
		carry = 0
		a = int(num1[i] - '0')
		for j := len(num2) - 1; j >= 0; j-- {
			b = int(num2[j] - '0')
			mul = a*b + carry
			k := (len(buf) - 1) - (len(num1) - 1 - i) - (len(num2) - 1 - j)
			sum := buf[k] + mul
			mod = sum % 10
			buf[k] = mod
			carry = sum / 10
			for carry > 0 {
				k--
				sum = buf[k] + carry
				mod = sum % 10
				carry = sum / 10
				buf[k] = mod
			}
		}
	}

	var f bool
	for i := 0; i < len(buf); i++ {
		if buf[i] == 0 && !f {
			continue
		}
		f = true
		str += fmt.Sprintf("%v", buf[i])
	}
	if str == "" {
		str = "0"
	}
	return str
}

func Test_multiply(t *testing.T) {
	num1 := "2"
	num2 := "3"
	assert.Equal(t, "6", multiply(num1, num2))

	num1 = "123"
	num2 = "456"
	assert.Equal(t, "56088", multiply(num1, num2))

	num1 = "12345"
	num2 = "789"
	assert.Equal(t, "9740205", multiply(num1, num2))

	num1 = "999"
	num2 = "9"
	assert.Equal(t, "8991", multiply(num1, num2))
}
