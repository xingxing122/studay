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
		"大黄",
	}
   //	var  stat map[string]int = map[string]int{}
   //	fmt.Println(stat)
    stat  := map[string]int{}
   	for  _,name := range  name {
             stat[name]++
	}
	fmt.Println(stat)
}
