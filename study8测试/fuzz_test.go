package _test

import (
	"fmt"
	"strconv"
	"testing"
)

func FuzzStrToNum(f *testing.F) {
	f.Fuzz(func(t *testing.T, a string) {
		b, _ := strconv.ParseInt(a, 10, 64)
		fmt.Printf("%d\n", b)
	})
}

func FuzzDivision(f *testing.F) {
	f.Fuzz(func(t *testing.T, a, b int) {
		fmt.Println(a / b)
	})
}

// 我们可以在模糊测试中，自己定义参数进去
func FuzzHello(f *testing.F) {
	f.Add(6)
	f.Fuzz(func(t *testing.T, num int) {
		if num != 6 {
			t.Errorf("expected 6,but got %d", num)
		}
	})
}
