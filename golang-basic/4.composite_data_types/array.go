package main

import "fmt"

func main() {

	// 创建并输出数组
	// 一个新数组中的元素初始值为元素类型的零值
	var a [3]int
	a[0] = 1
	fmt.Println(a[0])
	for i, v := range a {
		fmt.Println(i, v)
	}

	// 利用数组字面量的一组数初始化数组
	var b [3]int = [3]int{1, 2, 3}
	var c [3]int = [3]int{1, 2}
	fmt.Println(b, c)

	// 使用 ... 时，数组的长度由初始化数组的元素个数决定
	d := [...]int{1, 2, 3}
	fmt.Println(d)

	e := [3]int{1, 2, 3}
	// 以下会出现编译错误，
	// 因为数组的长度是数组类型的一部分，所以不同长度的数组是不同类型的
	// java 或 python 的数组变量可以指向不同长度的数组
	// e = [2]int{1, 2}
	fmt.Println(e)

	// 可以任意顺序初始化数组
	f := [...]string{2: "2",1: "1",3: "3",4: "4"}
	fmt.Println(f)

	// 定义了一个长度为 10 的数组，且第 8 位为 -1，第 9 位为 10
	g := [...]int{8:-1, 10}
	fmt.Println(g)

	// 可以使用 == 或 != 来判断两个数组是否完全相等
	h := [2]int{1, 2}
	i := [...]int{1, 2}
	j := [2]int{1, 3}
	fmt.Println(h == i, h == j)

	// 数组类型不同，无法比较
	// fmt.Println([3]int{1, 2} == h)
	
}

