package main

import (
	"fmt"
	"math/bits"
	"strconv"
)

func main() {

	v, err := div(1, 0);err == nil{
		fmt.Println(v)
	} else {
		fmt.Println("value is error", err)
	}

	fmt.Println(strconv.Itoa(123))
	fmt.Println(v, err)
	fmt.Println(div(2, 1))
	fmt.Println(divV2(1, 0))
	fmt.Println(divV2(2, 1))

}

