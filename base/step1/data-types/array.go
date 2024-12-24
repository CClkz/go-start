package main

import "fmt"

func tryArray() {
	// 数组需要指定长度或者[...]自动推断长度，[]就是切片了

	var arr1 [3]int = [3]int{1, 2, 3}
	fmt.Println(arr1)

	arr2 := [...]int{1, 2, 3}
	arr2[1] = 100
	fmt.Println(arr2)

	// 注意，go和ts是不同的，声明的时候也相当于赋值了零值
	// var arr3 [3]int可看作为 var arr3 [3]int = [3]int{0,0,0}
	// 所以arr3有数组的属性和方法
	var arr3 [3]int
	fmt.Println(len(arr3))

	// 切片
	slice := []int{1, 2, 3}
	fmt.Println(slice)

}
