package main

import "fmt"

func main() {

	var n1, n2,n3 int = 1, 8 ,-2
	var n4 uint = 2

	fmt.Println(n1 + n2)
	fmt.Println(n1 - n2)
	fmt.Println(n1 * n2)
	fmt.Println(n1 / n2)
	fmt.Println(n4,n3)

	n3++
	n4--
    fmt.Println("%T %T %T %T %T %T\n", n1,uint(1),n4,int(4),uint8(4),int64(n4))

	var f1 float32 = 3.1415926
	f1 += f1
	fmt.Println(f1)

	fmt.Printf("%T %T\n",f1,float64(f1))

}
