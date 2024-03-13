package find_the_index_of_the_first_occurrence_in_a_string

import "testing"

/*
示例 1：

输入：haystack = "sadbutsad", needle = "sad"
输出：0
解释："sad" 在下标 0 和 6 处匹配。
第一个匹配项的下标是 0 ，所以返回 0 。
示例 2：

输入：haystack = "leetcode", needle = "leeto"
输出：-1
解释："leeto" 没有在 "leetcode" 中出现，所以返回 -1 。

*/

func TestStrStr(t *testing.T) {
	haystack := "sadbutsad"
	needle := "sad"
	exceptResult := 0
	str := StrStr(haystack, needle)
	if str == exceptResult {
		t.Log("success")
	} else {
		t.Errorf("err except:%v,result:%v", exceptResult, str)
	}

	haystack = "leetcode"
	needle = "leeto"
	exceptResult = -1
	str = StrStr(haystack, needle)
	if str == exceptResult {
		t.Log("success")
	} else {
		t.Errorf("err except:%v,result:%v", exceptResult, str)
	}
}
