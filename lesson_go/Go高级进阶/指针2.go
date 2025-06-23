package main

import "fmt"

func main() {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	printSlice(numbers)
	fmt.Printf("len=%d cap=%d slice=%v\n", len(numbers), cap(numbers), numbers)
}

func printSlice(numbers []int) {
	if numbers == nil {
		fmt.Printf("切片是空的")
	}
	for i := 0; i < len(numbers); i++ {
		temp := numbers[i]
		numbers[i] = temp * 2
	}

}
