package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

// Product 定义了一个产品模型，用于映射数据库中的products表
// 嵌入gorm.Model获取基础字段：ID、CreatedAt、UpdatedAt、DeletedAt

func main() {
	// 配置MySQL连接信息
	// 格式：用户名:密码@协议(主机:端口)/数据库名?参数1&参数2
	dsn := "root:root@tcp(127.0.0.1:3306)/go?charset=utf8mb4&parseTime=True&loc=Local"

	// 连接数据库
	// mysql.Open(dsn) - 使用MySQL驱动打开连接
	// &gorm.Config{} - GORM配置，这里使用默认配置
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		// 连接失败时终止程序并输出错误
		panic(fmt.Sprintf("failed to connect database: %v", err))
	}
	// 清空测试数据（仅用于开发环境）
	// 先删除表，再重新创建，确保每次运行都是干净的环境
	db.Migrator().DropTable(&User{})
	db.AutoMigrate(&User{})

	user := User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}
	//result := db.Create(&user) // 通过数据的指针来创建
	//fmt.Println("result:", result)
	//result := db.Select("Name", "Age", "CreatedAt").Create(&user)
	result := db.Omit("Name", "Age", "CreatedAt").Create(&user)
	fmt.Println("result:", result)
}
