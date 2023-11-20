package main

import (
	"fmt"
	"net/http"
)

/*
w:给客户端回复数据
req:读取客户端发送的数据
*/

func HandConn(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Method:", request.Method)
	fmt.Println("URL:", request.URL)
	fmt.Println("RemoteAddr:", request.RemoteAddr)
	fmt.Println("Body:", request.Body)
	fmt.Println("UserAgent:", request.UserAgent())
	writer.Write([]byte("hello go"))
}
func main() {
	//注册处理函数，用户连接，自动调用指定的处理函数
	http.HandleFunc("/go", HandConn)

	//监听绑定
	http.ListenAndServe(":8888", nil)
}
