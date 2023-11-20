package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

/*
	func main() {
		//主协程退出，子协程也会退出
		go func() {
			i := 0
			for {
				i++
				fmt.Println("子协程：", i)
				time.Sleep(time.Second)
			}
		}() //函数调用
		i := 0
		for {
			i++
			fmt.Println("主协程：", i)
			time.Sleep(time.Second) //1000ms
			if i == 2 {
				break
			}
		}
	}
*/
/*func main() {
	go func() {
		study2_test()
		for i := 0; i < 5; i++ {
			fmt.Println("子协程——————》", i)
		}

	}()

	for i := 0; i < 3; i++ {
		//子协程来不及执行，主协程让出事件片，先让别的子协程执行
		//runtime
		runtime.Gosched()
		fmt.Println("主协程——————》", i)

	}
}*/

func test() {
	defer fmt.Println("aaaaaaaaa")
	//return
	runtime.Goexit() //终止此协程
	fmt.Println("cccccccc")
}

/*func main() {
	//n := runtime.GOMAXPROCS(1) //单核
	n := runtime.GOMAXPROCS(6) //6核
	fmt.Print("n=", n)

	for {
		go fmt.Print(0)
		fmt.Print(1)

	}
}
*/
//channel
var ch = make(chan int)

//通道没数据就会阻塞

func Printer(str string) {
	for _, data := range str {
		fmt.Printf("%c ", data)
		time.Sleep(time.Second)
	}
	fmt.Println()
}
func Print1() {
	Printer("hello")
	ch <- 666 //写数据，发送
}
func Print2() {
	<-ch //取数据，通道没数据就会阻塞
	Printer("world")

}

/*func main() {
	//多任务资源竞争问题
	go Print1()
	go Print2()

	for {
	} //不让主协程结束

}*/

/*
	func main() {
		//channel
		//创建有缓冲的，容量为10
		//ch1:=make(chan int,10)
		//ch:=make(chan int)
		ch := make(chan string)
		fmt.Println(len(ch), cap(ch))
		defer fmt.Println("主协程结束")
		go func() {
			defer fmt.Println("子协程结束")
			for i := 0; i < 2; i++ {
				fmt.Println("子协程：", i)
				time.Sleep(time.Second)
			}
			ch <- "子协程任务完成"
		}()
		str := <-ch
		fmt.Println(str)
	}
*/
func main1() {
	/*ch := make(chan int, 2)
	//新建子协程
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		//关闭通道，不能写数据，但是可以读数据
		close(ch)
	}()
	for {
		//值，状态
		if data, status := <-ch; status == true {
			fmt.Println(data)
		} else { //通道关闭
			break
		}
	}
	//自动退出
	for data := range ch {
		fmt.Println(data)
	}*/

	/*//单向通道
	c := make(chan int, 1) //默认双向
	//双向转单向
	var writeC chan<- int = c
	var readC <-chan int = c
	writeC <- 666
	<-readC

	//引用传递
	channel := make(chan int)
	//写入
	go Write(channel)
	//读出
	Read(channel)*/
}

func Read(in <-chan int) {
	for data := range in {
		fmt.Println(data)
	}
}

func Write(out chan<- int) {
	for i := 0; i < 5; i++ {
		out <- i * i
	}
	//关闭通道，不能写数据，但是可以读数据
	close(out)
}
func main2() {
	//创建定时器，设置2s后，往time通道写内容（当前时间）
	timer := time.NewTimer(2 * time.Second)
	fmt.Println("当前时间：", time.Now())

	//2s后，往timer.C写数据，有数据后，就可以读取
	t := <-timer.C //channel没有数据前阻塞
	fmt.Println(t)

	//实现延迟功能
	//1
	<-time.After(2 * time.Second) //定时2s，阻塞2s,2s后产生一个事件，往channel写内容
	fmt.Println("时间到")
	//2
	time.Sleep(2 * time.Second)
	fmt.Println("时间到")
	//3
	timer1 := time.NewTimer(2 * time.Second)
	timer1.Reset(5 * time.Second)
	//timer1.Stop()//停止
	/*go func() {
		<-timer1.C
	}()*/
	<-timer1.C
	fmt.Println("时间到")

}

/*
	func main() {
		//间隔触发的定时器
		ticker := time.NewTicker(time.Second)
		ticker.Reset(2*time.Second)
		i := 0
		for {
			<-ticker.C
			i++
			fmt.Println(i)
			if i == 3 {
				ticker.Stop()
				break
			}
		}
	}
*/
func main3() {
	//ch := make(chan int)    //数字通信
	var ch chan int
	ch = make(chan int)
	flag := make(chan bool) //判断程序是否结束
	//消费者，读取内容
	go func() {
		for i := 0; i < 7; i++ {
			fmt.Println(<-ch)
		}
		flag <- true
	}()
	//生产者，产生数字，写入
	fibonacci(ch, flag)
}

// ch只写，flag只读
func fibonacci(ch chan<- int, flag <-chan bool) {
	x, y := 1, 1
	for {
		select {
		case ch <- x:
			x, y = y, x+y
		case flag := <-flag:
			fmt.Println(flag)
			return
		}
	}
}

func main4() {
	ch := make(chan int)    //数字通信
	flag := make(chan bool) //判断程序是否结束

	go func() {
		for {
			select {
			case data := <-ch:
				fmt.Println(data)
			case <-time.After(3 * time.Second):
				fmt.Println("你已经3秒未操作，超时断开链接")
				flag <- true
			}
		}
	}()
	for i := 0; i < 5; i++ {
		ch <- i
		time.Sleep(time.Second)
	}
	<-flag

}

var wg sync.WaitGroup

func main5() {

	for i := 0; i < 10; i++ {
		wg.Add(1) //增加计数
		go f1(i)
	}
	wg.Wait() //等待计数归0
}

func f1(i int) {
	defer wg.Done() //任务执行完，计数-1
	time.Sleep(time.Duration(rand.Intn(3)))
	fmt.Println(i, time.Duration(rand.Intn(3)))
}

// 并发安全
func main6() {
	//GMP M:N goroutine初始化的栈大小为2k
	//M:子协程数
	//N:操作系统数
	runtime.GOMAXPROCS(2) //最大线程数
	wg.Add(2)
	go a()
	go b()
	wg.Wait()
}

func b() {
	for i := 0; i < 10; i++ {
		fmt.Println("B:", i)
	}
	wg.Done()
}

func a() {
	for i := 0; i < 10; i++ {
		fmt.Println("A:", i)
	}
	defer wg.Done()
}

// 互斥锁
var lock sync.Mutex
var x int //x线程不安全

func main7() {
	wg.Add(2)
	go fu()
	go fu()
	wg.Wait()
	//fmt.Println(x)//多次操作，结果可能不一样
	//sync.Mutex
	fmt.Println(x)
}

func fu() {
	for i := 0; i < 1000; i++ {
		lock.Lock()
		x += i
		lock.Unlock()
	}
	defer wg.Done()
}

// 读写互斥锁 适用用读大于写的操作
var (
	rwLock sync.RWMutex
)

func main8() {
	startTime := time.Now()

	for i := 0; i < 10; i++ {
		go write()
		wg.Add(1)
	}
	for i := 0; i < 1000; i++ {
		go read()
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println(time.Now().Sub(startTime))

}

var once sync.Once

func read() {
	defer wg.Done()
	rwLock.RLock()
	//fmt.Println(x)
	once.Do(do) //只能执行一次
	rwLock.RUnlock()
}
func do() {
	fmt.Println(x)

}
func write() {
	defer wg.Done()
	rwLock.Lock()
	x += 1
	rwLock.Unlock()
}

// sync.map
var m = sync.Map{} //开箱即用
//Store Load LoadOrStore Delete Range

func main9() {
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			m.Store(key, n)         //必须用store存储
			value, _ := m.Load(key) //必须用load取值
			fmt.Printf("key:%v,value:%v\n", key, value)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

// 原子性操作
var n int64

func main() {

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go fun()
	}
	wg.Wait()
	fmt.Println(n)
}

func fun() {
	a := atomic.AddInt64(&n, 2) //修改
	fmt.Println(a)
	atomic.StoreInt64(&n, 3000) //写入
	//fmt.Println(atomic.LoadInt64(&n))//读取
	//atomic.SwapInt64(&n, 5000)
	atomic.SwapInt64(&n, 1000)
	atomic.CompareAndSwapInt64(&n, 1000, 2000) //将n与1000对比，如果true则修改为2000
	defer wg.Done()
}
