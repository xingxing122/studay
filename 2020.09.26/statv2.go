package main

import "fmt"

func main() {
	name :=[]string{
		"小明",
		"小黑",
		"小红",
		"大黄",
		"大黄",
		"小明",
		"小黑",
		"小红",
		"大黄",
	}

	stat  := map[string]int{}
	for _,name := range name {
		stat[name]++
	}
	fmt.Println(stat)





}
