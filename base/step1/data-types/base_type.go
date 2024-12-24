package main

import "fmt"

func tryBase() {
	var a int = 1
	var b float32 = 1
	var c bool = true
	var d string = "abc"

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	// 短变量声明
	e := "abc"
	fmt.Println(e)

	// go var和js var
	// go var是块级作用域且没有变量提升，不像js var，更接近js let
}
