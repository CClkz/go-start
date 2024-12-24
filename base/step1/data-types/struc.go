package main

import (
	"fmt"
	"strconv"
)

type Point struct {
	X int
	Y int
}

// 重写String方法
// func (p Point) String() string {
// 	return fmt.Sprintf("Point{X:%d, Y:%d}", p.X, p.Y)
// }

func (p Point) String() string {
	// go支持string用 + 拼接，但不支持数值和string用 + 拼接
	return "Point{" + "X:" + strconv.Itoa(p.X) + ", Y:" + strconv.Itoa(p.Y) + "}"
}

func tryStruc() {
	point := Point{X: 1, Y: 2}
	fmt.Println(point)
}
