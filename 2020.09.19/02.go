package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {

	for a := 1; a <= 5; a++ {
		var a int
		fmt.Println("请输入数字: ")
		fmt.Scan(&a)

		rand.Seed(time.Now().Unix())
		b := (rand.Intn(100))
		fmt.Println(b)

		if a > b {
			fmt.Println("太大了")
		} else  if  a < b {
			fmt.Println("太小了")
		} else if a == b {
			fmt.Println("成功了")
		}

	}
	fmt.Println("太笨了")

}