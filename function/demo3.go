package main

import "fmt"

func sortIntAsc(slice []int) []int {

	for i := 0; i < len(slice); i++ {
		for j := i + 1; j < len(slice); j++ {
			if slice[i] > slice[j] {
				temp := slice[i]
				slice[i] = slice[j]
				slice[j] = temp
			}
		}
	}
	return slice
}

func main() {
	var sliceA = []int{12, 34, 47, 56, 9, 2}

	arr := sortIntAsc(sliceA)
	fmt.Println(arr)

}
