package main

import (
	"fmt"
	"reflect"
)

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
	ch := ChimeraBird{}
	t := reflect.TypeOf(ch)
	m, ok := t.MethodByName("Fly")
	if ok {
		fmt.Println(t.Name(), "has method", m.Name)
	} else {
		fmt.Println(t.Name(), "does NOT have method", "Fly")
	}
}
