package remove_duplicates_from_sorted_list

import (
	"fmt"
	"testing"
)

/*
输入：head = [1,1,2]
输出：[1,2]

输入：head = [1,1,2,3,3]
输出：[1,2,3]
*/
func TestDeleteDuplicates(t *testing.T) {
	s1 := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val:  2,
				Next: nil,
			},
		},
	}

	duplicates := DeleteDuplicates(&s1)
	if duplicates != nil {
		for {
			fmt.Println("val=", duplicates.Val)
			duplicates = duplicates.Next
			if duplicates == nil {
				break
			}
		}
	}

	s1 = ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val:  3,
						Next: nil,
					},
				},
			},
		},
	}

	duplicates = DeleteDuplicates(&s1)
	if duplicates != nil {
		for {
			fmt.Println("val=", duplicates.Val)
			duplicates = duplicates.Next
			if duplicates == nil {
				break
			}
		}
	}
}
