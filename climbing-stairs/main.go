package climbing_stairs

/*
70. 爬楼梯
假设你正在爬楼梯。需要 n 阶你才能到达楼顶。

每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
1 <= n <= 45

*/
func ClimbStairs(n int) int {
	s := [45]int{1, 2}
	for i := 2; i < n; i++ {
		s[i] = s[i-1] + s[i-2]
	}
	return s[n-1]
}

/*

 */
