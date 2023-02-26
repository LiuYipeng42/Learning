package main

import (
	"fmt"
	"image/color"
	"sync"
	"time"
)

// go run *.go
func main() {

	// 方法声明与指针接收者
	// pointerReceiver()
	// 结构体内嵌
	// structEmbedding()
	// 方法变量与方法表达式
	// methodValuesAndExpressions()
	// setter 和 getter 方法
	setterAndGetter()
}

func pointerReceiver() {
	// 方法的声明
	p := Point{1, 2}
	q := Point{4, 6}
	fmt.Println(Distance(p, q))
	fmt.Println(p.Distance(q))

	// 指针接收者
	(&p).ScaleBy(2)

	// 方法的实参接收者可以是 nil
	var list IntList
	// 第一个结点的 Tail 属性是 nil，第一次递归时的接收者是 nil
	fmt.Println(list.Sum())

	// 如果所有类型 T 方法的接收者是 T 自己（而非 *T），那么赋值它的实例是安全的
	// 但任何方法是指针的情况下，应该避免复制 T 的实例，因为这么做可能会破坏内部原本的数据

	// 方法调用表达式的三种形式：
	//  1. 实参接收者和形参接收者是同一个类型
	//  2. 实参接收者是 T 类型的变量而形参接收者是 *T 类型，编译器会隐式地获取变量的地址
	//  3. 实参接收者是 *T 类型的变量而形参接收者是 T 类型，编译器会隐式地解引用接收者，获得实际的取值
}

func structEmbedding() {
	// 结构体内嵌
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	var p = ColoredPoint{Point{1, 1}, red}
	var q = ColoredPoint{Point{5, 4}, blue}
	// Point 类型的方法都被纳入到 ColoredPoint 类型中
	// 但并不可以理解为 ColoredPoint 是 Point 的子类
	fmt.Println(p.Distance(q.Point))
	p.ScaleBy(2)
	q.ScaleBy(2)
	fmt.Println(p.Distance(q.Point))
	p.Point = q.Point
	fmt.Println(p.Point, q.Point)

	// 使用内嵌可以在未命名的结构体类型中声明方法
	// sync.Mutex 是内嵌的，Lock 和 Unlock 方法也整合进了 cache
	fmt.Println(Lookup("1"))
}

var cache = struct {
	sync.Mutex
	mapping map[string]string
}{
	mapping: make(map[string]string),
}

func Lookup(key string) string {
	cache.Lock()
	v := cache.mapping[key]
	cache.Unlock()
	return v
}

func methodValuesAndExpressions() {
	p := Point{1, 2}
	q := Point{4, 6}

	// 方法变量
	distanceFromP := p.Distance
	fmt.Println(distanceFromP(q))

	r := new(Rocket)
	r.name = "falcon heavy"
	time.AfterFunc(time.Second, r.Launch)
	time.Sleep(2 * time.Second)

	// 方法表达式
	// 将一个对象的方法变成一个函数，原来的接收者变成函数的第一个参数
	// 与方法变量不同的的是，方法表达式来自类的方法，而不是一个实例的方法
	distance := Point.Distance
	fmt.Println(distance(p, q))
	fmt.Printf("%T\n", distance)

	scale := (*Point).ScaleBy
	scale(&p, 2)
	fmt.Println(p)

	path := make(Path, 2)
	path[0] = p
	path[1] = q

	// 示例：非负整数位向量
	setA := IntSet{}
	setA.Add(1)
	setA.Add(80)

	setB := IntSet{}
	setB.Add(2)
	setB.Add(180)

	setA.UnionWith(&setB)

	fmt.Println(setA.Has(2))
	fmt.Println(setA.String())
	// fmt 包的函数会自动调用结构体的 String 方法
	// String 方法的接收者是指针，所以会正常输出
	fmt.Println(&setA)
	// setA 结构体并没有 String 方法，所以会直接输出结构体自身
	fmt.Println(setA)
}

func setterAndGetter() {
	// setter 和 getter 方法

	// l := Logger{}
	var l Logger
	l.SetFlags(1)
	l.SetPrefix("123")
	// getter 方法要省略 get
	fmt.Println(l.Flags(), l.Prefix())
}
