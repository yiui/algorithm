package longest_common_prefix

import "testing"

/*
编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。

*/

/*
示例 1：

输入：strs = ["flower","flow","flight"]
输出："fl"
示例 2：

输入：strs = ["dog","racecar","car"]
输出：""
解释：输入不存在公共前缀。
*/

func TestLongestCommonPrefix(t *testing.T) {
	strs := []string{"flower", "flow", "flight"}
	prefix := LongestCommonPrefix(strs)
	if prefix != "fl" {
		t.Fatal("error, except:fl, result:", prefix)
	} else {
		t.Log("success")
	}
	strs = []string{"dog", "racecar", "car"}
	prefix = LongestCommonPrefix(strs)
	if prefix != "" {
		t.Fatal("error, except:'', result:", prefix)
	} else {
		t.Log("success")
	}
}
