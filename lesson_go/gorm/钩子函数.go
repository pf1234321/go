package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"math/rand"
	"os"
	"time"
)

// User 用户模型，添加文章数量统计字段
type User struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	Email     string `gorm:"unique"`
	PostCount int    `gorm:"default:0"` // 文章数量统计
	Posts     []Post `gorm:"foreignKey:UserID"`
}

// Post 文章模型，添加评论状态字段
type Post struct {
	ID           uint `gorm:"primaryKey"`
	Title        string
	Content      string
	UserID       uint
	User         User
	CommentCount int       `gorm:"default:0"`     // 评论数量统计
	CommentState string    `gorm:"default:'有评论'"` // 评论状态
	Comments     []Comment `gorm:"foreignKey:PostID"`
}

// Comment 评论模型
type Comment struct {
	ID      uint `gorm:"primaryKey"`
	Content string
	UserID  uint
	PostID  uint
	Post    Post `gorm:"foreignKey:PostID"`
	User    User `gorm:"foreignKey:UserID"`
}

var db *gorm.DB

func main() {
	// 配置数据库连接
	dsn := "root:root@tcp(127.0.0.1:3306)/go?charset=utf8mb4&parseTime=True&loc=Local"

	// 创建自定义日志器
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	// 连接数据库
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(fmt.Sprintf("failed to connect database: %v", err))
	}
	defer closeDB()

	// 自动迁移表结构
	if err := db.AutoMigrate(&User{}, &Post{}, &Comment{}); err != nil {
		panic(fmt.Sprintf("failed to migrate database: %v", err))
	}

	// 注册钩子函数
	db.Callback().Create().After("gorm:create").Register("post:after_create", postAfterCreateHook)
	db.Callback().Delete().After("gorm:delete").Register("comment:after_delete", commentAfterDeleteHook)

	// 生成测试数据
	generateTestData()

	// 查询用户及其文章
	user, err := getUserWithPosts("User1")
	if err != nil {
		panic(fmt.Sprintf("查询用户失败: %v", err))
	}
	displayUserPosts(user)

	// 查询评论数量最多的文章
	topPost, err := getPostWithMostComments()
	if err != nil {
		panic(fmt.Sprintf("查询热门文章失败: %v", err))
	}
	displayTopPost(topPost)
}

// 关闭数据库连接
func closeDB() {
	sqlDB, err := db.DB()
	if err != nil {
		log.Printf("获取数据库连接失败: %v", err)
		return
	}
	if err := sqlDB.Close(); err != nil {
		log.Printf("关闭数据库连接失败: %v", err)
	}
}

// 文章创建后钩子函数
func postAfterCreateHook(db *gorm.DB) {
	if db.Error != nil {
		return
	}

	var post Post
	if err := db.First(&post).Error; err != nil {
		db.AddError(err)
		return
	}

	// 更新用户的文章数量
	//var user User
	if err := db.Model(&User{}).Where("id = ?", post.UserID).Update("post_count", gorm.Expr("post_count + 1")).Error; err != nil {
		db.AddError(err)
		return
	}

	fmt.Printf("文章创建后钩子: 已更新用户 %d 的文章数量为 %d\n", post.UserID, post.User.PostCount+1)
}

// 评论删除后钩子函数
func commentAfterDeleteHook(db *gorm.DB) {
	if db.Error != nil {
		return
	}

	var comment Comment
	if err := db.First(&comment).Error; err != nil {
		db.AddError(err)
		return
	}

	// 统计文章的评论数量
	var commentCount int64
	if err := db.Model(&Comment{}).Where("post_id = ?", comment.PostID).Count(&commentCount).Error; err != nil {
		db.AddError(err)
		return
	}

	// 更新文章的评论数量和状态
	updateData := map[string]interface{}{
		"comment_count": commentCount,
	}
	if commentCount == 0 {
		updateData["comment_state"] = "无评论"
	}

	if err := db.Model(&Post{}).Where("id = ?", comment.PostID).Updates(updateData).Error; err != nil {
		db.AddError(err)
		return
	}

	fmt.Printf("评论删除后钩子: 文章 %d 现在有 %d 条评论，状态: %s\n",
		comment.PostID, commentCount, updateData["comment_state"])
}

// 生成测试数据
func generateTestData() {
	// 清空现有数据
	db.Migrator().DropTable(&Comment{}, &Post{}, &User{})
	db.AutoMigrate(&User{}, &Post{}, &Comment{})

	rand.Seed(time.Now().UnixNano())

	// 生成用户
	users := []User{
		{Name: "User1", Email: "user1@example.com"},
		{Name: "User2", Email: "user2@example.com"},
	}
	db.Create(&users)

	// 生成文章
	posts := []Post{
		{Title: "Go语言入门指南", Content: "这是一篇关于Go语言的入门文章...", UserID: users[0].ID},
		{Title: "数据库设计最佳实践", Content: "本文介绍数据库设计的最佳实践...", UserID: users[0].ID},
		{Title: "微服务架构解析", Content: "微服务架构的详细解析...", UserID: users[1].ID},
	}
	db.Create(&posts)

	// 生成评论
	comments := []Comment{
		{Content: "非常有见地的文章，感谢分享！", UserID: users[1].ID, PostID: posts[0].ID},
		{Content: "请问作者，能否详细解释一下第三部分的内容？", UserID: users[0].ID, PostID: posts[0].ID},
		{Content: "我在实际项目中使用过类似技术，效果不错。", UserID: users[1].ID, PostID: posts[1].ID},
		{Content: "文章结构清晰，讲解深入浅出，赞！", UserID: users[0].ID, PostID: posts[2].ID},
	}
	db.Create(&comments)

	fmt.Println("测试数据生成完成")
}

// 查询用户及其文章
func getUserWithPosts(userName string) (User, error) {
	var user User
	err := db.Preload("Posts").Where("name = ?", userName).First(&user).Error
	return user, err
}

// 显示用户文章
func displayUserPosts(user User) {
	if user.ID == 0 {
		fmt.Printf("未找到名为 %s 的用户\n", user.Name)
		return
	}

	fmt.Printf("\n用户 %s 的文章信息:\n", user.Name)
	fmt.Printf("  文章数量: %d\n", user.PostCount)

	for i, post := range user.Posts {
		fmt.Printf("  文章 %d: %s\n", i+1, post.Title)
		fmt.Printf("    评论数量: %d, 状态: %s\n", post.CommentCount, post.CommentState)
	}
}

// 查询评论数量最多的文章
func getPostWithMostComments() (Post, error) {
	var post Post
	err := db.Table("posts p").
		Select("p.*, COUNT(c.id) as comment_count").
		Joins("LEFT JOIN comments c ON p.id = c.post_id").
		Group("p.id").
		Order("comment_count DESC").
		Limit(1).
		Scan(&post).
		Error
	return post, err
}

// 显示热门文章
func displayTopPost(post Post) {
	if post.ID == 0 {
		fmt.Println("\n没有找到文章")
		return
	}

	fmt.Println("\n评论数量最多的文章:")
	fmt.Printf("  标题: %s\n", post.Title)
	fmt.Printf("  评论数量: %d, 状态: %s\n", post.CommentCount, post.CommentState)
}
