// Pipeline1 demonstrates an infinite 3-stage pipeline.
package main

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go func() {
		for x := 0; x < 100; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	// Squarer
	go func() {

		// for {
		// 	// 如果通道关闭且读完，则返回 零值 和 false
		// 	x, ok := <-naturals
		// 	if !ok {
		// 		break
		// 	}
		// 	squares <- x * x
		// }
		// 以下代码是注释代码的简化版本
		// range 在通道上迭代，接收完最后一个值后关闭循环，
		// 使用 range 迭代的通道必须首先调用 close，否则会无法结束循环
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()

	// Printer (in main goroutine)
	for x := range squares {
		fmt.Println(x)
	}
}
