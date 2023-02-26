package main

import (
	"fmt"
	"time"
)

func main() {
	const_()
	constGenIota()
	nonTypeConst()
}

func const_() {
	const answer = 42

	const (
		e  = 2.71
		e1 // e1 与 e 的值相同
		pi = 3.1415926
		pi1
	)
	fmt.Println(e, e1)

	// 创建一个特定类型的常量
	const noDelay time.Duration = 0
	const timeout = 5 * time.Minute
	fmt.Printf("%T %[1]v\n", noDelay)
	fmt.Printf("%T %[1]v\n", timeout)
	fmt.Printf("%T %[1]v\n", time.Minute)
}

func constGenIota() {
	// 常量生成器

	// type 声明定义一个新的命名类型，它和某个已有类型使用同样的底层类型
	type Weekday int

	// iota 从 0 开始取值，逐项加一
	// Sunday 的值位 0，Monday 的值为 1，以此类推
	const (
		Sunday Weekday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Saturday
	)

	const (
		_   = 1 << (10 * iota)
		KiB // 1024
		MiB // 1048576
		GiB // 1073741824
		TiB
		PiB
		EiB
		ZiB
		YiB
	)
}

func nonTypeConst() {
	// 无类型常量
	// 常量可以是任何基本数据类型，从属类型待定的常量有 6 种：
	//   无类型布尔、无类型整数、无类型符号、无类型浮点数、无类型复数、无类型字符串
	// 无类型常量的精度要比基本类型的数字高
	// 精度至少达到 256 位
	// 利用推迟确定从属类型，可以暂时维持更高的精度和进行无需转换类型的运算

	const (
		_   = 1 << (10 * iota)
		KiB // 1024
		MiB // 1048576
		GiB // 1073741824
		TiB
		PiB
		EiB
		ZiB
		YiB
	)
	// YiB 与 ZiB 的值都过大，无法用其他的整型来存储
	fmt.Println(YiB / ZiB)

	// 等号右边的无类型常量转变为 float64 类型
	var f float64 = 3 + 0i
	f = 2
	f = 1e123
	f = 'a'
	fmt.Printf("%T\n", f)
	// 与以下等价
	// var f float64 = float64(3 + 0i)
	// f = float64(2)
	// f = float64(1e123)
	// f = float64('a')

	// 变量声明中假如没有现式指定类型，无类型常量会转换成默认类型
	a := 0
	b := '\000'
	c := 0.0
	d := 0i
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", d)

	// 将无类型常量转换为制定类型
	var i = int8(0)
	//或 var i int8 = 0
	fmt.Printf("%T\n", i)

}
