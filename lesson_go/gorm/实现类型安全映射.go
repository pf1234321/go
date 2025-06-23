package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

// Book 对应 books 表的结构体
type Book struct {
	gorm.Model         // 嵌入GORM内置模型，包含ID、时间戳和软删除字段
	ID         uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	Title      string  `gorm:"type:varchar(255);not null" json:"title"`
	Author     string  `gorm:"type:varchar(100);not null" json:"author"`
	Price      float64 `gorm:"type:decimal(10,2);not null" json:"price"`
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
	db.Migrator().DropTable(&Book{})
	db.AutoMigrate(&Book{})

	book := []Book{
		{Title: "Go语言编程1", Author: "张三1", Price: 159.50},
		{Title: "Go语言编程2", Author: "张三2", Price: 59.50},
		{Title: "Go语言编程3", Author: "张三3", Price: 9.50},
	}
	createResult := db.Create(&book)
	if createResult.Error != nil {
		fmt.Printf("create失败: %v\n", createResult.Error)
		return
	} else {
		fmt.Printf("create  success")
	}
	//编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。
	var books []Book
	result := db.Raw("SELECT ID,Title, Author, Price FROM books WHERE Price > ?", 50).Scan(&books)
	if result.Error != nil {
		fmt.Printf("查询失败: %v\n", result.Error)
		return
	}

	// 检查查询结果
	if len(books) > 0 {
		fmt.Printf("共查询到 %d 名技术部员工:\n", len(books))
		for _, emp := range books {
			fmt.Printf("ID: %d, Title: %s, Author: %s, Price: %.2f\n",
				emp.ID, emp.Title, emp.Author, emp.Price)
		}
	} else {
		fmt.Println("未找到")
	}

}
