package  main

import "fmt"

func main(){

     var  names [55]string

     var scores[2] int = [...]int {1:88}
     fmt.Printf("%T\n",names)
     fmt.Printf("%T\n",scores)

     fmt.Printf("%q\n",names)

     fmt.Println(names,scores)

     var  user [3]string = [3] string {"kk","蜗牛","abc"}
     bounds := [...]int{10,20,30,40,50}
     teachers := [5]string{"abcd"}
     nums := [10] int {1:10,3:30,5:50,8:80}

     fmt.Printf("%q\n",user)
     fmt.Println(bounds)
     fmt.Printf("%q\n",teachers)
     fmt.Println(nums)
     fmt.Println(nums[0])
     fmt.Println(nums[1])
     fmt.Println(nums[2])

     fmt.Println(len(nums))



     for  i :=0; i < len(nums) ; i ++ {
             fmt.Println(i,nums[i])
     }

       for   v := range  nums {
               fmt.Println(v)
       }

       for  i,v := range  nums {
               fmt.Println(i,v)
       }


       var  ms [3][2]int
       fmt.Println(ms)



}