package main

import "fmt"

func main() {
   nums := [...] int{1,2,3,4,5}
   /*nums2 := nums[1:3]
   fmt.Printf("%T\n",nums2)
   fmt.Println(nums2)

   fmt.Println(cap(nums2),len(nums2))
   */

   nums2 := nums[1:3:4]
   //cap = cap_end-start
   //len= end-start
   fmt.Println("%T\n",nums2)
   fmt.Println(nums2)
   fmt.Println(cap(nums2),len(nums2))





}
