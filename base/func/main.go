package main

import "fmt"

func main() {
	fmt.Println("main")

	a, b := swap("a", "b")
	fmt.Println(a, b)

}

/**
 * go 函数特点
 * 1. 可返回多个值
 */
func swap(a string, b string) (string, string) {
	return b, a
}
