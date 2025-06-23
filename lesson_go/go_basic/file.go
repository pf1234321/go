package main

import (
	"log"
	"os"
)

func main() {
	// 创建文件，如果文件已存在会被截断（清空）
	file, err := os.Create("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close() // 确保文件关闭
	//在 Go 语言里，defer 关键字的作用是延迟执行函数调用。
	//也就是说，被 defer 修饰的函数会在当前函数执行结束前才执行，不管当前函数是正常返回还是因为 panic 而终止
	//类似Java中的finally

	log.Println("文件创建成功")
}
