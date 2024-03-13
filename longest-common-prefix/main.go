package longest_common_prefix

import "strings"

/*
编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。

*/
func LongestCommonPrefix(strs []string) string {
	var builder strings.Builder
	s1 := strs[0]
	split := strings.Split(s1, "")
	for _, x := range split {
		for _, q := range strs {
			if !strings.HasPrefix(q, builder.String()+x) {
				return builder.String()
			}
		}
		builder.WriteString(x)
	}
	return builder.String()
}
