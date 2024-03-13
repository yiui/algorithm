package length_of_last_word

import (
	"regexp"
	"strings"
)

/*

58. 最后一个单词的长度
给你一个字符串 s，由若干单词组成，单词前后用一些空格字符隔开。返回字符串中 最后一个 单词的长度。

单词 是指仅由字母组成、不包含任何空格字符的最大
子字符串
。
*/
func LengthOfLastWord1(s string) int {
	compile, _ := regexp.Compile(`\w+`)
	allString := compile.FindAllString(s, -1)
	return len(allString[len(allString)-1])
}

// LengthOfLastWord2 有更好的其他方法
func LengthOfLastWord2(s string) int {
	compile, _ := regexp.Compile(`\s+`)
	allString := compile.Split(strings.TrimSpace(s), -1)
	return len(allString[len(allString)-1])
}

// LengthOfLastWord 探索更高效的方法
func LengthOfLastWord(s string) int {
	s = strings.TrimSpace(s)
	sLen := len(s)
	n := 0
	for i := sLen - 1; i >= 0; i-- {
		if string(s[i]) != " " {
			n += 1
		} else {
			if n != 0 {
				return n
			}
		}
	}
	return n
}
