package main

import (
	"io"
	"log"
	"net"
	"os"
)


// 同时在多个终端中运行都可以得到结果，因为 clock.go 中使用了 go 语法
func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	// 从服务器读取数据
	mustCopy(os.Stdout, conn)
}

// 从网络连接中读取，然后写到标准输出
func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
