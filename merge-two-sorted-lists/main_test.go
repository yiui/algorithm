package merge_two_sorted_lists

import (
	"fmt"
	"testing"
)

/*
输入：l1 = [1,2,4], l2 = [1,3,4]
输出：[1,1,2,3,4,4]
示例 2：

输入：l1 = [], l2 = []
输出：[]
示例 3：

输入：l1 = [], l2 = [0]
输出：[0]

*/
func TestMergeTwoLists(t *testing.T) {
	//l1 := &ListNode{
	//	Val: 1,
	//	Next: &ListNode{
	//		Val: 2,
	//		Next: &ListNode{
	//			Val:  4,
	//			Next: nil,
	//		},
	//	},
	//}
	//
	//l2 := &ListNode{
	//	Val: 1,
	//	Next: &ListNode{
	//		Val: 3,
	//		Next: &ListNode{
	//			Val:  4,
	//			Next: nil,
	//		},
	//	},
	//}
	//twoLists := MergeTwoLists(l1, l2)
	//fmt.Println(twoLists)
	//current := twoLists
	//for {
	//	fmt.Println("val=", current.Val)
	//	current = current.Next
	//	if current == nil {
	//		break
	//	}
	//}

	//var l1, l2 *ListNode
	//twoLists := MergeTwoLists(l1, l2)
	//fmt.Println(twoLists)
	//current := twoLists
	//for {
	//	if current == nil {
	//		break
	//	}
	//	fmt.Println("val=", current.Val)
	//	current = current.Next
	//
	//}
	//var l1, l2 *ListNode
	//l2 = &ListNode{
	//	Val:  0,
	//	Next: nil,
	//}
	//twoLists := MergeTwoLists(l1, l2)
	//fmt.Println(twoLists)
	//current := twoLists
	//for {
	//	if current == nil {
	//		break
	//	}
	//	fmt.Println("val=", current.Val)
	//	current = current.Next
	//
	//}
	/*
		list1 =[1,2,4]
		list2 =[1,3,4]
	*/
	var l1, l2 *ListNode
	l1 = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	l2 = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	twoLists := MergeTwoLists(l1, l2)
	fmt.Println(twoLists)
	current := twoLists
	for {
		if current == nil {
			break
		}
		fmt.Println("val=", current.Val)
		current = current.Next

	}

}
