package entity

import "fmt"

type User struct {
	//首字母大写代表public，别的包才能访问到
	Id int
	Name string
}

//func echo 测试
func Echo()  {
	fmt.Println("echo....")
}
