package plus_one

/*
66. 加一

给定一个由 整数 组成的 非空 数组所表示的非负整数，在该数的基础上加一。

最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。

你可以假设除了整数 0 之外，这个整数不会以零开头。
*/
func PlusOne1(digits []int) []int {
	digitsLen := len(digits)
	addnum := 1
	for i := digitsLen - 1; i >= 0; i-- {
		addnum = digits[i] + addnum
		if addnum <= 9 {
			digits[i] = addnum
			return digits
		} else {
			if addnum == 10 {
				digits[i] = 0
				addnum = 1
			} else if addnum > 10 {
				digits[i] = addnum - 10
				addnum = 1
			}
			if i == 0 {
				digits = append([]int{addnum}, digits...)
			}
		}
	}
	return digits
}

//更简洁的方法,但是使用内存更多
func PlusOne(digits []int) []int {
	digitsLen := len(digits)
	addendum := 1
	for i := digitsLen - 1; i >= 0; i-- {
		digits[i] += addendum
		digits[i] %= 10
		if digits[i] != 0 {
			return digits
		}
	}
	return append([]int{1}, digits...)
}
