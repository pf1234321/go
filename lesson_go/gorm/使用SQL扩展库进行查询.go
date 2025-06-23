package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

//假设你已经使用Sqlx连接到一个数据库，并且有一个 employees 表，包含字段 id 、 name 、 department 、 salary 。
//要求 ：
//编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
//编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。

type Employee struct {
	gorm.Model         // 嵌入GORM内置模型，包含ID、时间戳和软删除字段
	ID         uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name       string  `gorm:"type:varchar(255);not null" json:"name"`
	Department string  `gorm:"type:varchar(100)" json:"department"`
	Salary     float64 `gorm:"type:decimal(10,2);not null" json:"salary"`
}

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/go?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=True&loc=Local"

	// 配置日志，方便查看SQL执行情况
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n[GORM] ", log.LstdFlags),
		logger.Config{
			LogLevel: logger.Info,
		},
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(fmt.Sprintf("failed to connect database: %v", err))
	}

	// 清空并重建表
	db.Migrator().DropTable(&Employee{})
	db.AutoMigrate(&Employee{})

	// 创建员工
	employeesToCreate := []Employee{
		{Name: "张三", Department: "技术部", Salary: 10000.00},
		{Name: "李四", Department: "技术部", Salary: 12000.00},
		{Name: "王五", Department: "市场部", Salary: 8000.00},
	}

	createResult := db.Create(employeesToCreate)
	if createResult.Error != nil {
		fmt.Printf("create失败: %v\n", createResult.Error)
		return
	} else {
		fmt.Printf("create  success")
	}

	//
	var employees []Employee
	result := db.Raw("SELECT id, name, department,Salary FROM employees WHERE department = ?", "技术部").Scan(&employees)
	if result.Error != nil {
		fmt.Printf("查询失败: %v\n", result.Error)
		return
	}

	// 检查查询结果
	if len(employees) > 0 {
		fmt.Printf("共查询到 %d 名技术部员工:\n", len(employees))
		for _, emp := range employees {
			fmt.Printf("ID: %d, 姓名: %s, 部门: %s, 工资: %.2f\n",
				emp.ID, emp.Name, emp.Department, emp.Salary)
		}
	} else {
		fmt.Println("未找到技术部员工")
	}

	// 编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。
	var maxSalary float64 = 0
	maxSalaryResult := db.Raw("SELECT max(Salary) FROM employees").Scan(&maxSalary)
	if maxSalaryResult.Error != nil {
		fmt.Printf("查询maxSalaryResult失败: %v\n", maxSalaryResult.Error)
		return
	} else {
		fmt.Printf("最高工资: %.2f\n", maxSalary)
	}

}
