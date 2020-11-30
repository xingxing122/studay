package main

import "fmt"

func main() {
     var  score int
     fmt.Print("请输入成绩:")
     fmt.Scan(&score)

     switch  {

	 case score  >= 90:
	 	fmt.Println("A")
	 case score  >= 80:
	 	fmt.Println("B")
	 case score  >=70:
	 	fmt.Println("C")
	 default:
		  fmt.Println("E")
	 }

}
