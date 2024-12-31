package main

import (
	"fmt"
	"reflect"
)

func tryNil() {
	// go 数组和结构体是值类型，其它map、切片等js看起来是引用类型的均是引用类型
	// 引用类型的零值是nil，其余的看自己类型

	// 引用类型的属性，包含内部指针和该类型的一些属性
	// 以切片为例，var slice []int，声明未赋值，slice内部指针指向nil，但slice是切片类型也具备切片类型的属性len caps
	// 总结下，其余引用类型也是同理
	// 虽然 slice == nil 为 true，但 slice 并不等同于 nil：
	// nil 是 Go 语言中用来表示指针、接口、映射、切片、通道和函数类型零值的预定义标识符。
	// slice 是一个具体的切片变量，具有切片的类型信息和潜在的存储结构（即使当前处于零值状态）。
	// nil只是等于slice内部指针这个属性

	var str string
	var num int
	var num_f float32
	var arr [3]int
	var slice []int
	var map1 map[string]int
	var anytype interface{}

	fmt.Printf("str:%s\n", str)
	fmt.Printf("str:%d\n", num)
	fmt.Printf("str:%f\n", num_f)
	fmt.Printf("arr:%v\n", arr)
	fmt.Printf("slice:%v\n", slice) // []
	fmt.Printf("map:%v\n", map1)
	fmt.Printf("anytype:%v\n", anytype)

	var s []int
	fmt.Println(s)
	fmt.Println(s == nil) // 输出true

	// 判断变量类型
	slicetype := reflect.TypeOf(slice)
	fmt.Printf("slicetype:%s\n", slicetype)
}
