package main

import "fmt"

func main() {
	// 字面量初始化
	var  name map[string]string= map[string]string{"Go999":"开始"}
	//var name  map[string]string
	fmt.Printf("%T %v\n",name,name)

	//访问元素
	fmt.Printf("%q\n",name["Go888"])

	//添加元素
	name["Go888"] = "学习go"
	fmt.Println(name)
    //修改元素
    name["Go999"]="努力"
    fmt.Println(name)

    //删除元素
    delete(name,"Go888")
    fmt.Println(name)

    //遍历

    for k := range name {
    	fmt.Println(name[k])
	}

	for   k,v :=range name{
		fmt.Println(k,v)
	}

	var  scores = make(map[string]int)
	scores["Go3028"] = 90

	for k,v :=range  scores{
		fmt.Println(k,v)

	}



}
