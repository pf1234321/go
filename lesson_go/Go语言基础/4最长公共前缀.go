package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for _, s := range strs[1:] {
		for len(prefix) > 0 && len(s) < len(prefix) || s[:len(prefix)] != prefix {
			prefix = prefix[:len(prefix)-1]
		}
		if len(prefix) == 0 {
			return ""
		}
	}
	return prefix
}

func main() {
	// 测试示例
	testCases := [][]string{
		{"flower", "flow", "flight"},
		{"dog", "racecar", "car"},
		{"apple", "app", "application"},
		{},
		{"single"},
	}

	for _, tc := range testCases {
		fmt.Printf("输入: %v\n输出: %s\n\n", tc, longestCommonPrefix(tc))
	}
}
