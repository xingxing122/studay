package main

import "fmt"

func main() {
	left, right := "梨子","苹果"
	fmt.Println("left=",left,"right=",right)

	destktop := left
	left = right
	right  = destktop
	fmt.Println("left=", left, "right=", right )

	 left ,right = right,left

	fmt.Println("left=", left, "right=", right )


}
