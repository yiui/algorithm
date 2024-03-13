package sqrtx

/*
69. x 的平方根
给你一个非负整数 x ，计算并返回 x 的 算术平方根 。

由于返回类型是整数，结果只保留 整数部分 ，小数部分将被 舍去 。

注意：不允许使用任何内置指数函数和算符，例如 pow(x, 0.5) 或者 x ** 0.5 。
*/

/*
 */
func MySqrt2(x int) int {
	l, r := 1, x
	for r >= l {
		if x < l*l {
			return l - 1
		} else if x == l*l {
			return l
		} else if x < (l+1)*(l+1) {
			return l
		} else if x >= (r-1)*(r-1) {
			return r - 1
		} else if x >= r*r {
			return r
		}
		l += 1
		r -= 1
	}
	return 0
}

/*
二分法
*/
func MySqrt(x int) int {
	l, r := 1, x
	var mid int
	for r >= l {
		mid = (l + r) / 2
		if x < mid*mid {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return r
}
