package main

import "fmt"

func main() {
   //老婆的想法

   fmt.Println("老婆的想法")
   fmt.Println("买十个包子")
   var yes string
   fmt.Println("有卖西瓜的吗?")
   fmt.Scan(&yes)
   if yes == "Y" || yes == "y" {
      fmt.Println("买一个西瓜")
   }

   fmt.Println("老公的想法")

   if yes == "Y" || yes == "y" {
      fmt.Println("买一个包子")
   } else {
      fmt.Println("买十个包子")
   }

   var  score int
   fmt.Print("请输入成绩:")
   fmt.Scan(&score)

   if score >= 90 {
      fmt.Println("A")
   } else {
      if score >= 80 {
         fmt.Println("B")
      } else {
         if score <= 70 {
            fmt.Println("C")
         }
      }
   }

}






















