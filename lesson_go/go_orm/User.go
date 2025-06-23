package main

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model           // 嵌入GORM内置模型，包含ID、时间戳和软删除字段
	Name       string    // 产品编码，对应数据库中的code字段
	Age        uint      // 产品价格，对应数据库中的price字段(无符号整数)
	Birthday   time.Time `gorm:"autoUpdateTime"`
}
