package main

import "fmt"

func main() {
	// testStruct()
	testStr()
}

// 定义一个结构体
type T struct {
	name string
}

func (t T) method1() {
	// 结构体是值类型，传参和赋值时，不约束指针型就生成副本
	t.name = "new name1"
}

func (t *T) method2() {
	t.name = "new name2"
}

func testStruct() {
	t := T{"old name"}

	fmt.Println("method1 调用前 ", t.name)
	t.method1()
	fmt.Println("method1 调用后 ", t.name)
	fmt.Println("method2 调用前 ", t.name)
	t.method2()
	fmt.Println("method2 调用后 ", t.name)
}

func testStr() {
	s := "abc"
	fmt.Println(s)
	changeStr(&s)
	fmt.Println(s)

}

func changeStr(str *string) {
	*str = "str changed"
}
