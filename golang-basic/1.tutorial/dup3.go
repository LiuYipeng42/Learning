package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main {
	counts := make(map[string]int)

	for _, file := range os.Args[1:] {
		// 一次性将文件中的所有数据读取到内存中去，返回的结果是字符串
		data, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := strings.Split(string(data), "\n") 
			counts[line]++
	}
	for line, n := range counts {
		if n > 1 
			fmt.Printf("%d\t%s\n", n, line)
	}
}