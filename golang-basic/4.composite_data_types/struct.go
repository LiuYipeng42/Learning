package main

import (
	"encoding/json"
	"fmt"
	"time"
)

// 定义一个结构体
// 成员变量的顺序不同，则就不是一个结构体
// 如果成员变量的名称是首字母大写，则这个变量是可导出的，可以在其他包里使用
type Employee struct {
	ID            int
	Name, Address string
	DoB           time.Time
	Position      string
	Salary        int
	ManagerID     int
}

type tree struct {
	value       int
	left, right *tree
}

type Point struct{ X, Y int }

func main() {
	// 若初始化时省略表达式，则 go 会将变量初始值设为对应类型的零值
	// slice 和 map 的零值都是 nil，结构体的零值不是 nil，
	// 而是一个结构体的引用，并且所有成员变量都是对应类型的零值
	var dilbert Employee
	dilbert.Salary = 10000
	dilbert.Salary -= 5000

	// 获取成员变量的地址，然后通过指针来访问
	position := &dilbert.Position
	*position = "Senior " + *position

	// . 点号同样可以用在结构体指针上
	var employeeOfTheMonth *Employee = &dilbert
	employeeOfTheMonth.Position += " (proactive team player)"
	// (*employeeOfTheMonth.Position) += " (proactive team player)"

	// 通常是利用指针来操作结构体，较为简单的创建结构体指针的方法
	e := &Employee{}

	fmt.Println(dilbert, e)

	// ------------------------------------------------------------
	// 结构体字面量
	// 按照顺序写上成员变量的值
	p := Point{1, 2}
	// 可以指定每一个变量的值，可以不必按照顺序
	q := Point{X: 1}
	// 两种方式不能混合使用
	fmt.Println(p, q)

	// 结构体可以作为函数的参数和返回值，但是为了保证效率，
	// 使用结构体的指针作为参数或返回值的效果会更好
	p = Scale(p, 2)
	dilbert.Salary = Bonus(&dilbert, 20)

	// 如果结构体的所有成员变量都可以比较，那么这个结构体就是可比较的
	// 同时也可作为 map 的键
	fmt.Println(p == q)

	// ------------------------------------------------------------
	// 成员变量可以是不带名称的结构体类型，称为匿名成员
	// 一个结构体里不能有两个相同类型的匿名成员
	// 匿名成员是可以导出的
	type Circle struct {
		Point
		Radius int
	}
	type Wheel struct {
		Circle
		Spokes int
	}
	// 可以直接访问到匿名成员的成员变量
	var w Wheel
	w.X = 8
	w.Y = 8
	w.Radius = 5
	w.Spokes = 20
	// 也可以通过显式指定中间你名成员的方式访问变量
	w.Circle.Point.X = 2

	// 匿名成员初始化的方法
	wh := Wheel{Circle{Point{8, 8}, 5}, 20}
	whe := Wheel{
		Circle{
			Point:  Point{X: 8, Y: 8},
			Radius: 5,
		},
		20,
	}
	fmt.Printf("%#v\n", wh)
	fmt.Printf("%#v\n", whe)

	// ------------------------------------------------------------
	// 结构体与 JSON 字符串的转换
	// 结构体 --> JSON 字节切片
	data, err := json.Marshal(w)
	if err == nil {
		fmt.Printf("%s\n", data)
	}
	// 结构体 --> JSON 格式化后的字节切片
	// 参数：结构体，每行的前缀，缩进
	if data, err := json.MarshalIndent(w, "", "	 "); err == nil{
		fmt.Printf("%s\n", data)
	}
	// JSON 字节切片 --> 结构体
	t := &Wheel{}
	if err := json.Unmarshal(data, t); err == nil{
		fmt.Println(*t)
	}

	// ------------------------------------------------------------
	// 可以通过 text/template 和 html/template 将程序变量带入到文本或 HTML 模板中
	// 和 printf 类似，但功能更强大

}

func Scale(p Point, factor int) Point {
	return Point{p.X * factor, p.Y * factor}
}

func Bonus(e *Employee, percent int) int {
	return e.Salary * percent / 100
}
