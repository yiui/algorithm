package find_the_index_of_the_first_occurrence_in_a_string

import (
	"regexp"
)

/*
28. 找出字符串中第一个匹配项的下标
给你两个字符串 haystack 和 needle ，请你在 haystack 字符串中找出 needle 字符串的第一个匹配项的下标（下标从 0 开始）。如果 needle 不是 haystack 的一部分，则返回  -1 。

*/
func StrStr(haystack string, needle string) int {
	compile, _ := regexp.Compile(needle)
	loc := compile.FindStringIndex(haystack)
	if loc == nil {
		return -1
	}
	return loc[0]
}

//应该会有其他解法
