package main

import "fmt"

func main() {
	var name string  = "abc"

	fmt.Printf("%T, %s\n", name,name)

	// 索引
	//
	name  += "hi"

	name = "我爱中国"
	fmt.Printf("%T\n", name[0])
    fmt.Println(len(name))
	fmt.Println(name[0:4])


}
