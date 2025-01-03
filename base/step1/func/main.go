package main

import "fmt"

func main() {
	fmt.Println("main")

	a, b, err := swap("a", "b")
	c, d, _ := swap("c", "d") // 返回三个参数，指向接收两个，不接收的用 _

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(a, b)
	}
	fmt.Println(c, d)

}

/**
 * go 函数特点
 * 1. 可返回多个值
 */
// func swap(a string, b string) (string,  string) {
func swap(a string, b string) (res1 string, res2 string, err error) {
	if a == b {
		err = fmt.Errorf("输入的两个字符串相等，无法交换：%s", a)
		return // 函数的隐式返回，返回(res1,res2,err),未赋值，那就是("","",err)
		// 比较，ts可以约束函数返回类型，go既能约束类型还能定义参数直接函数里使用了
	} else {
		return b, a, nil
	}
}
