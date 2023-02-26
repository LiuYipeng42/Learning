package main

import "fmt"

// 函数声明

func main() {
	const freezingF, boilingF = 32.0, 212.0

	if t, ok := fToC(freezingF); ok {
		fmt.Printf("%g F = %g C\n", freezingF, t)
	}

	if t, ok := fToC(boilingF); ok {
		fmt.Printf("%g F = %g C\n", boilingF, t)
	}

	if t, ok := cToF(-100); ok {
		fmt.Printf("%g C = %g F\n", 100.0, t)
	}

}

// 函数名  参数列表       返回值列表
func fToC(f float64) (float64, bool) {
	// 函数的类型称为函数签名，当两个函数拥有相同的形参列表和返回列表时，
	// 则这两个函数的类型或函数签名相同

	// -459.669 绝对零度
	if f < -459.669 {
		return 0, false
	}
	return (f - 32) * 5 / 9, true
}

// 返回值列表可以声明变量，这些变量可以直接使用
func cToF(c float64) (f float64, ok bool) {

	if c < -273.15 {
		f, ok = 0, false
	} else {
		f, ok = c*9/5+32, true
	}

	// 如果函数有命名的返回值，则返回的参数可写可不写，称为裸返回
	// 裸返回是将每个命名返回结果按照顺序返回的快捷方式
	// return f, ok
	return
}

// 如果几个 形参 或者 返回值 的类型相同，那么类型只需要写一次
func add(x, y int) int {return x + y}
func sub(x, y int) (z int) {z = x - y;return }
func switchNum(x, y int) (x1, y1 int) {x1, y1 = y, x; return}

// go 语言的函数在调用时，传入的参数是原参数的副本，而不是原参数的引用，
// 所以在函数中修改一个数组参数，原数组并不会改变，
// 但是在转递较长的数组时，复制数组会很低效，所以可以使用指针的方式传递数组
func zero(ptr *[32]byte) {
	// 本函数可以清零一个长度为 32 的字节数组
	*ptr = [32]byte{}
}
