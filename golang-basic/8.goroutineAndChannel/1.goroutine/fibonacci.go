package main

import (
	"fmt"
	"time"
)

func main() {
	// 运行 main 函数的 goroutine 称为主 goroutine
	// 在函数或方法前加上 go 可以 goroutine
	go spinner(100 * time.Millisecond)
	const n = 42
	fibN := fib(n)
	fmt.Printf("\rFibonacci(%d) is %d\n", n, fibN)

	// 当 main 函数返回时，所有的 goroutine 都暴力地直接结束
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x - 1) + fib(x - 2)
}
