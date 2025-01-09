package main

import "fmt"

// 函数可返回多个值
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	// defer 延迟执行（当前函数执行完再执行），多个defer执行顺序类似栈，先入后出
	defer fmt.Println("defer1")
	defer fmt.Println("defer2")
	a, b := swap("Jim", "Green")
	fmt.Println(a, b)
	defer fmt.Println("defer3")
	defer fmt.Println("defer4")
}
