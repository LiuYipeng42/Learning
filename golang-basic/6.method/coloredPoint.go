package main

import "image/color"

type ColoredPoint struct {
	// 通过内嵌匿名字段来实现方法的整合
	Point
	// 也可以使用指针类型，两者用法相似，都可以直接调用对应的方法
	// 使用也更加灵活
	// *Point
	Color color.RGBA
}