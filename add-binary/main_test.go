package add_binary

import "testing"

/*
示例 1：

输入:a = "11", b = "1"
输出："100"
示例 2：

输入：a = "1010", b = "1011"
输出："10101"
*/

func TestAddBinary(t *testing.T) {
	a := "11"
	b := "1"
	except := "100"

	binary := AddBinary(a, b)
	if binary == except {
		t.Log("success")
	} else {
		t.Errorf("err except:%v,but result:%v\n", except, binary)
	}
	a = "1010"
	b = "1011"
	except = "10101"

	binary = AddBinary(a, b)
	if binary == except {
		t.Log("success")
	} else {
		t.Errorf("err except:%v,but result:%v\n", except, binary)
	}
}
