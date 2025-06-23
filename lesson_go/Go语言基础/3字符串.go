package main

// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
// 有效字符串需满足：
// 左括号必须用相同类型的右括号闭合。
// 左括号必须以正确的顺序闭合。
// 每个右括号都有一个对应的相同类型的左括号。
func main() {
	//str := "([])"
	str := "()[]"
	if isValid(str) {
		println(str, "有效")
	} else {
		println(str, "无效")
	}
}

func isValid(s string) bool {
	//var stack []rune  // 声明一个rune切片变量，默认初始化为nil
	//stack = make([]rune, 0)  // 使用make函数创建空切片
	stack := []rune{}
	pairs := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}
	for _, char := range s {
		switch char {
		case '(', '{', '[':
			stack = append(stack, char)
		case ')', '}', ']':
			if len(stack) == 0 || stack[len(stack)-1] != pairs[char] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
