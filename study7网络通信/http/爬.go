package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	//明确目标
	url := "https://pic.netbian.com/"
	//爬
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	defer resp.Body.Close()
	//读取网页内容
	buf := make([]byte, 1024*1000)

	for {
		n, err := resp.Body.Read(buf)
		if n == 0 {
			fmt.Println("Read err=", err)
			break
		}
	}
	fileName := "pic.html"
	file, err1 := os.Create(fileName)
	if err1 != nil {
		fmt.Println("err1=", err1)
		return
	}

	//file.Write(buf) //往文件写内容
	file.WriteString(string(buf))
	file.Close()
}
