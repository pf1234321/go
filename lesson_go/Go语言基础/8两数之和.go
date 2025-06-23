package main

import "fmt"

func main() {
	var numbers = [5]int{2, 7, 11, 15}
	target := 13
	for i, number := range numbers {
		for j, number2 := range numbers {
			if i != j && number+number2 == target {
				fmt.Println("numbers:", i, j)
				return
			}
		}
	}
}
