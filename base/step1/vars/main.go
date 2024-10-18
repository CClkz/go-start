package main

import "fmt"

// 常量
const C1 string = "hello"

const (
	Unknown = 0
	Female  = 1
	Male    = 2
	// 常量的值编译时确定，函数执行在运行时，故不可用自定义的函数，必须是内置函数
	// Male    = getMale()
)

const (
	azero = 0
	a     = iota // 1，iota 在检测到const时置为0，const内部常量声明每增一行iota++(第一行是0)
	b            // 2,未赋值常量继承上一行的值
	c            // 3
	d     = "ha" // 独立值，iota += 1
	e            // "ha"   iota += 1
	f     = 100  // iota +=1
	g            // 100  iota +=1
	h     = iota // 8,恢复计数
	i            // 9
)

// 开头大写的函数才能被到处，给其他包使用
func GetMale() int {
	return 2
}

/**
 * 程序的入口（特殊地位的函数），执行程序，然后退出，不需要返回类型
 */
func main() {
	// 最基本的 var v_name v_type
	var a1 int = 1
	// 类型推断 var v_name = xxx
	var a2 = 2
	// 简短形式，:=声明赋值变量，等同于 var c int =  a1 + a2
	a3 := a1 + a2
	fmt.Println(a1, a2, a3)

	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)

	// 这和js不同，let {a,b,c} = {} 或 let a=1,b=2,c=3
	vname1, vname2, vname3 := 1, 2, 3
	fmt.Println(vname1, vname2, vname3)

	var map1 map[string]string = map[string]string{
		"a": "a1",
	}
	map2 := map1
	map1["a"] = "a2"

	fmt.Println(map1, map2)

	fmt.Println("常量", a, b, c, d, e, f, g, h, i)

}
