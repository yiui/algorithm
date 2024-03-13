package add_binary

import (
	"strconv"
	"strings"
)

/*
67. 二进制求和
给你两个二进制字符串 a 和 b ，以二进制字符串的形式返回它们的和。
*/
func AddBinary(a string, b string) string {
	aSlice := strings.Split(a, "")
	bSlice := strings.Split(b, "")
	aLast := len(aSlice) - 1
	bLast := len(bSlice) - 1
	var rLast int
	var result []string
	if aLast >= bLast {
		rLast = aLast
		result = aSlice
	} else {
		rLast = bLast
		result = bSlice
	}
	yu := 0
	var aCurrent int
	var bCurrent int
	for {
		if aLast < 0 {
			aCurrent = 0
		} else {
			aCurrent, _ = strconv.Atoi(aSlice[aLast])
		}
		if bLast < 0 {
			bCurrent = 0
		} else {
			bCurrent, _ = strconv.Atoi(bSlice[bLast])
		}
		sum := aCurrent + bCurrent + yu
		if sum == 2 {
			result[rLast] = "0"
			yu = 1
		} else if sum == 1 {
			result[rLast] = "1"
			yu = 0
		} else if sum == 0 {
			result[rLast] = "0"
			yu = 0
		}
		if rLast == 0 {
			if yu == 0 {
				return strings.Join(result, "")
			} else {
				return "1" + strings.Join(result, "")
			}
		}
		if aLast >= 0 {
			aLast -= 1
		}
		if bLast >= 0 {
			bLast -= 1
		}
		rLast -= 1
	}
}
