package search_insert_position

import "testing"

/*
示例 1:

输入: nums = [1,3,5,6], target = 5
输出: 2
示例 2:

输入: nums = [1,3,5,6], target = 2
输出: 1
示例 3:

输入: nums = [1,3,5,6], target = 7
输出: 4

*/

func TestSearchInsert(t *testing.T) {
	nums := []int{1, 3, 5, 6}
	target := 5
	exceptResult := 2
	insert := SearchInsert(nums, target)
	if insert == exceptResult {
		t.Log("success")
	} else {
		t.Errorf("err except:%v,result:%v", exceptResult, insert)
	}

	nums = []int{1, 3, 5, 6}
	target = 2
	exceptResult = 1
	insert = SearchInsert(nums, target)
	if insert == exceptResult {
		t.Log("success")
	} else {
		t.Errorf("err except:%v,result:%v", exceptResult, insert)
	}
	nums = []int{1, 3, 5, 6}
	target = 7
	exceptResult = 4
	insert = SearchInsert(nums, target)
	if insert == exceptResult {
		t.Log("success")
	} else {
		t.Errorf("err except:%v,result:%v", exceptResult, insert)
	}
	nums = []int{1, 3, 5, 6}
	target = 0
	exceptResult = 0
	insert = SearchInsert(nums, target)
	if insert == exceptResult {
		t.Log("success")
	} else {
		t.Errorf("err except:%v,result:%v", exceptResult, insert)
	}
	nums = []int{1}
	target = 1
	exceptResult = 0
	insert = SearchInsert(nums, target)
	if insert == exceptResult {
		t.Log("success")
	} else {
		t.Errorf("err except:%v,result:%v", exceptResult, insert)
	}
	//nums := []int{1, 3}
	//target := 2
	//exceptResult := 1
	//insert := SearchInsert(nums, target)
	//if insert == exceptResult {
	//	t.Log("success")
	//} else {
	//	t.Errorf("err except:%v,result:%v", exceptResult, insert)
	//}
}
