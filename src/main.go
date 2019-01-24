package main

import (
	"entity"
	"fmt"
)

func main(){
	fmt.Println("hello world 20190123")

	//var user
	var tt = entity.User{1,"ddd"}
	fmt.Println(tt.Name)
	entity.Echo()


}
