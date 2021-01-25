package  main

import "fmt"

func main(){
	   //var name map[string]string
	   var  name map[string]string  = map[string]string{"G0001":"我爱祖国"}
	   fmt.Printf("%T  %v\n",name,name)

	   fmt.Printf("%q\n",name["Go3005"])
	   name["Go3005"] ="我爱中国"
	   fmt.Println(name)
	   name["Go3002"]="我爱家乡"
	   fmt.Println(name)

	   v, ok := name["Go00001"]
	   fmt.Println(v,ok)

	   delete(name,"Go3002")
	   fmt.Println(name)

       for k := range  name {
       	fmt.Println(k)
	   }

	   for v,k := range name {
	   	fmt.Println(v,k)
	   }

      var  int = make(map[string]int)
      int ["Go3002"] = 90
      int ["Go3003"] = 91
      int ["Go30004"] = 92

      for k,v := range int {
      	fmt.Println(k,v)
	  }
	  for v := range  int {
	  	fmt.Println(v)
	  }


}
