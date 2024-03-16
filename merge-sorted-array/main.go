package merge_sorted_array

import (
	"fmt"
	"sort"
)

/*
88. 合并两个有序数组

提示
给你两个按 非递减顺序 排列的整数数组 nums1 和 nums2，另有两个整数 m 和 n ，分别表示 nums1 和 nums2 中的元素数目。

请你 合并 nums2 到 nums1 中，使合并后的数组同样按 非递减顺序 排列。

注意：最终，合并后数组不应由函数返回，而是存储在数组 nums1 中。为了应对这种情况，nums1 的初始长度为 m + n，其中前 m 个元素表示应合并的元素，后 n 个元素为 0 ，应忽略。nums2 的长度为 n 。
*/
//方法一
func Merge1(nums1 []int, m int, nums2 []int, n int) {
	nums1ZeroPosition := m
	nums1Point := 0
	nums2Point := 0
	for {
		if nums2Point > n-1 {
			break
		}
		if nums1ZeroPosition == 0 {
			nums1[nums1Point] = nums2[nums2Point]
			nums2Point += 1
			nums1ZeroPosition += 1
		} else if nums1Point+1 == nums1ZeroPosition {
			if nums2[nums2Point] >= nums1[nums1Point] {
				nums1ZeroPosition += 1
				nums1[nums1Point+1] = nums2[nums2Point]
				nums2Point += 1
			} else {
				for j := nums1ZeroPosition; j > nums1Point; j-- {
					nums1[j] = nums1[j-1]
				}
				nums1[nums1Point] = nums2[nums2Point]
				nums1ZeroPosition += 1
				nums2Point += 1
			}
		} else if nums1Point+1 != nums1ZeroPosition {
			if nums2[nums2Point] < nums1[nums1Point] {
				for j := nums1ZeroPosition; j > nums1Point; j-- {
					nums1[j] = nums1[j-1]
				}
				nums1ZeroPosition += 1
				nums1[nums1Point] = nums2[nums2Point]
				nums2Point += 1
			} else if nums2[nums2Point] >= nums1[nums1Point] && nums2[nums2Point] < nums1[nums1Point+1] {
				for j := nums1ZeroPosition; j > nums1Point; j-- {
					nums1[j] = nums1[j-1]
				}
				nums1ZeroPosition += 1
				nums1[nums1Point+1] = nums2[nums2Point]
				nums2Point += 1
			} else {
				nums1Point += 1
			}
		}
	}
}

//方法二
func Merge2(nums1 []int, m int, nums2 []int, n int) {
	nums1ZeroPosition := m
	nums1Point := 0
	nums2Point := 0
	//var nums2Has = true
	var temp = append([]int{}, nums1...)
	tempEndPoint := 0
	//tempHas := false
	tempStartPoint := 0
	//最小的数在哪边
	maxNum := 2
	//0:temp ,nums1:1,nums2:2
	if n != 0 {
		for {
			if maxNum == 2 {
				if nums1ZeroPosition == nums1Point {
					nums1[nums1Point] = nums2[nums2Point]
					nums2Point += 1
					nums1ZeroPosition += 1
					nums1Point += 1
				} else if nums2[nums2Point] < nums1[nums1Point] {
					temp[tempEndPoint] = nums1[nums1Point]
					nums1[nums1Point] = nums2[nums2Point]
					tempEndPoint += 1
					nums1Point += 1
					nums2Point += 1
				} else {
					nums1Point += 1
				}
			} else if maxNum == 0 {
				if nums1ZeroPosition == nums1Point {
					nums1[nums1Point] = temp[tempStartPoint]
					tempStartPoint += 1
					nums1Point += 1
					nums1ZeroPosition += 1
				} else if temp[tempStartPoint] < nums1[nums1Point] {
					temp[tempEndPoint] = nums1[nums1Point]
					nums1[nums1Point] = temp[tempStartPoint]
					nums1Point += 1
					tempEndPoint += 1
					tempStartPoint += 1
				} else {
					nums1Point += 1
				}
			}
			if tempEndPoint-tempStartPoint > 0 && nums2Point < n {
				if temp[tempStartPoint] > nums2[nums2Point] {
					maxNum = 2
				} else {
					maxNum = 0
				}
			} else if tempEndPoint-tempStartPoint == 0 && nums2Point < n {
				maxNum = 2
			} else if tempEndPoint-tempStartPoint > 0 && nums2Point == n {
				maxNum = 0
			} else if tempEndPoint-tempStartPoint == 0 && nums2Point == n {
				break
			}
		}
	}
	fmt.Println(nums1)
}

//方法三
func Merge3(nums1 []int, m int, nums2 []int, n int) {
	copy(nums1[m:], nums2)
	sort.Ints(nums1)
	fmt.Println(nums1)
}

//方法四
func Merge(nums1 []int, m int, nums2 []int, n int) {
	nums1Point := m - 1
	nums2Point := n - 1
	newPoint := m + n - 1
	for i := m - 1; i < m; i-- {
		if nums2Point < 0 && nums1Point < 0 {
			break
		}
		if nums2Point < 0 && nums1Point >= 0 {
			nums1[newPoint] = nums1[nums1Point]
			newPoint -= 1
			nums1Point -= 1
		} else if nums2Point >= 0 && nums1Point < 0 {
			nums1[newPoint] = nums2[nums2Point]
			newPoint -= 1
			nums2Point -= 1
		} else if nums1[nums1Point] > nums2[nums2Point] {
			nums1[newPoint] = nums1[nums1Point]
			newPoint -= 1
			nums1Point -= 1
		} else if nums1[nums1Point] <= nums2[nums2Point] {
			nums1[newPoint] = nums2[nums2Point]
			newPoint -= 1
			nums2Point -= 1
		}
	}
	fmt.Println(nums1)
}
