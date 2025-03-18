package main

import "fmt"

func main() {

	/*	var arr1 [3]int
		var arr2 [4]int

		arr1[1] = 10

		fmt.Printf("arr1: %T  arr2: %T", arr1, arr2)

		fmt.Println(arr1)
	*/

	/*	var arr1 [3]int
		arr1[0] = 1
		arr1[1] = 100
		arr1[2] = 1000
		fmt.Println(arr1)
	*/
	/*	arr1 := [...]string{"php", "nodejs", "golang", "js", "java"}

		for _, k := range arr1 {
			fmt.Println(k)
		}*/

	/*	var arr = [...]int{12, 23, 100, 200, 400, 500}
		var sum = 0
		for _, v := range arr {
			sum += v
		}
		fmt.Println(sum)
		fmt.Printf("arr数组的和是: %v  平均值是： %.2f", sum, float64(sum)/float64(len(arr)))*/

	/*	var arr = [...]int{1, -10, 100, 96, 209}
		max := arr[0]
		for _, k := range arr {
			if max < k {
				max = k
			}
		}
		fmt.Printf("max: %d\n", max)
	*/

	//从数组 [1,3,5,7,8] 中找出和为8 的两个元素的下标分别为(0，3)和（1,2）
	var arr = [...]int{1, 3, 5, 7, 8}
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] == 8 {
				fmt.Printf("(%v ,%v)", i, j)
			}
		}
	}
}
