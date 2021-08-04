package main

import "fmt"

func addBase(base int)func(int)int {
	return  func(num int ) int {
		return  base + num
	}

}


func main() {
    add2 := addBase(2)
    fmt.Printf("%T\n",add2)
    fmt.Println(add2(10))


}




