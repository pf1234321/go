package main

import (
	"fmt"
	"strconv"
)

func main() {
	//var numbers = [5]int{2, 2, 3, 9, 9}
	var numbers = [1]int{9}
	str := ""
	for _, number := range numbers {
		str += fmt.Sprintf("%d", number)
	}
	fmt.Println(str)
	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("转换失败:", err)
		return
	}
	num += 1
	str2 := strconv.Itoa(num)          // "22336"
	numbers2 := make([]int, len(str2)) // 创建长度一致的切片

	for i, ch := range str2 {
		numbers2[i] = int(ch - '0') // 将每个字符转为 int
	}

	fmt.Println(numbers2) // 输出: [2 2 3 3 6]

}
