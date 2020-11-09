package  main

import "fmt"

func main(){

	var yes string

	fmt.Print("老婆的想法:")
	fmt.Print("有卖西瓜的吗?(Y/N):")
	fmt.Scan(&yes)

	fmt.Print("老婆的想法：")
	fmt.Print("十个包子")
	switch yes {
	case "y","Y":
		fmt.Println("一个西瓜")
	}

	fmt.Println("老公的想法:")
	switch yes {
	case "y","Y":
		fmt.Println("一个包子")
	default:
		fmt.Println("十个包子")
	}

	fmt.Print("请输入你的成绩:")
	var score int
	fmt.Scan(&score)
	switch  {
	case score >= 90:
		fmt.Println("A")
	case score >= 80:
		fmt.Println("C")
	case score <=60:
		fmt.Println("D")
	}

}
