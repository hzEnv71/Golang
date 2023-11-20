package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

type Client struct {
	C    chan string //用户发送数据的管道
	Name string      //用户名
	Addr string      //网络地址
}

var online map[string]Client //保存在线用户

var message = make(chan string) //通信管道

func main() {
	//监听
	listener, err := net.Listen("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	defer listener.Close()

	//子协程，转发消息给map中每个在线用户
	go Manager()

	//主协程，阻塞等待用户链接
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("err=", err)
			continue
		}
		go HandleConn(conn) //处理用户链接
	}

}

func Manager() {
	//分配空间
	online = make(map[string]Client)
	for {
		msg := <-message //没有消息会阻塞
		//遍历map，给每个map成员发送消息
		for _, cli := range online {
			cli.C <- msg
		}
	}
}

func HandleConn(conn net.Conn) {
	defer conn.Close()
	//获取客户端网络地址
	cliAddr := conn.RemoteAddr().String()
	//创建结构体，默认用户名=地址
	cli := Client{make(chan string), cliAddr, cliAddr}
	//把结构体添加到map
	online[cliAddr] = cli
	//子协程，给当前客户端发送消息
	go sendMsgTOClient(cli, conn)
	//广播某个在线
	message <- "[" + cli.Addr + "]" + cli.Name + ":login"
	//提示，我是谁
	cli.C <- "[" + cli.Addr + "]" + cli.Name + "I am here"

	isQuit := make(chan bool)  //对方是否主动退出
	hasData := make(chan bool) //对方是否有数据发送

	//新建协程，转发接收用户发来的数据
	go func() {
		buf := make([]byte, 1024)
		for {
			n, err := conn.Read(buf)
			if n == 0 { //对方断开或者出问题
				isQuit <- true
				fmt.Println("err=", err)
				return
			}
			msg := string(buf[:n-1])
			//查询用户在线人数
			if len(msg) == 3 && msg == "who" {
				//遍历map，给当前用户发送所有成员
				conn.Write([]byte("user list:\n"))
				for _, cli := range online {
					msg = cli.Addr + ":" + cli.Name + "\n"
					conn.Write([]byte(msg))
				}
			} else if //修改用户名
			len(msg) >= 8 && msg[:6] == "rename" {
				//rename|mike
				name := strings.Split(msg, "|")[1]
				cli.Name = name
				online[cliAddr] = cli
				conn.Write([]byte("rename ok\n"))
			} else {
				//转发此内容
				message <- "[" + cli.Addr + "]" + cli.Name + ":" + string(msg)
			}
			hasData <- true //代表有数据
		}
	}()
	//协程不能停
	for {
		//通过select检测channel的流动
		select {
		case <-isQuit:
			delete(online, cliAddr)                                   //移除当前用户
			message <- "[" + cli.Addr + "]" + cli.Name + ":login out" //广播谁下线了
			return
		case <-hasData:
		case <-time.After(30 * time.Second): //30s一直没数据进来就退出
			delete(online, cliAddr)
			message <- "[" + cli.Addr + "]" + cli.Name + ":time out leave out" //广播谁超时来
			return
		}
	}

}

func sendMsgTOClient(cli Client, conn net.Conn) {
	for msg := range cli.C { //有信息就写
		conn.Write([]byte(msg + "\n"))
	}
}
