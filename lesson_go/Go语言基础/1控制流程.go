package main

import "fmt"

// 给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
// 可以使用 for 循环遍历数组，结合 if 条件判断和 map 数据结构来解决，
// 例如通过 map 记录每个元素出现的次数，然后再遍历 map 找到出现次数为1的元素。
func main() {
	var numbers = [5]int{2, 2, 3, 3, 5}
	map1 := make(map[int]int)
	for i := 0; i < len(numbers); i++ {
		v2, ok := map1[numbers[i]]
		if ok {
			map1[numbers[i]] = v2 + 1
		} else {
			map1[numbers[i]] = 1
		}
	}
	for k, v := range map1 {
		if v == 1 {
			fmt.Println(k)
		}
	}
}
