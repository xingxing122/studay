package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var a int32 = 15

	fmt.Printf("num=%v 类型: %T\n", a, a)
	fmt.Println(unsafe.Sizeof(a))

}
