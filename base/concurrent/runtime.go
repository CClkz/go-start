package main

import (
	"fmt"
	"runtime"
	"sync"
)

func tryRuntime() {
	go func(s string) {
		for i := 0; i < 20; i++ {
			fmt.Println(s)
		}
	}("world")
	// 主协程
	for i := 0; i < 2; i++ {
		// 切一下，再次分配任务
		// 让当前协程暂停，其他等待的协程运行
		// 观察现象，不一定等子协程全部执行完，甚至字写成一次打印都无
		// 但相比不加runtime.Gosched()，确实提高了world打印的次数
		// 故想执行完子协程，该实现方法不好
		runtime.Gosched()
		fmt.Println("hello")
	}
}

// 优化下
func tryRuntime2() {
	// 等待组
	var wg sync.WaitGroup
	wg.Add(2) // 增加一个等待的 goroutine

	go func(s string) {
		defer wg.Done() // 在 goroutine 结束时减少计数器
		for i := 0; i < 10; i++ {
			fmt.Println(s)
		}
	}("world")

	go func(s string) {
		defer wg.Done() // 在 goroutine 结束时减少计数器
		for i := 0; i < 10; i++ {
			fmt.Println(s)
		}
	}("!")

	// 主协程
	for i := 0; i < 2; i++ {
		fmt.Println("hello")
		// 这里通常不需要 runtime.Gosched()，因为它对控制流的影响很小且不可预测
	}

	wg.Wait() // 等待所有增加的等待计数变为 0
}
