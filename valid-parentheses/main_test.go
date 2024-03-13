package valid_parentheses

import "testing"

/*
示例 1：

输入：s = "()"
输出：true
示例 2：

输入：s = "()[]{}"
输出：true
示例 3：

输入：s = "(]"
输出：false
*/

func TestIsValid(t *testing.T) {
	s := "()"
	valid := IsValid(s)
	if valid {
		t.Log("success")
	} else {
		t.Errorf("err except:%v,result:%v", true, valid)
	}
	s = "()[]{}"
	valid = IsValid(s)
	if valid {
		t.Log("success")
	} else {
		t.Errorf("err except:%v,result:%v", true, valid)
	}
	s = "(]"
	valid = IsValid(s)
	if valid == false {
		t.Log("success")
	} else {
		t.Errorf("err except:%v,result:%v", false, valid)
	}
	s = "([)]"
	valid = IsValid(s)
	if valid == false {
		t.Log("success")
	} else {
		t.Errorf("err except:%v,result:%v", false, valid)
	}

	s = "[([]])"
	valid = IsValid(s)
	if valid == false {
		t.Log("success")
	} else {
		t.Errorf("err except:%v,result:%v", false, valid)
	}
	//{[]}

	s = "{[]}"
	valid = IsValid(s)
	if valid == true {
		t.Log("success")
	} else {
		t.Errorf("err except:%v,result:%v", true, valid)
	}

	s = "(([]){})"
	valid = IsValid(s)
	if valid == true {
		t.Log("success")
	} else {
		t.Errorf("err except:%v,result:%v", true, valid)
	}
}
