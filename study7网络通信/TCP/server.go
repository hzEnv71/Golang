package main

import (
	"bufio"
	"fmt"
	"net"
	proto "study7网络通信/protocol"
)

func main1() {
	//监听
	listener, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	defer listener.Close()
	//阻塞等待用户链接
	con, err := listener.Accept()
	if err != nil {
		fmt.Println("err=", err)
		return
	}

	//接收用户的请求
	buf := make([]byte, 1024) //缓冲区大小
	n, err1 := con.Read(buf)
	if err1 != nil {
		fmt.Println("err1=", err1)
		return
	}
	fmt.Println("buf=", string(buf[:n])) //读多少打印多少

	defer con.Close()

}

func main() {
	//监听
	listener, err := net.Listen("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	defer listener.Close()
	//并发，接收多个用户
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("err=", err)
			return
		}
		//处理请求
		//新建协程
		go HandleConn(conn)
	}

}

func HandleConn(conn net.Conn) {
	defer conn.Close()
	//获取客户端的网络地址
	addr := conn.RemoteAddr().String()
	fmt.Println(addr, "connection successful")

	//buf := make([]byte, 1024)
	//不断接收用户数据
	/*for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("err=", err)
			return
		}
		fmt.Printf("[%s]:%s\n", addr, string(buf[:n]))
		fmt.Println(addr, string(buf[:n]))
		//结束
		//if "exit" == string(buf[:n-1]) { //去除\n ,nc测试
		if "exit" == string(buf[:n-1]) { //\n
			fmt.Println(addr, "exit")
			return
		}
		//把数据转换为大写，再给用户发送
		conn.Write([]byte(strings.ToUpper(string(buf[:n]))))
	}*/
	reader := bufio.NewReader(conn)
	for {

		str, err := proto.Decode(reader)
		if err != nil {
			fmt.Println("err=", err)
			return
		}
		fmt.Println("str=", str)
	}
}
