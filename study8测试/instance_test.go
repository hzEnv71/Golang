package _test

import (
	"fmt"
	"os"
	"testing"
)

func Hello() string {
	return "Hello World"
}

// Output 需要和上面打印的内容一致，否则测试失败
func ExampleHello() {
	fmt.Println(Hello())

	// Output: Hello World

}

func TestMain(m *testing.M) {
	fmt.Println("Before...")
	code := m.Run()
	fmt.Println("End...")
	os.Exit(code)
}
