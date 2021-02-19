package main

import "fmt"

func main() {
      left,right :="梨","苹果"
      fmt.Println("left=",left,"right=",right)

      desktop := left
      left = right
      right = desktop
      fmt.Println("left=",left,"right=",right)

      left,right = right,left
	fmt.Println("left=",left,"right=",right)





}
