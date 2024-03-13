package roman_to_integer

import "testing"

/*
示例 1:

输入: s = "III"
输出: 3
示例 2:

输入: s = "IV"
输出: 4
示例 3:

输入: s = "IX"
输出: 9
示例 4:

输入: s = "LVIII"
输出: 58
解释: L = 50, V= 5, III = 3.
示例 5:

输入: s = "MCMXCIV"
输出: 1994
解释: M = 1000, CM = 900, XC = 90, IV = 4.
*/
func TestRomanToInt(t *testing.T) {
	s := "III"
	toInt := RomanToInt(s)
	if toInt != 3 {
		t.Fatal("error, except:3, result:", toInt)
	} else {
		t.Log("success")
	}

	s = "IV"
	toInt = RomanToInt(s)
	if toInt != 4 {
		t.Fatal("error, except:4, result:", toInt)
	} else {
		t.Log("success")
	}
	s = "IX"
	toInt = RomanToInt(s)
	if toInt != 9 {
		t.Fatal("error, except:9, result:", toInt)
	} else {
		t.Log("success")
	}
	s = "MCMXCIV"
	toInt = RomanToInt(s)
	if toInt != 1994 {
		t.Fatal("error, except:1994, result:", toInt)
	} else {
		t.Log("success")
	}

}
