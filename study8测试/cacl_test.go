package cacl

import (
	"fmt"
	"os"
	"testing"
)

// 单元测试钩子
func before() {
	fmt.Println("Before all tests...")
}

func after() {
	fmt.Println("After all tests...")
}

func TestMain(m *testing.M) {
	before()
	code := m.Run()
	after()
	os.Exit(code)
}

// 单元测试
func TestAdd(t *testing.T) {
	if value := Add(1, 2); value != 3 {
		t.Errorf("1+2 expected be 3, but %d got", value)
	}

	if value := Mul(1, 2); value != 2 {
		t.Errorf("1*2 expected be 2, but %d got", value)
	}

}

// 嵌套测试
func TestMul(t *testing.T) {
	t.Run("pos", func(t *testing.T) {
		if Mul(2, 3) != 6 {
			t.Fatal("expected to get 6,but fail...")
		}
	})

	t.Run("neg", func(t *testing.T) {
		if Mul(2, -3) != -6 {
			t.Fatal("expected to get -6,but fail...")
		}
	})
}

// 测试组
func TestMul2(t *testing.T) {
	cases := []struct {
		Name           string
		A, B, Expected int
	}{
		{"pos", 2, 3, 6},
		{"neg", 2, -3, -6},
		{"zero", 2, 0, 0},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			if value := Mul(c.A, c.B); value != c.Expected {
				t.Fatalf("%d * %d expected %d,but %d got", c.A, c.B, c.Expected, value)
			}
		})
	}
}
