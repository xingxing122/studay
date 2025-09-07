package main

import (
	"fmt"
)

func main() {
	num := []int{1, 2, 3, 4}

	strs := []string{"张三", "lisi"}

	fmt.Println(num, strs)

	fmt.Printf("长度: %v, 容量: %v", len(num), cap(num))

}
