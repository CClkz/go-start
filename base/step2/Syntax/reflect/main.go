package main

import (
	"fmt"
	"reflect"
)

func main() {
	var i interface{}
	fmt.Println(reflect.TypeOf(i), reflect.ValueOf(i))
	i = 100
	fmt.Println(reflect.TypeOf(i), reflect.ValueOf(i))
	i = "hello"
	fmt.Println(reflect.TypeOf(i), reflect.ValueOf(i))

	value := reflect.ValueOf(i)
	convertValue := value.Interface().(string) // 已知类型，强制转换
	fmt.Println(convertValue)
	fmt.Println(reflect.TypeOf(convertValue), reflect.ValueOf(convertValue))

	fmt.Println("reflect设置值----------------------")
	strPtrValue := reflect.ValueOf(&i)
	strValue := strPtrValue.Elem()
	fmt.Println(strValue, strValue.Elem())
	if strValue.CanSet() {
		strValue.Set(reflect.ValueOf("world"))
		fmt.Println(i)
	} else {
		fmt.Println("The value is not settable.")
	}
}

// reflect 运行时进行反射操作，其核心功能就是获取变量的类型信息和值信息
// interface{} 是所有类型的超集，包含两个指针：1. 值 2. 类型，注意其他类型不包含两个指针
// 传参 interface{} 时，要经过断言才可使用，赋值后的话不需要断言了

// var i interface{}
// i = "hello"
// strPtrValue := reflect.ValueOf(&i)      // 获取 i 的指针的反射值
// strValue := strPtrValue.Elem()          // 获取 i 的反射值
// innerValue := strValue.Elem()           // 获取 i 中存储的动态值 "hello" 的反射值

// 为什么不直接reflect.ValueOf(i)？因为获取的是i副本的反射值
