package remove_duplicates_from_sorted_list

/*

83. 删除排序链表中的重复元素
给定一个已排序的链表的头 head ， 删除所有重复的元素，使每个元素只出现一次 。返回 已排序的链表 。
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func DeleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	current := head
	for {
		if current.Next == nil {
			return head
		}
		if current.Next.Val == current.Val {
			current.Next = current.Next.Next
		}
		if current.Next == nil {
			return head
		}
		if current.Next.Val != current.Val {
			current = current.Next
		}
	}
}
