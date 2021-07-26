package main

import (
	"fmt"
	"strconv"
)

func main() {
	//字符串<==> int
	i,err := strconv.Atoi("abc")
	fmt.Println(i,err)
	i, err = strconv.Atoi("-15")
	fmt.Println(i,err)

    fmt.Println(strconv.Itoa(45))


}
