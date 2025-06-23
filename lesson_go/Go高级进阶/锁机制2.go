package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var counter int64

func worker2(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Goroutine 完成时调用 Done()
	fmt.Printf("Worker %d started\n", id)

	for i := 1; i <= 1000; i++ {
		atomic.AddInt64(&counter, 1)
	}
	fmt.Printf("Worker %d finished\n", id)
}

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 10; i++ {
		wg.Add(1) // 增加计数器
		go worker2(i, &wg)
	}
	wg.Wait() // 等待所有 Goroutine 完成
	fmt.Println("All workers done")
	fmt.Println(counter)
}
