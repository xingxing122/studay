package main

import "fmt"


func main(){
	var  yes  string
	fmt.Print("有卖西瓜的吗?(Y/N):")
	fmt.Scan(&yes)

	fmt.Println("老婆的想法:")
	fmt.Println("十个包子")

	if   yes != "Y"  &&   yes  != "y" {

		   goto  END
	}
	fmt.Println("买一个西瓜")
	END:

		result := 0
		i  := 1
		START:
			if i > 100 {
				goto FOREND
			}
			result += i
			i++
			goto START
		FOREND:
			fmt.Println(result)

		sum := 0
		for idx := 1; idx <= 100; idx++{
			sum += idx
		}
       fmt.Println(sum)


}


