package main

import "fmt"

func main() {
	s := "hello 张三"

	/* for 循环遍历字符串 */
	//for i := 0; i < len(s); i++ {
	//	fmt.Printf("%T: %v", s[i], s[i])
	//}
	//使用range 遍历字符串
	for index, val := range s {
		fmt.Println(index, val)
	}

}
