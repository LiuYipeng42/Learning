package main

import "fmt"

type IntList struct {
	Value int
	Tail *IntList
}

func (list *IntList) Sum() int {
	if list == nil {
		fmt.Print("sum = ")
		return 0
	}
	return list.Value + list.Tail.Sum()
}

