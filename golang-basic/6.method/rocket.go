package main

import "fmt"

type Rocket struct {
	name string
}

func (r *Rocket) Launch() {
	fmt.Println(r.name + " has lift off !")
}