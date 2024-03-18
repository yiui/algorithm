package recursion

import (
	"testing"
)

func TestSum(t *testing.T) {
	sum := Sum(10)
	t.Log("sum=", sum)

	var re int
	r := Sum2(10, re)
	t.Log("sum=", r)

}

func TestFib(t *testing.T) {
	fib := Fib(5)
	except := 3
	if fib == except {
		t.Log("success")
	} else {
		t.Errorf("err except:%v,result:%v\n", except, fib)
	}
}

func TestForLoopRecur(t *testing.T) {
	//
	recur := ForLoopRecur(10)
	t.Log(recur)
}
