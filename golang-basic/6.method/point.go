package main

import (
	"math"
)

type Point struct {
	X, Y float64
}

func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// 将方法绑定到参数 p 对应的类型上
// 参数 p 成为方法的接收者，Go 语言的接收者不使用特殊名（ self 或 this 等），
// 而是可自定义的任意名，最常用的方法就是取类型名称的首字母
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p Point) Add(q Point) Point {return Point{p.X + q.X, p.Y + q.Y}}

func (p Point) Sub(q Point) Point {return Point{p.X - q.X, p.Y - q.Y}}

// ---------------------------------------------------------------
// 指针接收者的方法
// 习惯上要遵守一个类型的所有方法的接收者要么是指针接收者，要么都不是指针接收者
func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.X *= factor
}

// 不允许本身是指针的类型进行方法声明
type P *int
// func (P) f() {}


