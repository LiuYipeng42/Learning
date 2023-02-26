package main

import "fmt"


// 只能发送的通道：chan<-
func counter(out chan<- int) {
	for x := 0; x < 100; x++ {
		out <- x
	}
	close(out)
}

// 只能接收的通道：<-chan
func squarer(out chan<- int, in <-chan int) {
	for v := range in {
		out <- v * v
	}
	close(out)
	// 只能接收的通道不能被关闭，因为 close 操作是说明不会有数据发送到此通道上
	// close(in)
}

func printer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// 双向通道转换成单向通道是允许的，但反过来不行
	go counter(naturals)
	go squarer(squares, naturals)
	printer(squares)
}
