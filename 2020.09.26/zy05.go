package main

import "fmt"

func main() {
	// 找出最大的数字
	zy := []int{108, 107, 105, 109, 103, 102}

	for j := 0;j<len(zy)-1;j++{
		for i :=0;i<len(zy)-1;i++{
			if zy[i]>zy[i+1]{
				zy[i],zy[i+1]= zy[i+1],zy[i]
			}
		}
	}
  fmt.Println(zy)




}