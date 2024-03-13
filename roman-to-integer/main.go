package roman_to_integer

import (
	"strings"
)

/*
罗马数字包含以下七种字符: I， V， X， L，C，D 和 M。

字符          数值
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
例如， 罗马数字 2 写做 II ，即为两个并列的 1 。12 写做 XII ，即为 X + II 。 27 写做  XXVII, 即为 XX + V + II 。

通常情况下，罗马数字中小的数字在大的数字的右边。但也存在特例，例如 4 不写做 IIII，而是 IV。数字 1 在数字 5 的左边，所表示的数等于大数 5 减小数 1 得到的数值 4 。同样地，数字 9 表示为 IX。这个特殊的规则只适用于以下六种情况：

I 可以放在 V (5) 和 X (10) 的左边，来表示 4 和 9。
X 可以放在 L (50) 和 C (100) 的左边，来表示 40 和 90。
C 可以放在 D (500) 和 M (1000) 的左边，来表示 400 和 900。
给定一个罗马数字，将其转换成整数。

*/

func RomanToInt1(s string) int {
	spec := map[string]int{
		"IV": 4,
		"IX": 9,
		"XL": 40,
		"XC": 90,
		"CD": 400,
		"CM": 900,
	}
	if k, ok := spec[s]; ok {
		return k
	}
	split := strings.Split(s, "")
	var sum int = 0
	rule := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000}
	var skip bool
	sl := len(split)
	for i, x := range split {
		if skip {
			skip = false
			continue
		}
		var num int = 0
		switch x {
		case "I":
			if i+1 < sl && rule[split[i]] < rule[split[i+1]] {
				num = spec[split[i]+split[i+1]]
				skip = true
			} else {
				num = 1
			}
		case "V":
			num = 5
		case "X":
			if i+1 < sl && rule[split[i]] < rule[split[i+1]] {
				num = spec[split[i]+split[i+1]]
				skip = true
			} else {
				num = 10
			}
		case "L":
			num = 50
		case "C":
			if i+1 < sl && rule[split[i]] < rule[split[i+1]] {
				num = spec[split[i]+split[i+1]]
				skip = true
			} else {
				num = 100
			}
		case "D":
			num = 500
		case "M":
			num = 1000
		}
		sum += num
	}
	return sum
}
func RomanToInt(s string) int {
	split := strings.Split(s, "")
	var sum int = 0
	rule := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000}
	sl := len(split)
	for i, _ := range split {
		if i+1 < sl && rule[split[i]] < rule[split[i+1]] {
			sum -= rule[split[i]]
		} else {
			sum += rule[split[i]]
		}
	}
	return sum
}
