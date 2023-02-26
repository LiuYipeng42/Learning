package main

import (
	"flag"
	"fmt"
	"strings"
)

// 标识变量

// 创建一个标识变量，参数分别为： 
//    变量名、默认值、提示（赋值非法、-h 和 -help 时会出现）
// 返回值时一个指向标识变量的指针
// 在命令行中可以使用 -n 来赋值
var n = flag.Bool("n", false, "omit trailing newline")
// 在命令行中可以使用 -s 来赋值
var sep = flag.String("s", " ", "separator")

func main() {
	// 使用标识前，必须调用 Parse 函数来更新标识变量的默认值
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))

	if !*n {
		fmt.Println()
	}
}
