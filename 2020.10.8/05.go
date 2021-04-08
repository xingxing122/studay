package main

func main(){
	defer func() {
	  if err := recover();err != nil

	}()
	panic("我是错误")

}

