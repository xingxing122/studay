package main

import "fmt"

func main() {
	//解决大量的if  elase if 的判断
	score := "B"
	switch score {
	case "A":
		fmt.Println("优秀")
	case "B":
		fmt.Println("中等")
	case "C":
		fmt.Println("下等")
	default:
		fmt.Println("不及格")

	}

}
