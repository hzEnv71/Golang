package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

/*
{
	"company":"go"
	"subjects":[
		"GO"
		"Java"
		"Python"
	],
	"status":true
	"price":77.88
}
*/
//成员变量名大写

type IT struct {
	Company  string   `json:"company"` //二次编码
	Subjects []string `json:"subjects"`
	Status   bool     `json:"status"`
	Price    float64  `json:"price"`
}

func main() {
	/*//结构体初始化
	it := IT{"go", []string{"GO", "Java", "Python"}, true, 77.88}
	//编码，根据内容生成json文本
	//buf, err := json.Marshal(it)
	buf, err := json.MarshalIndent(it, "  ", "") //格式化编码
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(buf))*/

	/*//map生成json
	m := make(map[string]interface{}, 4)
	m["company"] = "Go"
	m["subjects"] = []string{"Go", "Java", "Python"}
	m["status"] = true
	m["price"] = 77.88
	//编码成json
	//result, err := json.Marshal(m)
	result, err := json.MarshalIndent(m, " ", " ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(result))*/

	/*	jsonBuf := `
		{
		"company":"go",
		"subjects":[
			"GO",
			"Java",
			"Python"
		],
		"status":true,
		"price":77.88
	}`*/
	/*
		//json解析到结构体
		var it IT
		err := json.Unmarshal([]byte(jsonBuf), &it) //参数要地址传递，要不然改变不了
		if err != nil {
			fmt.Println(err)
			return
		}
		//fmt.Println(it)
		fmt.Printf("%+v", it)*/

	/*//json解析到map
	m := make(map[string]interface{}, 4)
	err := json.Unmarshal([]byte(jsonBuf), &m) //第二个参数要地址传递
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", m)
	fmt.Println("==============")
	str := m["price"]
	fmt.Printf("%T\n", str)
	fmt.Println("==============")
	//类型断言
	for key, value := range m {
		fmt.Println(key, value)
		switch data := value.(type) {
		case string:
			fmt.Println(data)
		case bool:
			fmt.Println(data)
		case interface{}:
			fmt.Println(data)

		}
	}*/

	/*//IO
	os.Stdout.Write([]byte("输出"))
	//os.Stdout.Close() //关闭后无法输出
	fmt.Println("输出")

	//os.Stdin.Close() //关闭后无法输入
	var a string
	os.Stdin.Read([]byte(a))
	fmt.Println("请输入a：")
	fmt.Scan(&a)
	fmt.Println(a)*/

	//path := "./demo.txt"
	//WriteFile(path)
	//ReadFile(path)

	//拷贝文件
	// go run study5json和IO/src/server.go pic.jpg pic1.jpg
	list := os.Args //获取命令行参数
	if len(list) != 3 {
		fmt.Println("srcFile  dstFile")
		return
	}
	srcFileName := list[1]
	dstFileName := list[1]
	/*var srcFileName string
	var dstFileName string
	fmt.Println("请输入源文件：")
	fmt.Scan(&srcFileName)
	fmt.Println("请输入目的文件：")
	fmt.Scan(&dstFileName)*/
	if srcFileName == dstFileName {
		fmt.Println("源文件和目的文件不能相同")
		return
	}
	//只读方式打开源文件
	sF, err1 := os.Open(srcFileName)
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	//新建目的文件
	dF, err2 := os.Create(dstFileName)
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	//操作完毕，需要关闭文件
	defer sF.Close()
	defer dF.Close()
	//核心处理，从源文件读取内容，往目的文件写，读多少写多少
	buf := make([]byte, 4*1024) //临时缓冲区
	for {
		n, err := sF.Read(buf) //从源文件读取内容
		if err != nil {
			fmt.Println(n, err)
			if err == io.EOF {
				break
			}
		}
		//往目的文件写，读多少写多少
		dF.Write(buf[:n])
	}
}

func WriteFile(path string) {
	//打开文件，新建文件
	f, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	//使用完毕，需要关闭文件
	defer f.Close()
	var buf string
	for i := 0; i < 10; i++ {
		buf = fmt.Sprint(i)
		n, err := f.WriteString(buf)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(n)
	}
}
func ReadFile(path string) {
	//打开文件
	//f, err := os.Create(path)
	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	//使用完毕，需要关闭文件
	defer f.Close()
	buf := make([]byte, 1024) //1k
	//n代表从文件读取内容的长度
	n, err := f.Read(buf)
	if err != nil && err != io.EOF {
		fmt.Println(err)
		return
	}
	fmt.Println(string(buf[:n]))
}

// 每次读取一行
func ReadFileLine(path string) {
	//打开文件
	//f, err := os.Create(path)
	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	//使用完毕，需要关闭文件
	defer f.Close()
	//新建一个缓冲区，把内容先放到缓冲区
	r := bufio.NewReader(f)
	for {
		//遇到哦'\n'结束读取，但是'\n'也读取
		buf, err := r.ReadBytes('\n')
		if err != nil {
			if err == io.EOF { //结束
				break
			}
			fmt.Println(err)
		}
		fmt.Println(string(buf))
	}

}
