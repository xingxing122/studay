package main

import "fmt"

func main() {

	for i := 0; i < 2; i++ {
		for j := 0; j < 10; j++ {
			if j == 3 {
				break
			}
			fmt.Printf("i=%v j=%v\n", i, j)
		}

	}
}
