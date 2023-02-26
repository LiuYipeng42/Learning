package main

import "fmt"

// 变量声明

func main() {
	// 变量声明的格式：
	// var name type = expression
	// type 类型 和 expression 表达式 可以省略一个
	// 如类型省略，则变量的类型由初始化的表达式决定
	// 如表达式省略，则初始值为对应类型的零值
	var a, b, c int
	var d, e, f = true, 2.3, "four"

	// 短变量声明，:= 表示声明，
	// 如果 := 左边的变量已经声明过，则 := 只是赋值
	// 如果 := 左边的变量不能全是已经声明的变量，会报错
	//  = 左边的变量可以是全是已声明的变量
	g := 42
	g, h := 1, 2
	// a, b := 1, 2
	a, b = b, a //交换 a, b 的值

	fmt.Println(a, b, c, d, e, f, g, h)

	// 指针
	x := 1
	// 利用 & 运算获取变量 x 的地址，p 的类型是 *int
	p := &x 
	// 利用 * 运算访问指针指向的地址
	fmt.Println(*p)
	*p = 2
	fmt.Println(x)

	// new 函数创建变量
	// new(T) 可以创建一个未命名的 T 类型变量，初始化为 T 类型的零值，并返回其地址
	// 创建一个 *int 类型的指针
	p1 := new(int) 
	fmt.Println(*p1)
	*p1 = 42
	fmt.Println(*p1)
}

// 返回的变量 v 的地址即使在函数返回之后也是安全可用的
// 每次调用都会返回不同的地址
func f() *int {
	v := 1
	return &v
}
