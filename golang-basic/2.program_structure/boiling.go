package main

import "fmt"

// 全局变量

// 包级别的变量，只要是在 main 包中的函数都可以访问
const boilingF = 212.0

func main() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %g F or %g C\n", f, c)
}