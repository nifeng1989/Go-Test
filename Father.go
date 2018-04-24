package main

import "fmt"

type Father interface {
	speak()
}

type Son struct {
	age int
	Father
}

func (Son) speak() {
	fmt.Print("Son")
}