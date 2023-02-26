package main

import (
	"io"
	"log"
	"net"
	"os"
)

// Netcat is a simple read/write client for TCP servers.
func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// 获取服务器的数据
	go mustCopy(os.Stdout, conn)

	// 向服务器发送数据
	// io.Copy 会一直复制，直到遇到 EOF 或者出错
	// 因为 os.Stdin 没有 EOF，所以再不出错的情况下会一直复制，出现阻塞
	// io.Copy(os.Stdout, os.Stdin) 会不断读取一个输入然后输出
	mustCopy(conn, os.Stdin)
	
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
