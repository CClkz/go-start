package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 把 sum 发送到通道 c
}

func tryChannel() {
	s := []int{7, 2, 8, 9, 4, 0}

	// 创建两个独立的通道
	c1 := make(chan int)
	c2 := make(chan int)
	c3 := make(chan interface{})

	// 启动两个 goroutine 分别计算数组的前半部分和后半部分
	go sum(s[:len(s)/2], c1)
	go sum(s[len(s)/2:], c2)

	// 从两个通道中接收结果，程序会阻塞，等待值过来
	x, y := <-c1, <-c2

	// 打印结果
	fmt.Println(x, y, x+y)

	// 关闭通道（可选，但在实际应用中是个好习惯）
	close(c1)
	close(c2)
	close(c3)
}

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	// for select，一直阻塞，等待条件满足，条件下的return退出for select
	// case 条件为通道发送数据或接收数据
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

func trySelect() {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}
