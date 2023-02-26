package main

import (
    "fmt"
    "tempconv"
)

func main() {

    boilingF := tempconv.CToF(tempconv.BoilingC)
    fmt.Println(boilingF)

    var c tempconv.Celsius
    var f tempconv.Fahrenheit

    // 虽然两个新的类型使用相同的底层类型，但不是相同的类型
    // 两个类型的变量不能使用算术表达式进行比较和合并
    // 以下运算会报错
    // fmt.Println(boilingF - tempconv.FreezingC)
    // 对于每一个类型 T，都有一个对应的类型转换操作 T(x) 将 x 转换为类型 T 
    fmt.Println(c == tempconv.Celsius(f))

    fmt.Println(tempconv.CToF(tempconv.AbsoluteZeroC))
}