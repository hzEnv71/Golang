package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	//监听
	listener, err := net.Listen("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	defer listener.Close()
	//阻塞等待用户连接
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	//接收文件名
	buf := make([]byte, 1024)
	var n int
	n, err = conn.Read(buf)
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	fileName := string(buf[:n])
	//回复OK
	conn.Write([]byte("OK"))
	//接收文件内容
	receiveFile(fileName, conn)

}

func receiveFile(fileName string, conn net.Conn) {
	//新建文件
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	buf := make([]byte, 1024)
	//收多少，写多少
	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("文件接收完毕")
			} else {
				fmt.Println("conn.Read err=", err)
			}
			return
		}
		file.Write(buf[:n])
	}

}
