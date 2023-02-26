package main

import (
	"fmt"
	"math"
)

// 浮点数
// 浮点数有两种类型，float32 和 float64

func main (){
	// float32 的有效数组大约 6 位
	fmt.Printf("float32 的最大值：%f" , math.MaxFloat32)
	// float64 的有效数组大约 15 位
	fmt.Printf("float64 的最大值：%f" , math.MaxFloat64)

	var z float64

	fmt.Println(z, -z, 1/z, -1/z, z/z)

	nan := math.NaN()
	fmt.Println(nan, nan == nan)
}