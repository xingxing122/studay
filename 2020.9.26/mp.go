package main

import (
	"fmt"
)

func main() {
	heights := []int{10, 6, 7, 9, 5}
	for j := 0; j < len(heights)-1; j++ {
		fmt.Println("第",j,"轮")
		for i := 0; i < len(heights)-1; i++ {
			if heights[i] > heights[i+1] {
				fmt.Println("交换:", heights[i])
				tmp := heights[i]
				heights[i] = heights[i+1]
				heights[i+1] = tmp
			}
			fmt.Println(i, "交换完毕", heights)

		}

	}
}