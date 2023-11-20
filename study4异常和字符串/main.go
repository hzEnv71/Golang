package main

import (
	"fmt"
	"strconv"
)

func testa() {
	fmt.Println("aaaaaaaaaaaaaaaaaaaaaaaaa")
}
func testb() {
	//fmt.Println("bbbbbbbbbbbbbbbbbbbbbbbbb")
	//显式调用panic，中断
	panic("this is panic study2_test")
	fmt.Println("bbbbbbbbbbbbbbbbbbbbbbbbb")
}
func testc(x int) {

	defer func() {
		//recover(),打印panic的错误信息
		if err := recover(); err != nil {
			//fmt.Println(recover())
			fmt.Println(err)
		}
	}()
	var c [10]int
	c[x] = 100

	//fmt.Println("ccccccccccccccccccccccc")
}

func main() {
	/*
		//error 接口
		err1 := fmt.Errorf("%s", "this is normal err")
		err2 := errors.New("this is normal err2")
		fmt.Println(err1, err2)
		//testa()
		//testb()
		testc(11)
	*/

	/*//字符串操作
	//Contains
	fmt.Println(strings.Contains("hellogo", "hello"))
	fmt.Println(strings.Contains("hellogo", "hs"))
	//Join
	s := []string{"abd", "hello", "mike", "go"}
	buf := strings.Join(s, "_")
	fmt.Println(buf)
	//Index
	fmt.Println(strings.Index("asdasfdsf", "da"))
	fmt.Println(strings.Index("asdasfdsf", "dd")) //不包含返回 -1
	//Repeat
	buf = strings.Repeat("go", 4)
	fmt.Println(buf)
	//Spilt
	buf = "hello&sd&ao&kk"
	s2 := strings.Split(buf, "&")
	s3 := strings.SplitN(buf, "&", 2)
	fmt.Println(s2, s3)
	//Trim 去掉两头的字符
	buf = strings.Trim("77777asd7u7dsjkfh7dsf7777", "7")
	fmt.Println(buf)
	//Fields 返回[]string 去掉空字符
	s4 := strings.Fields("   asd  u  dsjkfh dsf   ")
	fmt.Println(s4)*/

	//转换为字符串后追加到字节数组 切片
	slice := make([]byte, 0, 1024)
	slice = strconv.AppendBool(slice, true)
	slice = strconv.AppendInt(slice, 9999, 10)
	slice = strconv.AppendFloat(slice, 3.12, 'f', -1, 64)
	slice = strconv.AppendQuote(slice, "asdasfdsa")
	fmt.Println(string(slice))

	//其它类型转换为字符串
	var str string
	str = strconv.FormatBool(true)
	str = strconv.FormatFloat(3.23, 'f', -1, 64)
	fmt.Println(str)
	//整数转字符串，常用
	str = strconv.Itoa(777)
	fmt.Println(str)
	//字符串转其它类型
	var flag bool
	var err error
	flag, err = strconv.ParseBool("tr7ue")
	if err == nil {
		fmt.Println(flag)
	} else {
		fmt.Println(err)
	}
	//字符串转整型
	a, _ := strconv.Atoi("876")
	fmt.Println(a)
}
