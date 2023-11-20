package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Background()和TODO()

//WithCancel

func main1() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() //取完数后调用
	for n := range f1(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}

}

func f1(ctx context.Context) <-chan int {
	dst := make(chan int)
	n := 1
	fmt.Println("=================")
	go func() {
		for {
			select {
			case <-ctx.Done():
				return //结束子协程
			case dst <- n:
				fmt.Println("===", n)
				n++
			}
		}

	}()
	fmt.Println("-----------------")
	return dst
}

// WithDeadline
func main2() {
	d := time.Now().Add(50 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), d)

	// 尽管ctx会过期，但在任何情况下调用它的cancel函数都是很好的实践。
	// 如果不这样做，可能会使上下文及其父类存活的时间超过必要的时间。
	defer cancel()
	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done(): //Done方法需要返回一个Channel，这个Channel会在当前工作完成或者上下文被取消之后关闭，多次调用Done方法会返回同一个Channel；
		fmt.Println(ctx.Err(), 11)
	}
}

// WithTimeOut
var wg1 sync.WaitGroup

func worker1(ctx context.Context) {
LOOP:
	for {
		fmt.Println("db connecting ...")
		time.Sleep(time.Millisecond * 10) // 假设正常连接数据库耗时10毫秒
		select {
		case <-ctx.Done(): // 50毫秒后自动调用
			break LOOP //跳出for
		default:
		}
	}
	fmt.Println("worker done!")
	wg1.Done()
}

func main3() {
	// 设置一个50毫秒的超时
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*50)
	wg.Add(1)
	go worker1(ctx)
	time.Sleep(time.Second * 5)
	cancel() // 通知子goroutine结束
	wg1.Wait()
	fmt.Println("over")
}

// WithValue
// 键必须是可比较的，并且不应该是string类型或任何其他内置类型，以避免使用上下文在包之间发生冲突
type TraceCode string

var wg sync.WaitGroup

func worker(ctx context.Context) {
	key := TraceCode("TRACE_CODE")
	traceCode, ok := ctx.Value(key).(string) // 在子goroutine中获取trace code,断言
	if !ok {
		fmt.Println("invalid trace code")
	}
LOOP:
	for {
		fmt.Printf("worker, trace code:%s\n", traceCode)
		time.Sleep(time.Millisecond * 10) // 假设正常连接数据库耗时10毫秒
		select {
		case <-ctx.Done(): // 50毫秒后自动调用
			break LOOP
		default:
		}
	}
	fmt.Println("worker done!")
	wg.Done()
}

func main() {
	// 设置一个50毫秒的超时
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*50)
	// 在系统的入口中设置trace code传递给后续启动的goroutine实现日志数据聚合
	ctx = context.WithValue(ctx, TraceCode("TRACE_CODE"), "12512312234")
	wg.Add(1)
	go worker(ctx)
	time.Sleep(time.Second * 5)
	cancel() // 通知子goroutine结束
	wg.Wait()
	fmt.Println("over")
}
