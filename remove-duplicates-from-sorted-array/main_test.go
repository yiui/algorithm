package remove_duplicates_from_sorted_array

import (
	"testing"
)

/*
示例 1：

输入：nums = [1,1,2]
输出：2, nums = [1,2,_]
解释：函数应该返回新的长度 2 ，并且原数组 nums 的前两个元素被修改为 1, 2 。不需要考虑数组中超出新长度后面的元素。
示例 2：

输入：nums = [0,0,1,1,1,2,2,3,3,4]
输出：5, nums = [0,1,2,3,4]
解释：函数应该返回新的长度 5 ， 并且原数组 nums 的前五个元素被修改为 0, 1, 2, 3, 4 。不需要考虑数组中超出新长度后面的元素。

*/

func TestRemoveDuplicates(t *testing.T) {
	//nums := []int{1, 1, 2}
	//duplicates := RemoveDuplicates(nums)
	//if duplicates == 2 {
	//	t.Log("success")
	//} else {
	//	t.Errorf("err except:%v,but result:%v", 2, duplicates)
	//}

	//nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	//duplicates := RemoveDuplicates(nums)
	//if duplicates == 5 {
	//	t.Log("success")
	//} else {
	//	t.Errorf("err except:%v,but result:%v", 5, duplicates)
	//}
	nums := []int{1, 1, 2}
	exceptResult := 2
	duplicates := RemoveDuplicates(nums)
	if duplicates == exceptResult {
		t.Log("success")
	} else {
		t.Errorf("err except:%v,but result:%v", exceptResult, duplicates)
	}
}
