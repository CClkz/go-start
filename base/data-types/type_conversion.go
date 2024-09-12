package main

import (
	"fmt"
	"strconv"
)

func transeType() {
	// 原类型
	// 数值
	var n1 int = 1
	var f1 float32 = float32(n1)
	fmt.Println(f1)

	// 字符串
	var s2 string = "10"
	// 字符串特定的函数去转换，也好理解，String.xx()
	n2, err := strconv.Atoi(s2)
	fmt.Println(n2, err)

	// 接口
	var i interface{} = "Hello, World"
	str, ok := i.(string)
	if ok {
		fmt.Printf("'%s' is a string\n", str)
	} else {
		fmt.Println("conversion failed")
	}
}
