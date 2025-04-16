package main

import "fmt"

func main() {
	map1 := make(map[int]int, 10)
	map1[10] = 100
	map1[1] = 13
	map1[2] = 56
	map1[8] = 90

	for key, val := range map1 {
		fmt.Println(key, val)
	}

}
