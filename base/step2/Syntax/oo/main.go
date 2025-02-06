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
	// 一
	fmt.Println(s)
	fmt.Println(&s)
	// 二
	changeStr(&s)
	fmt.Println(s)
	fmt.Println(&s)
	// 三
	changeStr2(s)
	fmt.Println(s)
	fmt.Println(&s)
	// 四，二和四殊途同归，都是创建新字符串，s再存储新字符串的地址，原字符串还在内存里
	s = "def"
	fmt.Println(s)
	fmt.Println(&s)

	// 代码运行分析：
	// 1. abc 分配空间 0x111，s分配空间 0x222, 0x222 存储的值为 0x111
	// 2. str changed 分配空间 0x333，0x222 存值变为 0x333
	// 3. 值类型传参不用指针，生成副本，不影响原 s
	// 4. def 分配空间 0x444，0x222 存值变为 0x444

	// 只读数据段 - 主要存储常量(const声明的常量和函数里的字符串)，程序编译时就创建了
	// 垃圾回收 - 不处理只读数据段

	// abc，不可访问了，又不会被垃圾回收，造成内存泄漏怎么办：
	// 1. 常量的共享特性，一样的字符串不会重复创建，例如 s1:= "abc"，s2 = "abc"，s1 和 s2 存储的是同一个地址
	// 2. 只读数据段大小相对固定

}

func changeStr(str *string) {
	*str = "str changed"
}
func changeStr2(str string) {
	str = "str2"
}
