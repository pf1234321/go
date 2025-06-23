package main

import "fmt"

func main() {
	var numbers = [5]int{2, 2, 3, 3, 5}
	map1 := make(map[int]int)
	for i, number := range numbers {
		map1[number] = i
	}

	numbers2 := make([]int, 0, len(map1)) // 预分配容量提升性能
	for key := range map1 {
		numbers2 = append(numbers2, key)
	}

	fmt.Println("numbers2:", numbers2)            // 输出类似：[2 3 5]
	fmt.Println("numbers2 count:", len(numbers2)) // 输出类似：[2 3 5]
}
