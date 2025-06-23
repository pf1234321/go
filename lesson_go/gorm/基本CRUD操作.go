package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model        // 嵌入GORM内置模型，包含ID、时间戳和软删除字段
	ID         uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name       string `gorm:"type:varchar(255);not null" json:"name"`
	Age        int    `gorm:"type:int;not null" json:"age"`
	Grade      string `gorm:"type:varchar(50)" json:"grade"`
}

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/go?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=True&loc=Local"
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		//Logger: gorm.Logger.Default.LogMode(gorm.Info), // 启用Info级别日志
	})

	if err != nil {
		panic(fmt.Sprintf("failed to connect database: %v", err))
	}

	// 清空并重建表
	db.Migrator().DropTable(&Student{})
	db.AutoMigrate(&Student{})

	// 插入学生记录
	student := Student{Name: "张三", Age: 20, Grade: "三年级"}
	result := db.Create(&student)
	if result.Error != nil {
		fmt.Printf("插入失败: %v\n", result.Error)
		return
	}
	fmt.Println("插入成功")
	//========================================================================================
	// 查询年龄大于18岁的学生
	var students []Student
	query := db.Where("age > ?", 18).Find(&students) // 注意：查询结果应存入切片

	// 打印查询结果
	fmt.Println("查询结果:")
	fmt.Printf("SQL语句: %v\n", query.Statement.SQL.String())
	fmt.Printf("查询结果数量: %d\n", query.RowsAffected)

	if query.RowsAffected > 0 {
		for i, s := range students {
			fmt.Printf("学生 %d:\n", i+1)
			fmt.Printf("  ID: %d\n", s.ID)
			fmt.Printf("  姓名: %s\n", s.Name)
			fmt.Printf("  年龄: %d\n", s.Age)
			fmt.Printf("  年级: %s\n", s.Grade)
			fmt.Printf("  创建时间: %v\n", s.CreatedAt)
			fmt.Printf("  更新时间: %v\n", s.UpdatedAt)
		}
	} else {
		fmt.Println("未找到符合条件的学生记录")
	}
	//================================编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。========================================================
	// 更新操作
	update := db.Model(&Student{}).
		Where("name = ?", "张三").
		Update("grade", "四年级")

	if update.Error != nil {
		fmt.Printf("更新失败: %v\n", update.Error)
		return
	}

	fmt.Printf("更新成功，影响行数: %d\n", update.RowsAffected)

	//	编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。db.Where("email LIKE ?", "%jinzhu%").Delete(&Email{})
	delete := db.Where("age < ?", "15").Delete(&Student{})
	if delete.Error != nil {
		fmt.Printf("delete失败: %v\n", delete.Error)
		return
	}

	fmt.Printf("delete成功，影响行数: %d\n", delete.RowsAffected)

}
