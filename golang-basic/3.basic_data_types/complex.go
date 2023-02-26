package main

import (
	"fmt"
	"math/cmplx"
)

// 复数
// 复数有两种类型:
//   complex64 和 complex128，二者分别由 float32 和 float64 构成

func main() {
	// 可以使用内置的 complex 函数根据给定的实部和虚部创建复数
	var x complex128 = complex(1, 2)
	var y complex128 = complex(3, 4)

	fmt.Println(x*y)
	// real 函数可以提取复数的实部
	fmt.Println(real(x*y))
	// imag 函数（imaginary）可以提取复数的实部
	fmt.Println(imag(x*y))

	// 数字后加上 i ，就变成了一个实部为 0 的复数
	fmt.Println(1i * 1i)

	// 复数可以与其他数想加，用于创建复数
	a := 1 + 2i
	b := 1.2 + 4i
	fmt.Println(a + b)

	// 可以利用 == 与 != 判断复数是否相等
	fmt.Println(a == b, a != b)

	// math/cmplx 包提供了复数运算所需的库函数
	fmt.Println(cmplx.Sqrt(-1))

}