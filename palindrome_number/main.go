package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

/*
示例 1：

输入：x = 121
输出：true
示例 2：

输入：x = -121
输出：false
解释：从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
示例 3：

输入：x = 10
输出：false
解释：从右向左读, 为 01 。因此它不是一个回文数。
*/
func isPalindrome(x int) bool {
	var temp = x
	if x >= 0 && x < 10 {
		return true
	}
	for temp > 0 {
		last := temp % 10
		var Len = len(strconv.Itoa(temp))
		first := temp / (int(math.Pow(10, float64(Len-1))))
		if first != last {
			return false
		}
		temp = temp - int(math.Pow(10, float64(Len-1)))*first
		temp = (temp - last) / 10
		if temp >= 0 && temp < 10 {
			return true
		}
	}
	return false
}
func isPalindrome2(x int) bool {
	if x >= 0 && x < 10 {
		return true
	}
	s := strconv.Itoa(x)
	split := strings.Split(s, "")
	splitLen := len(split)
	for i := 0; i < splitLen/2; i++ {
		if split[i] != split[splitLen-i-1] {
			return false
		}
	}
	return true
}
func isPalindrome3(x int) bool {
	if x >= 0 && x < 10 {
		return true
	}
	s := strconv.Itoa(x)

	split := strings.Split(s, "")
	splitLen := len(split)
	for i := 0; i < splitLen/2; i++ {
		if split[i] != split[splitLen-i-1] {
			return false
		}
	}
	return true
}
func main() {
	num := 10
	palindrome := isPalindrome2(num)
	fmt.Println(palindrome)
}
