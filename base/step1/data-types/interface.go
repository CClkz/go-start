package main

import "fmt"

// go ts 的interface，二者设计理念不同
// 第一点
// go interface 只约束方法，ts interface 约束属性和方法
// 第二点，
// 接口实现方式，go是【隐式】【接口实现】，ts是【显式】【接口实现】
// go 结构体实现接口，但结构体定义的时候不受到interface约束，这个过程看不到interface的存在，调动函数、传参是才会校验类型
// ts 类实现接口，class implements interface，class定义的时候就收到interface的约束

// 定义接口
type Drawer interface {
	Draw()
}

// 实现接口，用结构体
type Circle struct {
	Radius float64
}

func (c Circle) Draw() {
	// 这里是绘制圆形的具体实现逻辑，例如打印一些表示圆形的信息
	fmt.Println("Drawing a circle with radius", c.Radius)
}

// 接口用途：
// 1 - 多态
func DrawShapes(ds ...Drawer) { // 可接受Draw类型的参数，可以是Circle、Square、Triangle等
	for _, d := range ds {
		d.Draw()
	}
}

// 2 - 解耦与模块化
type DataStorer interface {
	Store(data interface{}) error
	Retrieve(id interface{}) (interface{}, error)
}

// 类型断言（也包含了类型转换）
// 就是断言接口变量储存的实例是不是某个具体类型
func ProcessShape(shape Drawer) {
	if c, ok := shape.(Circle); ok {
		// 如果是Circle类型，可以进行一些特定于Circle的操作，如获取半径
		fmt.Println("It's a circle with radius", c.Radius)
	}
}

func tryInterface() {

	circle := Circle{Radius: 10}
	ProcessShape(circle)

	DrawShapes(circle)

}

// shape.(Circle) 什么意思
// v,ok := interfaceVariable.(ConcreteType) 是类型断言的语法
// interfaceVariable - 接口变量，存储实现了这个接口的实例
// shape可能是Circle的实例，也可能是其他实现了接口Drawer的结构体（例如Tri）的实例
// 若断言成功，即shape是Circle的实例
// 第一个参数返回shape储存的实例转换后为Circle类型的值，ok true
// 【shape确实已经是Circle类型了，但还是会转换下】，因为shape只是Drawer类型，直接用会导致v也是Drawer类型而不是更具体的Circle类型
// go的类型涉及的事情比ts要多，不能用ts的类型思维来理解
// 若断言失败
// 第一个参数返回Circle的零值，第二个参数返回false
