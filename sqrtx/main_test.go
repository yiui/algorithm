package sqrtx

import "testing"

/*
示例 1：

输入：x = 4
输出：2
示例 2：

输入：x = 8
输出：2
解释：8 的算术平方根是 2.82842..., 由于返回类型是整数，小数部分将被舍去。

*/
func TestMySqrt(t *testing.T) {
	x := 3
	except := 1
	sqrt := MySqrt(x)
	if sqrt == except {
		t.Log("success")
	} else {
		t.Errorf("err except: %v,result:%v\n", except, sqrt)
	}
	x = 8
	except = 2
	sqrt = MySqrt(x)
	if sqrt == except {
		t.Log("success")
	} else {
		t.Errorf("err except: %v,result:%v\n", except, sqrt)
	}
	//x := 5
	//except := 2
	//sqrt := MySqrt(x)
	//if sqrt == except {
	//	t.Log("success")
	//} else {
	//	t.Errorf("err except: %v,result:%v\n", except, sqrt)
	//}
}
