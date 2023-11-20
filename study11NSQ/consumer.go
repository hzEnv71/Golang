package main

import (
	"fmt"
	"github.com/nsqio/go-nsq"
)

type Handler struct{}

// HandleMessage 实现的处理消息的方法
func (m *Handler) HandleMessage(msg *nsq.Message) (err error) {
	addr := msg.NSQDAddress
	message := string(msg.Body)
	fmt.Println(addr, message)
	return
}
func NewConsumers(t string, c string, addr string) error {
	conf := nsq.NewConfig()
	consumer, err := nsq.NewConsumer(t, c, conf)
	if err != nil {
		fmt.Println("create consumer failed err ", err)
		return err
	}
	handle := &Handler{}
	consumer.AddHandler(handle) //增加处理
	// 连接nsqlookupd
	if err := consumer.ConnectToNSQLookupd(addr); err != nil {
		fmt.Println("connect nsqlookupd failed ", err)
		return err
	}
	return nil
}
func main() {
	// 这是nsqlookupd的地址
	addr := "127.0.0.1:4161"
	err := NewConsumers("topic-demo1", "channel-aa", addr)
	if err != nil {
		fmt.Println("new nsq consumer failed", err)
		return
	}
	select {}
}
