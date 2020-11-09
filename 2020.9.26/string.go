package main

import "fmt"
func main() {

	var name string = `xingxing`
	var abc string = `\n`
	fmt.Printf("%T,%s\n",name,name)
	fmt.Printf("%T,%s\n",abc,abc)

	name += "hi"
	fmt.Println(name)

	fmt.Printf(name[0:4])




}
