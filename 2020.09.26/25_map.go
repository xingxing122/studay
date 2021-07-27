package main

import (
	"fmt"
)

func main() {

	m1 := map[int]string{1:"mike",2:"yoyo",3:"go"}

	fmt.Println("m1==",m1)
	m1[1] = "c++"
	fmt.Println("m1=",m1)
    m1[3] = "go"
    fmt.Println("m2==",m1)

    for key, value := range m1 {
    	fmt.Println("%d=====> %s\n", key,value)
	}
	//如何判断一个key 值是否存在

	value, ok := m1[1]
	if  ok == true {
		fmt.Println("m1[1]==",value)
	}else {
		fmt.Println("key 不存在")
	}

	delete(m1,1)
	fmt.Println(m1)
	test(m1)
	fmt.Println()


}
