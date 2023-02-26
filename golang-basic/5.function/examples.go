package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"sort"
	"time"
)

func square(n int) int     { return n * n }
func negative(n int) int   { return -n }
func product(m, n int) int { return m * n }

func add1(r rune) rune { return r + 1 }

func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x) // 当 x = 0 时会宕机
	defer fmt.Printf("defer %d\n", x)
	f(x - 1)
}

func printStack() {
	// 输出栈的信息
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	os.Stdout.Write(buf[:n])
}

func triple(x int) (result int) {
	defer func() { result += x }()
	return x + x
}

func bigSlowOperation() {
	// trace() 函数并不会延迟调用，trace() 函数返回的函数会被延迟调用
	defer trace("bigSlowOperation")()
	time.Sleep(2 * time.Second)
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() {
		log.Printf("exit %s (%s)", msg, time.Since(start))
	}
}

func readFile(filename string) ([]byte, error) {
	// defer 后的函数或表达式是在 defer 所在函数结束后运行，
	// 函数结束可以是正常返回时，也可以是发生宕机时

	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	// defer 代码的执行顺序与最终的执行顺序是反向的，像栈一样
	defer f.Close()
	defer fmt.Println("--------------")
	defer fmt.Println("read file over")

	info, err := f.Stat()
	if (err != nil) {
		fmt.Println(err)
	}
	buf := make([]byte, info.Size())
	f.Read(buf)

	return buf, nil
}

func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

func squares() func() int {
	var x int
	// 里层函数可以使用外层函数中的变量
	return func() int {
		x++
		return x * x
	}
}

// 拓扑排序
func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)

	// 当一个匿名函数需要进行递归，必须先声明一个函数变量，再进行赋值
	var visitAll func(items []string)
	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}

	var keys []string

	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	// fmt.Println(keys)
	visitAll(keys)
	return order
}
