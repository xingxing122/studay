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
		v, ok :=stat[name]
		if !ok {
			stat[name] = 1
		} else {
			stat[name] = v+1
		}
	}
	fmt.Println(stat)





}
