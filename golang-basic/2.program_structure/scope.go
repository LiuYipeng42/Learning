package main

import (
	"fmt"
	"log"
	"os"
)

// 作用域

// 语法块：
//   是由大括号围起来的一个语句序列，语法块内部声明的变量对外部不可见
// 词法块：
//   将块的概念推广到其他没有显示包含在大括号中的声明代码
// 全局块：
//   包含了全部源代码的词法块

var cwd string

func init() {
	// 因为 cwd 与 err 在 init 函数块的内部都尚未声明，
	// 所以 := 语句将他们都声明为局部变量
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("%v", err)
	}
	fmt.Println(cwd)
}

func main() {
	x := "hello"
	// for 循环创建了两个词法块
	// 一个程序可以包含多个同名的声明，前提是它们在不同词法块中
	// 当编译器遇到一个引用时，将从最内层的封闭词法块到全局块寻找其声明
	// 如果在内层和外层块都存在某个声明（x变量），则内层声明将覆盖外部声明
	for _, x := range x {
		x := x + 'A' - 'a'
		fmt.Printf("%c", x)
	}

	// 有三个词法块
	// 一个包含整个 if 语句
	// 两个是两组大括号
	if f, err := os.Open("./scope.go"); err != nil {
		fmt.Println(err)
	} else {
		// f 与 err 在这里可见
		f.Stat()
		f.Close()
	}
}
