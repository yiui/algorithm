package merge_two_sorted_lists

/*
将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeTwoLists1(list1 *ListNode, list2 *ListNode) *ListNode {
	List1Current := list1
	List2Current := list2
	var ReListNode *ListNode
	var ListCurrent *ListNode
	var temp *ListNode
	for {
		if List1Current == nil && List2Current == nil {
			return ReListNode
		}
		if List1Current == nil && List2Current != nil {
			temp = List2Current
			if ListCurrent == nil {
				ListCurrent = temp
				ReListNode = ListCurrent
			} else {
				ListCurrent.Next = temp
				ListCurrent = ListCurrent.Next
			}
			return ReListNode
		}
		if List1Current != nil && List2Current == nil {
			temp = List1Current
			if ListCurrent == nil {
				ListCurrent = temp
				ReListNode = ListCurrent
			} else {
				ListCurrent.Next = temp
				ListCurrent = ListCurrent.Next
			}
			return ReListNode
		}
		if List2Current.Val >= List1Current.Val {
			temp = List1Current
			List1Current = List1Current.Next
			temp.Next = nil
			if ListCurrent == nil {
				ListCurrent = temp
				ReListNode = ListCurrent
			} else {
				ListCurrent.Next = temp
				ListCurrent = ListCurrent.Next
			}
		} else {
			temp = List2Current
			List2Current = List2Current.Next
			temp.Next = nil
			if ListCurrent == nil {
				ListCurrent = temp
				ReListNode = ListCurrent
			} else {
				ListCurrent.Next = temp
				ListCurrent = ListCurrent.Next
			}
		}
	}
}

//优化
func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	List1Current := list1
	List2Current := list2
	var ReListNode *ListNode
	var ListCurrent *ListNode
	var temp *ListNode
	for {
		if List1Current == nil && List2Current == nil {
			return ReListNode
		}
		if List1Current == nil && List2Current != nil {
			temp = List2Current
			List2Current = nil
		}
		if List1Current != nil && List2Current == nil {
			temp = List1Current
			List1Current = nil
		}
		if List1Current != nil && List2Current != nil {
			if List2Current.Val >= List1Current.Val {
				temp = List1Current
				List1Current = List1Current.Next
				temp.Next = nil
			} else {
				temp = List2Current
				List2Current = List2Current.Next
				temp.Next = nil
			}
		}
		if ListCurrent == nil {
			ListCurrent = temp
			ReListNode = ListCurrent
		} else {
			ListCurrent.Next = temp
			ListCurrent = ListCurrent.Next
		}
	}
}
