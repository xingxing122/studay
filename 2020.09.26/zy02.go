
package main

import "fmt"

func main() {
	// 找出最大的数字
	num := []int{108,107,105,109,103,102}
	max := num[0]

	for _,v := range  num {
		if v > max {
			max = v
		}
	}
  fmt.Println(max)

}