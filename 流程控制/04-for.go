package main

import "fmt"

func main() {
	var sum = 0
	var count = 0

	for i := 1; i <= 100; i++ {
		if i%9 == 0 {
			fmt.Println(i)
			sum += i
			count++
		}
	}
	fmt.Println(sum)
	fmt.Println(count)
}
