package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, file := range files {
			// 返回一个打开的文件（*os.File）和 error 类型的值
			f, err := os.Open(file) 
			if err != nil {
				// 在标准错误流上输出
				// %v 可以使用默认格式显示任意的值
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}

}

func countLines(f *os.File, counts map[string]int) {
	// 用流式的方式逐行读取文件，不会将所有数据放入内存
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
