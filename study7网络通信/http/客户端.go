package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("http://127.0.0.1:8888/go")
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Status=", resp.Status)
	fmt.Println("StatusCode=", resp.StatusCode)
	fmt.Println("Header=", resp.Header)
	fmt.Println("Body=", resp.Body)
	buf := make([]byte, 1024)
	var tmp string
	for {
		n, err := resp.Body.Read(buf)
		if n == 0 {
			fmt.Println("Read err=", err)
			break
		}
		tmp += string(buf[:n])
	}
	fmt.Println("tmp=", tmp)
}
