package main

import (
	"fmt"
	"os"
)

func main(){

	// 声明变量 变量名 变量类型
	var s, sep string
	// 声明变量的 4 种方式
	// s := ""
	// var s string
	// var s = ""
	// var s string = ""

	// := 短变量声明
	// i++ 是语句，不是表达式
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	// 每一次迭代，range 会产生一对值：索引和这个索引所处的值 
	// for _, arg := range os.Args[1:] {
	// 	s += sep + arg
	// 	sep = " "
	// }
	fmt.Println(s)
	
	// 效果与以上代码相同
	// fmt.Println(strings.Join(os.Args[1:], " "))

}