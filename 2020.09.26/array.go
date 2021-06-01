package main

import "fmt"

func main() {
        var names [55] string
        var  scores = [...]int{1:99}

        fmt.Printf("%T\n", names)
        fmt.Printf("%T\n",scores)

        fmt.Println(names,scores)

        var nums [2]int=[...]int{100,88}
        //零值
        fmt.Println(nums == scores)
        fmt.Println(nums[0])
        fmt.Println(nums[1])

        nums[0]= 100
        nums[1]=101
        fmt.Println(nums)
        fmt.Println(len(nums))

        for i:=0; i< len(nums); i++ {
                fmt.Println(i,nums[i])
        }

        for v := range  nums {
                fmt.Println(v,nums[v])
        }

        for i,v := range  nums{
                fmt.Println(i,v)
        }

       var mms = [...][2]int{
               1: [2]int{1,2},
               2: [2]int{3,4},
               3: [2]int{5,6},

       }
       fmt.Printf("%T\n",mms)
       fmt.Println(mms)

}
