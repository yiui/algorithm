package recursion

import "container/list"

// Sum 普通递归求和
func Sum(n int) int {
	if n == 1 {
		return 1
	}
	return n + Sum(n-1)
}

// Sum2 尾递归
func Sum2(n, res int) int {
	if n == 0 {
		return res
	}
	return Sum2(n-1, res+n)
}

/*
设斐波那契数列
f(1)=0,f(2)=1
f(n)=f(fn-1)+f(n-2)
*/
//递归树
func Fib(n int) int {
	if n == 1 || n == 2 {
		return n - 1
	}
	return Fib(n-1) + Fib(n-2)
}

//通过迭代模拟递归

func ForLoopRecur(n int) int {
	//使用一个显式的栈来模拟系统调用栈
	stack := list.New()
	res := 0
	//递进：递归调用
	for i := n; i > 0; i-- {
		//通过入栈操作模拟递
		stack.PushBack(i)
	}

	//回归：返回结果
	for stack.Len() != 0 {
		//通过出栈操作模拟 归
		res += stack.Back().Value.(int)
		stack.Remove(stack.Back())
	}
	return res
}
