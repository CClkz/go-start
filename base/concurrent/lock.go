package main

import (
	"fmt"
	"sync"
	"sync/atomic"
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

// 读写锁
func write(count int) {
	rwlock.Lock() // 加写锁
	x = x + 1
	time.Sleep(10 * time.Millisecond) // 假设读操作耗时10毫秒
	rwlock.Unlock()                   // 解写锁
	fmt.Println("write count", count)
	wg.Done()
}

func read(count int) {
	rwlock.RLock()               // 加读锁
	time.Sleep(time.Millisecond) // 假设读操作耗时1毫秒
	rwlock.RUnlock()             // 解读锁
	fmt.Println("read count", count)
	wg.Done()
}

// 原子操作版加函数
func atomicAdd() {
	// 原子操作是更底层的、硬件级别的同步机制，效率高于互斥锁
	atomic.AddInt64(&x, 1)
	wg.Done()
}

func tryLock() {
	// wg.Add(2)

	// 基本操作
	// go add()
	// go add()

	// 加互斥锁
	// go addByMutex()
	// go addByMutex()

	// 加读写锁
	// for i := 0; i < 10; i++ {
	// 	wg.Add(1)
	// 	go write(i)
	// }
	// for i := 0; i < 10; i++ {
	// 	wg.Add(1)
	// 	go read(i)
	// }

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go atomicAdd()
	}

	wg.Wait()
	// 期望是10000，但输出结果不一定是10000
	fmt.Println(x)
}

// int 和 int64区别
// int 32微系统是32位，64系统是64位
// int64 确定为64位
// for i := 0; i < 10; i++ 此处i为int类型
// for i := int64(0); i < 10; i++ 此处为int64类型
// int64(i) 将一个数值转为int64,同理int(i)
