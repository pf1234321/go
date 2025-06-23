package main

//
//import (
//	"fmt"
//	"gorm.io/driver/mysql"
//	"gorm.io/gorm"
//	"gorm.io/gorm/logger"
//	"log"
//	"os"
//	"time"
//)
//
//// User 用户模型
//type User struct {
//	ID    uint `gorm:"primaryKey"`
//	Name  string
//	Email string `gorm:"unique"`
//	Posts []Post `gorm:"foreignKey:UserID"`
//}
//
//// Post 文章模型
//type Post struct {
//	ID       uint `gorm:"primaryKey"`
//	Title    string
//	Content  string
//	UserID   uint
//	User     User
//	Comments []Comment `gorm:"foreignKey:PostID"`
//}
//
//// Comment 评论模型
//type Comment struct {
//	ID      uint `gorm:"primaryKey"`
//	Content string
//	UserID  uint
//	PostID  uint
//	User    User `gorm:"foreignKey:UserID"`
//}
//
//// 评论数量最多的文章查询结果结构体
//type PostWithCommentCount struct {
//	Post             // 嵌入Post结构体
//	CommentCount int `gorm:"column:comment_count"`
//}
//
//var db *gorm.DB
//
//func main() {
//	// 配置数据库连接
//	dsn := "root:root@tcp(127.0.0.1:3306)/go?charset=utf8mb4&parseTime=True&loc=Local"
//
//	// 创建自定义日志器
//	newLogger := logger.New(
//		log.New(os.Stdout, "\r\n", log.LstdFlags),
//		logger.Config{
//			SlowThreshold: time.Second,
//			LogLevel:      logger.Info,
//			Colorful:      true,
//		},
//	)
//
//	// 连接数据库
//	var err error
//	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
//		Logger: newLogger,
//	})
//	if err != nil {
//		panic(fmt.Sprintf("failed to connect database: %v", err))
//	}
//	defer closeDB() // 程序结束时关闭数据库连接
//
//	// 查询用户及其文章和评论
//	userName := "User1" // 要查询的用户名
//	user, err := getUserWithPostsAndComments(userName)
//	if err != nil {
//		panic(fmt.Sprintf("用户文章查询失败: %v", err))
//	}
//	displayUserPosts(user)
//
//	// 查询评论数量最多的文章
//	topPost, err := getPostWithMostComments()
//	if err != nil {
//		panic(fmt.Sprintf("热门文章查询失败: %v", err))
//	}
//	displayTopPost(topPost)
//}
//
//// 关闭数据库连接
//func closeDB() {
//	sqlDB, err := db.DB()
//	if err != nil {
//		log.Printf("获取数据库连接失败: %v", err)
//		return
//	}
//	if err := sqlDB.Close(); err != nil {
//		log.Printf("关闭数据库连接失败: %v", err)
//	}
//}
//
//// 查询用户及其文章和评论
//func getUserWithPostsAndComments(userName string) (User, error) {
//	var user User
//	err := db.Preload("Posts.Comments.User"). // 预加载文章及评论的用户信息
//							Where("name = ?", userName).
//							First(&user).
//							Error
//	return user, err
//}
//
//// 显示用户文章及评论
//func displayUserPosts(user User) {
//	if user.ID == 0 {
//		fmt.Printf("未找到名为 %s 的用户\n", user.Name)
//		return
//	}
//
//	fmt.Printf("\n用户 %s 的文章及评论信息:\n", user.Name)
//	for i, post := range user.Posts {
//		fmt.Printf("  文章 %d: %s (ID: %d)\n", i+1, post.Title, post.ID)
//		fmt.Printf("    内容: %s\n", post.Content)
//		fmt.Printf("    评论数量: %d\n", len(post.Comments))
//
//		for j, comment := range post.Comments {
//			fmt.Printf("      评论 %d: %s (用户: %s)\n",
//				j+1, comment.Content, comment.User.Name)
//		}
//	}
//}
//
//// 查询评论数量最多的文章
//func getPostWithMostComments() (PostWithCommentCount, error) {
//	var topPost PostWithCommentCount
//	err := db.Table("posts p").
//		Select("p.*, u.name as user_name, COUNT(c.id) as comment_count").
//		Joins("JOIN users u ON p.user_id = u.id").
//		Joins("LEFT JOIN comments c ON p.id = c.post_id").
//		Group("p.id, u.name").
//		Order("comment_count DESC").
//		Limit(1).
//		Scan(&topPost).
//		Error
//
//	// 加载评论详情
//	if err == nil && topPost.ID > 0 {
//		if err := loadCommentsForPost(&topPost.Post); err != nil {
//			return topPost, fmt.Errorf("加载评论详情失败: %w", err)
//		}
//	}
//	return topPost, err
//}
//
//// 加载文章的评论详情
//func loadCommentsForPost(post *Post) error {
//	return db.Preload("Comments.User").First(post, post.ID).Error
//}
//
//// 显示热门文章信息
//func displayTopPost(topPost PostWithCommentCount) {
//	if topPost.ID == 0 {
//		fmt.Println("\n没有找到文章")
//		return
//	}
//
//	fmt.Println("\n评论数量最多的文章:")
//	fmt.Printf("ID: %d\n", topPost.ID)
//	fmt.Printf("标题: %s\n", topPost.Title)
//	fmt.Printf("内容: %s\n", topPost.Content)
//	fmt.Printf("作者: %s\n", topPost.User.Name)
//	fmt.Printf("评论数量: %d\n", topPost.CommentCount)
//
//	if len(topPost.Comments) > 0 {
//		fmt.Printf("评论详情 (%d 条):\n", len(topPost.Comments))
//		for i, comment := range topPost.Comments {
//			fmt.Printf("  %d. %s (用户: %s)\n",
//				i+1, comment.Content, comment.User.Name)
//		}
//	}
//}
