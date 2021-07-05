package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	for i :=1; i<=5; i++ {
		var a int
		fmt.Print("请输入数字")
        fmt.Scan(&a)
		rand.Seed(time.Now().Unix())
		sj:=rand.Intn(100)
		fmt.Println("打印随机数",sj)
		if a > sj {
			fmt.Println("提示太大了")
		}else if a < sj {
			fmt.Println("提示太小了")
		}else {
			fmt.Println("猜对了")
			break
		}

		}
}
