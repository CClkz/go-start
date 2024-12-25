package main

import "fmt"

func tryMap() {
	// 创建map实例
	m := map[string]int{
		"a": 1,
		"b": 2,
	}

	// 访问map字段
	m["c"] = 3

	fmt.Println(m)

	// value类型可以不同吗
	// 可以，用interface{}，map[string]interface{}{}
}
