package main

import "fmt"

func main() {

	var num = []int{9, 8, 3, 6, 5, 4, 11, 14}

	for i := 0; i < len(num); i++ {
		for j := i + 1; j < len(num); j++ {
			if num[i] < num[j] {
				temp := num[i]
				num[i] = num[j]
				num[j] = temp
			}
		}
	}
	fmt.Println(num)
}
