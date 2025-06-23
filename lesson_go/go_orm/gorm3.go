package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func main() {
	// 配置MySQL连接信息
	// 格式：用户名:密码@协议(主机:端口)/数据库名?参数1&参数2
	dsn := "root:root@tcp(127.0.0.1:3306)/go?charset=utf8mb4&parseTime=True&loc=Local"

	// 连接数据库
	// mysql.Open(dsn) - 使用MySQL驱动打开连接
	// &gorm.Config{} - GORM配置，这里使用默认配置
	//db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{
	//	QueryFields: true,
	//	注意 在 QueryFields 模式中, 所有的模型字段（model fields）都会被根据他们的名字选择。
	//})

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		// 连接失败时终止程序并输出错误
		panic(fmt.Sprintf("failed to connect database: %v", err))
	}
	// 清空测试数据（仅用于开发环境）
	// 先删除表，再重新创建，确保每次运行都是干净的环境
	db.Migrator().DropTable(&User{})
	db.AutoMigrate(&User{})

	user := User{Name: "NewUser", Age: 20, Birthday: time.Now()}
	user2 := User{Name: "NewUser2", Age: 220, Birthday: time.Now()}
	result := db.Create(&user)
	result2 := db.Create(&user2)
	fmt.Println("result:", result)
	fmt.Println("result:", result2)

	db.Where(&User{Name: "jinzhu", Age: 20}).First(&user)

	// 查询第一条记录
	//queryResult := db.First(&user)
	//db.Select("name", "age").Find(&users)
	// SELECT name, age FROM users;
	//db.Select([]string{"name", "age"}).Find(&users)
	var users []User
	queryResult := db.Find(&users)
	if queryResult.Error != nil {
		if queryResult.Error == gorm.ErrRecordNotFound {
			fmt.Println("No records found")
		} else {
			fmt.Println("Query error:", queryResult.Error)
		}
	} else {
		fmt.Println("Query success, rows affected:", queryResult.RowsAffected)
		// 打印查询到的所有用户信息
		for _, u := range users {
			fmt.Printf("User ID: %d, Name: %s, Age: %d, Birthday: %s\n", u.ID, u.Name, u.Age, u.Birthday.Format(time.RFC3339))
		}
	}

}
