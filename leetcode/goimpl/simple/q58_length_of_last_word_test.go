package simple

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func lengthOfLastWord(s string) int {
	var (
		left, right = len(s) - 1, len(s) - 1
	)
	for left >= 0 {
		// right指向第一个非空字符
		if s[right] == ' ' {
			right--
			left--
			continue
		}
		// left指向单词后第一个空字符
		if s[left] == ' ' {
			break
		}
		left--
	}
	return right - left
}

// 简短直观方法
func lengthOfLastWord2(s string) int {
	idx, count := len(s)-1, 0
	for idx >= 0 {
		if s[idx] != ' ' {
			count++
		} else if count > 0 {
			break
		}
		idx--
	}
	return count
}

func Test_lengthOfLastWord(t *testing.T) {
	s := "Hello World"
	assert.Equal(t, 5, lengthOfLastWord(s))

	s = "   fly me   to   the moon  "
	assert.Equal(t, 4, lengthOfLastWord(s))

	s = "luffy is still joyboy"
	assert.Equal(t, 6, lengthOfLastWord(s))
}

func Test_lengthOfLastWord2(t *testing.T) {
	s := "Hello World"
	assert.Equal(t, 5, lengthOfLastWord2(s))

	s = "   fly me   to   the moon  "
	assert.Equal(t, 4, lengthOfLastWord2(s))

	s = "luffy is still joyboy"
	assert.Equal(t, 6, lengthOfLastWord2(s))
}
