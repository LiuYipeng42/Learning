package main

import "fmt"

func main() {
	ch := make(chan int, 1)
	// select 语句会运行第一个可以运行的情况，若都不能运行，则会阻塞
	// 如果多个情况同时满足，select 会随机选择一个
	// select 可以有 default 情况，当所有的情况都不可以运行时会运行 default，此时并不会阻塞
	for i := 0; i < 10; i++ {
		select {
		// 1. 循环开始时，通道内没有元素，所以从管道接收数据时会阻塞，此种情况不会运行
		// 2. 第二次循环时，通道内有元素 0，可以运行，取出后并且输出 0
		case x := <-ch:
			fmt.Println(x)
		// 1. 循环开始时，通道内没有元素，i=0，此情况可以运行，将 0 放入通道
		case ch <- i:
		// default:
		}
	}
}
