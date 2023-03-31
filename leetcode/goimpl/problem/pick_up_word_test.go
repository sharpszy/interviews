package problem

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// 给你定一个包含单词和#分割的字符串，提取其中的单词并反向输出
// 输入：#Please##show####me#the####code
// 输出：code the me show Please

func pickUpWords(s string) string {
	var (
		str  string
		l, r = len(s) - 1, len(s) // 指向单词的左右两端非单词字符位置
	)
	for l >= -1 {
		if l == -1 || s[l] == '#' {
			if r-l > 1 {
				if len(str) == 0 {
					str = s[l+1 : r]
				} else {
					str += " " + s[l+1:r]
				}
			}
			r = l
		}
		l--
	}
	return str
}

func Test_(t *testing.T) {
	str := "#Please##show####me#the####code"
	assert.Equal(t, "code the me show Please", pickUpWords(str))

	str = "Please##show####me#the####code#"
	assert.Equal(t, "code the me show Please", pickUpWords(str))

	str = "Please##show####me#the####code"
	assert.Equal(t, "code the me show Please", pickUpWords(str))

	str = "code"
	assert.Equal(t, "code", pickUpWords(str))

	str = "###"
	assert.Equal(t, "", pickUpWords(str))

	str = ""
	assert.Equal(t, "", pickUpWords(str))
}
