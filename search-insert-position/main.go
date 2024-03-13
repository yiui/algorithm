package search_insert_position

/*
35. 搜索插入位置
给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
请必须使用时间复杂度为 O(log n) 的算法。
*/
func SearchInsert1(nums []int, target int) int {
	numsLen := len(nums)
	for i := 0; i < numsLen; i++ {
		if i == 0 && target <= nums[i] {
			return 0
		}
		if i != numsLen-1 {
			if target == nums[i] {
				return i
			} else if target > nums[i] && target < nums[i+1] {
				return i + 1
			}
		}
		if i == numsLen-1 {
			if target == nums[i] {
				return i
			} else if target > nums[i] {
				return i + 1
			}

		}
	}
	return numsLen
}

//二分法
func SearchInsert(nums []int, target int) int {
	numsLen := len(nums)
	left, right := 0, numsLen-1
	for right >= left {
		if target < nums[left] {
			return 0
		} else if target == nums[left] {
			return left
		} else if target > nums[left] && target < nums[right] {
			if left+1 <= right && target > nums[left] && target < nums[left+1] {
				return left + 1
			}
			if left <= right-1 && target > nums[right-1] && target < nums[right] {
				return right
			}
		} else if target == nums[right] {
			return right
		}
		left += 1
		right -= 1
	}
	return numsLen
}
