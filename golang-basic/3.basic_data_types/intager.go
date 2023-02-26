package main

import "fmt"

// 整数

// 类型
//  有符号整数：int、int8、int16、int32、int64
//  无符号整数：uint、uint8、uint16、uint32、uint64
//  其中 int 和 uint 根据 CPU 的总线位数不同而不同，现在通常是 64 位，但与 int64 不是等价的类型
//  rune 类型是和 int32 等价的类型，通常用于表示一个 Unicode 码点
//  byte 类型 和 uint8 类型是等价类型，byte 类型一般用于强调数值是一个原始的数据而不是一个小的整数
//  uintptr 类型也是一种无符号的整数类型，没有指定具体的 bit 大小但是足以容纳指针，uintptr 类型只有在底层编程时才需要

func main() {

	// 运算
	//  % 取模仅能用于整数，余数的正负号总是与被除数一致
	fmt.Println(3%2, -3%2)

	// 除法的操作数都为整数时，结果为整数
	fmt.Println(1/2, 1.0/2, 1/2.0)

	// 有符号数的右移操作是按符号位的值补左边空位，无符号数的右移操作是用 0 补左边空位
	// 左移操作是用 0 补右边空位
	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2

	// 08 表示用 0 补够 8 位，%b 表示以二进制形式输出
	fmt.Printf("%08b\n", x)
	fmt.Printf("%08b\n", y)
	fmt.Printf("%08b\n", x&y)	// AND
	fmt.Printf("%08b\n", x|y)	// OR
	fmt.Printf("%08b\n", x^y)	// XOR
	fmt.Printf("%08b\n", x&^y)	// AND NOT (位清除， x 与非 y 进行与运算)

	// % 之后的 [n] 表明重复使用第 n 个操作数
	fmt.Printf("%b %[1]b %[1]b", x)
}
