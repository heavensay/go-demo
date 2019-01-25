package syntax

import (
	"testing"
	"fmt"
)

type User struct{
	Id int
	Name string
}

func (user User) getId() (int){
	return user.Id;
}

func (user User) setId(id int){
	user.Id = id
}

/**
测试值得修改，和指针对应值修改效果
 */
func Test01(t *testing.T){
	var user1 = User{1,"tom"}
	//user2 := User{Id:2,Name:"tom"}
	fmt.Println("原始值",user1.Id,user1.Name)

	changeStructValue(user1);
	fmt.Println("change struct value之后：",user1.Id,user1.Name)

	changeStructPointer(&user1);
	fmt.Println("change struct point之后：",user1.Id,user1.Name)
}

/**
安装传统java bean set方式调用，发现外部bean属性没有被改变
struct传递到函数中的的还是struct值copy
 */
func Test02(t *testing.T){
	user := User{1,"tom"}
	fmt.Println(user.Id)
	fmt.Println(user.getId())
	//User.getId();
	user.setId(2)
	fmt.Println(user.getId())
	fmt.Println(User.getId(user))
}

/**
struct结构传的也是值得拷贝；
函数内部对sturct的修改不会影响外部值
 */
func changeStructValue(user User){
	user.Id = 15;
	fmt.Println(user.Id,user.Name)
}

/**
指针传的是内存地址；
函数内部对sturct的修改会影响外部值
 */
func changeStructPointer(userPoint *User){
	userPoint.Id = 20;
	fmt.Println(userPoint.Id,userPoint.Name)
}
