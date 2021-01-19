package main

import (
	"fmt"

)

func main() {
	var name map[string]string = map[string]string{"Go3007":"中国"}
	fmt.Printf("%T %v\n",name,name)
	fmt.Printf("%q\n",name["G036"])

	name["Go3006"]="学习go"
	fmt.Println(name)

	//判断key 是否存在
	 v,ok  := name["Go0002"]
	 fmt.Println(v,ok)

	 for k,v :=range  name {
	 	fmt.Println(v,k)
	 }

	 var score = make(map[string]int)
	 score["Go3028"]  = 90
	 score["Go3018"]  = 91
	 score["Go3008"]  = 92
	 score["Go3098"]  = 92

	 for k,v :=range  score {
	 	fmt.Println(k,v)
	 }






}
