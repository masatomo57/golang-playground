package main

import "fmt"

type Flyer1 struct{}

func (b Flyer1) Fly() {
	fmt.Println("Flyer1 fly")
}

type Flyer2 struct{}

func (b Flyer2) Fly() {
	fmt.Println("Flyer2 fly")
}

type ChimeraBird struct {
	Flyer1
	Flyer2
}

func main() {
	crow := ChimeraBird{}
	crow.Fly()
}
