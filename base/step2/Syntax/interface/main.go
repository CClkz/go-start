package main

import "fmt"

func main() {
	typeAssert(123)
}

func typeAssert(a interface{}) string {
	// go所有类型都实现了interface{}，此处参数被转成了interface{}类型
	// 类型断言 a.(T)
	value, ok := a.(string)
	if !ok {
		// 断言失败，value会赋值断言类型的零值，此处为空字符串
		fmt.Println("It is not ok for type string")
		return ""
	}
	fmt.Println("The value is ", value)

	return value
}

func typeSwitch(a interface{}) {
	switch t := a.(type) { // 类型开关（type switch），也是类型断言，a.(type)只可在switch使用，只返回类型一个参数
	default:
		fmt.Printf("unexpected type %T", t) // %T prints whatever type t has
	case bool:
		fmt.Printf("boolean %t\n", t) // t has type bool
	case int:
		fmt.Printf("integer %d\n", t) // t has type int
	case *bool:
		fmt.Printf("pointer to boolean %t\n", *t) // t has type *bool
	case *int:
		fmt.Printf("pointer to integer %d\n", *t) // t has type *int
	}
}
