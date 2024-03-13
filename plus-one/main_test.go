package plus_one

import (
	"reflect"
	"testing"
)

/*
示例 1：

输入：digits = [1,2,3]
输出：[1,2,4]
解释：输入数组表示数字 123。
示例 2：

输入：digits = [4,3,2,1]
输出：[4,3,2,2]
解释：输入数组表示数字 4321。
示例 3：

输入：digits = [0]
输出：[1]

输入：digits = [9,8,9]
输出：[9,9,0]
*/

func TestPlusOne(t *testing.T) {
	digits := []int{9}
	exceptResult := []int{1, 0}
	one := PlusOne(digits)
	if reflect.DeepEqual(exceptResult, one) {
		t.Log("success")
	} else {
		t.Errorf("err except:%v,result:%v\n", exceptResult, one)
	}
	digits = []int{1, 2, 3}
	exceptResult = []int{1, 2, 4}
	one = PlusOne(digits)
	if reflect.DeepEqual(exceptResult, one) {
		t.Log("success")
	} else {
		t.Errorf("err except:%v,result:%v\n", exceptResult, one)
	}

	digits = []int{4, 3, 2, 1}
	exceptResult = []int{4, 3, 2, 2}
	one = PlusOne(digits)
	if reflect.DeepEqual(exceptResult, one) {
		t.Log("success")
	} else {
		t.Errorf("err except:%v,result:%v\n", exceptResult, one)
	}

	digits = []int{0}
	exceptResult = []int{1}
	one = PlusOne(digits)
	if reflect.DeepEqual(exceptResult, one) {
		t.Log("success")
	} else {
		t.Errorf("err except:%v,result:%v\n", exceptResult, one)
	}
	digits = []int{9, 8, 9}
	exceptResult = []int{9, 9, 0}
	one = PlusOne(digits)
	if reflect.DeepEqual(exceptResult, one) {
		t.Log("success")
	} else {
		t.Errorf("err except:%v,result:%v\n", exceptResult, one)
	}
}
