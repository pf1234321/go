package main

import "fmt"

// User 定义用户结构体
type User struct {
	Name string
	Age  int
}

// 值接收者：修改副本，不影响原结构体
func (u User) SetNameValue(name string) {
	u.Name = name // 修改的是副本的 Name 字段
	fmt.Printf("值接收者内部: Name=%s, Age=%d\n", u.Name, u.Age)
}

// 指针接收者：直接修改原结构体
func (u *User) SetNamePointer(name string) {
	u.Name = name // 通过指针修改原结构体的 Name 字段
	fmt.Printf("指针接收者内部: Name=%s, Age=%d\n", u.Name, u.Age)
}

func main() {
	// 创建用户实例
	user := User{Name: "Alice", Age: 30}

	fmt.Println("修改前:", user)
	// 输出: 修改前: {Alice 30}

	// 调用值接收者方法
	user.SetNameValue("Bob")
	fmt.Println("值接收者修改后:", user)
	// 输出: 值接收者内部: Name=Bob, Age=30
	// 输出: 值接收者修改后: {Alice 30}（原结构体未被修改）

	// 调用指针接收者方法
	user.SetNamePointer("Charlie")
	fmt.Println("指针接收者修改后:", user)
	// 输出: 指针接收者内部: Name=Charlie, Age=30
	// 输出: 指针接收者修改后: {Charlie 30}（原结构体被修改）

	// 直接通过指针调用方法（无需显式解引用）
	userPtr := &user
	userPtr.SetNamePointer("David")
	fmt.Println("通过指针调用修改后:", user)
	// 输出: 通过指针调用修改后: {David 30}
}
