package main

import "fmt"
 
func main() {
    // 创建一个空的 Map
    m := make(map[string]int)
    fmt.Println(m)
    // 使用字面量创建 Map
    m = map[string]int{
        "apple": 1,
        "banana": 2,
        "orange": 3,
    }
    fmt.Println(m)
	// 访问 Map 中的值
	v1 := m["apple"]
	fmt.Println("Value for 'apple':", v1)
	// v2, ok := m["banana"] 
	// if ok {
	// 	fmt.Println("Value for 'banana':", v2)
	// } else {
	// 	fmt.Println("'pear' not found in map")
	// }


	// // 修改键值对
	// m["apple"] = 5
	// fmt.Println("Updated value for 'apple':", m["apple"])
	//  fmt.Println(m)


	 // 获取 Map 的长度
	len := len(m)
	fmt.Println("Length of map:", len)



	// 遍历 Map
	for k, v := range m {
		fmt.Printf("key=%s, value=%d\n", k, v)
	}

	// 删除键值对
	delete(m, "banana")
	fmt.Println("After deleting 'banana':", m)

}