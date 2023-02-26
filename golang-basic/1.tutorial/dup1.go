package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// map[数据类型]数据类型 ：创建一个字典，第一个数据类型是键，第二个是值
	// make 函数可以用来创建 map
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	// Scan 函数会读取一行，并曲调换行符，若有新的数据，则返回 true ，否则返回 false
	for input.Scan() {
		// 若是新创建的键值对，会自动对值赋上默认的初始值
		// Text 函数可以获得 Scan 函数读取到的内容
		counts[input.Text()]++
	}
	// map 的迭代次序是随机的，每次的运行结果都不一样
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}

}
