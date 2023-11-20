package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"reflect"
	"runtime"
)

// 结构体
// 定义结构体类型
type Student struct {
	id   int
	name string
	age  int

	Id   int
	Name string
	Age  int
}

// import . "fmt"//调用函数 ，无需通过包名
// import i "fmt"//给包名起别名
// import _ "fmt"//忽略此包

// 多个形参,要放在最后一个
func test03(i ...int) {
	fmt.Println(i)
}

// 闭包只要使用变量，就一直存在
// 返回一个函数类型
func test02(i int) func() int {
	i = 0
	return func() int {
		i++
		return i * i
	}
}

// 普通函数
func test01() int {
	i := 0
	i++
	return i * i //函数调用时分配空间，完毕自动释放
}

// 函数类型
type FuncType func(int, int) int

// 实现多态
func Add(a, b int) int {
	return a + b
}
func Minus(a, b int) (result int) {
	result = a - b
	return
}

// 回调函数调用
func Cal(a, b int, fun FuncType) (result int) {
	result = fun(a, b) //可扩展性
	return
}

func func01(i int, j float64, k string) (int, float64, string) {
	return i, j, k
}

// Go官方推荐
func func02(i, j, k int) (i1, j1, k1 int) {
	i1 = i
	j1 = j
	k1 = k
	return
}

// 匿名变量
func test() (i, j, k int) {
	return 1, 2, 3
}

var a1 = "skljdfgh"

func main() {

	/*fmt.Println("Hello, World!")
	a2 := 87 //自动推导类型,只能在函数里面使用
	// fmt.Println(a1)
	fmt.Println(a2)*/

	/*var temp int
	a,b:=10,20
	temp=a
	a=b
	b=temp
	fmt.Println(a,b)*/

	/*//多重赋值
	var a,b=10,20;
	//交换
	b,a=a,b;
	fmt.Println(a,b)*/

	/*// 匿名变量
	var i,_,_=study2_test()
	_ ,j,k:=study2_test()
	fmt.Println(i,j,k);*/

	/*//常量
	const a int =666
	// a=777
	fmt.Println(a)
	const i,j,k=666,12.24484953732,"sdjhfgb"
	fmt.Println(i,j,k)
	fmt.Printf("i的类型是%T，j的类型是%T，k的类型是%T",i,j,k)*/

	// //多个  不同类型  变量的声明
	// var q1,q2,q3=23,2.345,"sdffg"
	// var(
	//    a int
	//    b float64
	//    c string
	// )
	// //多个  不同类型  变量声明并初始化
	// var (
	//    i=34
	//    j="sjkdfg"
	//    k=2.34256
	// )
	// a,b,c=12,3.55,"sdfjkhg"
	// fmt.Println(q1,q2,q3)
	// fmt.Println(a,b,c)
	// fmt.Println(i,j,k)
	// // const
	// const l,m,n=23,3.444,"sjhmdnfgb"
	// fmt.Println(l,m,n)

	/*//自定义形参个数
	test03(1, 2, 3)
	test03(1, 2)*/

	// //iota枚举
	// const a,b,c =iota,iota,iota//一行的数值一样
	// fmt.Println(a,b,c)
	// const(
	//    i=iota
	//    j=iota
	//    k=iota
	// )//逐个+1
	// const l=iota//iota遇到const时重置为0
	// fmt.Println(i,j,k,l)
	// const(
	//    i1 =iota
	//    j1
	//    k1
	// )//可以省略
	// fmt.Println(i1,j1,k1)

	// //布尔
	// var flag bool=true;
	// fmt.Println(flag)

	// //字符型
	//byte int32 ascll
	// var c byte='a'
	// fmt.Println(c,c-32)
	// fmt.Printf("小写:%c,大写%c\n",c,c-32)
	// fmt.Printf("hello %c",'\n')
	// // fmt.Print("\n")相当于fmt.Println()
	// fmt.Println("world\n")
	// fmt.Print("go")
	//rune int32 utf-8
	//r := "面积是"
	//run := []rune(r)
	//run[0] = '是'
	//fmt.Printf("%T", run)
	//fmt.Println(string(run))

	// //字符串 后面隐藏了'\0'结束符
	// var str string="smhjdfg"
	// fmt.Println(len(str))//内建函数，测字符串长度
	// fmt.Println(str[0],str[1])

	// //复数
	//  var t complex128
	//  t=2+4i
	//  fmt.Println(t)
	//  tt:=2+4i
	//  fmt.Printf("%T",tt)
	//  fmt.Println(real(tt),imag(tt))//内建函数，实部，虚部

	// //键入
	// var a int
	// fmt.Scanf("%d",&a)
	// fmt.Scan(&a)
	// fmt.Println(a)

	//for          range
	/*i:=0
	for{//无限循环
	   i++
	   fmt.Println(i)
	}
	for i:=0;i<10;i++{
	   fmt.Println(i)
	}
	var s="dfjgh"
	for i,data:=range s{
	   fmt.Println(i,data)
	}*/

	// //调用
	// var i,j,k=func01(1,2.2234,"xsjfdhg")
	// fmt.Println(i,j,k)
	// fmt.Println(func02(2,6,8))

	// 回调函数调用
	// fmt.Println(Cal(1,1,Add))
	// fmt.Println(Cal(1,1,Minus))

	//匿名函数
	// f:=func(x,y int)(result int){
	//    result=x+y
	//    return
	// }
	// fmt.Println(f(10,200))
	// x,y:=func(x,y int)(max,min int){
	//    if x>y {
	//       max=x
	//       min=y
	//    }else{
	//       max=y
	//       min=x
	//    }
	//    return
	// }(10,20)//()后面括号代表调用此函数
	// fmt.Println(x,y)

	// //闭包以引用方式捕获外部变量
	// a:=10
	// b:="shdmfg"
	// func(){
	//    a=777
	//    b="jsdhf"
	// }()
	// fmt.Println(a,b)//改变了原值

	// //普通函数和返回函数类型函数对比
	// fmt.Println(test01())
	// fmt.Println(test01())
	// fmt.Println(test01())
	// f:=test02(0)//返回函数类型，通过f调用返回的匿名函数
	// fmt.Println(f())
	// fmt.Println(f())
	// fmt.Println(f())

	// //defer 遇到错误也会执行其他
	// a:=100
	// b:=200
	// //参数已经传了，只是main结束前才调用
	// defer func(a,b int){
	//    fmt.Println(a,b)
	// }(a,b)
	// a=1000
	// b=2000
	// fmt.Println(a,b)

	// //获取命令行参数
	// a:=os.Args
	// fmt.Printf("%T\n",a)
	// fmt.Println(a[0])

	/*//指针,默认为nil
	a := 100
	var p *int = &a
	*p = 1000
	fmt.Println(a, *p)
	//申请空间
	p1 := new(int)
	*p1 = 666
	fmt.Println(*p1)*/

	/*//数组,切片[]里面为空，或者...
	var a1 []int//切片

	var a2 = []int{1, 2, 3}//切片

	arr := []int{1, 2, 3, 4, 5}//切片
	fmt.Println(a1, a2[2:3], arr[:3])
	arr1 := [][]int{}//切片
	arr2 := [][]int{{4, 5, 6}, {1, 2, 3}}//切片
	arr3 := [3][4]int{{3, 45, 67}, {21, 45, 56}, {1, 54, 54}}
	fmt.Println(arr1, arr2[1:], arr3[1:3][1:2])*/

	/*//随机数必须设置种子
	rand.Seed(time.Now().UnixNano())
	var a [10]int
	n := len(a)
	for i := 0; i < n; i++ {
		a[i] = rand.Intn(100)
		fmt.Print(a[i], ",")
	}
	fmt.Println()
	//冒泡排序
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
	for i, data := range a {
		fmt.Println(i, data)
	}*/

	/*//go语言中数组是值传递
	//数组指针
	a := [5]int{1, 2, 3, 4, 5}
	testp(a) //值传递
	//testq(&a) //地址传递
	fmt.Println(a)
	//切片：地址传递
	a1 := []int{1, 2, 3, 4, 5}
	testp1(a1) //地址传递
	fmt.Println(a1)*/

	/*//切片
	a := []int{1, 2, 3, 4, 5}
	slice := a[0:3:5] //起始索引，长度，容量
	fmt.Println(slice, len(slice), cap(slice))
	//自动推导类型
	s1 := []int{1, 2, 3, 4}
	fmt.Println(s1, len(s1), cap(s1))
	//借助make函数 格式make(切片类型，长度，容量)
	s2 := make([]int, 5, 10)
	fmt.Println(len(s2), cap(s2))
	//不指定容量，容量和长度一样
	s3 := make([]int, 0)
	fmt.Println(len(s3), cap(s3))
	//append：如果超过容量，扩容一般2倍
	s3 = append(s3, 1, 2, 3)
	fmt.Println(len(s3), cap(s3))
	//copy(数组，被拷贝的数组)
	copy(s3, s2)
	fmt.Println(s3)*/

	/*//map  没有cap（）键不能是切片，函数等带有引用的
	var m1 map[int]string
	var m2 = make(map[int]string)
	var m3 = make(map[int]string, 5)
	m3[0] = "golang"
	m3[1] = "gin"
	fmt.Println(len(m1), len(m2), len(m3))
	//初始化,如果已存在键，则覆盖
	m4 := map[int]string{
		1: "q",
		2: "w",
	}
	m4[2] = "t"
	fmt.Println(m4)
	//遍历
	for key, value := range m4 {
		fmt.Println(key, value)
	}
	//判断是否存在key值 值，是否存在
	value, status := m4[1]
	fmt.Println(value, status)
	//删除
	delete(m4, 1)
	value1, status1 := m4[1]
	fmt.Println(value1, status1)
	testm(m4) //引用传递
	fmt.Println(m4)*/

	/*//结构体普通变量
	var s1 Student = Student{1, "hz", 18}
	//指定成员初始化，其他默认为0
	s3 := Student{id: 3, name: "hhz"}
	fmt.Println(s1, s3)
	//结构体指针变量
	var s2 *Student = &Student{2, "hh", 19}
	s4 := &Student{id: 4, name: "hzz"}
	fmt.Println(*s2, *s4)
	//
	var s Student
	var p *Student
	p = &s
	p.age = 19
	(*p).name = "dsjf"
	s.id = 8
	fmt.Println(s, *p)
	pp := new(Student)
	pp.name = "na"
	fmt.Println(*pp)
	//同类型结构体直接赋值
	p2 := new(Student)
	p2 = p
	ss := s
	fmt.Println(ss, *p2)
	var s6 Student
	tests1(s6)  //值传递
	testp3(&s6) //引用传递
	fmt.Println(s6)*/

	/*//time类
	fmt.Println(time.Now().UTC())
	fmt.Println(time.Now())
	now := time.Now()
	fmt.Println(now.Unix(), now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	fmt.Println(time.Now().Format("2006/01/02 15:04:05PM"))
	sub, err := time.Parse("2006-01-02", "2020-01-01")
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	subloc, err1 := time.ParseInLocation("2006-01-02", "2020-01-01", loc)
	if err1 != nil {
		fmt.Println("err=", err)
		return
	}
	fmt.Println(time.Now().Sub(sub)) //减
	fmt.Println(time.Now().Sub(subloc))
	fmt.Println(sub.UTC().Sub(time.Now()))                      //本地时区
	fmt.Println(time.Now().Equal(time.Now()))                   //比较
	fmt.Println(time.Now().Before(sub), sub.Before(time.Now())) //判断在  之前
	fmt.Println(time.Now().After(sub), sub.After(time.Now()))   //判断在  之后
	fmt.Println("==========================")
	//var d time.Duration
	fmt.Println(time.December, time.Second, time.Microsecond, time.Nanosecond)
	fmt.Println(time.Now().Format(time.Layout), time.Now().Add(24*time.Hour))
	fmt.Println(time.Now().Format(time.ANSIC))
	unix := now.Unix()
	t := time.Unix(unix, 0)
	fmt.Println(t.Year(), t.Second(), time.Unix(now.Unix(), 0).Day())
	*/
	//日志库
	log.Println("日志库==")
	file1, err := os.OpenFile("./deemo.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644) //追加｜创建｜写
	if err != nil {
		fmt.Println("file err=", err)
		return
	}
	log.SetOutput(file1) //设置输出路径
	log.Println("日志输出")
	//runtime.Caller()
	//pc, file, line, ok := runtime.Caller(0) //调用者所在的行号，隔层调用
	//if !ok {
	//	fmt.Println("runtime.Caller() failer\n")
	//	return
	//}
	//fmt.Println(runtime.FuncForPC(pc).Name(), file, line, ok) //函数名，路径，行号，状态
	//f()

	//reflect
	var a float64 = 3.14

	var b int64 = 100
	var s Student = Student{19, "hz", 1, 19, "hz", 11}
	slic := make([]byte, 10)
	m := map[int]string{1: "666"}
	//类型
	fmt.Println(reflect.TypeOf(a), reflect.TypeOf(b), reflect.TypeOf(s), reflect.TypeOf(s.name), reflect.TypeOf(s).Kind())
	//值
	fmt.Println(reflect.ValueOf(a), reflect.ValueOf(b), reflect.ValueOf(s), reflect.ValueOf(s.age), reflect.ValueOf(s).Kind())
	//设置值 Elem()获取指针对应的值
	reflect.ValueOf(&a).Elem().SetFloat(6.7777)
	reflect.ValueOf(&b).Elem().SetInt(666)

	fmt.Println(a, b, s)
	//地址传：isNil()  值传：isValid()
	fmt.Println(reflect.ValueOf(a).IsValid(), reflect.ValueOf(s).FieldByName("name").IsValid())
	fmt.Println(reflect.ValueOf(slic).IsNil(), reflect.ValueOf(m).MapIndex(reflect.ValueOf(1)).IsValid())
	//结构体反射
	fmt.Println(reflect.TypeOf(s).Field(0).Name, reflect.TypeOf(s).Field(1).Name, reflect.TypeOf(s).Field(1).Index)
	fmt.Println(reflect.TypeOf(s).FieldByName("id"))
	fmt.Println(reflect.TypeOf(s).NumField())
	//reflect.ValueOf(&s).Elem().FieldByName("id").Set(reflect.ValueOf(21))
	reflect.ValueOf(&s).Elem().Field(4).SetString("hzhzhz")                //只能获得开头大写字段
	reflect.ValueOf(&s).Elem().FieldByName("Id").Set(reflect.ValueOf(666)) //
	fmt.Println(s.Id, s.Name)
}
func f() {
	//pc, file, line, ok := runtime.Caller(1) //调用者所在的行号，隔层调用
	pc, file, line, ok := runtime.Caller(0) //自己所在的行号，隔层调用
	if !ok {
		fmt.Println("runtime.Caller() failer\n")
		return
	}
	fmt.Println(runtime.FuncForPC(pc).Name(), file, line, ok) //函数名，路径，行号，状态
	fmt.Println(runtime.FuncForPC(pc).Name(), path.Base(file), line, ok)
}

func testp3(p3 *Student) {
	p3.age = 99
}

func tests1(s6 Student) {
	s6.age = 77
}

// map
func testm(m4 map[int]string) {
	delete(m4, 2)
}

// 数组
func testp(a [5]int) {
	a[0], a[1] = a[1], a[0]
	fmt.Println(a[0:])
}
func testq(p *[5]int) {
	(*p)[0], (*p)[1] = (*p)[1], (*p)[0]
	fmt.Println(*p)
}

// 切片
func testp1(a []int) {
	a[0], a[1] = a[1], a[0]
	fmt.Println(a[:])
}
