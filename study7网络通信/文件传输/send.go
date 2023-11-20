package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	//提示输入文件
	fmt.Println("请输入需要传输的文件")
	var path string
	fmt.Scan(&path)
	//获取文件名
	info, err := os.Stat(path)
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	//主动连接服务器
	conn, err1 := net.Dial("tcp", "127.0.0.1:9999")
	if err1 != nil {
		fmt.Println("net Dial err1=", err1)
		return
	}
	defer conn.Close()

	//给接收方发送文件名
	_, err = conn.Write([]byte(info.Name()))
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	//接收对方的回复，如果回复OK，说明对方准备好，可以发文件
	var n int
	buf := make([]byte, 1024)
	n, err = conn.Read(buf)
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	if "OK" == string(buf[:n]) {
		//发送文件内容
		sendFile(path, conn)
	}

}

func sendFile(path string, conn net.Conn) {
	//以只读方式打开文件
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	defer file.Close()
	buf := make([]byte, 1024)
	//读多少，发多少
	for {
		n, err := file.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("文件发送完毕")
			} else {
				fmt.Println("err=", err)
			}
			return
		}
		conn.Write(buf[:n]) //给服务器发送内容
	}
}
