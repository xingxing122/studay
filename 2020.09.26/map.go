package  main

import "fmt"

func main(){
	// 字面量初始化
	var name map[string]string = map[string]string{"go3037":"李雄"}

	//map[keytype]valuetype{k:v1,k:v2}


	fmt.Printf("%T %v\n",name,name)
	name["Go3038"] ="王星星"
	fmt.Println(name)

	v, ok := name["Go0001"]
	fmt.Println(v,ok)

	v, ok = name["Go3038"]
	fmt.Println(v,ok)

	delete(name,"Go3038")
	fmt.Println(name)

	fmt.Println("_____________________")

	for k := range name {
		fmt.Println(k,name[k])
	}

	for v ,k := range  name {
		fmt.Println(v,k)
	}


	var score = make(map[string]int)
	score["01"] = 90
	score["02"] = 80
	score["03"] = 87
	score["04"] = 79
	score["05"] = 78
	score["06"] = 76
	score["07"] = 75
	score["08"] = 60
	fmt.Println(score)

	for v,k := range score{
		fmt.Println(v,k)
	}

}
