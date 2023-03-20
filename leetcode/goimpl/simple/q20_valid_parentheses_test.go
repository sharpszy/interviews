package simple

import (
	"goimpl/common"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Q20: 有效的括号
// https://leetcode.cn/problems/valid-parentheses/
// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
// 有效字符串需满足：
// 左括号必须用相同类型的右括号闭合。
// 左括号必须以正确的顺序闭合。
// 每个右括号都有一个对应的相同类型的左括号。
// 补充：括号可以嵌套
//
// 示例 1：
// 输入：s = "()"
// 输出：true
//
// 示例 2：
// 输入：s = "()[]{}"
// 输出：true
//
// 示例 3：
// 输入：s = "{[()]}"
// 输出：true

func validParentheses(s string) bool {
	l := len(s)
	if l%2 == 1 {
		return false
	}

	pairs := map[byte]byte{'(': ')', '[': ']', '{': '}'}

	stack := common.NewStack()
	for _, c := range []byte(s) {
		if v, ok := pairs[c]; ok {
			stack.Push(v)
		} else {
			if v, ok := stack.Pop(); ok && c == v {
				continue
			} else {
				return false
			}
		}
	}
	return stack.IsEmpty()
}

func Test_validParentheses(t *testing.T) {
	str := "()"
	assert.True(t, validParentheses(str))

	str = "()[]{}"
	assert.True(t, validParentheses(str))

	str = "{[()]}"
	assert.True(t, validParentheses(str))

	str = "{{[]]}"
	assert.False(t, validParentheses(str))
}
