package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	// 创建一个 listener 对象，监听一个网络端口的连接
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		// Accept() 方法被阻塞，直到监听的端口有连接
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn) // 并发处理连接
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		// time.Now().Format() 的参数是一个模板
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}