package main

import (
	"fmt"
	"strconv"
)

// 定义结构体类型
// 结构体是有顺序的
// 可见性规则：首字母大写的属性才可在其他包内访问到
type Point struct {
	X int
	Y int
	// PublicField int
}

// 重写String方法
// func (p Point) String() string {
// 	return fmt.Sprintf("Point{X:%d, Y:%d}", p.X, p.Y)
// }

// 定义结构体方法
// 结构体默认有String方法，此处是重写了
func (p Point) String() string {
	// go支持string用 + 拼接，但不支持数值和string用 + 拼接
	return "Point{" + "X:" + strconv.Itoa(p.X) + ", Y:" + strconv.Itoa(p.Y) + "}"
}

func (p Point) GetInfo() string {
	return fmt.Sprintf("X:%d, Y:%d", p.X, p.Y)
}

func tryStruc() {
	// 实例化结构体
	// 1.指定字段名
	point := Point{X: 1, Y: 2}
	// 2.不指定字段名，按顺序
	point1 := Point{3, 4}
	// 3.new(type) 注意Point是type，new(type)这在ts里是不能的
	point2 := new(Point)
	point2.X = 5
	point2.Y = 6
	// 4.var name type
	var point3 Point

	fmt.Println(point)
	fmt.Println(point1)
	fmt.Println(point2)
	fmt.Println(point3)

	// 访问结构体字段
	fmt.Println(point2.X)
	// 访问结构体方法
	fmt.Println(point2.GetInfo())

}
