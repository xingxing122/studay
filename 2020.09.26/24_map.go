package main

import "fmt"

func main() {
/*	var m1 map[int]string
	fmt.Println("m1==",m1)
	fmt.Println("len=",len(m1))

	m2 := make(map[int]string)
	fmt.Println("m2==",m2)
	fmt.Println("len=",len(m2))
*/

	var   scores  map[string]int
	fmt.Printf("%T %#v\n", scores,scores)
	fmt.Println(scores==nil)

	scores = map[string]int{"数据":1,"abcd":2}
	fmt.Println(scores)

	scores = make(map[string]int)
	fmt.Println(scores)

}
