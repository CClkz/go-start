package main

import (
	"fmt"
)

func main() {
	// channel()
	selectChannel()
}

func channelType() {
	var ch1 chan int       // ch1是一个正常的channel，是双向的
	var ch2 chan<- float64 // ch2是单向channel，只用于写float64数据
	var ch3 <-chan int     // ch3是单向channel，只用于读int数据
	fmt.Println(ch1, ch2, ch3)
}

func channel() {
	c := make(chan int)

	go func() {
		defer fmt.Println("子go程结束")

		fmt.Println("子go程正在运行……")

		c <- 666 //666发送到c
	}()

	num := <-c //从c中接收数据，并赋值给num

	fmt.Println("num = ", num)
	fmt.Println("main go程结束")
}
func selectChannel() {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 6; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()

	fibonacci(c, quit)
}

func fibonacci(c, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}
