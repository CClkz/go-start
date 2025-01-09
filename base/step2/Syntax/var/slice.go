package main

import (
	"fmt"
)

func declareSlice() {
	// 指针，引用类型
	// cap 容量，容量的作用是分配固定的空间
	// 切片append添加元素超过cap时，会扩容，并且扩容不是每次容量+1
	// 观察到的规律，初始容量一直*2来扩容，若初始空切片0 1 2 4 8这样来扩容
	// 不绝对，较大的切片（超过1024 ）可能就增加1/4
	numbers := []int{1, 2, 3, 4, 5} // 长度5，容量,5
	var emptySlice []int
	slice := make([]int, 3, 4) // 长度3，容量5

	printSlice(numbers)
	printSlice(emptySlice)
	fmt.Println(emptySlice == nil)
	printSlice(slice)

	fmt.Println("添加第一个元素")
	slice1 := append(emptySlice, 4)
	fmt.Println("添加元素后的原切片slice")
	printSlice(slice)
	fmt.Println("添加元素后赋值的新切片slice1")
	printSlice(slice1)
	fmt.Println("开始循环添加元素")
	for i := 5; i <= 15; i++ {
		slice1 = append(slice1, i)
		printSlice(slice1)
	}
}

func declareArray() {
	// 数组，值类型
	numbers := [...]int{1, 2, 3}
	var emptyArray [3]int

	printArray(numbers)
	printArray(emptyArray)
	fmt.Println(emptyArray == [3]int{0, 0, 0})
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}

func printArray(x [3]int) {
	// 数组len cap一直相等，很少会用到cap
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
