package main

import "fmt"

func main(){
	var name map[string]string  = map[string]string{"Go9898":"天天向上"}
	//一定要初始化

	fmt.Printf("%T %v\n",name,name)

	fmt.Println(name["Go3009"])
	name["Go3009"]="学习go"
	fmt.Println(name)
	v,ok  := name["Go3009"]
	fmt.Println(v,ok)
	//删除
	delete(name,"Go3009")
	fmt.Println(name)

	for  k,v := range  name {
		fmt.Println(k,v)
	}

    // make赋值
	//var scores  map[string]int = make(map[string]int)
	var scores = make(map[string]int)
	scores["Go3028"]=80
	scores["Go3027"]=84
	scores["Go3025"]=88
	scores["Go3029"]= 89
	scores["Go3030"]=81
	scores["Go3027"]= 100

	for k, v := range scores {
		fmt.Println(k,v)
	}

}

