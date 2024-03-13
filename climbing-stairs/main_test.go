package climbing_stairs

import "testing"

/*
示例 1：

输入：n = 2
输出：2
解释：有两种方法可以爬到楼顶。
1. 1 阶 + 1 阶
2. 2 阶
示例 2：

输入：n = 3
输出：3
解释：有三种方法可以爬到楼顶。
1. 1 阶 + 1 阶 + 1 阶
2. 1 阶 + 2 阶
3. 2 阶 + 1 阶
*/
func TestClimbStairs(t *testing.T) {
	n := 3
	except := 3
	stairs := ClimbStairs(n)
	if stairs == except {
		t.Log("success")
	} else {
		t.Errorf("err except:%v,result:%v\n", except, stairs)
	}

	n = 3
	except = 3
	stairs = ClimbStairs(n)
	if stairs == except {
		t.Log("success")
	} else {
		t.Errorf("err except:%v,result:%v\n", except, stairs)
	}
	n = 5
	except = 8
	stairs = ClimbStairs(n)
	if stairs == except {
		t.Log("success")
	} else {
		t.Errorf("err except:%v,result:%v\n", except, stairs)
	}
	n = 6
	except = 13
	stairs = ClimbStairs(n)
	if stairs == except {
		t.Log("success")
	} else {
		t.Errorf("err except:%v,result:%v\n", except, stairs)
	}
	n = 8
	except = 34
	stairs = ClimbStairs(n)
	if stairs == except {
		t.Log("success")
	} else {
		t.Errorf("err except:%v,result:%v\n", except, stairs)
	}
}

/*
4
1 1 1 1
2 1 1
1 2 1
1 1 2
2 2

1  1
2  2
3  3
4  5
5  8
6  13
7  21

*/
