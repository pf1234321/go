package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Product 定义了一个产品模型，用于映射数据库中的products表
// 嵌入gorm.Model获取基础字段：ID、CreatedAt、UpdatedAt、DeletedAt
type Product struct {
	gorm.Model        // 嵌入GORM内置模型，包含ID、时间戳和软删除字段
	Code       string // 产品编码，对应数据库中的code字段
	Price      uint   // 产品价格，对应数据库中的price字段(无符号整数)
}

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

	// 自动迁移Schema
	// 根据Product结构体定义创建或更新数据库表结构
	// 注意：只会添加新字段，不会删除或修改已有字段
	if err := db.AutoMigrate(&Product{}); err != nil {
		panic(fmt.Sprintf("failed to migrate schema: %v", err))
	}

	// 清空测试数据（仅用于开发环境）
	// 先删除表，再重新创建，确保每次运行都是干净的环境
	db.Migrator().DropTable(&Product{})
	db.AutoMigrate(&Product{})

	// 【创建操作】插入一条产品记录
	// db.Create() - 插入记录到数据库
	// result包含操作结果：影响行数、错误信息等
	result := db.Create(&Product{Code: "D42", Price: 100})
	if result.Error != nil {
		panic(fmt.Sprintf("failed to create product: %v", result.Error))
	}
	// 获取插入记录的自增ID（GORM自动设置到结构体的ID字段）
	fmt.Printf("Created product with ID: %d\n", result.Statement)

	// 【查询操作】按ID查找记录
	// db.First() - 根据主键查询第一条记录
	// &product - 查询结果将填充到这个对象
	var product Product
	if err := db.First(&product, 1).Error; err != nil {
		panic(fmt.Sprintf("failed to find product by ID: %v", err))
	}
	// 输出查询结果（%+v会显示字段名和值）
	fmt.Printf("Found product by ID: %+v\n", product)

	// 【查询操作】按条件查找记录
	// db.First() - 带条件的查询，使用问号占位符防止SQL注入
	var productByCode Product
	if err := db.First(&productByCode, "code = ?", "D42").Error; err != nil {
		panic(fmt.Sprintf("failed to find product by Code: %v", err))
	}
	fmt.Printf("Found product by Code: %+v\n", productByCode)

	// 【更新操作】更新单个字段
	// db.Model() - 指定要更新的对象
	// Update() - 更新单个字段，自动更新UpdatedAt时间戳
	if err := db.Model(&product).Update("Price", 200).Error; err != nil {
		panic(fmt.Sprintf("failed to update product price: %v", err))
	}
	fmt.Println("Updated product price to 200")

	// 验证更新结果
	var updatedProduct Product
	db.First(&updatedProduct, product.ID)
	fmt.Printf("After update: %+v\n", updatedProduct)

	// 【更新操作】更新多个字段
	// 方式1：使用结构体更新（注意：零值字段不会更新）
	db.Model(&product).Updates(Product{Price: 310, Code: "F42"})

	// 方式2：使用map更新（所有字段都会更新）
	// db.Model(&product).Updates(map[string]interface{}{"Price": 300, "Code": "F42"})

	// 【删除操作】软删除记录
	// 软删除：不会真正删除记录，而是设置DeletedAt字段
	// 被软删除的记录在查询时会自动过滤
	if err := db.Delete(&product).Error; err != nil {
		panic(fmt.Sprintf("failed to delete product: %v", err))
	}
	fmt.Printf("Soft-deleted product with ID: %d\n", product.ID)

	// 验证删除结果（软删除的记录查询不到）
	var deletedProduct Product
	if err := db.First(&deletedProduct, product.ID).Error; err == nil {
		fmt.Println("Error: Product still exists after soft-delete")
	} else {
		fmt.Println("Verification: Product not found after soft-delete")
	}

	// 【强制删除】物理删除记录（慎用！）
	// db.Unscoped().Delete(&product)
}
