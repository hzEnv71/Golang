package main

import (
	"fmt"
	"net"
	proto "study7网络通信/protocol"
)

func main2() {
	//主动链接服务器
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	defer conn.Close()
	//发送数据
	conn.Write([]byte("hello,go"))

}

func main() {
	//主动链接服务器
	conn, err := net.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	defer conn.Close()

	/*go func() {
		//键盘键入内容，给服务器发送
		str := make([]byte, 1024)
		for {
			n, err := os.Stdin.Read(str)
			if err != nil {
				fmt.Println("err=", err)
				return
			}
			//发送到服务器
			conn.Write(str[:n])
		}
	}()*/
	go func() {
		for i := 0; i < 10; i++ {
			//调用协议编码数据
			b, err := proto.Encode("hello,go")
			if err != nil {
				fmt.Println("err=", err)
				return
			}
			conn.Write(b)
		}
	}()

	//接收服务器回复的消息
	//切片缓冲
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("err=", err)
			return
		}
		fmt.Println(string(buf[:n]))
	}
}
