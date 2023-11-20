package proto

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
)

// 编码
func Encode(msg string) ([]byte, error) {
	// 读取消息长度，转成 int32 【刚好 4个字节】
	var length = int32(len(msg))
	var pkg = new(bytes.Buffer)
	// 写入消息头 【这里简单处理：整个消息头就只存储了 消息实体的长度】
	err := binary.Write(pkg, binary.LittleEndian, length)
	if err != nil {
		fmt.Println("write error:", err)
		return nil, errors.New("write error")
	}

	// 写入消息实体
	e := binary.Write(pkg, binary.LittleEndian, []byte(msg))
	if err != nil {
		fmt.Println("write entity failed:", e)
		return nil, errors.New("write entity error")
	}

	return pkg.Bytes(), nil
}

// 解码
func Decode(reader *bufio.Reader) (string, error) {
	// 读取消息的长度
	lengthByte, _ := reader.Peek(4) // 读取前4个字节的数据
	lengthBuff := bytes.NewBuffer(lengthByte)
	var length int32
	err := binary.Read(lengthBuff, binary.LittleEndian, &length) // 把 数据实体的长度 读出来赋给 length
	if err != nil {
		return "", err
	}
	// Buffered返回缓冲中现有的可读取的字节数。
	if int32(reader.Buffered()) < length+4 {
		return "", err
	}

	// 读取真正的消息数据
	pack := make([]byte, int(4+length))
	_, err = reader.Read(pack)
	if err != nil {
		return "", err
	}
	return string(pack[4:]), nil
}
