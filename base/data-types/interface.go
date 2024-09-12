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

type Dog struct {
}

func (cat Dog) color() string {
	fmt.Println("dog color black")
	return "black"
}

func tryInterface() {
	var animal Animal

	animal = new(Cat)
	fmt.Println(animal.color())

	animal = Dog{}
	fmt.Println(animal.color())

}
