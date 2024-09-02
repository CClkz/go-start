package main

import "fmt"

func tryPointer() {
	var a int = 20   /* 声明实际变量 */
	var ip *int = &a /* 声明指针变量 */

	// ip := &a /* 指针变量的存储地址 */

	fmt.Printf("a 变量的地址是: %x\n", &a)

	/* 指针变量的存储地址 */
	fmt.Printf("ip 变量储存的指针地址: %x\n", ip)

	/* 使用指针访问值 */
	fmt.Printf("*ip 变量的值: %d\n", *ip)

	fmt.Printf("ip 变量的地址: %x\n", &ip)
}

// 变量存在内存地址，变量a内存地址赋值给指针变量p
// 变量本身没有指针概念，不是a存在指针p，是存在指针p指向变量a
