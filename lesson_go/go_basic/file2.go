package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 1. 创建文件
	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}

	// 2. 声明 defer 语句（此时 file.Close() 未执行）
	//关键要点：defer 语句只是 “注册” 了一个延迟执行的函数调用，不会阻塞后续代码的执行。
	defer file.Close()

	// 3. 创建缓冲写入器并写入内容
	writer := bufio.NewWriter(file)
	fmt.Fprintln(writer, "Hello, World!")
	writer.Flush() // 刷新缓冲区，确保数据写入文件

	// 4. main 函数返回时，才会执行 defer 中的 file.Close()
}
