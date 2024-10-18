package main

import (
	"fmt"
	"sync"
	"time"
)

// GoroutinePool 是一个简单的goroutine池
type GoroutinePool struct {
	wg         sync.WaitGroup
	jobs       chan func()
	shutdown   chan bool
	maxWorkers int
}

// NewGoroutinePool 创建一个新的goroutine池
func NewGoroutinePool(maxWorkers int) *GoroutinePool {
	return &GoroutinePool{
		jobs:       make(chan func()),
		shutdown:   make(chan bool),
		maxWorkers: maxWorkers,
	}
}

// Start 启动goroutine池
func (p *GoroutinePool) Start() {
	for i := 0; i < p.maxWorkers; i++ {
		p.wg.Add(1)
		go func() {
			defer p.wg.Done()
			for {
				select {
				case job := <-p.jobs:
					job()
				case <-p.shutdown:
					return
				}
			}
		}()
	}
}

// Stop 停止goroutine池
func (p *GoroutinePool) Stop() {
	close(p.shutdown)
	p.wg.Wait()
}

// Submit 提交一个任务到goroutine池
func (p *GoroutinePool) Submit(job func()) {
	p.jobs <- job
}

func tryPool1() {
	// 创建一个最大工作goroutine数为5的池
	pool := NewGoroutinePool(5)
	pool.Start()

	// 提交一些任务
	for i := 0; i < 10; i++ {
		idx := i
		pool.Submit(func() {
			fmt.Printf("Processing job %d\n", idx)
			time.Sleep(time.Second) // 模拟耗时任务
		})
	}

	// 等待一段时间，以便观察输出
	time.Sleep(5 * time.Second)

	// 停止goroutine池
	pool.Stop()
	fmt.Println("Goroutine pool stopped.")
}
