package main

//题目 ：编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
//考察点 ： go 关键字的使用、协程的并发执行。

import (
	"fmt"
	"strconv"
	"sync"
)

func worker(id int, wg *sync.WaitGroup, s [10]int) {
	defer wg.Done() // Goroutine 完成时调用 Done()
	fmt.Printf("Worker %d started\n", id)
	for i := 0; i < 10; i++ {
		if s[i]%2 == 0 && id == 2 {
			fmt.Printf(strconv.Itoa((s[i])))
		}
		if s[i]%2 != 0 && id == 1 {
			fmt.Printf(strconv.Itoa((s[i])))
		}
	}
	fmt.Printf("Worker %d finished\n", id)
}

func main() {
	var wg sync.WaitGroup
	var s [10]int /* n 是一个长度为 10 的数组 */
	for i := 0; i < 10; i++ {
		s[i] = i
	}
	//fmt.Print(s)
	for i := 1; i <= 2; i++ {
		wg.Add(1) // 增加计数器
		go worker(i, &wg, s)
	}

	wg.Wait() // 等待所有 Goroutine 完成
	fmt.Println("All workers done")
}
