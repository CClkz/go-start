package main

import "fmt"

// 定义一个名为 Person 的结构体
type Person struct {
	Name   string
	Age    int
	Gender string
}

func main() {
	var a int = 1
	var b float32 = 1
	var c bool = true
	var d string = "abc"
	var e [2]int = [2]int{1, 2}
	var f Person = Person{
		Name:   "1",
		Age:    12,
		Gender: "male",
	}
	var g map[string]int = map[string]int{
		"Alice": 25,
		"Bob":   30,
	}
	g["add"] = 1
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f.Name)
	fmt.Println(g["add"])
}
