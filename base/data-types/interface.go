package main

import "fmt"

type Animal interface {
	color() string
}

type Cat struct {
}

func (cat Cat) color() string {
	fmt.Println("cat color white")
	return "white"
}

func tryInterface() {
	var animal Animal
	animal = new(Cat)
	fmt.Println(animal.color())

}
