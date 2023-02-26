package main

import (
	"fmt"
	"strings"
)

// go run function.go examples.go
func main () {
	// 函数变量
	// funcVar()
	// 匿名函数
	// anonymousFunc()
	// 变长函数
	// variadicFunc()
	// 延迟函数
	// deferredFunc()
	// 宕机
	panic()
	
}

func funcVar() {
	// 函数变量
	f := square
	fmt.Println(f(3))

	// 当两个函数拥有相同的形参列表和返回列表时，则这两个函数的类型或函数签名相同
	// 所以 square 函数和 negative 函数的类型相同
	f = negative
	fmt.Println(f(3))
	fmt.Printf("%T\n", f)

	// 类型不同，编译错误
	// f = product

	var fun func(int) int
	// 函数变量可以与 nil 比较
	fmt.Println(fun == nil)
	// 函数变量之间不可比较，也不可以作为 map 的键
	// fmt.Println(fun == f)

	// 函数的行为可以当作参数进行传递
	// strings.Map 对字符串中的每一个字符使用一个函数
	fmt.Println(strings.Map(add1, "HAL-9000"))
}

func anonymousFunc() {
	// 匿名函数
	// func 关键字之后没有函数名称，只是一个表达式
	fmt.Println(strings.Map(func(r rune) rune { return r + 1 }, "HAL-9000"))

	// squares() 返回一个函数
	sq := squares()
	// 函数中的 x 变量的生命周期不是由它的作用域决定的，
	// 变量 x 在 squares() 返回后依旧存在
	fmt.Println(sq())
	fmt.Println(sq())
	fmt.Println(sq())

	// e.g. 每个课程的前置课程
	var prereqs = map[string][]string{
		"algorithms": {"data structures"},
		"calculus":   {"linear algebra"},
		"compilers": {
			"data structures",
			"formal languages",
			"computer organization",
		},
		"data structures":       {"discrete math"},
		"databases":             {"data structures"},
		"discrete math":         {"intro to programming"},
		"formal languages":      {"discrete math"},
		"networks":              {"operating systems"},
		"operating systems":     {"data structures", "computer organization"},
		"programming languages": {"data structures", "computer organization"},
	}

	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}

	// 捕获迭代变量
	var prints []func()
	// 在 for 循环开始时声明一次 i 变量，之后都是赋值
	for i := 0; i < 4; i++ {
		// 在 for 循环的作用域之内再声明一个同名变量
		// 这样每次循环都会创建一个新的变量，每个匿名函数中的 i 变量都不相同
		// i := i
		prints = append(prints, func() {
			// 匿名函数中的 i 变量的值会随着循环而变化
			// 所以所有的匿名函数的输出值都是 4
			fmt.Println(i)
		})
	}
	for _, f := range prints {
		f()
	}
}

func variadicFunc() {
	// 变长函数
	// 在函数的参数列表最后的类型名称之前使用省略号 ... 表明声明一个变长函数
	// 调用这个函数的时候可以传递该类型任意数目的参数
	fmt.Println(sum(1, 2, 3))
	values := []int{1, 2, 3, 4}
	fmt.Println(sum(values...))
}

func deferredFunc() {
	// 延迟函数
	buf, _ := readFile("./function.go")
	fmt.Println(string(buf)[:10])

	// 延迟调用可以用来调用一个较为复杂的函数
	// 使用延迟调用实现在函数的开头和结尾进行输出
	// bigSlowOperation()

	// 延迟执行的函数在 return 语句之后执行，并且可以更新函数的结果变量
	fmt.Println(triple(2))
}

func panic() {
	// 宕机
	// 当宕机发生时，所有的延迟函数以倒序执行，从栈最上面的函数开始一直返回至 main 函数
	// 宕机时还会输出宕机消息和栈转储信息，也可以使用 runtime 包主动转储栈的信息
	// defer printStack()
	// f(3)

	// 宕机恢复
	// recover 函数在延迟函数的内部调用时，若外部函数发生宕机，
	// recover 函数会终止当前的宕机状态并返回宕机的值
	f := func() {
		defer func() {
			if p := recover(); p != nil {
				err := fmt.Errorf("internal error: %v", p)
				if err != nil {
					fmt.Println(err)
				}
			}
		}()
		n := 0
		fmt.Println(1/n)
	}
	f()

}
