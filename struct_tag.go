package main

import (
    "fmt"
    "reflect" // 这里引入reflect模块
)

type User struct {
    Name   string `flag:"user name" cfg:"id"` //这引号里面的就是tag
    Passwd string `"user passsword"`
}

func main() {
    user := &User{"Johnsun", "helloworld"}
    s := reflect.TypeOf(user).Elem() //通过反射获取type定义
    for i := 0; i < s.NumField(); i++ {
        fmt.Println(s.Field(i).Tag.Get("flag"),s.Field(i).Tag.Get("cfg")) //将tag输出出来
        fmt.Println(s.Field(i).Tag)
    }
}

//go里面StructTag是个string，格式为tagName:"tagValue"
//如果直接field(i).Tag是取到整个tag内容
//Get函数可以拿到tag里的某一个参数，可以定制化