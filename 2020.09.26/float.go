package main

import "fmt"

func main() {
	height := 1.70
	fmt.Printf("%T,%f,%g,%e\n",height,height,height,height)

	fmt.Println(1.2+1.1 )
    fmt.Println(1.2-1.1 )
	fmt.Println(1.2*1.1)
	fmt.Println(1.2/1.1)


    height--
    fmt.Println(height)

	height++
	fmt.Println(height)




}
