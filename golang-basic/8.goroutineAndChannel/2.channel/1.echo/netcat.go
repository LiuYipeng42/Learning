package main

import (
	"io"
	"log"
	"net"
	"os"
	"time"
)

// 无缓冲通道
// 1.向无缓冲通道发送数据会导致 goroutine 的阻塞，直到另一个 goroutine 接收了此数据
// 	 如果在 goroutine 中先执行接收操作，但通道中还没有数据，此 goroutine 就会阻塞，
//   直到直到另一个 goroutine 发送了数据，所以无缓冲通道可以用来同步多个 goroutine。
// 2.若存在一直被阻塞的 goroutine，这种情况叫做 goroutine 泄漏，
//   泄漏的 goroutine 不会自动回收，所以要确保 goroutine 在不再需要时可以自动结束。
func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// 使用 make 创建无缓冲通道，类型是 struct{}
	done := make(chan struct{})
	
	go func() {
		// 连接关闭时，服务器才会看到 EOF，并发送回来，之后才会停止复制
		io.Copy(os.Stdout, conn)
		log.Println("done")

		time.Sleep(2 * time.Second)
		// 向通道发送数据
		// 当通信没有携带额外信息时，叫做事件，通常使用的是 struct{} 类型的通道
		done <- struct{}{}
	}()

	mustCopy(conn, os.Stdin)
	// 关闭连接
	conn.Close()
	// 接收通道里的数据，可以丢弃结果
	// 因为在接收数据的 goroutine 中，在结束接收信息后会暂停 2 秒，
	// 为了让 goroutine 在主goroutine结束前结束，使用通道进行阻塞
	// x := <- done
	<-done
	// 关闭通道，并不是必须的，垃圾回收机制可以回收
	// 关闭后的通道里的数据可以接收，当数据接收完后，进行接收操作时得到的是对应类型的零值
	close(done)
	log.Println("finish")
	
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
