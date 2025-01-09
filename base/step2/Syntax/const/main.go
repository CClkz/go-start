package main

import "fmt"

// 常量只能是基础类型
// 常量赋值不能是自定义函数（内置函数可以），即使函数返回基础类型，因为常量在编译时就确定了，不会涉及到函数执行
const (
	C1        = 1
	C2 string = "abc"
)

const (
	Apple, Banana     = iota + 1, iota + 3 // iota在const内部第一行是0，每行加1
	Cherimoya, Durian                      // 等同于Cherimoya, Durian = iota + 1, iota + 3
)

func main() {
	fmt.Println("const C1 C2", C1, C2)
	fmt.Println(Apple, Banana, Cherimoya, Durian)
}
