package main

import "fmt"

func main() {
	for i := 1; i <= 5; i++ {
		for j := 2; j <= 4; j++ {
			fmt.Printf("i : %v, j ; %v \n", i, j)
			for i == 2 && j == 2 {
				break
			}
		}
	}
}
