package length_of_last_word

import "testing"

/*

示例 1：

输入：s = "Hello World"
输出：5
解释：最后一个单词是“World”，长度为5。
示例 2：

输入：s = "   fly me   to   the moon  "
输出：4
解释：最后一个单词是“moon”，长度为4。
示例 3：

输入：s = "luffy is still joyboy"
输出：6
解释：最后一个单词是长度为6的“joyboy”。
*/

func TestLengthOfLastWord(t *testing.T) {
	s := "Hello World"
	exceptResult := 5
	word := LengthOfLastWord(s)
	if word == exceptResult {
		t.Log("success")
	} else {
		t.Errorf("err except:%v,but result:%v\n", exceptResult, word)
	}

	s = "   fly me   to   the moon  "
	exceptResult = 4
	word = LengthOfLastWord(s)
	if word == exceptResult {
		t.Log("success")
	} else {
		t.Errorf("err except:%v,but result:%v\n", exceptResult, word)
	}
	s = "luffy is still joyboy"
	exceptResult = 6
	word = LengthOfLastWord(s)
	if word == exceptResult {
		t.Log("success")
	} else {
		t.Errorf("err except:%v,but result:%v\n", exceptResult, word)
	}
}
