package main

import "fmt"

func main() {

    var names [] string
    fmt.Println(names)

    names = []string{"go","中国"}

    fmt.Printf("%T\n",names)
    fmt.Println("%T\n",names)

    names = append(names,"陕西")
    fmt.Println(names)
    fmt.Println(len(names))

    fmt.Printf("%q\n",names[1:4])

    names = names[1:len(names)]
    fmt.Printf("%q\n",names)
    names = names[0:len(names)-1]
    fmt.Printf("%q\n",names)
    // 删除中间元素
    nums := []int{0,1,2,3,4,5}
    nums2 :=[]int{10,11,12,13,14}

    copy(nums,nums2)
    fmt.Println(nums,nums2)

    copy (nums[3:len(nums)],nums[4:len(nums)])

    fmt.Println(nums)



}
