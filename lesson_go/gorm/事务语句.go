package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

// Account 对应 accounts 表
type Account struct {
	gorm.Model         // 嵌入GORM内置模型（包含ID、CreatedAt、UpdatedAt、DeletedAt）
	ID         uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	Balance    float64 `gorm:"type:decimal(10,2);not null" json:"balance"` // 账户余额，使用decimal类型避免精度丢失
}

// Transaction 对应 transactions 表
type Transaction struct {
	gorm.Model            // 嵌入GORM内置模型
	ID            uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	FromAccountID uint    `gorm:"not null" json:"from_account_id"`           // 转出账户ID
	ToAccountID   uint    `gorm:"not null" json:"to_account_id"`             // 转入账户ID
	Amount        float64 `gorm:"type:decimal(10,2);not null" json:"amount"` // 转账金额
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
	db.Migrator().DropTable(&Account{})
	db.AutoMigrate(&Account{})

	db.Migrator().DropTable(&Transaction{})
	db.AutoMigrate(&Transaction{})

	// 更安全的方式插入账户数据
	fromID, toID := uint(1), uint(2)
	transferAmount := float64(100)

	// 插入账户记录
	account1 := Account{ID: fromID, Balance: 100}
	account2 := Account{ID: toID, Balance: 200}
	db.Create(&account1)
	db.Create(&account2)
	fmt.Println("账户数据插入成功")

	// 开始事务
	tx := db.Begin()
	defer func() {
		// 最终确保事务被正确提交或回滚
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 1. 检查转出账户余额并扣除金额
	var fromAccount Account
	if err := tx.First(&fromAccount, fromID).Error; err != nil {
		tx.Rollback()
		fmt.Printf("查询转出账户失败: %v\n", err)
		return
	}

	if fromAccount.Balance < transferAmount {
		tx.Rollback()
		fmt.Println("转出账户余额不足，事务回滚")
		return
	}

	if err := tx.Model(&fromAccount).Update("Balance", fromAccount.Balance-transferAmount).Error; err != nil {
		tx.Rollback()
		fmt.Printf("扣除转出账户余额失败: %v\n", err)
		return
	}

	// 2. 向转入账户增加金额
	var toAccount Account
	if err := tx.First(&toAccount, toID).Error; err != nil {
		tx.Rollback()
		fmt.Printf("查询转入账户失败: %v\n", err)
		return
	}

	if err := tx.Model(&toAccount).Update("Balance", toAccount.Balance+transferAmount).Error; err != nil {
		tx.Rollback()
		fmt.Printf("增加转入账户余额失败: %v\n", err)
		return
	}

	// 3. 记录转账交易
	transaction := Transaction{
		FromAccountID: fromID,
		ToAccountID:   toID,
		Amount:        transferAmount,
	}
	if err := tx.Create(&transaction).Error; err != nil {
		tx.Rollback()
		fmt.Printf("记录转账交易失败: %v\n", err)
		return
	}

	// 提交事务
	tx.Commit()
	fmt.Println("转账事务执行成功")

	// 验证转账结果
	var afterFromAccount, afterToAccount Account
	tx.First(&afterFromAccount, fromID)
	tx.First(&afterToAccount, toID)
	fmt.Printf("转账后账户 %d 余额: %.2f\n", fromID, afterFromAccount.Balance)
	fmt.Printf("转账后账户 %d 余额: %.2f\n", toID, afterToAccount.Balance)
}
