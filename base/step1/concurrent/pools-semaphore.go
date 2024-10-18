package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/semaphore"
)

func trySemaphore() {
	// 创建一个初始权重为3的信号量
	sem := semaphore.NewWeighted(3)

	// 启动5个goroutine，但只允许3个同时执行
	for i := 0; i < 5; i++ {
		go func(id int) {
			// 尝试获取信号量
			if err := sem.Acquire(context.Background(), 1); err == nil {
				defer sem.Release(1)
				fmt.Printf("Goroutine %d is running\n", id)
				time.Sleep(1 * time.Second) // 模拟耗时操作
				fmt.Printf("Goroutine %d is done\n", id)
			} else {
				fmt.Printf("Goroutine %d failed to acquire semaphore: %v\n", id, err)
			}
		}(i)
	}

	// 等待一段时间以观察输出，确保所有goroutine都有机会执行
	time.Sleep(6 * time.Second)
}
