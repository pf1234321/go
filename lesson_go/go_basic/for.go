package main

import "fmt"
 

func main() {
    // 外部循环，i 从 1 到 3
    for i := 1; i <= 3; i++ {
        fmt.Printf("不使用标记,外部循环, i = %d\n", i)
        // 内部循环，j 从 5 到 10
        for j := 5; j <= 10; j++ {
            fmt.Printf("不使用标记,内部循环 j = %d\n", j)
            // 终止内部循环，外部循环不受影响
            break
        }
    }
}