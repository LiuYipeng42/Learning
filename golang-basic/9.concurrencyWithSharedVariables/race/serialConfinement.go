package main

import "fmt"


// 串行受限
type Cake struct{ i int; state string }

func baker(cooked chan<- *Cake) {
	i := 0
	for {
		cake := new(Cake)
		cake.i = i
		cake.state = "cooked"
		i ++
		// 在把变量地址传给下一步后，就不再访问此变量
		// 这样所有对这个变量的访问都是串行的，此变量也不会在多个 goroutine 中同时访问
		cooked <- cake
	}
}

func icer(iced chan<- *Cake, cooked <-chan *Cake) {
	for cake := range cooked {
		cake.state = "iced"
		iced <- cake
	}
}

func main() {
	cakes := make(chan *Cake, 3)
	finished :=  make(chan *Cake, 3)

	go baker(cakes)
	go icer(finished, cakes)

	for cake := range finished {
		fmt.Println(cake)
	}

}
