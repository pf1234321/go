package main

import (
	"fmt"
	"strconv"
)

// 考察：数字操作、条件判断
// 题目：判断一个整数是否是回文数
func main() {
	var number uint64 = 12210
	if checkNumber(number) {
		fmt.Println(number, "是回文数")
	} else {
		fmt.Println(number, "不是回文数")
	}
}

func checkNumber(number uint64) bool {
	str := strconv.FormatUint(number, 10)
	reverseStr := reverse(str)
	if str == reverseStr {
		return true
	}
	return false
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
