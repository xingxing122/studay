package main

import "fmt"




func main() {
    users := make(map[string]map[string]string)
    fmt.Println("欢迎添加用户")
    for {
        var op string
        fmt.Print(`
        1. 新建用户
        2. 修改用户
        3. 删除用户
        4. 查询用户
        5. 退出
        请输入指令:`)
        fmt.Scan(&op)
        fmt.Println("请输入指令")
    }



}
