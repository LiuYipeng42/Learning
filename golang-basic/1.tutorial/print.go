package main

import "fmt"

/*
%d 				十进制整数
%x %o %b		十六进制、八进制、二进制整数
%f %g %e		浮点数
%t				布尔型
%c				字符
%s				字符串
%q				带引号字符串，"abc" 或 'a'
%v				内置格式的任何值
%T				任何值的类型
%%				百分号本身
*/


func main() {
	num := 100.42
	// %f 无指数形式
	// %g 自动保持足够的精度
	// %e 指数形式的浮点数
	fmt.Printf("%f %[1]g %[1]e", num)
}