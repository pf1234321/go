package main

import (
	"fmt"
	"sync"
	"time"
)

// Task 表示一个可执行的任务
type Task struct {
	ID       int
	Name     string
	Execute  func() // 任务执行函数
	Duration time.Duration
}

// TaskScheduler 任务调度器
type TaskScheduler struct {
	wg      sync.WaitGroup
	tasks   []*Task
	results chan *Task
}

// NewTaskScheduler 创建新的任务调度器
func NewTaskScheduler() *TaskScheduler {
	return &TaskScheduler{
		results: make(chan *Task, 100), // 缓冲通道，避免阻塞
	}
}

// AddTask 添加任务到调度器
func (s *TaskScheduler) AddTask(task *Task) {
	s.tasks = append(s.tasks, task)
}

// Run 并发执行所有任务
func (s *TaskScheduler) Run() {
	// 启动工作协程执行任务
	for _, task := range s.tasks {
		s.wg.Add(1)
		go func(t *Task) {
			defer s.wg.Done()

			start := time.Now()
			t.Execute()
			t.Duration = time.Since(start)

			s.results <- t // 将结果发送到通道
		}(task)
	}

	// 等待所有任务完成后关闭结果通道
	go func() {
		s.wg.Wait()
		close(s.results)
	}()
}

// PrintResults 打印所有任务的执行时间
func (s *TaskScheduler) PrintResults() {
	for task := range s.results {
		fmt.Printf("任务 %d (%s) 执行时间: %v\n",
			task.ID, task.Name, task.Duration)
	}
}

func main() {
	scheduler := NewTaskScheduler()

	// 添加多个模拟任务
	scheduler.AddTask(&Task{
		ID:   1,
		Name: "数据库查询",
		Execute: func() {
			time.Sleep(500 * time.Millisecond) // 模拟耗时操作
		},
	})

	scheduler.AddTask(&Task{
		ID:   2,
		Name: "文件读取",
		Execute: func() {
			time.Sleep(300 * time.Millisecond)
		},
	})

	scheduler.AddTask(&Task{
		ID:   3,
		Name: "网络请求",
		Execute: func() {
			time.Sleep(700 * time.Millisecond)
		},
	})

	// 启动任务调度
	scheduler.Run()

	// 打印结果
	scheduler.PrintResults()
}
