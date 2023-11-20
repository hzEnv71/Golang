package _test

import (
	"bytes"
	"testing"
)

func Benchmark_Add(b *testing.B) {
	var n int
	//b.N 是基准测试框架提供的，表示循环的次数，没有具体的限制
	for i := 0; i < b.N; i++ {
		n++
	}
}

func BenchmarkConcatStringByAdd(b *testing.B) {
	//有些测试需要一定的启动和初始化时间，如果从 Benchmark() 函数开始计时会很大程度上影响测试结果的精准性。
	//testing.B 提供了一系列的方法可以方便地控制计时器，从而让计时器只在需要的区间进行测试
	elem := []string{"1", "2", "3", "4", "5"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ret := ""
		for _, v := range elem {
			ret += v
		}
	}
	b.StopTimer()
}

func BenchmarkConcatStringByByteBuffer(b *testing.B) {
	elems := []string{"1", "2", "3", "4", "5"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		for _, elem := range elems {
			buf.WriteString(elem)
		}
	}
	b.StopTimer()
}
