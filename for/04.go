package main

import "fmt"

// >=90 A
// >=80 B
// >=70 C
// >=60 D
// 其他 >= E

func main() {
	fmt.Print("请输入成绩:")
	var score int
	fmt.Scan(&score)

	if score >= 90 {
		fmt.Println("A")
	} else if score >= 80 {
		fmt.Println("B")
	} else if score >= 70 {
		fmt.Println("C")
	} else if score >= 60 {
		fmt.Println("D")
	} else {
		fmt.Println("E")
	}

}
