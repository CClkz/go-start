package main

import "fmt"

func tryArray() {
	// 数组需要指定长度或者[...]自动推断长度，[]就是切片了
	arr := [...]int{1, 2, 3}
	arr[1] = 100
	fmt.Println(arr)

	// 切片
	slice := []int{1, 2, 3}
	fmt.Println(slice)

}
