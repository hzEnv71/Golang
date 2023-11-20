package main

import "fmt"

type Person struct {
	name string
	age  int
}

/*
	type Student struct {
		//person Person
		Person        //匿名字段，继承Person的成员
		name   string //同名字段，就近原则
		id     int
		int    //基础类型匿名字段
	}
*/
/*type Student struct {
	*Person        //指针类型
	name   string //同名字段，就近原则
	id     int
}*/

/*
	func main() {
		//顺序初始化
		var s1 Student = Student{
			Person: Person{"hz", 19},
			id:     8,
		}
		//自动推导类型
		s2 := Student{
			Person: Person{"hh", 10},
			id:     9,
		}
		s3 := Student{Person{"sdd", 22}, 1}
		fmt.Printf("s1=%+v\ns2=%+v\ns3=%+v\n", s1, s2, s3)
	}
*/
/*func main() {
	s1 := Student{}
	s1.id = 9
	s1.name = "sj"
	s1.Person.age = 19 //显式
	fmt.Printf("%+v\n%+v\n", s1.Person.name, s1.name)
	s1.int = 666
	fmt.Printf("%+v\n", s1.int)
}
*/
/*func main() {
	s1 := Student{&Person{"hz", 19}, "hh", 19}
	fmt.Println(*s1.Person, s1.name, s1.id)
	s2 := Student{}
	s2.Person = new(Person)
	s2.age = 19
	s2.name = "sd"//同名字段，就近原则
	s2.id = 11
	fmt.Println(*s2.Person, s2.name, s2.age)
	s2.Person.name = "aa"
	fmt.Println(*s2.Person, s2.name, s2.age)
}
*/

// 自定义类型 添加方法
// 自定义类型名称不同，即使方法相同
type MyInt int
type myInt int
type mInt int

func (a myInt) test(b myInt) myInt {
	return a + b
}
func (a MyInt) test(b MyInt) MyInt {
	return a + b
}

// 省略参数名
func (mInt) test(b mInt) mInt {
	return b + 1
}

/*//基类型T不能是接口或者指针，可以是T 或*T
type T *int
func (T) study2_test(b T) T {
	return b + 1
}*/
/*
//不支持重载

	func (a MyInt) study2_test(b MyInt,c MyInt) MyInt {
		return a + b
	}
*/
/*func main() {
	//方法
	var a MyInt //默认为0
	fmt.Println(a, a.study2_test(10))
	var b myInt //默认为0
	fmt.Println(b, b.study2_test(6))
	var c mInt
	fmt.Println(c, c.study2_test(7))

}
*/

type Student struct {
	Person        //继承成员和方法
	name   string //同名字段，就近原则
	id     int
}

func (p Person) method01() {
	fmt.Println("结构体普通变量")
}
func (p *Person) method02() {
	fmt.Println("结构体指针变量Person")
}
func (s *Student) method02() {
	fmt.Println("结构体指针变量Student")
}

/*func main() {
	//方法集
	p := &Person{"hjh", 10}
	p.method01() //自动转换 p转为(*p)
	p.method02()
	q := Person{"hh", 29}
	q.method01()
	q.method02() //自动转换 q转为&q
	s := Student{}
	s.method01()
	s.method02()        //s转换为(&s)
	s.Person.method02() //就近原则
	//方法值
	pFunc := p.method01 //自动转换 p转为(*p)
	pFunc()             //等价于p.method01(),隐藏了接收者
	//方法表达式
	f := (*Student).method01
	f(&s) //显式把接收者传递

}
*/

type Inter interface { //子集
	print()
}
type FInter interface { //超集
	Inter //接口继承
	sing(str string)
}

type Teacher struct {
	name string
	age  int
}
type Mystring string

func (t *Teacher) print() {
	fmt.Println(*t, "print")
}
func (t *Teacher) sing(str string) {
	fmt.Println(*t, str, "sing")
}
func (str Mystring) print() {
	fmt.Println(str)
}

/*
func main() {

	t := &Teacher{name: "dhsfh"}
	var str Mystring = "dfjhg"
	//创建接口切片
	x := make([]Inter, 2)
	x[0] = t
	x[1] = str //自动转换 str转为&str
	for _, i := range x {
		i.print()
	}
	//继承
	var f FInter = &Teacher{}
	f.print()
	f.sing("sd")
	//超集可以转换成子集，反过来不可以
	var i Inter
	i = f
	i.print()
	//空接口可以保存任意类型的值
	var ii interface{}
	ii = 1
	ii = "an"
	ii = 9.12
	fmt.Println(ii)

}
*/
func main() {
	i := make([]interface{}, 3)
	i[0] = 19
	i[1] = "djsf"
	i[2] = Teacher{"sjdf", 19}
	//类型断言
	for index, data := range i {
		//值:判断结果的真假
		if value, status := data.(int); status == true {
			fmt.Printf("x[%d]类型为int,内容为%d\n", index, value)
		} else if value, status := data.(string); status == true {
			fmt.Printf("x[%d]类型为string,内容为%s\n", index, value)
		} else if value, status := data.(Teacher); status == true {
			fmt.Printf("x[%d]类型为Student,内容为name=%s,age=%d\n", index, value.name, value.age)
		}
	}
	for index, data := range i {
		//switch case
		switch value := data.(type) {
		case int:
			fmt.Printf("x[%d]类型为int,内容为%d\n", index, value)
		case string:
			fmt.Printf("x[%d]类型为string,内容为%s\n", index, value)
		case Teacher:
			fmt.Printf("x[%d]类型为Student,内容为name=%s,age=%d", index, value.name, value.age)
		}
	}
}
