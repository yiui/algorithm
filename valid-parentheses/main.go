package valid_parentheses

import (
	"strings"
)

/*
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
每个右括号都有一个对应的相同类型的左括号。
*/
func IsValid(s string) bool {
	rule := map[string]int{
		"(": 1,
		")": -1,
		"{": 2,
		"}": -2,
		"[": 3,
		"]": -3,
	}
	split := strings.Split(s, "")
	sl := len(split)
	var temp = []int{}
	var sum int
	var skipInsert bool = false
	for i, x := range split {
		if i == 0 && rule[x] < 0 {
			return false
		}
		if i == sl-1 && rule[x] > 0 {
			return false
		}
		if len(temp) > 0 && rule[x] < 0 && temp[len(temp)-1] > 0 {
			if rule[x]+temp[len(temp)-1] == 0 {
				skipInsert = true
				temp = temp[:len(temp)-1]
			} else {
				return false
			}
		}
		if skipInsert == false {
			temp = append(temp, rule[x])
		}
		if skipInsert == true {
			skipInsert = false
		}
		sum += rule[x]
	}
	return sum == 0
}
