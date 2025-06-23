package main

import "fmt"

// Person 定义人的基本信息结构体
type Person struct {
	Name string // 姓名
	Age  int    // 年龄
}

// Employee 定义员工结构体，通过组合嵌入Person
type Employee struct {
	EmployeeID int // 员工ID
	Person         // 组合Person结构体
}

// PrintInfo 打印员工的完整信息
func (e *Employee) PrintInfo() {
	fmt.Printf("员工ID: %d, 姓名: %s, 年龄: %d\n",
		e.EmployeeID, e.Name, e.Age)
}

func main() {
	// 创建员工实例并初始化
	emp := &Employee{
		Person: Person{
			Name: "Tom",
			Age:  30,
		},
		EmployeeID: 1001,
	}

	// 调用方法打印信息
	emp.PrintInfo()
}
