package main

import "fmt"

func main() {
   name := []string{
   	"小明",
   	"小红",
   	"小花",
   	"小黑",
   	"大黄",
   	"小明",
   	"大黄",
   	"大黄",
   }

   var stat map[string]int= map[string]int{}
   for _,name := range  name {
   	v,ok := stat[name]
   	if !ok {
   		stat[name] = 1
	}else {
		stat[name] = v + 1
	}
   }
   fmt.Println(stat)

}