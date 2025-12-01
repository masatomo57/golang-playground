package main

import "fmt"

type Flyer struct{}

func (b Bird) Fly() {
	fmt.Println("Flyer fly")
}

type Bird struct {
	Flyer
}

func (b Bird) Fly() {
	fmt.Println("Bird fly")
}

func main() {
	crow := Bird{}
	crow.Fly()
}
