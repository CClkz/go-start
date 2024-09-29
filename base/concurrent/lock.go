package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	x      int64
	wg     sync.WaitGroup
	lock   sync.Mutex
	rwlock sync.RWMutex
)

func add() {
	for i := 0; i < 5000; i++ {
		x = x + 1
	}
	wg.Done()
}

// 互斥锁Mutex
func addByMutex() {
	for i := 0; i < 5000; i++ {
		lock.Lock()
		x = x + 1
		lock.Unlock()
	}
	wg.Done()
}

func write(count int64) {
	rwlock.Lock() // 加写锁
	x = x + 1
	time.Sleep(10 * time.Millisecond) // 假设读操作耗时10毫秒
	rwlock.Unlock()                   // 解写锁
	fmt.Println("write count", count)
	wg.Done()
}

func read(count int64) {
	rwlock.RLock()               // 加读锁
	time.Sleep(time.Millisecond) // 假设读操作耗时1毫秒
	rwlock.RUnlock()             // 解读锁
	fmt.Println("read count", count)
	wg.Done()
}

func tryLock() {
	// wg.Add(2)
	// go add()
	// go add()
	// go addByMutex()
	// go addByMutex()

	// 当读的操作很多时，用读写锁，让读操作可并发运行
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write(i)
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go read(i)
	}

	wg.Wait()
	// 期望是10000，但输出结果不一定是10000
	fmt.Println(x)
}
