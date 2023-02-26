package main

import (
	"bytes"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
)

// 创建一个接口，可以使用 type 为创建的接口命名
type Human interface {
	// 可以通过组合已有接口得到新接口
	// 此语法称为嵌入式接口，与嵌入式结构类似
	// 嵌入的接口可以有相同的方法
	Mouth
	Hands
	Feet
}

type Mouth interface{ Speak(words string) }
type Hands interface{ Touch(sth string) }
type Feet interface{ Walk(length int) }
type test interface{test(i int)}

// 实现以上接口
type Man struct{ name string }

func (m Man) Speak(words string) { fmt.Println("speak: " + words) }
func (m Man) Touch(sth string)   { fmt.Println("touch " + sth) }
func (m Man) Walk(length int)    { fmt.Println("walk " + fmt.Sprintf("%d", length) + " meter") }

type Woman struct{ name string }

func (m *Woman) Speak(words string) { fmt.Println("speak: " + words) }

func main() {
	// 接口创建，实现与赋值
	// Interface()
	// 接口值
	// InterfaceValues()
	// examples()
	TypeAssertions()

}

func Interface() {
	// 接口创建，实现与赋值
	var man Human = new(Man)
	man.Speak("42")

	var h Hands = new(Man)
	h.Touch("blackhole")
	// 只有接口暴露的方法才可以调用
	h.Speak("42")

	// women 实现了 Speak 方法，但方法的接收者是一个指针
	var woman Woman
	// 即使实际的接收者不是指针，此方法可以正常运行，
	// 编译器针对此情况作了优化，这只是一个语法糖
	woman.Speak("hhh")

	// 但因为接口默认的接收者不是指针，此时编译器并没有对应的优化，所以不能赋值
	var m Mouth = woman

	// 空接口
	var any interface{}

	// 可以将任何类型赋值给空接口，但不能调用任何方法
	any = true
	any = 12.34
	any = "hello"
	fmt.Println(any)
}

func InterfaceValues() {
	// 接口值
	// 接口值：
	// 	1. 动态类型：一个具体的类型
	// 	2. 动态值：具体类型的一个值

	// 动态类型：nil	动态值：nil
	var w io.Writer
	// 动态类型：*os.File	动态值：os.Stdout 的一个副本
	// 是一个隐式的类型转换，等价于 io.Writer(os.Stdout)
	w = os.Stdout
	fmt.Printf("%T\n", w)
	// 动态类型：*bytes.Buffer	动态值：指向新分配缓冲区的指针
	w = new(bytes.Buffer)
	fmt.Printf("%T\n", w)
	w = nil
	fmt.Println(w)

	// 接口值可以使用 == 和 != 操作符来做比较，可以作为 map 的键和switch 语句的操作数
	// 当两个接口值的动态类型与动态值都相等时两个接口值才相等
	// 动态类型相同（都为 Man），但动态值不同
	var a Human = new(Man)
	var b Human
	fmt.Println(a == b)

	// 空的接口值 与 动态值为 nil 的接口值不相等
	// 此变量时一个指针类型，值为 nil
	var c *Man
	fmt.Println(c == nil)
	// 此变量是一个接口，动态值虽然为 nil，但动态类型不为空，所以此变量不为 nil
	// 所以当调用此接口的方法时，此方法的接收值就为 nil，可能会出错
	var d Human = c
	fmt.Println(d == nil)
	var e Human
	// 空的接口值
	fmt.Println(e == nil)
}

func examples () {
	// 示例：排序
	// StringSlice 实现了 sort.Interface 接口的三个方法
	names := StringSlice{"c", "a", "b"}
	sort.Sort(names)
	fmt.Println(names)
	// 字符串排序可以直接使用 sort.Strings()

	// 示例：表达式计算器
	TestEval()
}

type StringSlice []string

func (p StringSlice) Len() int           { return len(p) }
func (p StringSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p StringSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func TestEval() {
	tests := []struct {
		expr string
		env  Env
		want string
	}{
		{"sqrt(A / pi)", Env{"A": 87616, "pi": math.Pi}, "167"},
		{"pow(x, 3) + pow(y, 3)", Env{"x": 12, "y": 1}, "1729"},
		{"pow(x, 3) + pow(y, 3)", Env{"x": 9, "y": 10}, "1729"},
		{"5 / 9 * (F - 32)", Env{"F": -40}, "-40"},
		{"5 / 9 * (F - 32)", Env{"F": 32}, "0"},
		{"5 / 9 * (F - 32)", Env{"F": 212}, "100"},
		//!-Eval
		// additional tests that don't appear in the book
		{"-1 + -x", Env{"x": 1}, "-2"},
		{"-1 - x", Env{"x": 1}, "-2"},
		//!+Eval
	}
	var prevExpr string
	for _, test := range tests {
		// Print expr only when it changes.
		if test.expr != prevExpr {
			fmt.Printf("\n%s\n", test.expr)
			prevExpr = test.expr
		}
		expr, err := Parse(test.expr)
		if err != nil {
			fmt.Println(err)
			continue
		}
		got := fmt.Sprintf("%.6g", expr.Eval(test.env))
		fmt.Printf("\t%v => %s\n", test.env, got)
		if got != test.want {
			fmt.Printf("%s.Eval() in %v = %q, want %q\n",
				test.expr, test.env, got, test.want)
		}
	}
}


func TypeAssertions() {
	// 类型断言是作用在接口值上的操作
	var h Human
	h = Man{"hhh"}
	// 类型断言会检查作为操作数的动态类型是否满足指定的断言类型
	// 如果检查成功，则返回接口值的动态值
	m := h.(Man)
	fmt.Println(m)
	// 如果检查不成功，会宕机
	// w := h.(*Man)
	// fmt.Println(w)
 
	var hand Hands = new(Man)
	// 如果检查的类型是一个接口类型，则返回值还是一个接口值，
	// 接口值的类型变为所检查类型，相当于进行类型转换
	man := hand.(Human)
	// 接口值可以进行比较
	fmt.Println(man == nil)
	// 返回的接口值的方法通常会变多
	man.Walk(42)

	// 断言可以返回两个返回值，第二个用于判断断言是否成功，此时断言失败不会宕机
	t, ok := hand.(test)
	fmt.Println(t, ok)

	// 类型分支
	// 判断接口值的动态类型
	fmt.Println(getType(1.0))
}

func getType(x interface{}) string {

	// x.(type) 是一种在 switch 语句中的简写
	// 是对，if _, ok := x.(int) 等语句的简化
	switch x := x.(type) {
	case nil:
		return "nil"
	case int, uint:
		return "int " + fmt.Sprintf("%d", x)
	case float32, float64:
		return "float64 " + fmt.Sprintf("%f", x)
	case bool:
		if x {
			return "bool true"
		}
		return "bool false"
	case string:
		return "string " + x
	default:
		return "unknow type"
	}
}
