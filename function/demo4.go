package main

import (
	"fmt"
	"sort"
)

func mapSort(map1 map[string]string) string {

	var slicekey []string
	for k, _ := range map1 {
		slicekey = append(slicekey, k)

	}
	sort.Strings(slicekey)

	var str string
	for _, v := range slicekey {

		str += fmt.Sprintf("%v=>%v", v, map1[v])
	}
	return str
}

func main() {

	m1 := map[string]string{
		"username": "zhangsan",
		"age":      "20",
		"sex":      "男",
		"height":   "180",
		"salt":     "xxxxx",
	}

	str := mapSort(m1)
	fmt.Println(str)

}
